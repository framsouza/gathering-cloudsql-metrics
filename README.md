# Gathering Google Cloud SQL metrics

This code will connect to the Cloud Monitoring API and collect the metrics from Google Cloud SQL, at the moment, the metrics are:

- Disk Utilization
- Memory Utilization
- CPU Utilization
- Active Connections 
- Total memory

For more metrics, check [Google Cloud SQL metrics](https://cloud.google.com/monitoring/api/metrics_gcp#gcp-cloudsql)

### Usage

First, make sure you have the `GOOGLE_APPLICATION_CREDENTIALS` environment variable set as per the [google docs](https://cloud.google.com/docs/authentication/production).

It will require one argument which is `project`.

You can choose to build the binary by running
```
go build -o gathering-cloudsql-metrics
```

and run it as
```
./gathering-cloudsql-metrics -project=<PROJECTNAME>
```

or without building the binary
```
go run main.go -project=<PROJECTNAME>
```

## Output expected
You will have an output like the following:

```
CloudSQL CPU Utilization (%)
name: my-project:db-dev value: 2.701801066450571
name: my-project:db01 value: 2.419767372093702

CloudSQL memory total usage
name my-project:db-dev value: 0 GB
name my-project:db01 value: 0 GB

CloudSQL memory total size
name my-project:db-dev value: 8 GB
name my-project:db01 value: 8 GB

CloudSQL Active Connections
Name: my-project:db01 Value: 4
name: my-project:db-dev value: 2
name: my-project:db-dev value: 0
name: my-project:db-dev value: 0

CloudSQL Disk Utilization
name: my-project:db-dev value: 0 GB
name: my-project:db01 value: 1 GB
```

Where `my-project` is the Google project name, `db01` and `db-dev` are the Cloud SQL instances running.