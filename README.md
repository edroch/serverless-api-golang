# serverless-api-golang


### curl commands

* curl --header "Content-Type: application/json" --request POST --data '{ "email": "edrocha@gmail.com", "firstName": "ed", "lastName": "rocha" }' https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev

* curl -X GET https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev

* curl -X GET https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev\?email\=edrocha@gmail.com

* curl --header "Content-Type: application/json" --request PUT --data '{ "email": "edrocha@gmail.com", "firstName": "lalala", "lastName": "lalala" }' https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev

* curl -X DELETE https://{apig-resource-id}.execute-api.us-east-1.amazonaws.com/dev\?email\=edrocha@gmail.com