package nginx_http

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/grafana/agent/pkg/integrations"
	"github.com/nginxinc/nginx-prometheus-exporter/client"
	nc "github.com/nginxinc/nginx-prometheus-exporter/collector"
)

type Config struct {
	// NginxAddr is the URI to Nginx stub status page.
	NginxAddr string `yaml:"scrape_uri,omitempty"`

	// NginxNamespace is the prefix for the metrics names.
	NginxNamespace string

	// HttpInsecure is the flag to skip SSL verification.
	HttpInsecure bool `yaml:"insecure,omitempty"`

	// HttpTimeout is the timeout for the HTTP client.
	HttpTimeout time.Duration
}

// DefaultConfig holds the default settings for the nginx_http integration
var DefaultConfig = Config{
	NginxAddr:      "http://localhost/nginx_status",
	NginxNamespace: "nginx",
	HttpInsecure:   false,
	HttpTimeout:    0,
}

// UnmarshalYAML implements yaml.Unmarshaler for Config
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	*c = DefaultConfig

	type plain Config
	return unmarshal((*plain)(c))
}

// Name returns the name of the integration this config is for.
func (c *Config) Name() string {
	return "nginx_http"
}

// InstanceKey returns the addr of the nginx server.
func (c *Config) InstanceKey(agentKey string) (string, error) {
	u, err := url.Parse(c.NginxAddr)
	if err != nil {
		return "", err
	}
	return u.Host, nil
}

// NewIntegration converts the config into an integration instance.
func (c *Config) NewIntegration(logger log.Logger) (integrations.Integration, error) {
	return New(logger, c)
}

func init() {
	integrations.RegisterIntegration(&Config{})
}

// New creates a new nginx_http integration. The integration scrapes metrics
// from a Nginx HTTP server.
func New(logger log.Logger, c *Config) (integrations.Integration, error) {
	_, err := url.ParseRequestURI(c.NginxAddr)
	if err != nil {
		level.Error(logger).Log("msg", "scrape_uri is invalid", "err", err)
		return nil, err
	}

	httpTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.HttpInsecure},
	}
	httpClient := &http.Client{
		Timeout:   c.HttpTimeout,
		Transport: httpTransport,
	}

	nginxClient := client.NewNginxClient(httpClient, c.NginxAddr)
	nginxCollector := nc.NewNginxCollector(nginxClient, c.NginxNamespace, nil, logger)

	return integrations.NewCollectorIntegration(
		c.Name(),
		integrations.WithCollectors(nginxCollector),
	), nil
}
