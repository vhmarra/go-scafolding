package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func CreateSqsClient(auth *AWSCredentials) *sqs.SQS {
	sqsSession := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(auth.Region),
		Credentials: credentials.NewStaticCredentials(auth.Id, auth.Secret, auth.Token),
		Endpoint:    aws.String(auth.Url),
	}))
	sqsService := sqs.New(sqsSession)
	return sqsService
}
