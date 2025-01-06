package main

import (
	"go-scafolding/infra/aws"
	"go-scafolding/infra/presenter"
)

func main() {
	awsAuth := aws.Credentials{ //nolint:govt
		Id:     "localstack",
		Secret: "localstack",
		Token:  "localstack",
		Region: "us-east-1",
		Url:    "http://localhost:4566",
	}

	sqsClient := aws.CreateSqsClient(&awsAuth)
	for {
		go presenter.Handle(sqsClient)
	}

}
