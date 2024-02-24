---
aliases:
  - /docs/grafana-cloud/agent/flow/reference/components/prometheus.exporter.nginx/
  - /docs/grafana-cloud/monitor-infrastructure/agent/flow/reference/components/prometheus.exporter.nginx/
  - /docs/grafana-cloud/monitor-infrastructure/integrations/agent/flow/reference/components/prometheus.exporter.nginx/
  - /docs/grafana-cloud/send-data/agent/flow/reference/components/prometheus.exporter.nginx/
canonical: https://grafana.com/docs/agent/latest/flow/reference/components/prometheus.exporter.nginx/
description: Learn about prometheus.exporter.nginx
title: prometheus.exporter.nginx
---

# prometheus.exporter.nginx

The `prometheus.exporter.nginx` component embeds
[`nginx_prometheus_exporter`](https://github.com/nginxinc/nginx-prometheus-exporter) for
collecting [stub_status](https://nginx.org/en/docs/http/ngx_http_stub_status_module.html#stub_status)
statistics from a nginx server.

## Usage

```river
prometheus.exporter.nginx "LABEL" {
}
```

## Arguments

The following arguments can be used to configure the exporter's behavior.
All arguments are optional. Omitted fields take their default values.

| Name         | Type     | Description                               | Default                         | Required |
|--------------|----------|-------------------------------------------|---------------------------------|----------|
| `scrape_uri` | `string` | URI to Nginx stub status page.            | `http://localhost/nginx_status` | no       |
| `insecure`   | `bool`   | Ignore server certificate if using https. | 0                               | false    |

## Exported fields

{{< docs/shared lookup="flow/reference/components/exporter-component-exports.md" source="
agent" version="<AGENT_VERSION>" >}}

## Component health

`prometheus.exporter.nginx` is only reported as unhealthy if given
an invalid configuration. In those cases, exported fields retain their last
healthy values.

## Debug information

`prometheus.exporter.nginx` does not expose any component-specific
debug information.

## Debug metrics

`prometheus.exporter.nginx` does not expose any component-specific
debug metrics.

## Example

This example uses a [`prometheus.scrape` component][scrape] to collect metrics
from `prometheus.exporter.nginx`:

```river
prometheus.exporter.nginx "example" {
  scrape_uri = "http://web.example.com/nginx_status"
}

// Configure a prometheus.scrape component to collect nginx metrics.
prometheus.scrape "demo" {
  targets    = prometheus.exporter.nginx.example.targets
  forward_to = [prometheus.remote_write.demo.receiver]
}

prometheus.remote_write "demo" {
  endpoint {
    url = PROMETHEUS_REMOTE_WRITE_URL

    basic_auth {
      username = USERNAME
      password = PASSWORD
    }
  }
}
```

Replace the following:

- `PROMETHEUS_REMOTE_WRITE_URL`: The URL of the Prometheus remote_write-compatible server
  to send metrics to.
- `USERNAME`: The username to use for authentication to the remote_write API.
- `PASSWORD`: The password to use for authentication to the remote_write API.

[scrape]: {{< relref "./prometheus.scrape.md" >}}

<!-- START GENERATED COMPATIBLE COMPONENTS -->

## Compatible components

`prometheus.exporter.nginx` has exports that can be consumed by the following components:

- Components that consume [Targets]({{< relref "../compatibility/#targets-consumers" >}})

{{< admonition type="note" >}}
Connecting some components may not be sensible or components may require further
configuration to make the connection work correctly.
Refer to the linked documentation for more details.
{{< /admonition >}}

<!-- END GENERATED COMPATIBLE COMPONENTS -->
