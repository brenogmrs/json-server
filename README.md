# json-server

### this is a simplistic implementation of javascripts json-server using Go! It uses the the file called `db.json` at the root of the project to create a Web Server using the resources of a JSON file.

## JSON example:

```
{
  "books": [
    {
      "id": 1,
      "author": "author",
      "name": "the hobbit"
    }
  ],
  "users": [
    {
      "id": 1,
      "age": 24,
      "email": "email@email.com",
      "name": "Name",
      "password": "password"
    }
  ]
}
```

### This JSON example would generate 5 routes with CRUD operations for each resource, the resources being: `books` and `users` or any other resource following the example above.
