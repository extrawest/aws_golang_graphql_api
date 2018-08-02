package manager

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws_golang_graphql_api/src/models"
	"log"
)

type ec2Manager struct {
	Client *ec2.EC2 `inject:""`
}

func NewEc2Manager() models.Ec2Manager {
	return new(ec2Manager)
}

func (m *ec2Manager) DescribeInstances(ids ...string) (*ec2.DescribeInstancesOutput, error) {
	input := new(ec2.DescribeInstancesInput)
	input.InstanceIds = ids

	req := m.Client.DescribeInstancesRequest(input)
	resp, err := req.Send()
	if err != nil {
		return nil, err
	}
	log.Printf("Describe Instances Response %+v", resp)
	return resp, nil
}

func (m *ec2Manager) StartInstance(ids ...string) error {
	input := new(ec2.StartInstancesInput)
	input.InstanceIds = ids

	req := m.Client.StartInstancesRequest(input)
	resp, err := req.Send()
	if err != nil {
		return err
	}

	log.Printf("Start Instance Response %+v", resp)
	return nil
}

func (m *ec2Manager) StopInstance(ids ...string) error {
	input := new(ec2.StopInstancesInput)
	input.InstanceIds = ids

	req := m.Client.StopInstancesRequest(input)
	resp, err := req.Send()
	if err != nil {
		return nil
	}

	log.Printf("Stop Instance Response %+v", resp)
	return nil
}
