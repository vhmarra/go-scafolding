package adapter

import (
	_ "go-scafolding/infra/aws"
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
