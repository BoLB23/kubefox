package kit

import (
	"context"
	"io"
	"net/http"
	"net/url"

	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/kit/env"
	"github.com/xigxog/kubefox/logkf"
)

type EventHandler func(ktx Kontext) error

type Kit interface {
	// Start connects to the Broker passing the Component's Service Account
	// Token to authenticate. Once connected Kit will accept incoming request
	// Events. If an error occurs the program will exit with a status code of 1.
	// Start is a blocking call.
	Start()

	// Route registers an EventHandler for the specified rule. If an incoming
	// Event matches the rule the Broker will route it to the Component and Kit
	// will call the EventHandler.
	//
	// Rules are written in a simple predicate based language that matches parts
	// of an Event. Some predicates accept inputs which should be surrounded
	// with back ticks. The boolean operators '&&' (and) and '||' (or) can be
	// used to combined predicates. Predicates can be negated with the '!' (not)
	// operator. The following predicates are supported:
	//
	//   All()
	//     Matches all Events.
	//
	//   Header(`key`, `value`)
	//     Matches if a header `key` exists and is equal to `value`.
	//
	//   Host(`example.com`)
	//     Matches if the domain (host header value) is equal to input.
	//
	//   Method(`GET`, ...)
	//     Matches if the request method is one of the given methods (GET, POST,
	//     PUT, DELETE, PATCH, HEAD)
	//
	//   Path(`/path`)
	//     Matches if the request path is equal to given input.
	//
	//   PathPrefix(`/prefix`)
	//     Matches if the request path begins with given input.
	//
	//   Query(`key`, `value`)
	//     Matches if a query parameter `key` exists and is equal to `value`.
	//
	//   Type(`value`)
	//     Matches if Event type is equal to given input.
	//
	// Predicate inputs can utilize regular expressions to match and optionally
	// extract parts of an Event to a named parameter. Regular expression use
	// the format '{[REGEX]}' or '{[NAME]:[REGEX]}' to extract the matching part
	// to a parameter.
	//
	// Additionally, environment variables can be utilized in predicate inputs.
	// They are resolved at request time with the value specified in the
	// VirtualEnv. Environment variables can be used with the format
	// '{{.Env.[NAME]}}'.
	//
	// For example, the following will match Events of type 'http' that are
	// 'GET' requests and have a path with three parts. The first part of the
	// path must equal the value of the environment variable 'SUB_PATH', the
	// second part must equal the string literal 'orders', and the third part
	// can be one or more lower case letters or numbers. The third part of the
	// path is extracted to the parameter 'orderId' which can be used by the
	// EventHandler:
	//
	//   kit.Route("Type(`http`) && Method(`GET`) && Path(`/{{.Env.SUB_PATH}}/orders/{orderId:[a-z0-9]+}`)",
	//     func(ktx kit.Kontext) error {
	//       return ktx.Resp().SendStr("The orderId is ", ktx.Param("orderId"))
	//     })
	Route(rule string, handler EventHandler)

	// Default registers a default EventHandler. If Kit receives an Event from
	// the Broker that does not match any registered rules the default
	// EventHandler is called.
	Default(handler EventHandler)

	// EnvVar registers an environment variable dependency with given options.
	// The returned EnvVarDep can be used by EventHandlers to retrieve the value
	// of the environment variable at request time.
	//
	// For example:
	//
	//   v := kit.EnvVar("SOME_VAR")
	//   kit.Route("Any()", func(ktx kit.Kontext) error {
	//       return ktx.Resp().SendStr("the value of SOME_VAR is ", ktx.Env(v))
	//   })
	EnvVar(name string, opts ...env.VarOption) EnvVarDep

	// Component registers a Component dependency. The returned ComponentDep can
	// be used by EventHandlers to invoke the Component at request time.
	//
	// For example:
	//
	//   b := kit.Component("backend")
	//   kit.Route("Any()", func(ktx kit.Kontext) error {
	//       r, _ := ktx.Req(backend).Send()
	//       return ktx.Resp().SendStr("the resp from backend is ", r.Str())
	//   })
	Component(name string) ComponentDep

	// HTTPAdapter registers a dependency on the named HTTP Adapter. The
	// returned ComponentDep can be used by EventHandlers to invoke the Adapter
	// at request time.
	//
	// For example:
	//
	//   h := kit.HTTPAdapter("httpbin")
	//   kit.Route("Any()", func(ktx kit.Kontext) error {
	//       r, _ := ktx.HTTP(h).Get("/anything")
	//       return ktx.Resp().SendReader(r.Header.Get("content-type"), r.Body)
	//   })
	HTTPAdapter(name string) ComponentDep

	// Title sets the Component's title.
	Title(title string)

	// Description sets the Component's description.
	Description(description string)

	// Log returns a pre-configured structured logger for the Component.
	Log() *logkf.Logger
}

type Kontext interface {
	EventReader

	// Env returns the value of the given environment variable as a string. If
	// the environment variable does not exist or cannot be converted to a
	// string, empty string is returned. To check if an environment exists use
	// EnvV() and check if the returned api.Val's is 'Nil'.
	Env(v EnvVarDep) string

	// EnvV returns the value of the given environment variable as an api.Val.
	// It is guaranteed the returned api.Val will not be nil. If the environment
	// variable does not exist the api.ValType of the returned api.Val will be
	// 'Nil'.
	EnvV(v EnvVarDep) *api.Val

	// EnvDef returns the value of the given environment variable as a string.
	// If the environment variable does not exist, is empty, or cannot be
	// converted to a string, then the given 'def' string is returned.
	EnvDef(v EnvVarDep, def string) string

	// EnvDefV returns the value of the given environment variable as an
	// api.Val. If the environment variable does not exist, is an empty string,
	// or an empty array, then the given 'def' api.Val is returned. If the
	// environment variable exists and is a boolean or number, then it's value
	// will be returned.
	EnvDefV(v EnvVarDep, def *api.Val) *api.Val

	// Resp returns a Resp object that can be used to send a response Event to
	// the source of the current request.
	Resp() Resp

	// Req returns an empty Req object that can be used to send a request Event
	// to the given target Component.
	Req(target ComponentDep) Req

	// Forward returns a clone of the current request Event as a Req object that
	// can be used to send it to the given target Component.
	Forward(target ComponentDep) Req

	// HTTP returns a native Go http.Client. Any requests made with the client
	// are sent to the given target Component. The target should be capable of
	// processing HTTP requests.
	HTTP(target ComponentDep) *http.Client

	// HTTP returns a native go http.RoundTripper. This is useful to integrate
	// with HTTP based libraries.
	Transport(target ComponentDep) http.RoundTripper

	// Context returns a context.Context with it's duration set to the TTL of
	// the current request.
	Context() context.Context

	// Log returns a pre-configured structured logger for the current request.
	Log() *logkf.Logger
}

type Req interface {
	EventWriter

	// SendStr sends the request to the target Component and returns the
	// response. The given string is used as the content of the request Event,
	// content-type is set to 'text/plain'.
	SendStr(s string) (EventReader, error)

	// SendHTML sends the request to the target Component and returns the
	// response. The given HTML is used as the content of the request Event,
	// content-type is set to 'text/html'.
	SendHTML(h string) (EventReader, error)

	// SendJSON sends the request to the target Component and returns the
	// response. The given object is marshalled to JSON and the output is used
	// as the content of the request Event, content-type is set to
	// 'application/json'.
	SendJSON(v any) (EventReader, error)

	// SendBytes sends the request to the target Component using the given
	// content-type and content and returns the response.
	SendBytes(contentType string, content []byte) (EventReader, error)

	// SendReader sends the request to the target Component and returns the
	// response. All data is read from the given reader and is used as the
	// content of the request Event. If the reader implements io.ReadCloser then
	// it will be automatically closed.
	SendReader(contentType string, reader io.Reader) (EventReader, error)

	// Send sends the request to the target Component and returns the response.
	Send() (EventReader, error)
}

type Resp interface {
	EventWriter

	// Forward sends the response to the source Component of the current
	// request. The response Event is a clone of the given Event.
	Forward(evt EventReader) error

	// SendStr sends the response to the source Component of the current
	// request. The given string is used as the content of the response Event,
	// content-type is set to 'text/plain'.
	SendStr(s ...string) error

	// SendHTML sends the response to the source Component of the current
	// request. The given HTML is used as the content of the response Event,
	// content-type is set to 'text/html'.
	SendHTML(h string) error

	// SendJSON sends the response to the source Component of the current
	// request. The given object is marshalled to JSON and the output is used as
	// the content of the response Event, content-type is set to
	// 'application/json'.
	SendJSON(v any) error

	// SendAccepts sends the response to the source Component using one of the
	// given values based on the 'Accept' header of the request Event.
	SendAccepts(json any, html, str string) error

	// SendBytes sends the response to the source Component of the current
	// request using the given content-type and content.
	SendBytes(contentType string, b []byte) error

	// SendReader sends the response to the source Component of the current
	// request. All data is read from the given reader and is used as the
	// content of the response Event. If the reader implements io.ReadCloser
	// then it will be automatically closed.
	SendReader(contentType string, reader io.Reader) error

	// Send sends the response to the source Component of the current request.
	Send() error
}

type EventReader interface {
	EventType() api.EventType

	Param(key string) string
	ParamV(key string) *api.Val
	ParamDef(key string, def string) string

	URL() (*url.URL, error)

	Query(key string) string
	QueryV(key string) *api.Val
	QueryDef(key string, def string) string
	QueryAll(key string) []string

	Header(key string) string
	HeaderV(key string) *api.Val
	HeaderDef(key string, def string) string
	HeaderAll(key string) []string

	Status() int
	StatusV() *api.Val

	Bind(v any) error
	Str() string
	Bytes() []byte
}

type EventWriter interface {
	EventReader

	SetParam(key, value string)
	SetParamV(key string, value *api.Val)

	SetURL(u *url.URL)
	TrimPathPrefix(prefix string)

	SetQuery(key, value string)
	SetQueryV(key string, value *api.Val)
	DelQuery(key string)

	SetHeader(key, value string)
	SetHeaderV(key string, value *api.Val)
	AddHeader(key, value string)
	DelHeader(key string)

	SetStatus(code int)
	SetStatusV(val *api.Val)
}

type EnvVarDep interface {
	Name() string
	Type() api.EnvVarType
}

type ComponentDep interface {
	Name() string
	Type() api.ComponentType
	EventType() api.EventType
}

type route struct {
	api.RouteSpec

	handler EventHandler
}

type dependency struct {
	typ  api.ComponentType
	name string
}

func (c *dependency) Name() string {
	return c.name
}

func (c *dependency) Type() api.ComponentType {
	return c.typ
}

func (c *dependency) EventType() api.EventType {
	switch c.typ {
	case api.ComponentTypeHTTP:
		return api.EventTypeHTTP
	default:
		return api.EventTypeKubeFox
	}
}
