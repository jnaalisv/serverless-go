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

type Item struct {
	Id string`json:"id"`
	Data string`json:"data"`
	// ExpirationTime string`json:"ExpirationTime"`
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

func Handler(ctx context.Context, item Item) (string, error) {

	av, err := dynamodbattribute.MarshalMap(item)

	condition := "attribute_not_exists(id)"

	inputCommand := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("Movies"),
		ConditionExpression: &condition,
	}

	_, err = withDynamoSession().PutItem(inputCommand)

	if err == nil {
		return fmt.Sprintf("Successfully added new item with id %v", item.Id), nil
	} else {
		return fmt.Sprintf("Got error %v, calling PutItem with id %v", err.Error(), item.Id), nil
	}
}

func main() {
	lambda.Start(Handler)
}