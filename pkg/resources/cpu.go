package resources

import (
	"context"
	"fmt"
	"log"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"google.golang.org/api/iterator"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func CpuUtlizaiton(projectId string) {
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	req := &monitoringpb.QueryTimeSeriesRequest{
		Name:  fmt.Sprintf("projects/%s", projectId), // optional
		Query: (`fetch cloudsql_database :: cloudsql.googleapis.com/database/cpu/utilization | within 5m`),
	}

	it := c.QueryTimeSeries(ctx, req)

	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		value := resp.GetPointData()[0].GetValues()[0].GetDoubleValue() * 100

		getstart := resp.GetPointData()[2].GetTimeInterval()
		getend := resp.GetPointData()[1].GetTimeInterval()
		starttime := getstart.StartTime.AsTime().Format(time.RFC3339)
		endtime := getend.EndTime.AsTime().Format(time.RFC3339)

		fmt.Println("starttime:", starttime, "endttime:", endtime, "name:", resp.GetLabelValues()[2].GetStringValue(), "value:", value)

		//publisher.Publish("elastic-support", "test-cloudmetrics", msg)
	}

}
