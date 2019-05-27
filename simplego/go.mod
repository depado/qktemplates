module {{ .gitserver }}/{{ .organization }}/{{ .name }}

go 1.12
{{ if .cli }}
require (
	github.com/onrik/logrus v0.3.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
)
{{ end }}