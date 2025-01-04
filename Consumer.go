package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

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
		MaxNumberOfMessages: aws.Int64(10),
		VisibilityTimeout:   aws.Int64(25),
		WaitTimeSeconds:     aws.Int64(0),
	})

	if err != nil {
		return nil, err
	}
	if len(messages.Messages) == 0 {
		return make([]*sqs.Message, 0), nil
	}
	return messages.Messages, nil
}

func processMessage(messages []*sqs.Message) {
	if len(messages) == 0 {
		return
	}

	for i := range messages {
		log.Default().Println("Message body : ", *messages[i].Body)
		log.Default().Println("Message MD5Body: ", *messages[i].MD5OfBody)
		log.Default().Println("Message ID: ", *messages[i].MessageId)
		//PROCESS MESSAGE
	}

}
