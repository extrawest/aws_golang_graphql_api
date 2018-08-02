package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws_golang_graphql_api/src/config"
	"github.com/aws_golang_graphql_api/src/manager"
	"github.com/aws_golang_graphql_api/src/router"
	"github.com/facebookgo/inject"
	"log"
)

func main() {
	server := router.NewServer()
	err := inject.Populate(
		NewEc2Client(),
		manager.NewEc2Manager(),
		server,
	)
	if err != nil {
		log.Fatalf("Can't inject values %s", err.Error())
	}

	server.Start()
}

func NewEc2Client() *ec2.EC2 {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		log.Fatalf("Can't load AWS config %s", err.Error())
	}

	cfg.Region = config.GetConfig().Region
	cfg.Credentials = aws.StaticCredentialsProvider{Value: aws.Credentials{
		AccessKeyID: config.GetConfig().AccessKeyID, SecretAccessKey: config.GetConfig().SecretAccessKey}}

	return ec2.New(cfg)
}
