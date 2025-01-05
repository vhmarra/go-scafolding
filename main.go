package main

import (
	"go-scafolding/infra/aws"
	"go-scafolding/infra/presenter"
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
		go presenter.Handle(sqsClient)
	}

}
