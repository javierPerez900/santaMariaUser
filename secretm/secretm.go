package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/javier/santaMariaUser/awsgo"
	"github.com/javier/santaMariaUser/models"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido secreto " + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	err = json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	if err != nil {
		fmt.Println("Error al interpretar el contenido del secreto " + nombreSecret)
		return datosSecret, err
	}

	fmt.Println(" > Lectura Secret OK " + nombreSecret)
	return datosSecret, nil
}
