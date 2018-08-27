package main

import (
	"bitbucket.org/credomobile/ArchitecturePOCs/jenkinsgolambda/handler"
	"github.com/aws/aws-lambda-go/lambda"
)

var version string

func main() {
	lambda.Start(handler.HandleApiGatewayRequest)
}
