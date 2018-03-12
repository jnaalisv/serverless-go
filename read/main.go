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

type GetItemEvent struct {
	Id string`json:"id"`
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

func Handler(ctx context.Context, event GetItemEvent) (map[string]*dynamodb.AttributeValue, error) {

	av, err := dynamodbattribute.MarshalMap(event)

	getItemCommand := &dynamodb.GetItemInput{
		TableName: aws.String("Movies"),
		Key: av,
	}

	pGetItemOutput, err := withDynamoSession().GetItem(getItemCommand)

	if err != nil {
		fmt.Println("Got error calling GetItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return pGetItemOutput.Item, nil
}

func main() {
	lambda.Start(Handler)
}
