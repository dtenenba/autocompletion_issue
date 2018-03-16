package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/external"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)
	// On the next line, try typing "cfg." and see if any autocompletion happens

}
