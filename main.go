package main

import (
	"flag"
	"fmt"
	"os"

	resources "github.com/framsouza/gathering-metrics-gcp/pkg/resources"
)

var (
	projectId = flag.String("project", "", "Enter the Project ID")
)

func main() {
	flag.Parse()

	if *projectId == "" {
		fmt.Fprintln(os.Stderr, "Missing project")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Print("\nCloudSQL CPU Utilization (%)\n")
	resources.CpuUtlizaiton(*projectId)
	fmt.Print("\nCloudSQL memory total usage\n")
	resources.MemUtlizaiton(*projectId)
	fmt.Print("\nCloudSQL memory total size\n")
	resources.MemUTotal(*projectId)
	fmt.Print("\nCloudSQL Active Connections\n")
	resources.MySQLConnections(*projectId)
	resources.PGSQLConnections(*projectId)
	fmt.Print("\nCloudSQL Disk Utilization\n")
	resources.DiskUtil(*projectId)

}
