package awsgo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config

func InicializoAWS() error {
	Ctx = context.TODO()

	cfg, err := config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		return fmt.Errorf("error al cargar la configuración de AWS: %w", err)
	}

	Cfg = cfg
	return nil
}