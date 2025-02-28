package kit

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/build"
	"github.com/xigxog/kubefox/core"
	"github.com/xigxog/kubefox/grpc"
	"github.com/xigxog/kubefox/kit/env"
	"github.com/xigxog/kubefox/logkf"
	"github.com/xigxog/kubefox/utils"
)

const (
	maxAttempts = 5
)

// TODO also support declarative routes? Example:
//
//	kit.RouteBuilder().
//	    Header("host", "google.com").
//	    Query("param1", "fish").
//	    Handler(myHandler)

type kit struct {
	spec *api.ComponentDetails

	routes     []*route
	defHandler EventHandler

	brk          *grpc.Client
	numWorkers   int
	maxEventSize int64

	export bool

	log *logkf.Logger
}

func New() Kit {
	svc := &kit{
		routes: make([]*route, 0),
		spec: &api.ComponentDetails{
			ComponentDefinition: api.ComponentDefinition{
				Type:             api.ComponentTypeKubeFox,
				Routes:           make([]api.RouteSpec, 0),
				VirtualEnvSchema: make(map[string]*api.VirtualEnvVarDefinition),
				Dependencies:     make(map[string]*api.Dependency),
			},
		},
	}

	comp := &core.Component{Id: core.GenerateId()}

	var help bool
	var brokerAddr, healthAddr, logFormat, logLevel string
	flag.StringVar(&comp.Name, "name", "", "Component name. (required)")
	flag.StringVar(&comp.Commit, "commit", "", "Commit the Component was built from. (required)")
	flag.StringVar(&brokerAddr, "broker-addr", "127.0.0.1:6060", "Address of the Broker gRPC server.")
	flag.StringVar(&healthAddr, "health-addr", "127.0.0.1:1111", `Address and port the HTTP health server should bind to, set to "false" to disable.`)
	flag.Int64Var(&svc.maxEventSize, "max-event-size", api.DefaultMaxEventSizeBytes, "Maximum size of event in bytes.")
	flag.IntVar(&svc.numWorkers, "num-workers", runtime.NumCPU(), "Number of worker threads to start, default is number of logical CPUs.")
	flag.StringVar(&logFormat, "log-format", "console", "Log format. [options 'json', 'console']")
	flag.StringVar(&logLevel, "log-level", "debug", "Log level. [options 'debug', 'info', 'warn', 'error']")
	flag.BoolVar(&svc.export, "export", false, "Exports component configuration in JSON and exits.")
	flag.BoolVar(&help, "help", false, "Show usage for component.")
	flag.Parse()

	if help {
		fmt.Fprintf(flag.CommandLine.Output(), `
Flags can be set using names below.

Flags:
`)
		flag.PrintDefaults()
		os.Exit(0)
	}

	if !svc.export {
		utils.CheckRequiredFlag("name", comp.Name)
		utils.CheckRequiredFlag("commit", comp.Commit)

		// TODO also check app (root) hash
		if comp.Commit != build.Info.Commit {
			fmt.Fprintf(os.Stderr, "commit '%s' does not match build info commit '%s'", comp.Commit, build.Info.Commit)
			os.Exit(1)
		}
	} else {
		logLevel = "error"
	}

	logkf.Global = logkf.
		BuildLoggerOrDie(logFormat, logLevel).
		WithComponent(comp)

	svc.log = logkf.Global
	svc.log.DebugInterface("build info:", build.Info)

	svc.brk = grpc.NewClient(grpc.ClientOpts{
		Component:     comp,
		BrokerAddr:    brokerAddr,
		HealthSrvAddr: healthAddr,
	})

	svc.log.Info("kit created 🦊")

	return svc
}

func (svc *kit) Log() *logkf.Logger {
	return svc.log
}

func (svc *kit) L() *logkf.Logger {
	return svc.log
}

func (svc *kit) Title(title string) {
	svc.spec.Title = title
}

func (svc *kit) Description(description string) {
	svc.spec.Title = description
}

func (svc *kit) Route(rule string, handler EventHandler) {
	r := &core.Route{
		RouteSpec: api.RouteSpec{
			Id:   len(svc.routes),
			Rule: rule,
		},
	}
	r.BuildEnvSchema()

	kitRoute := &route{
		RouteSpec: r.RouteSpec,
		handler:   handler,
	}
	svc.routes = append(svc.routes, kitRoute)
	svc.spec.Routes = append(svc.spec.Routes, r.RouteSpec)
}

func (svc *kit) Default(handler EventHandler) {
	svc.defHandler = handler
	svc.spec.DefaultHandler = handler != nil
}

func (svc *kit) EnvVar(name string, opts ...env.VarOption) EnvVarDep {
	if name == "" {
		svc.log.Fatal("environment variable name is required")
	}

	envSchema := &api.VirtualEnvVarDefinition{}
	for _, o := range opts {
		o(envSchema)
	}
	if envSchema.Type == "" {
		envSchema.Type = api.EnvVarTypeString
	}
	svc.spec.VirtualEnvSchema[name] = envSchema

	return env.NewVar(name, envSchema.Type)
}

func (svc *kit) Component(name string) ComponentDep {
	return svc.dependency(name, api.ComponentTypeKubeFox)
}

func (svc *kit) HTTPAdapter(name string) ComponentDep {
	return svc.dependency(name, api.ComponentTypeHTTP)
}

func (svc *kit) dependency(name string, typ api.ComponentType) ComponentDep {
	c := &dependency{
		typ:  typ,
		name: name,
	}
	svc.spec.Dependencies[name] = &api.Dependency{Type: typ}

	return c
}

func (svc *kit) Start() {
	defer logkf.Global.Sync()

	if err := svc.start(); err != nil {
		os.Exit(1)
	}
}

func (svc *kit) start() (err error) {
	if svc.export {
		c, _ := json.MarshalIndent(svc.spec, "", "  ")
		fmt.Println(string(c))
		os.Exit(0)
	}

	svc.log.DebugInterface("component spec:", svc.spec)

	if err = svc.brk.StartHealthSrv(); err != nil {
		svc.log.Errorf("error starting health server: %v", err)
		return
	}

	go svc.brk.Start(&svc.spec.ComponentDefinition, maxAttempts)

	var wg sync.WaitGroup
	wg.Add(svc.numWorkers)
	svc.log.Infof("starting %d workers", svc.numWorkers)

	for i := 0; i < svc.numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case req := <-svc.brk.Req():
					svc.recvReq(req)
				case err = <-svc.brk.Err(): // Sets start() err
					svc.log.Errorf("broker error: %v", err)
					return
				}
			}
		}()
	}
	wg.Wait()

	return
}

func (svc *kit) recvReq(req *grpc.ComponentEvent) {
	req.Event.ReduceTTL(req.ReceivedAt)

	ctx, cancel := context.WithTimeout(context.Background(), req.Event.TTL())
	defer cancel()

	log := svc.log.WithEvent(req.Event)

	ktx := &kontext{
		Event: req.Event,
		kit:   svc,
		env:   req.Env,
		start: time.Now(),
		ctx:   ctx,
		log:   log,
	}

	var err error
	switch {
	case req.RouteId == api.DefaultRouteId:
		if svc.defHandler == nil {
			err = core.ErrNotFound(fmt.Errorf("default handler not found"))
		} else {
			err = svc.defHandler(ktx)
		}

	case req.RouteId >= 0 && req.RouteId < int64(len(svc.routes)):
		err = svc.routes[req.RouteId].handler(ktx)

	default:
		err = core.ErrNotFound(fmt.Errorf("invalid route id %d", req.RouteId))
	}

	if err != nil {
		log.Error(err)

		errEvt := core.NewErr(err, core.EventOpts{})
		if err := ktx.Resp().Forward(errEvt); err != nil {
			log.Error(err)
		}
	}
}
