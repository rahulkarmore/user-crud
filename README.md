
# user-crud

The project is for managed the users i.e. Fetch, create, updat, delete users


## Tech Stack

**Server:** Golang.


## API Reference

#### Get all items

```http
  GET /user-crud/get-users
```

#### Add single user

```http
  Post /user-crud/add-user
```

#### Update user by userID

```http
 Put /user-crud/update-user/{userID}
```

#### Delete user by userID

```http
 Put /user-crud/delete-user/{userID}
```
## Installation

```bash
 - Clone the project
 - go mod tidy
 - go run main.go
```
    
## Roadmap

Designing this project with beginners in mind.
- V1 = Creating Server and write CRUD api without using DB
- V2 = Used Database for created CRUD API's.
- V3 = Created Middlerware
- V4 = More improvement....
