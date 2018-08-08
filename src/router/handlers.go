package router

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws_golang_graphql_api/src/models"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	"net/http"
)

const (
	ID    = "id"
	Query = "query"

	errorMsg   = "error"
	successMsg = "message"

	contextParam = "output"
)

var (
	instanceType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Instance",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
			"type": &graphql.Field{
				Type: graphql.String,
			},
			"launchtime": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

	rootQuery = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"instance": &graphql.Field{
				Type: instanceType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"operation": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if !ok {
						return nil, errors.New("ID not found")
					}

					operation, ok := p.Args["operation"].(string)
					if !ok {
						return nil, errors.New("operation not found")
					}

					if ec2manager, ok := p.Context.Value(contextParam).(models.Ec2Manager); !ok {
						return nil, errors.New("can't cast to value")
					} else {
						return prepareGraphQLOutput(id, operation, ec2manager)
					}
				},
			},
		},
	})

	schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: rootQuery,
	})
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

func (s *server) graphQLHandler(c *gin.Context) {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: c.Query(Query),
		Context:       context.WithValue(context.Background(), contextParam, s.Ec2Manager),
	})
	c.JSON(http.StatusOK, result)
}
