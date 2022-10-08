package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/devalexandre/openbook-api/adapters/database"
	"github.com/devalexandre/openbook-api/domain/services"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	//var books []repo.Books

	db, err := database.NewConnect()

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	bookService := services.Books{Database: db}

	_, books, err := bookService.List(ctx, 0, 10)

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	booksResponse, err := json.Marshal(books)

	if err != nil {
		return Response{StatusCode: 500}, err
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            string(booksResponse),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
