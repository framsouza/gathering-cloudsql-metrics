# Gathering Google Cloud SQL metrics

This code will connect to the Cloud Monitoring API, gather the metrics from Google Cloud SQL and send it to a Cloud Pub/Sub. This function executes each 60 seconds, right now, the metrics that are being collected are:

- disk utilization
- memory utilization
- cpu utilization
- active connections 
- total memory

For more metrics, please check [Google Cloud SQL metrics](https://cloud.google.com/monitoring/api/metrics_gcp#gcp-cloudsql).

### Usage

First, make sure you have the `GOOGLE_APPLICATION_CREDENTIALS` environment variable set as per the [google docs](https://cloud.google.com/docs/authentication/production).

It will require one argument which is `project` and `topic`. Make sure to create the topic before run this code, it won't automatically create it.

You can choose to build the binary by running
```
go build -o gathering-cloudsql-metrics
```

and run it as
```
./gathering-cloudsql-metrics -project=<PROJECTNAME> -topic=test-topic
```

or without building the binary
```
go run main.go -project=<PROJECTNAME>
```

## Output expected
You will have an output like the following:

```
Connecting with topic: pull-cloud-monitoring-db01-sub

CloudSQL CPU Utilization (%)
starttime: 2022-08-22T07:28:00Z endttime: 2022-08-22T07:29:00Z name: my-gcp-project:db-dev value: 2.5237555640205755
Published a message; msg ID: 5371724353696475
starttime: 2022-08-22T07:28:00Z endttime: 2022-08-22T07:29:00Z name: my-gcp-project:db01 value: 2.344554992491794
Published a message; msg ID: 5371720838171288

CloudSQL memory total usage
starttime: 2022-08-22T07:28:00Z endttime: 2022-08-22T07:29:00Z name my-gcp-project:db-dev value: 0 GB
Published a message; msg ID: 5371723455172154
starttime: 2022-08-22T07:28:00Z endttime: 2022-08-22T07:29:00Z name my-gcp-project:db01 value: 0 GB
Published a message; msg ID: 5371720838169768

CloudSQL memory total size
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name my-gcp-project:db-dev value: 8 GB
Published a message; msg ID: 5371721896997627
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name my-gcp-project:db01 value: 8 GB
Published a message; msg ID: 5371721398656734

CloudSQL Active Connections
starttime: 2022-08-22T07:28:00Z endttime: 2022-08-22T07:29:00Z name: my-gcp-project:db01 value: 4
Published a message; msg ID: 5371721741725995
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name: my-gcp-project:db-dev value: 2
Published a message; msg ID: 5371722466560998
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name: my-gcp-project:db-dev value: 0
Published a message; msg ID: 5371721954124203
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name: my-gcp-project:db-dev value: 0
Published a message; msg ID: 5371721968467596

CloudSQL Disk Utilization
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name: my-gcp-project:db-dev value: 0 GB
Published a message; msg ID: 5371723698374426
starttime: 2022-08-22T07:27:00Z endttime: 2022-08-22T07:28:00Z name: my-gcp-project:db01 value: 1 GB
Published a message; msg ID: 5371722447240568
```

Where `my-gcp-project` is the Google project name, `db01` and `db-dev` are the Cloud SQL instances running. You will see the same output on a Pub/Sub you have configured.