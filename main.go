package main

import "go-scafolding/infra"

func main() {
	awsAuth := infra.AWSCredentials{
		"localstack",
		"localstack",
		"localstack",
		"us-east-1",
		"http://localhost:4566",
		"",
	}

	sqsClient := infra.CreateSqsClient(&awsAuth)
	for {
		go infra.Handle(sqsClient)
	}

}
