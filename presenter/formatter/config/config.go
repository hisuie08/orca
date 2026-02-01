package config

import (
	"bytes"
	"orca/model/config"
	"text/template"
)

var orcayamltmpl = `name: {{.Name}}
volume:
    volume_root: {{if eq .Volume.VolumeRoot nil -}}
	null
	{{- else -}}
	{{- .Volume.VolumeRoot -}}
{{end}}
    ensure_path: {{.Volume.EnsurePath}}
network:
    enabled: {{.Network.Enabled}}
    internal: {{.Network.Internal}}
    name: {{.Network.Name}}
`

func FmtConfig(r config.OrcaConfig) string {
	b := bytes.NewBuffer([]byte{})
	tmpl, err := template.New("config").Parse(orcayamltmpl)
	if err != nil {
		panic(err)
	}
	tmpl.Execute(b, r)
	return b.String()
}
