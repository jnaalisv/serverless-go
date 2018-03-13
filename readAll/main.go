package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"os"
	"context"
)

type MyDataItem struct {
	Id string`json:"id"`
	Data string`json:"data"`
}

func withDynamoSession() (*dynamodb.DynamoDB){
	mySession, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	if err != nil {
		fmt.Println("Got error calling createSession:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return dynamodb.New(mySession)
}

func Handler(ctx context.Context) ([]MyDataItem, error) {

	output, err := withDynamoSession().Scan(&dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	})

	if err != nil {
		panic(fmt.Sprintf("Got error calling Scan, %v", err))
	}

	items := []MyDataItem{}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
	}

	return items, nil
}

func main() {
	lambda.Start(Handler)
}