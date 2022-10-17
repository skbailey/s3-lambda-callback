package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	svc := s3.New(session.New(), &aws.Config{Region: aws.String("us-east-1")})
	for _, record := range s3Event.Records {
		s3record := record.S3
		bucket := s3record.Bucket.Name
		key := s3record.Object.Key
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, bucket, key)

		input := &s3.GetObjectInput{
			Bucket: &bucket,
			Key:    &key,
		}

		result, err := svc.GetObject(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case s3.ErrCodeNoSuchKey:
					fmt.Println(s3.ErrCodeNoSuchKey, aerr.Error())
				case s3.ErrCodeInvalidObjectState:
					fmt.Println(s3.ErrCodeInvalidObjectState, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			return
		}

		fmt.Printf("Here's the s3 object: %#v\n", result)
	}
}

func main() {
	// Make the handler available for Remote Procedure Call by AWS Lambda
	lambda.Start(handler)
}
