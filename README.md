# darxio food backend

## Build Instructions

```
docker-compose up -d
```

or (to rebuild)

```
docker-compose up --build -d
```

## API
`500 Internal Server Error` is always possible.

Method   | Path     | Body                                          | Response                               | Response Body |
-------- | -------- | --------------------------------------------- | -------------------------------------- | ----- |
POST | /users | {"username":\<string\>, "password":\<string\>} | 201 Created, 400 Bad Request, 409 Conflict | {"message":\<string\>} |
POST | /session | {"username":\<string\>, "password":\<string\>} | 200 OK, 400 Bad Request, 404 Not Found | {"message":\<string\>} |
DELETE | /session | | 200 OK, 401 Unauthorized | {"message":\<string\>} |