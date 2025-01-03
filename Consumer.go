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
