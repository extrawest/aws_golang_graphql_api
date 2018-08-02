package router

import (
	"github.com/aws_golang_graphql_api/src/models"
	"github.com/gin-gonic/gin"
	"log"
)

type server struct {
	Ec2Manager models.Ec2Manager `inject:""`
}

const (
	startURL    = "/start"
	stopURL     = "/stop"
	describeURL = "/describe"
)

func NewServer() *server {
	return new(server)
}

func (s *server) Start() {
	r := gin.Default()
	r.GET(startURL, s.startInstanceHandler)
	r.GET(stopURL, s.stopInstanceHandler)
	r.GET(describeURL, s.describeInstancesHandler)
	log.Fatal(r.Run())
}
