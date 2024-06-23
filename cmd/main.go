package main

import (
	"github.com/ProovGroup/worker-claim-declaration/pkg"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(pkg.Handler)
}
