package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type AWSCredentials struct {
	Id     string
	Secret string
	Token  string //TODO acho que nao precisa disso
	Region string
	Url    string
}

func poolMessage(client *sqs.SQS) ([]*sqs.Message, error) {
	queue := "http://sqs.us-east-1.localhost.localstack.cloud:4566/000000000000/go-scafolding" //ENV

	messages, err := client.ReceiveMessage(&sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            &queue,
		MaxNumberOfMessages: aws.Int64(1),
		VisibilityTimeout:   aws.Int64(1),
		WaitTimeSeconds:     aws.Int64(1),
	})

	if err != nil {
		return nil, err
	}
	if len(messages.Messages) == 0 {
		return make([]*sqs.Message, 0), nil
	}
	return messages.Messages, nil
}

func createSqsClient(auth *AWSCredentials) *sqs.SQS {
	sqsSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(auth.Region),
		Credentials: credentials.NewStaticCredentials(auth.Id, auth.Secret, auth.Token),
		Endpoint:    aws.String(auth.Url),
	}))
	sqsService := sqs.New(sqsSession)
	return sqsService
}

func processMessage(messages []*sqs.Message) {
	for i := range messages {
		if nil == messages[i].Body {
			continue
		}
		log.Default().Println("Message body : ", *messages[i].Body)
		log.Default().Println("Message MD5Body: ", *messages[i].MD5OfBody)
		log.Default().Println("Message ID: ", *messages[i].MessageId)
		//PROCESS MESSAGE
	}

}

func handle(sqsC *sqs.SQS) {
	for {
		messages, err := poolMessage(sqsC)
		if nil == messages && nil == err {
			panic("Shit...")
		}
		if len(messages) == 0 {
			continue
		}
		processMessage(messages)
	}
}

func main() {
	awsAuth := AWSCredentials{
		"localstack",
		"localstack",
		"localstack",
		"us-east-1",
		"http://localhost:4566",
	}

	sqsClient := createSqsClient(&awsAuth)
	for {
		go handle(sqsClient)
	}

}
