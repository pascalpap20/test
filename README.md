# mini-project-2

api contract dapat dilihat [disini](https://documenter.getpostman.com/view/11280646/2s93sZ8EaD)

# setup

1. Install Go version 1.20
2. Use GoLand (recommended)
3. Download dependencies with command go mod tidy and go vendor

# Architecture and Design

Architecture and Design
this service using onion architecture, there are 5 layers
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:


Entity: this layer contains the domain model or entities
of the system. These are the core objects that
represent the business concepts and rules.

Repository: This layer provides an interface for the
application to access and manipulate the entities.
It encapsulates the data access logic and provides
a way to abstract the database implementation details.

Use case : This layer contains the business logic
or use cases of the system. It defines the operations
that can be performed on the entities and orchestrates
the interactions between the entities and the repository layer.

Controller: This layer handles the HTTP requests and
responses. It maps the incoming requests to the appropriate
use case and returns the response to the client.

Request handler: This layer is responsible for handling
the incoming HTTP requests and passing them on to
the controller layer.

# ERD

![erd](https://gitlab.com/pascalpanatagama/bootcampbri-miniproject1/-/raw/main/MiniProject1-Page-2.drawio.png)
