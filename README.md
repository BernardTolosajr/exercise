# The Chatter API
It’s a side chatter service. As a user of this service, you would be able to comment on any GitHub organization on this particular platform.
###### Hackmd link (https://hackmd.io/TPctUBDiRtKr3i5yJK3eOQ?both)
## Project Structure
* db - Abstraction for the database
* handler - Handler for the http
* models - DB models
* repositories - Persistence layer.
* router -  dispatcher to their respective handler
* services - Bussiness logic
* main.go - Entry point
* Dockerfile
* k8s - Kubernetes deployment manifest
## Requirements
* Go 1.11 or higher
* MongoDB 2.6 and higher
* Kubernetes
    * https://github.com/kubernetes/minikube
    * https://kind.sigs.k8s.io/docs/user/quick-start/
## Kubernetes Deployment
Inside the root dir, apply the k8s manifest
```
kubectl apply -f k8s/
```
## Running locally
Inside the root dir, make sure go modules enabled  https://github.com/golang/go/wiki/Modules
```
go run main go
```
## How to use (locally via httpie)
You can you httpie(https://httpie.org) or curl command
### Create an Organization
##### POST http://localhost:8080/orgs
##### Parameter
| Name | Type | Description |
| -------- | -------- | -------- |
| login     | string     | Required. The organization's username.     |
| profile_name     | string     | The organization's display name.     |
| admin     | string     | Required. The login of the user who will manage this organization.     |
##### httpie command
```
http post localhost:8080/orgs login=userlogin profile_name=profile admin=admin
```
##### Response 200 OK
```
{
    "id": "2370werks32423"
}
```
### Create a Comment
##### POST http://localhost:8080/orgs/<org-name>/comments
##### Parameter
| Name | Type | Description |
| -------- | -------- | -------- |
| comment     | string     | Required. Comment.     |
|  member_id    | string     | Required. The commenter user id|
##### httpie command
```
http post localhost:8080/orgs/xendit/comments comment=awesome member_id=1
```
##### Response 200 OK
```
{
    "id": "2370werks32423"
}
```
##### Response 422 Unprocessable Entity
```
{
    "message": "reason"
}
```
### Get a Comments
##### GET http://localhost:8080/orgs/<org-name>/comments
##### httpie command
```
http get localhost:8080/orgs/xendit/comments
```
##### Response 200 OK
```
{
    "comments": [{
        member_id: 1,
        comment: "awesome company!"
    }]
}
```
##### Response 422 Unprocessable Entity
```
{
    "message": "reason"
}
```
### Delete an Comments
##### DELETE http://localhost:8080/orgs/<org-name>/comments
##### httpie command
```
http delete localhost:8080/orgs/xendit/comments
```
##### Response 200 OK
```
{}
```
##### Response 422 Unprocessable Entity
```
{
    "message": "reason"s
}
```
### Create an Member
##### POST http://localhost:8080/orgs/<org-name>/members
##### Parameter
| Name | Type | Description |
| -------- | -------- | -------- |
| login     | string     | Required. The member username.     |
| avatar_url    | string     | Required. The member avatar url |
##### httpie command
```
http post localhost:8080/orgs/xendit/members login=bt avatar_url=http://foo
```
##### Response 200 OK
```
{
    "id": "2370werks32423"
}
```
##### Response 422 Unprocessable Entity
```
{
    "message": "reason"
}
```
### Get an organizatin members
##### GET http://localhost:8080/orgs/<org-name>/members
##### httpie command
```
http localhost:8080/orgs/xendit/members
```
##### Response 200 OK
```
{
    "members": [{
        login: "bt",
        avatar_url: "http://some.com/bt",
        followers: 10,
        following: 15
    }]
}
```
##### Response 422 Unprocessable Entity
```
{
    "message": "reason"
}
```
