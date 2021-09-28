# Gitpod's Mixin

Gitpod's mixin is based on the [Prometheus Monitoring Mixins project](https://github.com/monitoring-mixins/docs/blob/master/design.pdf). Mixins are jsonnet packages that bundles together [Prometheus Alerts](https://prometheus.io/docs/prometheus/latest/configuration/alerting_rules/), [Prometheus Recording Rules](https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/) and [Grafana Dashboards](https://grafana.com/grafana/).

## Table of contents

* [Folders and Teams](#Folders-and-Teams)
* [How to develop Dashboards](#How-to-develop-Dashboards)
    * [Grafonnet](#Grafonnet)
    * [Exporting JSONs from Grafana UI](#Exporting-JSONs-from-Grafana-UI)
* [How to develop Prometheus Alerts and Rules](#How-to-develop-Prometheus-Alerts-and-Rules)
* [Frequently asked Questions](#FAQ)
    * [How is our mixin consumed?](#How-is-our-mixin-consumed)
    * [How do I review dashboards before merging PRs?](#How-do-I-review-dashboards-before-merging-PRs)


## Folders and Teams

Folder are organized following Gitpod as an organization, while also adding an extra folder for dashboards and alerts that involves multiple teams (Good place for broad overviews and SLOs):
* Meta
* Workspace
* IDE
* Cross-Teams

We've organized our mixins to make it easier to each team to own their own dashboards and alerts. Every team has its own folder with a `mixin.libsonnet` file, which imports all dashboards and alerts from the subfolders.

It doesn't matter how the imports inside the subfolders work, it is only important that all dashboards end up in a `grafanaDashboars` object, all alerts in the `prometheusAlerts` object and all recording rules in the `prometheusRules` object. [Read more about jsonnet objects](https://jsonnet.org/ref/language.html).

From past experiences, the platform team suggests that dashboards and alerts get split by component inside the subfolders because, so far, we haven't implemented metrics the involves more than a single component operation.


## How to develop Dashboards

### Grafonnet

Grafana provides a jsonnet library, called [Grafonnet](https://github.com/grafana/grafonnet-lib/tree/master/grafonnet), that can help us develop Grafana dashboards while using amount of code low.

Instead of creating a gigantic JSON, you can use grafonnet and make a dashboard that is a lot easier to review in Pull Requests. For example:

```jsonnet
// content of my-team/dashboards/my-new-dashboard.libsonnet
local grafana = import 'grafonnet/grafana.libsonnet';
local dashboard = grafana.dashboard;
local row = grafana.row;
local prometheus = grafana.prometheus;
local template = grafana.template;
local graphPanel = grafana.graphPanel;
local template = grafana.template;

local datasourceVariable = {
  hide: 0,
  name: 'datasource',
  query: 'prometheus',
  refresh: 1,
  type: 'datasource',
};

local runningWorkspacesGraph =
  graphPanel.new(
    'Running Workspaces',
    datasource='$datasource',
    format='none',
    stack=false,
    fill=1,
    fillGradient=5,
    min=0,
  )
  .addTarget(
      prometheus.target(
          |||
            sum(
              gitpod_ws_manager_workspace_phase_total{phase="RUNNING"}
            ) by (type)
          |||, legendFormat='{{ type }}'
      )
  );

{
    'my-new-dashboard.json': dashboard.new(
    'My new dashboard!',
    time_from='now-1h',
    timezone='utc',
    refresh='30s',
    )
    .addTemplate(datasourceVariable)
    .addPanel(runningWorkspacesGraph)
    ,
}




/*******************************/
// content of my-team/dashboards.libsonnet
{
    grafanaDashboards+:: {
        // other dashboards
        ...
        ...
        ...
        'my-new-dashboard.json': (import 'my-team/dashboards/my-new-dashboard.libsonnet'),
    }
}
```

To make sure your jsonnet code compiles and is well-formated, you can always run `cd operations/observability/mixins && make lint`.

You can also use our [preview environments to make sure the dashboard really looks like what you imagined](#How-do-I-review-dashboards-before-merging-PRs).

### Exporting JSONs from Grafana UI

## How to develop Prometheus Alerts and Rules

## FAQ

### How is our mixin consumed?

### How do I review dashboards before merging PRs?

There is a couple o ways to trigger a werft job that will deploy a preview environment with Prometheus+Grafana with your changes:

1. You can open a Pull Request with the following line in the description:
```
/werft with-observability

# Just in case your PR requires extra configuration on Prometheus side
# (and you have a new branch on https://github.com/gitpod-io/observability with such changes)
# you can add the line below
/werft withObservabilityBranch="<my-branch>"
```

2. After opening a Pull Request, you can add a Github comment:
```
/werft run with-observability
/werft run withObservabilityBranch="<my-branch>"
```

3. Inside your workspace, run:
```
werft run github -a with-observability="" -a withObservabilityBranch="<my-branch>"
```

As mentioned in [How is our mixin consumed?](#How-is-our-mixin-consumed), please remember that a commit must be done for us to update monitoring-satellite with the dashboards/alerts/recording rule changes.

Please remember that the annotation `withObservabilityBranch` is completely optional, and most of the times you won't need it at all.