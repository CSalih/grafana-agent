package nginx

import (
	"time"

	"github.com/grafana/agent/component"
	"github.com/grafana/agent/component/prometheus/exporter"
	"github.com/grafana/agent/pkg/integrations"
	"github.com/grafana/agent/pkg/integrations/nginx_http"
)

func init() {
	component.Register(component.Registration{
		Name:    "prometheus.exporter.nginx",
		Args:    Arguments{},
		Exports: exporter.Exports{},

		Build: exporter.New(createExporter, "nginx"),
	})
}

func createExporter(opts component.Options, args component.Arguments, defaultInstanceKey string) (integrations.Integration, string, error) {
	a := args.(Arguments)
	return integrations.NewIntegrationWithInstanceKey(opts.Logger, a.Convert(), defaultInstanceKey)
}

// DefaultArguments holds the default settings for the nginx prometheus exporter
var DefaultArguments = Arguments{
	NginxAddr:      nginx_http.DefaultConfig.NginxAddr,
	NginxNamespace: nginx_http.DefaultConfig.NginxNamespace,
	HttpTimeout:    nginx_http.DefaultConfig.HttpTimeout,
}

// Arguments controls the nginx exporter.
type Arguments struct {
	// NginxAddr is the URI to Nginx stub status page.
	NginxAddr string `river:"scrape_uri,attr,optional"`

	// NginxNamespace is the prefix for the metrics names.
	NginxNamespace string

	// HttpInsecure is the flag to skip SSL verification.
	HttpInsecure bool `yaml:"insecure,attr,optional"`

	// HttpTimeout is the timeout for the HTTP client.
	HttpTimeout time.Duration
}

// SetToDefault implements river.Defaulter.
func (a *Arguments) SetToDefault() {
	*a = DefaultArguments
}

func (a *Arguments) Convert() *nginx_http.Config {
	return &nginx_http.Config{
		NginxAddr:      a.NginxAddr,
		NginxNamespace: a.NginxNamespace,
		HttpInsecure:   a.HttpInsecure,
		HttpTimeout:    a.HttpTimeout,
	}
}
