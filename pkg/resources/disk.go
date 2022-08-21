package resources

import (
	"context"
	"fmt"
	"log"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"cloud.google.com/go/pubsub"
	"github.com/framsouza/gathering-metrics-gcp/pkg/publisher"
	"google.golang.org/api/iterator"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func DiskUtil(projectId, topic string) {
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	req := &monitoringpb.QueryTimeSeriesRequest{
		Name:  fmt.Sprintf("projects/%s", projectId), // optional
		Query: (`fetch cloudsql_database :: cloudsql.googleapis.com/database/disk/bytes_used | within 60m`),
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

		getstart := resp.GetPointData()[2].GetTimeInterval()
		getend := resp.GetPointData()[1].GetTimeInterval()
		starttime := getstart.StartTime.AsTime().Format(time.RFC3339)
		endtime := getend.EndTime.AsTime().Format(time.RFC3339)
		name := resp.GetLabelValues()[2].GetStringValue()
		value := resp.GetPointData()[0].GetValues()[0].GetInt64Value() * 8 / 8000000000

		fmt.Println("starttime:", starttime, "endttime:", endtime, "name:", name, "value:", value, "GB")

		msg := fmt.Sprintf("%s %s %s %d", starttime, endtime, name, value)

		client, err := pubsub.NewClient(ctx, projectId)
		if err != nil {
			fmt.Println(err)
		}

		if err := publisher.Publish(client, topic, msg); err != nil {
			log.Fatal("Failed to publish: %^v", err)
		}

	}
}
