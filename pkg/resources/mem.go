package resources

import (
	"context"
	"fmt"
	"log"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	"google.golang.org/api/iterator"
	monitoringpb "google.golang.org/genproto/googleapis/monitoring/v3"
)

func MemUtlizaiton(projectId string) {
	//projectID = projectId
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	req := &monitoringpb.QueryTimeSeriesRequest{
		Name:  fmt.Sprintf("projects/%s", projectId), // optional
		Query: (`fetch cloudsql_database :: cloudsql.googleapis.com/database/memory/total_usage  | within 60m`),
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

		value := resp.GetPointData()[0].GetValues()[0].GetInt64Value() * 8 / 8000000000

		fmt.Println("name", resp.GetLabelValues()[2].GetStringValue(), "value:", value, "GB")

	}
}

func MemUTotal(projectId string) {
	//projectId := "elastic-apps-163815"
	ctx := context.Background()
	c, err := monitoring.NewQueryClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	req := &monitoringpb.QueryTimeSeriesRequest{
		Name:  fmt.Sprintf("projects/%s", projectId), // optional
		Query: (`fetch cloudsql_database :: cloudsql.googleapis.com/database/memory/quota  | within 60m`),
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

		value := resp.GetPointData()[0].GetValues()[0].GetInt64Value() * 8 / 8000000000

		fmt.Println("name", resp.GetLabelValues()[2].GetStringValue(), "value:", value, "GB")

	}
}
