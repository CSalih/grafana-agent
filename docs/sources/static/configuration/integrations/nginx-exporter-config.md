---
aliases:
- ../../../configuration/integrations/nginx-exporter-config/
- /docs/grafana-cloud/monitor-infrastructure/agent/static/configuration/integrations/nginx-exporter-config/
- /docs/grafana-cloud/send-data/agent/static/configuration/integrations/nginx-exporter-config/
canonical: https://grafana.com/docs/agent/latest/static/configuration/integrations/nginx-exporter-config/
description: Learn about nginx_http_config
title: nginx_http_config
---

# nginx_http_config

The `nginx_http_config` block configures the `nginx_http` integration,
which is an embedded version of
[`nginx_prometheus_exporter`](https://github.com/nginxinc/nginx-prometheus-exporter). This allows the collection of Nginx [stub_status](https://nginx.org/en/docs/http/ngx_http_stub_status_module.html#stub_status) statistics via HTTP.

Full reference of options:

```yaml
  # Enables the nginx_http integration, allowing the Agent to automatically
  # collect metrics for the specified nginx http servers.
  [enabled: <boolean> | default = false]

  # Sets an explicit value for the instance label when the integration is
  # self-scraped. Overrides inferred values.
  #
  # The default value for this integration is inferred from the hostname portion
  # of api_url.
  [instance: <string>]

  # Automatically collect metrics from this integration. If disabled,
  # the nginx_http integration will be run but not scraped and thus not
  # remote-written. Metrics for the integration will be exposed at
  # /integrations/nginx_http/metrics and can be scraped by an external
  # process.
  [scrape_integration: <boolean> | default = <integrations_config.scrape_integrations>]

  # How often should the metrics be collected? Defaults to
  # prometheus.global.scrape_interval.
  [scrape_interval: <duration> | default = <global_config.scrape_interval>]

  # The timeout before considering the scrape a failure. Defaults to
  # prometheus.global.scrape_timeout.
  [scrape_timeout: <duration> | default = <global_config.scrape_timeout>]

  # Allows for relabeling labels on the target.
  relabel_configs:
    [- <relabel_config> ... ]

  # Relabel metrics coming from the integration, allowing to drop series
  # from the integration that you don't care about.
  metric_relabel_configs:
    [ - <relabel_config> ... ]

  # How frequent to truncate the WAL for this integration.
  [wal_truncate_frequency: <duration> | default = "60m"]

  #
  # Exporter-specific configuration options
  #

  # URI to nginx stub status page.
  [scrape_uri: <string> | default = "http://localhost/nginx_status"]

  # Ignore server certificate if using https.
  [insecure: <boolean> | default = false]
```
