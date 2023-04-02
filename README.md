
# Task Manager API

A REST API that allows users to manage their tasks, including creating, updating, deleting, and querying tasks.



## API Reference

#### List Routes

```http
POST /auth/login - authenticate user and return JWT token
POST /auth/register - create a new user account
POST /tasks - create a new task
GET /tasks - get a list of tasks
GET /tasks/:id - get a single task by ID
PUT /tasks/:id - update a task by ID
DELETE /tasks/:id - delete a task by ID
GET /users/:id/tasks - get a list of tasks assigned to a user
POST /tasks/:id/comments - add a comment to a task
GET /tasks/:id/comments - get a list of comments for a task
POST /tasks/:id/attachments - upload an attachment for a task
GET /tasks/:id/attachments/:filename - download an attachment for a task
POST /users/:id/avatar - upload a user profile picture
GET /users/:id - get user profile information
PUT /users/:id - update user profile information
PUT /users/:id/password - update user password
GET /admin/users - get a list of all users (admin only)
GET /admin/users/:id/tasks - get a list of tasks assigned to a user (admin only)
PUT /admin/users/:id - update user profile information (admin only)
DELETE /admin/users/:id - delete a user account (admin only)

```


| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API path |

#### Post a todo

```http
  POST /api/v1/tasks
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `tile , email`      | `string` | **Required**. Your API path |


#### Update a tasks

```http
  PUT /api/v1/todo/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `todo_id` | `int` | **Required**. Id of todo to update |

#### Delete todo

```http
  DELETE /api/v1/todo/:id
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `todo_id` | `int` | **Required**. Id of todo to DELETE |





## Run

Step to run this api to your local machine

```bash
  mkdir todo_list
  cd todo_list
  git clone https://github.com/titoyudha/golang_todo_list.git
  change .env files with your configuration
  run go mod tidy
  go run main.go
```
    
## License

[MIT](https://choosealicense.com/licenses/mit/)


## Run By Docker

Pull the project to your local images

```bash
  docker pull titoyb/go_todo:v1      
```



Start the Container

```bash
 docker run -d -p 80:80 --name web titoyb/go_todo:v1

```
