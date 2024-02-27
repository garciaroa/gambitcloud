package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/garciaroa/gambitcloud/awsgo"
	"github.com/garciaroa/gambitcloud/bd"
	"github.com/garciaroa/gambitcloud/handlers"
)

func main() {
	lambda.Start(EjecutoLambda)
}
func EjecutoLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

	awsgo.InicializoAWS()

	if !ValidoParametros() {
		panic("Error en los parametros. Debe enviar 'secretName', 'UrlPrefix' ")
	}

	var res *events.APIGatewayProxyResponse
	prefix := os.Getenv("UrlPrefix")
	path := strings.Replace(request.RawPath, prefix, "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()
	status, message := handlers.Manejadores(path, method, body, header, request)

	headersResp := map[string]string{
		"Content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       "aqui" + string(message),
		Headers:    headersResp,
	}

	return res, nil
}

func ValidoParametros() bool {
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}
	/*

	   //ya no se van a utilizar
	   	_, traeParametro = os.LookupEnv("UserPoolId")
	   	if !traeParametro {
	   		return traeParametro
	   	}

	   	_, traeParametro = os.LookupEnv("Region")
	   	if !traeParametro {
	   		return traeParametro
	   	}*/

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}
	return traeParametro
}
