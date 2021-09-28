# Gitpod's Mixin

Gitpod's mixin is based on the [Prometheus Monitoring Mixins project](https://github.com/monitoring-mixins/docs/blob/master/design.pdf). Mixins are jsonnet packages that bundles together [Prometheus Alerts](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/), [Prometheus Recording Rules](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) and [Grafana Dashboards](https://grafana.com/grafana/).

## Table of contents

* [Folders and Teams](#Folders-and-Teams)
* [Frequently asked Questions](#FAQ)
    * [How is our mixin consumed?](#How-is-our-mixin-consumed)
    * [How do I changes to dashboards before merging my PR?](#How-do-I-changes-to-dashboards-before-merging-my-PR)


## Folders and Teams

Folder are organized following Gitpod as an organization, while also adding an extra folder for dashboards and alerts that involves multiple teams (Good place for broad overviews and SLOs):
* Meta
* Workspace
* IDE
* Cross-Teams

We've organized our mixins to make it easier to each team to own their own dashboards and alerts. Every team has its own folder with a `mixin.libsonnet` file, which imports all dashboards and alerts from the subfolders.

It doesn't matter how the imports inside the subfolders work, it is only important that all dashboards end up in a `grafanaDashboars` object, all alerts in the `prometheusAlerts` object and all recording rules in the `prometheusRules` object. [Read more about jsonnet objects](https://jsonnet.org/ref/language.html).

From past experiences, the platform team suggests that dashboards and alerts get split by component inside the subfolders because, so far, we haven't implemented metrics the involves more than a single component operation.





## FAQ

### How is our mixin consumed?

### How do I changes to dashboards before merging my PR?