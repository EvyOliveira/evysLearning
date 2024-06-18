<h1 align="center">
  Evy´s Learning: Course Platform
</h1>
</br>

## 💻 Project Description

This project consists of a platform called Evy's Learning with the aim of offering distance learning courses with activities for students who want to learn english and programming in Brazil. An Entity-Relationship diagram was drawn to illustrate this solution:
</br>

![homescreen](diagram.png)

---

## ⚙️ Features

- [x] Create, edit, delete and update a course
- [x] Create, edit, delete and update a class
- [x] Create, edit, delete and update an activity
- [x] Create, edit, delete and update an activity
- [x] Unit and integration tests with native Go packages
---

## 🚀 Project Execution

### Pre-requisites
Before you begin, you will need to have the following tools installed on your machine:
[Git](https://git-scm.com), [Go](https://go.dev/).
In addition, it is good to have an editor to work with the code, such as [VSCode](https://code.visualstudio.com/). We will adopt containerization with [Docker](https://www.docker.com/) and relational persistence with the [PostgreSQL](https://www.postgresql.org/) database.

#### 🧭 Running the application

```bash

# Clone this repository

# Access the project folder via the following terminal
$ cd evys-learning

# Initialize module in Go

# To synchronize code dependencies
$ go mod tidy

# Start and run containers in background
$ docker-compose up -d

# Allows command execution in the bash shell inside the container
$ docker-compose exec postgres bash

# Connect PostgreSQL database server running on local machine using psql command line too
$ psql -h localhost -p 5432 -U postgres

# Create a new Docker network
$ docker network inspect <network name>

# List Docker networks known to Docker running on your system
$ docker network ls

# Connect an existing Docker container to a specific network
$ docker network connect <network name> <container name>

# Run an interactive shell inside an already running Docker container
$ docker exec -it <container name> /bin/sh
```

---

## 🛠 Technologies

The following tools were used during the development of the project:
([Go](https://go.dev/) + HTML5 + [Docker](https://www.docker.com/) + [PostgreSQL](https://www.postgresql.org/))
<br>

---

## 🦸 Author

[![Linkedin Badge](https://img.shields.io/badge/-evelyncristinioliveira-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/evelyncristinioliveira/)](https://www.linkedin.com/in/evelyncristinioliveira/)