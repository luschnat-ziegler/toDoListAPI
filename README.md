# ToDo-List API

## Setup Instructions

### MongoDB 

The lists are persisted using mongoDB Atlas (Local instances of mongoDB will work, as well).
In order to connect to your own mongoDB Atlas cluster, add a .env file in the root directory of your repository and insert your connection url as `DB_URL`. Example:

`DB_URL=mongodb+srv://abc:<password>@cluster0.z1fxp.mongodb.net/<dbname>?retryWrites=true&w=majority`

where "dbname" should be set to "todo". The database and the collection "lists" will be created on first insert.

### Run or build

Install the Go binaries, following the instructions on [golang.org](https://golang.org/dl/).

Clone this repository using your preferred auth method

Run the application: `go run main.go`  
Build the application: `go build`

The server is set up to run on `http://localhost:8000`. 

### API

There are five endpoints:

#### Get all lists:
GET `http://localhost:8000/todos`: Returns an array of all todo-lists.  

#### Save a new list:
POST `http://localhost:8000/todos`: Saves a list. The request body (JSON) can look like this:

```json
{
  "id": "someID",
  "name": "My ToDo List", 
  "description": "This is my first ToDo_list",
  "tasks": [
    {
      "id": "someID", 
      "name": "My first task", 
      "description": "Task Description"
    },
    {
      "id": "someOtherID",
      "name": "My second task",
      "description": "Task Description"
    }
  ] 
}
``` 
Both `id` and `task->id` can be submitted or omitted. In the former case they will be ignored and reset. `name` and `task->name` are required fields, as opposed to `description` and `task->description` which can be included or omitted, in which case they will be set to `null`. Thus, the following is also a valid request body:

```json
{
  "name": "My ToDo List", 
  "tasks": [
    {
      "name": "My first task", 
      "description": "Task Description"
    },
    {
      "name": "My second task"
    }
  ] 
}
```
The response body will look like this:

```json
{
    "id": "601d68d2b69d07127cb97eff",
    "name": "My ToDo List",
    "description": "The description can also be omitted",
    "tasks": [
        {
            "id": "5f0546be-9325-4076-9f32-c9b70d99037c",
            "name": "My first task",
            "description": "Task Description"
        },
        {
            "id": "2c2d0eee-bfcb-485f-917d-ad2d135be203",
            "name": "My second task",
            "description": null
        }
    ]
}
``` 
Requests are validated. If validation fails, an error message is returned:

```json
{
    "invalid_fields": {
        "name": "required",
        "tasks[0].name": "required"
    }
}
```

#### Get one list by ID:
GET `http://localhost:8000/todos/{id}`: Returns one list.  

#### Update one list by ID:
PUT `http://localhost:8000/todos/{id}`: Overwrites an existing list and - on success - returns the new list. Request and response are similar to saving a new list. Note: Task-IDs, if submitted, will be assigned anew.

#### Delete one list by ID:
DELETE `http://localhost:8000/todos/{id}`: Deletes the list, if ID exists. Returns status code `204` on success and no response body.  