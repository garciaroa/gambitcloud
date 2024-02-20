package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/garciaroa/gambitcloud/awsgo"
	"github.com/garciaroa/gambitcloud/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println("> Pido Secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{SecretId: aws.String(nombreSecret)})
	if err != nil {
		fmt.Println("> 20-secretm:" + err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" > Lectura Secret ok " + nombreSecret)

	return datosSecret, nil
}
