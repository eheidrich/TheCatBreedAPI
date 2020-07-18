# TheCatBreedAPI
A RestAPI, in Golang, using Gin, JWT and MySQL in Docker.

The code works so far, but it is still a work in progress. It still needs to:

* Run fully on Docker, migration and the API still have to be started manually.
* Implement some unit tests.
* Have an API for signing up and handling user Authentication.
* Document the API with Swagger

So far there is an Endpoint to consult Cat Breeds by name. It gets its information from The Cat API (https://docs.thecatapi.com) and caches it in a database by search term. If the same exact term is used once again it gets the information from the database and return to the client.

The Endpoint to search cat breeds is '/breeds' with a query string with the search term '?name='. Example below.
```
domain.com/breeds?name=sibirian
```

The API is protected with a JWT Authentication, there for it is necessary to pass a Bearer Token along with the call. Otherwise it will be rejected. There is no signing up so far.

The authentication is fixed to a login and password, which fits the purpose it was built for, using JWT, but there should be implemented a users model, with a database to save the necessary data and CRUD to manage users.

In order to get the JWT Token, this a API that should be called:
```
domain.com/auth
```
The call should contain a JSON object in the body with the login and password. Example bellow.
```
{
    "password" : "password",
    "username" : "user"
}
```
