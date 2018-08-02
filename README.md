# aws_golang_graphql_api
GraphQL server on Golang to manage AWS EC2 instances.

## Setup

### Dependencies

```bash
# install go dependency management tool 
go get -u github.com/golang/dep/cmd/dep
  
# download and install docker-compose
https://docs.docker.com/compose/install

# install dependencies
dep init

# create ".env" file
create ".env" file with the necessary environment variables (see example .env.sample)
```
### Launch
```bash
docker-compose up
```
### Examples 
*/start?id="ID"* - start instance with ID

*/stop?id="ID"* - stop instance with ID

*/describe?id="ID"* - get info of instance with ID

for testing on localhost: port 9999 (*localhost:9999/start?id=1111111111*)
