package main

import (
	"flag"
	"fmt"
	"os"

	resources "github.com/framsouza/gathering-metrics-gcp/pkg/resources"
)

var (
	projectId = flag.String("project", "", "Enter the Project ID")
	topic     = flag.String("topic", "", "Enter Pub/Sub topic name")
)

func main() {
	flag.Parse()

	if *projectId == "" {
		fmt.Fprintln(os.Stderr, "Missing project")
		flag.Usage()
		os.Exit(2)
	}
	if *topic == "" {
		fmt.Fprintln(os.Stderr, "Missing topic name")
		flag.Usage()
		os.Exit(2)
	}
	fmt.Println("Connecting with topic %s\n", topic)
	fmt.Print("\nCloudSQL CPU Utilization (%)\n")
	resources.CpuUtlizaiton(*projectId, *topic)
	fmt.Print("\nCloudSQL memory total usage\n")
	resources.MemUtlizaiton(*projectId, *topic)
	fmt.Print("\nCloudSQL memory total size\n")
	resources.MemUTotal(*projectId, *topic)
	fmt.Print("\nCloudSQL Active Connections\n")
	resources.MySQLConnections(*projectId, *topic)
	resources.PGSQLConnections(*projectId, *topic)
	fmt.Print("\nCloudSQL Disk Utilization\n")
	resources.DiskUtil(*projectId, *topic)

}
