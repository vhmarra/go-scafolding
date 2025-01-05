package main

import (
	"go-scafolding/infra/adapter"
	_ "go-scafolding/infra/adapter"
	"go-scafolding/infra/aws"
)

func main() {
	awsAuth := aws.AWSCredentials{
		"localstack",
		"localstack",
		"localstack",
		"us-east-1",
		"http://localhost:4566",
		"",
	}

	sqsClient := aws.CreateSqsClient(&awsAuth)
	for {
		go adapter.Handle(sqsClient)
	}

}
