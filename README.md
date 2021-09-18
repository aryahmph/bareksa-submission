# Bareksa Internship Submission

Registration to take part in the internship selection at the Bareksa company

Here is the link from the wireframe : https://whimsical.com/bareksa-internship-submission-DdpNScFC3JthJVD6jcs3Vj

See API Specification in bellow.

## Getting Started

1. Start with cloning this repo on your local machine :

```
$ git clone https://github.com/aryahmph/bareksa_aryayunanta
$ cd bareksa_aryayunanta
```

2. Install docker
3. Run this command in terminal :

```
$ docker build -t bareksa-internship-api .
$ docker-compose -f docker-compose.yml up
```

4. Run migrations, see documentation [here](https://github.com/golang-migrate/migrate) :

```
$ migrate -database "mysql://root:aryahmph@tcp(localhost:3306)/bareksa_aryayunanta?parseTime=true" -path app/migrations up
```

Connect to http://localhost:8080

## API Specification

### Authentication

All API except static images must use this authentication

Request :

- Header :
  - X-API-KEY : "BAREKSA_INTERNSHIP"

### Create Topic

Request :

- Method : POST
- Endpoint : `/api/topics`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string, unique"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "number",
    "name": "string"
  }
}
```

### Get Topic

Request :

- Method : GET
- Endpoint : `/api/topics/:topicId`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "number",
      "title": "string",
      "description": "string",
      "date": "string",
      "image_url": "string"
    }
  ]
}
```

### List Topic

Request :

- Method : GET
- Endpoint : `/api/topics`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "number",
      "name": "string"
    }
  ]
}
```

### Create Tag

Request :

- Method : POST
- Endpoint : `/api/tags`
- Header :
  - Content-Type: application/json
  - Accept: application/json
- Body :

```json
{
  "name": "string, unique"
}
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "number",
    "name": "string"
  }
}
```

### List Tag

Request :

- Method : GET
- Endpoint : `/api/tags`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "number",
      "name": "string"
    }
  ]
}
```

### Create News

Request :

- Method : POST
- Endpoint : `/api/news`
- Header :
  - Content-Type: multipart
  - Accept: application/json
- Body :

```
title         : "string"
topic_name    : "string"
short_desc    : "string"
content       : "string"
writer        : "string"
tags          : "string1,string2,string3 (comma seperated)"
image         : "image binary"
```

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": {
    "id": "number",
    "title": "string",
    "description": "string",
    "content": "string",
    "topic": "string",
    "writer": "string",
    "tags": "string",
    "image_url": "string"
  }
}
```

### Get News

Request :

- Method : GET
- Endpoint : `/api/news/:newsId`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "number",
      "title": "string",
      "date": "string",
      "writer": "string",
      "image_url": "string",
      "content": "string"
    }
  ]
}
```

### List News

Request :

- Method : GET
- Endpoint : `/api/news`
- Header :
  - Accept: application/json

Response :

```json
{
  "code": "number",
  "status": "string",
  "data": [
    {
      "id": "number",
      "title": "string",
      "description": "string",
      "date": "string",
      "image_url": "string"
    }
  ]
}
```