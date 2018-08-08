# ec2manager
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
for testing on localhost: port 9999 (*localhost:9999/*** *)

#### REST
*/start?id="ID"* - start instance with ID

*/stop?id="ID"* - stop instance with ID

*/describe?id="ID"* - get info of instance with ID

for testing on localhost: port 9999 (*localhost:9999/start?id=1111111111*)

#### GraphQL
for these queries we can pass the fields that interest us in {}:

**'id'** for **start,stop** and **describe** operations

 **'type','launchtime','state'** only for **describe** operation

*/graphql?query={instance(id:"ID",operation:"start"){id}}*

*/graphql?query={instance(id:"ID",operation:"stop"){id}}*

*/graphql?query={instance(id:"ID",operation:"describe"){id,type,launchtime,state}}*
