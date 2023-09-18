package utils

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

// GetSecret - uses a secretsmanager client to fetch from AWS
func GetSecret(name, region string) (string, error) {
	ctx := context.Background()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", err
	}
	cfg.Region = region

	client := secretsmanager.NewFromConfig(cfg)
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(name),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := client.GetSecretValue(ctx, input)
	if err != nil {
		return "", err
	}

	return *result.SecretString, nil
}
