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
POST | /users | {"username":\<string\>, "password":\<string\>} | 201 Created, 400 Bad Request (len(username or password) < 3), 409 Conflict | {"id" : \<int\>, "username":\<string\>, "password":""} or {"message":\<string\>} | - |
POST | /session | {"username":\<string\>, "password":\<string\>} | 200 OK, 400 Bad Request, 404 Not Found | {"id" : \<int\>, "username":\<string\>, "password":""} or {"message":\<string\>} | - |
DELETE | /session | | 200 OK, 401 Unauthorized | {"message":\<string\>} | + |
 | | | | | |
GET | /groups | | 200 OK | [{"id":\<int>\,"name":\<string\>, "about": \<string\>}] | - |
GET | /groups/:name_or_id | | 200 OK, 404 Not Found | {"id":\<int>\,"name":\<string\>, "about": \<string\>} | - |
 | | | | | |
GET | /user/groups | | 200 OK, 401 Unauthorized | [{"id":\<int>\,"name":\<string\>, "about": \<string\>}] | + |
GET | /user/groups/:name_or_id | | 200 OK, 404 Not Found, 401 Unauthorized | {"id":\<int>\,"name":\<string\>, "about": \<string\>} | + |
POST | /user/groups/:name_or_id/add | | 200 OK, 404 Not Found, 409 Conflict, 401 Unauthorized | user's current groups after adding a new group: [{"id":\<int>\,"name":\<string\>, "about": \<string\>}] or {"message":\<string\>} | + |
DELETE | /user/groups/:name_or_id | | 200 OK, 401 Unauthorized |  user's current groups after deleting a group: [{"id":\<int>\,"name":\<string\>, "about": \<string\>}] | + |
 | | | | | |
GET | /ingredients | | 200 OK | [{"id":\<int>\,"name":\<string\>, "about": \<string\>, "type": \<string\>}] | - |
GET | /ingredients/:name_or_id | | 200 OK, 404 Not Found | {"id":\<int>\,"name":\<string\>, "about": \<string\>, "type": \<string\>} | - |
GET | /ingredients/:group_name_or_id/groups | | 200 OK, 404 Not Found | [{"id":\<int>\,"name":\<string\>, "about": \<string\>, "type": \<string\>}] | - |
 | | | | | |
GET | /user/ingredients | | 200 OK, 401 Unauthorized | [{"id":\<int>\, "name":\<string\>, "about": \<string\>}] | + |
POST | /user/ingredients/:name_or_id | | 200 OK,  404 Not Found, 409 Conflict, 401 Unauthorized | user's current excluded ingredients after adding a new ingredient: [{"id":\<int>\, "name":\<string\>, "about": \<string\>}] or {"message":\<string\>} | + |
DELETE | /user/ingredients/:name_or_id | | 200 OK, 401 Unauthorized |  user's current excluded ingredients after deleting an ingredient: [{"id":\<int>\, "name":\<string\>, "about": \<string\>}] | + |
 | | | | | |
GET | /products | | 200 OK |  [{"id":\<int>\,"name":\<string\>, "barcode": int}] | - |
GET | /products/:barcode | | 200 OK, 404 Not Found | {"id":\<int>\,"name":\<string\>, "barcode": int} | - |