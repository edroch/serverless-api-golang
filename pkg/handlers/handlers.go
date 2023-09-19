package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/edroch/serverless-api-golang/pkg/user"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

// GetUser intermediates dataflow between datasource (dynamodb) and api response
func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {

	id := req.PathParameters["id"]
	if len(id) > 0 {
		result, err := user.FetchUser(id, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)

}

// CreateUser intermediates dataflow between datasource (dynamodb) and api response
func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := user.CreateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}

// UpdateUser intermediates dataflow between datasource (dynamodb) and api response
func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	result, err := user.UpdateUser(req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

// DeleteUser intermediates dataflow between datasource (dynamodb) and api response
func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamodbiface.DynamoDBAPI) (
	*events.APIGatewayProxyResponse, error,
) {
	err := user.DeleteUser(req, tableName, dynaClient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, nil)
}

// UnhandledMethod for unsupported api requests
func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
