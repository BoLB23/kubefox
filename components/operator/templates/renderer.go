package templates

import (
	"embed"
	"encoding/json"
	"path"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/xigxog/kubefox/utils"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/yaml"
)

//go:embed all:*
var EFS embed.FS

func Render(name string, data *Data) ([]*unstructured.Unstructured, error) {
	rendered, err := renderStr("list.tpl", name+"/*", data)
	if err != nil {
		return nil, err
	}

	resList := &ResourceList{}
	if err := yaml.Unmarshal([]byte(rendered), resList); err != nil {
		return nil, err
	}

	return resList.Items, nil
}

func renderStr(tplFile, path string, data *Data) (string, error) {
	tpl := template.New(tplFile).Option("missingkey=zero")
	initFuncs(tpl, data)

	if _, err := tpl.ParseFS(EFS, "helpers.tpl", path); err != nil {
		return "", err
	}

	var buf strings.Builder
	if err := tpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return strings.ReplaceAll(buf.String(), "<no value>", ""), nil
}

func initFuncs(tpl *template.Template, data *Data) {
	funcMap := sprig.FuncMap()

	funcMap["include"] = func(name string, data interface{}) (string, error) {
		var buf strings.Builder
		err := tpl.ExecuteTemplate(&buf, name, data)
		return buf.String(), err
	}

	funcMap["toYaml"] = func(v any) string {
		data, err := yaml.Marshal(v)
		if err != nil {
			return ""
		}
		s := strings.TrimSuffix(string(data), "\n")
		if s == "null" {
			s = ""
		}
		return s
	}

	funcMap["toJSON"] = func(v any) string {
		data, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		s := strings.TrimSuffix(string(data), "\n")
		if s == "null" {
			s = ""
		}
		return s
	}

	funcMap["file"] = func(name string) string {
		if s, err := renderStr(path.Base(name), name, data); err != nil {
			return ""
		} else {
			return s
		}
	}

	funcMap["cleanLabel"] = utils.CleanLabel
	funcMap["name"] = data.Name
	funcMap["namespace"] = data.Namespace
	funcMap["platformFullName"] = data.PlatformFullName
	funcMap["platformVaultName"] = data.PlatformVaultName
	funcMap["componentFullName"] = data.ComponentFullName
	funcMap["componentVaultName"] = data.ComponentVaultName
	funcMap["homePath"] = data.HomePath

	tpl.Funcs(funcMap)
}
