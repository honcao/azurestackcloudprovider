// Our first program will print the classic "hello world"
// message. Here's the full source code.
package main

import (
	"fmt"

	"azureprovider/azure"
	"bufio"
    "os"
)

func main() {
	fmt.Println("hello world")
	file,_ := os.Open("cloudConfig.yml")
	var configReader = bufio.NewReader(file)
	cloud,_ := azure.NewCloud(configReader)
	fmt.Println(cloud.ProviderName())
	az, _ := cloud.(*azure.Cloud)
	nextlun,_ := az.GetNextDiskLun("k8s-agentpool1-15485654-0");
	fmt.Println(nextlun)
}
