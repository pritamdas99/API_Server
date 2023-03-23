# api-server #
Simple api server for CRUD operation <br>
## Technology Used ##
- Golang
- JWT Authentication
- Cobra CLI
## Running the server ##
### Running the server from direct source code ##
```git clone git@github.com:pritamdas99/API_server.git``` <br>

Go to the api-server directory and run <br>
```go mod tidy``` <br>
```go mod vendor``` <br>
```go run . start``` or ```go run . start -p <choosen port>```

## Data model ##
<pre><code>
// For string a cricket players basic information
type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	GenreIds []int  `json:"genreids"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
</code></pre>

## api calls ##

| Method |                     url                      | payload | actions                                    |
|--------|:--------------------------------------------:|--------:|--------------------------------------------|
| GET    |         ```http://localhost:8080/```         |      No | call home page and get new token           |
| GET    |    ```http://localhost:8080/getallbook```    |      No | call all the books exist in data model     |
| GET    |   ```http://localhost:8080/getallgenre```    |      No | call all the genres exist in data model    |
| POST   |     ```http://localhost:8080/addbook```      |     Yes | add new book in the data model             |
| POST   |     ```http://localhost:8080/adduser```      |     Yes | add new user in the data model             |
| POST   |     ```http://localhost:8080/addgenre```     |     Yes | add new book in the data model             |
| PUT    | ```http://localhost:8080/updatebook/{id}```  |     Yes | update a specific book by calling with id  |
| PUT    | ```http://localhost:8080/updategenre/{id}``` |     Yes | update a specific genre by calling with id |
| POST   | ```http://localhost:8080/updateuser/{id}```  |     Yes | add new user in the data model             |
| POST   |    ```http://localhost:8080/loginuser```     |     Yes | login a user in the data model             |
| DELETE | ```http://localhost:8080/deletebook/{id}```  |      No | delete a specific book by calling its id   |
| DELETE | ```http://localhost:8080/deletegenre/{id}``` |      No | delete a specific book by calling its id   |


## CAUTIONS ##
- All the api call is authorized by JWT token. So we need to pass token as a cookie while calling methods(except for homepage request) by postman or curl.
- A token can be generated by call homepage initially.