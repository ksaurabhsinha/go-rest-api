## Simple REST API in GoLang

This is a simple REST API in GoLang. This is just for educational purpose.
I do not use any database and everything is saved in a slice.

### Endpoints

#### List All People
**Request**
```markdown
/people - GET
```
**Response**
```markdown
[
    {
        "id": "1",
        "first_name": "User 1 First Name",
        "last_name": "User 1 Last Name",
        "address": {
            "city": "City 1",
            "state": "State 1"
        }
    },
    {
        "id": "2",
        "first_name": "User 2 First Name",
        "last_name": "User 2 Last Name",
        "address": {
            "city": "City 2",
            "state": "State 2"
        }
    },
    {
        "id": "3",
        "first_name": "User 3 First Name",
        "last_name": "User 3 Last Name"
    }
]
```

#### Get single person by Id
**Request**
```markdown
/people/{id} - GET
```
**Response**
```markdown
{
    "id": "1",
    "first_name": "User 1 First Name",
    "last_name": "User 1 Last Name",
    "address": {
        "city": "City 1",
        "state": "State 1"
    }
}
```

#### Add New Person
**Request**
```markdown
/people/{id} - POST
```
**Payload**
```markdown
{
    "first_name": "User 4 First Name",
    "last_name": "User 4 Last Name",
    "address": {
        "city": "City 4",
        "state": "State 4"
    }
}
```
**Response**
```markdown
[
    {
        "id": "1",
        "first_name": "User 1 First Name",
        "last_name": "User 1 Last Name",
        "address": {
            "city": "City 1",
            "state": "State 1"
        }
    },
    {
        "id": "2",
        "first_name": "User 2 First Name",
        "last_name": "User 2 Last Name",
        "address": {
            "city": "City 2",
            "state": "State 2"
        }
    },
    {
        "id": "3",
        "first_name": "User 3 First Name",
        "last_name": "User 3 Last Name"
    },
    {
        "id": "4",
        "first_name": "User 4 First Name",
        "last_name": "User 4 Last Name",
        "address": {
            "city": "City 4",
            "state": "State 4"
        }
    }
]
```

#### Delete a Person
**Request**
```markdown
/people/{id} - DELETE
```
**Response**
```markdown
[
    {
        "id": "1",
        "first_name": "User 1 First Name",
        "last_name": "User 1 Last Name",
        "address": {
            "city": "City 1",
            "state": "State 1"
        }
    },
    {
        "id": "2",
        "first_name": "User 2 First Name",
        "last_name": "User 2 Last Name",
        "address": {
            "city": "City 2",
            "state": "State 2"
        }
    },
    {
        "id": "3",
        "first_name": "User 3 First Name",
        "last_name": "User 3 Last Name"
    }
]
```
