package main

func main() {
	awsAuth := AWSCredentials{
		"localstack",
		"localstack",
		"localstack",
		"us-east-1",
		"http://localhost:4566",
		"",
	}

	sqsClient := createSqsClient(&awsAuth)
	for {
		go handle(sqsClient)
	}

}
