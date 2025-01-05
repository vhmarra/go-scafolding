package infra

import (
	"log"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func Handle(sqsC *sqs.SQS) {
	for {
		messages, err := poolMessage(sqsC)
		if nil == messages && nil == err {
			panic("Shit...")
		}
		if len(messages) == 0 {
			log.Print("No messages")
		}
		processMessage(messages)
	}
}
