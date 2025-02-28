package kit

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/xigxog/kubefox/api"
	"github.com/xigxog/kubefox/core"
	"github.com/xigxog/kubefox/logkf"

	"google.golang.org/protobuf/types/known/structpb"
)

type kontext struct {
	*core.Event

	kit *kit
	env map[string]*structpb.Value

	start time.Time
	ctx   context.Context
	log   *logkf.Logger
}

type respKontext struct {
	*core.Event

	ktx *kontext
}

type reqKontext struct {
	*core.Event

	ktx *kontext
}

func (k *kontext) Context() context.Context {
	return k.ctx
}

func (k *kontext) Log() *logkf.Logger {
	return k.log
}

func (k *kontext) Env(v EnvVarDep) string {
	return k.EnvV(v).String()
}

func (k *kontext) EnvV(v EnvVarDep) *api.Val {
	val, _ := api.ValProto(k.env[v.Name()])
	return val
}

func (k *kontext) EnvDef(v EnvVarDep, def string) string {
	if val := k.Env(v); val == "" {
		return def
	} else {
		return val
	}
}

func (k *kontext) EnvDefV(v EnvVarDep, def *api.Val) *api.Val {
	if val := k.EnvV(v); val.IsEmpty() {
		return def
	} else {
		return val
	}
}

func (k *kontext) Resp() Resp {
	return &respKontext{
		Event: core.NewResp(core.EventOpts{
			Parent: k.Event,
			Source: k.kit.brk.Component,
			Target: k.Event.Source,
		}),
		ktx: k,
	}
}

func (k *kontext) Req(target ComponentDep) Req {
	return &reqKontext{
		Event: core.NewReq(core.EventOpts{
			Type:   target.EventType(),
			Parent: k.Event,
			Source: k.kit.brk.Component,
			Target: &core.Component{Name: target.Name()},
		}),
		ktx: k,
	}
}

func (k *kontext) Forward(target ComponentDep) Req {
	return &reqKontext{
		Event: core.CloneToReq(k.Event, core.EventOpts{
			Parent: k.Event,
			Source: k.kit.brk.Component,
			Target: &core.Component{Name: target.Name()},
		}),
		ktx: k,
	}
}

func (k *kontext) HTTP(target ComponentDep) *http.Client {
	return &http.Client{
		Transport: k.Transport(target),
	}
}

func (k *kontext) Transport(target ComponentDep) http.RoundTripper {
	return &EventRoundTripper{
		req: &reqKontext{
			Event: core.NewReq(core.EventOpts{
				Type:   target.EventType(),
				Parent: k.Event,
				Source: k.kit.brk.Component,
				Target: &core.Component{Name: target.Name()},
			}),
			ktx: k,
		},
	}
}

func (resp *respKontext) Forward(evt EventReader) error {
	resp.Event = core.CloneToResp(evt.(*core.Event), core.EventOpts{
		Parent: resp.ktx.Event,
		Source: resp.ktx.kit.brk.Component,
		Target: resp.ktx.Event.Source,
	})

	return resp.Send()
}

func (resp *respKontext) SendStr(val ...string) error {
	c := fmt.Sprintf("%s; %s", api.ContentTypePlain, api.CharSetUTF8)
	return resp.SendBytes(c, []byte(strings.Join(val, "")))
}

func (resp *respKontext) SendHTML(val string) error {
	c := fmt.Sprintf("%s; %s", api.ContentTypeHTML, api.CharSetUTF8)
	return resp.SendBytes(c, []byte(val))
}

func (resp *respKontext) SendJSON(val any) error {
	if err := resp.SetJSON(val); err != nil {
		return err
	}

	return resp.Send()
}

func (resp *respKontext) SendAccepts(json any, html, str string) error {
	a := strings.ToLower(resp.ktx.Header("accept"))
	switch {
	case strings.Contains(a, "application/json"):
		return resp.SendJSON(json)

	case strings.Contains(a, "text/html"):
		return resp.SendHTML(html)

	default:
		return resp.SendStr(str)
	}
}

func (resp *respKontext) SendReader(contentType string, reader io.Reader) error {
	if closer, ok := reader.(io.ReadCloser); ok {
		defer closer.Close()
	}

	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	return resp.SendBytes(contentType, bytes)
}

func (resp *respKontext) SendBytes(contentType string, b []byte) error {
	resp.Event.ContentType = contentType
	resp.Event.Content = b

	return resp.Send()
}

func (resp *respKontext) Send() error {
	return resp.ktx.kit.brk.SendResp(resp.Event, resp.ktx.start)
}

func (req *reqKontext) SendStr(val string) (EventReader, error) {
	c := fmt.Sprintf("%s; %s", api.ContentTypePlain, api.CharSetUTF8)
	return req.SendBytes(c, []byte(val))
}

func (req *reqKontext) SendHTML(val string) (EventReader, error) {
	c := fmt.Sprintf("%s; %s", api.ContentTypeHTML, api.CharSetUTF8)
	return req.SendBytes(c, []byte(val))
}

func (req *reqKontext) SendJSON(val any) (EventReader, error) {
	if err := req.SetJSON(val); err != nil {
		return nil, err
	}

	return req.Send()
}

func (req *reqKontext) SendReader(contentType string, reader io.Reader) (EventReader, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return req.SendBytes(contentType, bytes)
}

func (req *reqKontext) SendBytes(contentType string, b []byte) (EventReader, error) {
	req.Event.ContentType = contentType
	req.Event.Content = b

	return req.Send()
}

func (req *reqKontext) Send() (EventReader, error) {
	return req.ktx.kit.brk.SendReq(req.ktx.ctx, req.Event, req.ktx.start)
}
