module {{ .gitserver }}/{{ .organization }}/{{ .name }}

go 1.12
require (
	{{- if .prometheus }}
	github.com/Depado/ginprom v1.1.1
	{{- end }}
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/gin v1.4.0
	github.com/onrik/logrus v0.3.0
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
)
