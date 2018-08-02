package main

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/ec2manager/src/config"
	"github.com/ec2manager/src/manager"
	"log"
)

func main() {
	client, err := createEc2Client()
	if err != nil {
		log.Fatalf("Can't create ec2 client")
	}

	ec2manager := manager.NewEc2Manager(client)
	err = ec2manager.StartInstance(config.GetConfig().InstanceID)
	if err != nil {
		log.Fatalf("Can't start instance with ID %s. Error %s", config.GetConfig().InstanceID, err.Error())
	} else {
		log.Printf("Instance with ID %s successfully started", config.GetConfig().InstanceID)
	}
}

func createEc2Client() (*ec2.EC2, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}

	cfg.Region = config.GetConfig().Region
	cfg.Credentials = aws.StaticCredentialsProvider{Value: aws.Credentials{
		AccessKeyID: config.GetConfig().AccessKeyID, SecretAccessKey: config.GetConfig().SecretAccessKey}}

	return ec2.New(cfg), nil
}
