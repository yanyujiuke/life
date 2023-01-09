package global

{{- if .HasGlobal }}

import "life/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}