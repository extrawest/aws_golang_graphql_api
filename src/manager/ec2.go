package manager

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

type ec2Manager struct {
	client *ec2.EC2
}

func NewEc2Manager(client *ec2.EC2) *ec2Manager {
	ec2 := new(ec2Manager)
	ec2.client = client
	return ec2
}

func (m *ec2Manager) DescribeInstances(ids ...string) error {
	input := new(ec2.DescribeInstancesInput)
	input.InstanceIds = ids

	req := m.client.DescribeInstancesRequest(input)
	resp, err := req.Send()
	if err != nil {
		return err
	}
	log.Printf("Describe Instances Response %+v", resp)
	return nil
}

func (m *ec2Manager) StartInstance(ids ...string) error {
	input := new(ec2.StartInstancesInput)
	input.InstanceIds = ids

	req := m.client.StartInstancesRequest(input)
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

	req := m.client.StopInstancesRequest(input)
	resp, err := req.Send()
	if err != nil {
		return nil
	}

	log.Printf("Stop Instance Response %+v", resp)
	return nil
}
