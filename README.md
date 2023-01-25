# Graphqt Sample

API graphql test with mongodb, and JWT auth.

## TODO

configure a kb8 machine with mongo.

## Auth

To authenticate you need firt use the <base_uri>/api/register (POST), with the body below:

```json
{
    "username": "<your user name>",
    "password": "<your password>"
}

```

The response (OK) must be:

```json
{
    "message": "registration success"
}
```

After that you must have a token to access anothers paths, to get a token valid for 1 hour use: <base_uri>/api/login (GET), with the body below:

```json
{
    "username": "<your user name>",
    "password": "<your password>"
}

```

The response (OK) must be:

```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzQ2NjA5MjYsInVzZXJfaWQiOiI2M2QxMzE1ZTAwMmRlOTQ0ODA4YjJjZWMifQ.x9YqKrjR6fslVXaHbq7V8CLV37p5Sp07unVVmDxUvEk"
}
```

This token can be used like a querystring ?token=<token> or in Header Authorization: bearer <token>

