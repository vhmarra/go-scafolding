package main

import "github.com/aws/aws-sdk-go/service/sqs"

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
