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

Method   | Path     | Body                                          | Response                               | Response Body | Requires Auth |
-------- | -------- | --------------------------------------------- | -------------------------------------- | ----- | --- |
POST | /users | {"username":\<string\>, "password":\<string\>} | 201 Created, 400 Bad Request, 409 Conflict | {"message":\<string\>} | - |
POST | /session | {"username":\<string\>, "password":\<string\>} | 200 OK, 400 Bad Request, 404 Not Found | {"message":\<string\>} | - |
DELETE | /session | | 200 OK, 401 Unauthorized | {"message":\<string\>} | + |
* | | | | | |
GET | /groups | | 200 OK | [{"id":\<int>\,"group":\<string\>, "about": \<string\>}] | - |
GET | /groups/:name_or_id | | 200 OK, 404 Not Found | {"id":\<int>\,"group":\<string\>, "about": \<string\>} | - |
* | | | | | |
GET | /user/groups | {"username":\<string\>} | 200 OK, 401 Unauthorized | [{"id":\<int>\,"group":\<string\>, "about": \<string\>}] | + |
GET | /user/groups/:name_or_id | {"username":\<string\>} | 200 OK, 404 Not Found, 401 Unauthorized | {"id":\<int>\,"group":\<string\>, "about": \<string\>} | + |
POST | /user/groups/:name_or_id | {"username":\<string\>} | 200 OK, 404 Not Found, 409 Conflict, 401 Unauthorized | user's current groups after adding a new group: {"id":\<int>\,"group":\<string\>, "about": \<string\>} or {"message":\<string\>} | + |
DELETE | /user/groups/:name_or_id | {"username":\<string\>} | 200 OK, 404 Not Found, 401 Unauthorized |  user's current groups after deleting a group: {"id":\<int>\,"group":\<string\>, "about": \<string\>} | + |