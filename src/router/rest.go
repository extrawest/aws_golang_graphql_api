package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	ID = "id"

	errorMsg   = "error"
	successMsg = "message"
)

func (s *server) startInstanceHandler(c *gin.Context) {
	id := c.Query(ID)
	if err := s.Ec2Manager.StartInstance(id); err != nil {
		renderResponse(c, http.StatusForbidden, errorMsg, err.Error())
	} else {
		renderResponse(c, http.StatusOK, successMsg, fmt.Sprintf("Successfully started instance with ID %s", id))
	}
}

func (s *server) stopInstanceHandler(c *gin.Context) {
	id := c.Query(ID)
	if err := s.Ec2Manager.StopInstance(id); err != nil {
		renderResponse(c, http.StatusForbidden, errorMsg, err.Error())
	} else {
		renderResponse(c, http.StatusOK, successMsg, fmt.Sprintf("Successfully stoped instance with ID %s", id))
	}
}

func (s *server) describeInstancesHandler(c *gin.Context) {
	id := c.Query(ID)
	if result, err := s.Ec2Manager.DescribeInstances(id); err != nil {
		renderResponse(c, http.StatusForbidden, errorMsg, err.Error())
	} else {
		c.JSON(http.StatusOK, &result)
	}
}

func renderResponse(c *gin.Context, status int, title, message string) {
	c.JSON(status, gin.H{
		title: message,
	})
}
