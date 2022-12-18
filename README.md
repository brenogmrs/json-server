# json-server

### this is a simplistic implementation of javascripts json-server using Go! It uses the the file called `db.json` at the root of the project to create a Web Server using the resources of it.

---

### This JSON example would generate 5 routes with CRUD operations for each resource.

### `db.json` example:

```
{
  "books": [
    {
      "id": 1,
      "author": "author",
      "name": "book"
    }
  ],
  "users": [
    {
      "id": 1,
      "age": 24,
      "email": "email@email.com",
      "name": "name",
      "password": "password"
    }
  ]
}
```

## TODO List

- create a generic field validator for POST and PUT requests
- create a generic filtering by multiple query params for GET request
