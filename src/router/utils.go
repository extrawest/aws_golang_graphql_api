package router

import (
	"errors"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws_golang_graphql_api/src/models"
	"github.com/gin-gonic/gin"
)

func renderResponse(c *gin.Context, status int, title, message string) {
	c.JSON(status, gin.H{
		title: message,
	})
}

func prepareGraphQLOutput(id, operation string, manager models.Ec2Manager) (*models.GraphQLResult, error) {

	result := new(models.GraphQLResult)
	switch operation {
	case "start":
		if err := manager.StartInstance(id); err != nil {
			return nil, err
		} else {
			result.ID = id
			return result, err
		}
	case "stop":
		if err := manager.StopInstance(id); err != nil {
			return nil, err
		} else {
			result.ID = id
			return result, err
		}
	case "describe":
		if output, err := manager.DescribeInstances(id); err != nil {
			return nil, err
		} else {
			return prepareDescribeInstanceOutput(output), nil
		}
	default:
		err := errors.New("operation not supported")
		return nil, err
	}
}

func prepareDescribeInstanceOutput(output *ec2.DescribeInstancesOutput) *models.GraphQLResult {
	result := new(models.GraphQLResult)
	if output == nil {
		return result
	}

	if len(output.Reservations) <= 0 {
		return result
	}

	if len(output.Reservations[0].Instances) <= 0 {
		return result
	}

	if ID := output.Reservations[0].Instances[0].InstanceId; ID == nil {
		result.ID = "not available"
	} else {
		result.ID = *ID
	}

	result.Type = (string)(output.Reservations[0].Instances[0].InstanceType)

	if launchTime := output.Reservations[0].Instances[0].LaunchTime; launchTime == nil {
		result.LaunchTime = "not available"
	} else {
		result.LaunchTime = launchTime.String()
	}

	if state := output.Reservations[0].Instances[0].State; state == nil {
		result.State = "not available"
	} else {
		result.State = (string)(state.Name)
	}

	return result
}
