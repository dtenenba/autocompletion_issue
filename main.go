package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func main() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	cfg.Region = endpoints.UsWest2RegionID
	svc := s3.New(cfg)
	input := &s3.ListObjectsV2Input{
		Bucket:  aws.String("mybucket"),
		MaxKeys: aws.Int64(1000),
	}
	// var keys []string

	for {
		req := svc.ListObjectsV2Request(input)
		result, err := req.Send()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		// on the next line, type `result.` and see if there are autocompletions

	}

}
