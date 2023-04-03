
# Task Manager API

A REST API that allows users to manage their tasks, including creating, updating, deleting, and querying tasks.



## Tech Stack

#### JWT for user authentication and authorization
#### Gin-Gonic for routing and middleware
#### GORM for object-relational mapping (ORM)
#### Viper for configuration management
#### GoMail for sending email notifications
#### AWS S3 SDK for uploading and downloading attachments


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


| Parameter | Type     | Params               |
| :-------- | :------- | :------------------------- |
| `api_key` | `string` | **Required**. Your API path |
| `file_path` | `string` | **Required**. Your File path |

#### User

```http
/api/v1/users
```

| Path | Type     | Method                | Required
| :-------- | :------- | :------------------------- | :------------ |
| `todo_id` | `int` | **Required**. Id of todo to update | user id |


#### Task

```http
/api/v1/task/
```

| Path | Type     | Method                | Required
| :-------- | :------- | :------------------------- | :------------ |
| `todo_id` | `int` | **Required**. Id of todo to update | task id |




