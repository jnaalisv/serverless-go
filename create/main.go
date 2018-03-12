package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"fmt"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

type Item struct {
	Id string`json:"id"`
	Data string`json:"data"`
}

func Handler() (Response, error) {

	item := Item{
		Id: "my-id",
		Data: "The Big New Movie",
	}

	av, err := dynamodbattribute.MarshalMap(item)


	input := &dynamodb.PutItemInput{
		Item: av,
		TableName: aws.String("Movies"),
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1")},
	)
	svc := dynamodb.New(sess)

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully added 'The Big New Movie' (2015) to Movies table")

	return Response{
		Message: "Successfully added 'The Big New Movie' (2015) to Movies table",
	}, nil
}

func main() {
	lambda.Start(Handler)
}