package presenter

import (
	"go-scafolding/infra/adapter"
	"log"

	"github.com/aws/aws-sdk-go/service/sqs"
)

func Handle(sqsC *sqs.SQS) {
	for {
		messages, err := adapter.PoolMessage(sqsC)
		if nil == messages && nil == err {
			panic("Shit...")
		}
		if len(messages) == 0 {
			log.Print("No messages")
		}
		adapter.ProcessMessage(messages)
	}
}
