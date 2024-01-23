package secretm

import (
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"encoding/json"
	"fmt"
	"github.com/ElisaCalderon/gambituser1/awsgo"
	"github.com/ElisaCalderon/gambituser1/models"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func GetSecret(nombreSecret string) (models.SecretRDSJson, error) {
	var datosSecret models.SecretRDSJson
	fmt.Println(" > Pido secreto" + nombreSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nombreSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println(" >Lectura secret Ok " + nombreSecret)
	return datosSecret, nil
}
