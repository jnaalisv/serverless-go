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

	var itemsSlice []MyDataItem

	err := withDynamoSession().ScanPages(&dynamodb.ScanInput{
		TableName: aws.String("Movies"),
	}, func(page *dynamodb.ScanOutput, last bool) bool {
		recs := []MyDataItem{}

		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &recs)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal Dynamodb Scan Items, %v", err))
		}

		itemsSlice = append(itemsSlice, recs...)
		return true // keep paging
	})

	if err != nil {
		fmt.Println("Got error calling Scan:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return itemsSlice, nil
}

func main() {
	lambda.Start(Handler)
}