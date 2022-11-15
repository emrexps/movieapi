# MovieApi
Follow the instructions below to run the project

#### Download the project
```sh
git clone https://gitlab.com/emrexps/movieapi.git
```
#### Change Directory
```sh
cd movieapi
```
#### Run the project
```sh
go run .
```

The project will be running on port 8080. You can play with the following endpoints with Postman or any rest client.

| Endpoints |  |
| ------ | ------ |
| GET /movies | get all movies |
| GET /movies/{id} | get movie with id |
| POST /movies | insert movie object |

## Docker Build
In the project directory run the build command to create new docker image
```sh
docker build -t movieapi .
```

Enjoy and have a good day







