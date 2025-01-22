# Viswals Backend Test
This is a monorepo project that implements a producer-consumer architecture. The system uses RabbitMQ for message queuing, PostgreSQL for database storage, and Redis for caching. Below is a detailed overview of the project, including folder structure, features, and API documentation.

## Project Overview

### Producer
- Reads data from a CSV file.
- Sanitizes and cleans the data.
- Publishes the processed data to a RabbitMQ queue.

### Consumer
- Reads data from the RabbitMQ queue.
- Hosts an HTTP server to expose the data via APIs.
- Uses Redis for caching frequently accessed user data.
- NOTE: For now invalid data will be kept it the rabbitmq queue only. We can handle this data based on our business logic.
- NOTE: I have used squirrel package for building queries and applying advance filtering capabilities to PostgreSQL.

## Folder Structure
- core/
    - Contains shared packages and utilities that are common across all services.
    - Wrappers around PostgreSQL, Redis, and other infrastructure components.
    - Reusable utilities for logging, configuration management, and more.

- usecase/
    - Contains the service-specific business logic for each module.
    - Interacts with database layer, cache layer or any dependent services to perform business logic.
    - Current setup can be extended by extending domain specific business logic in separate folders.

- repository/
    - Containes the database layer and cache layer related to each module.

- controller/
    - Manages the API controllers for HTTP or gRPC endpoints.
    - All logic related to the API controllers is defined in the controller folder.

- config/
    - Handles application configuration and environment management.

- build/
    - Contains Dockerfiles and build-related scripts.

## API Documentation

1. Get List of Users
- Endpoint: GET /users?page=1&page_size=2
- Query Parameters:
    - page (optional, default: 0): The page number for pagination.
    - page_size (optional, default: 25): The number of users to fetch per page.
    - sort (optional, default: none): Sort response by particular field ( for eg. `sort=id:ASC` will sort users data by id )
    - id:min ( optional, default: none): Fetch users data whose id is greater than or equal to id:min.
    - id:max ( optional, default: none): Fetch users data whose id is less than or equal to id:max.
- Description: Fetches a paginated list of users from the database.
- Response:
```
{
    "data": [
        {
            "id": 3,
            "email": "FelipeKim@gmail.com",
            "firstname": "Felipe",
            "lastname": "Kim",
            "created_at": "2013-09-11T06:41:08Z"
        },
        {
            "id": 4,
            "email": "SantiagoBrown@gmail.net",
            "firstname": "Santiago",
            "lastname": "Brown",
            "created_at": "2013-09-16T15:09:09Z"
        }
    ],
    "pagination": {
        "page": 1,
        "page_size": 2,
        "total_records": 507026,
        "total_page": 253513
    }
}
```

2. Get User by ID
- Endpoint: GET /users/:id
- Path Parameters:
    - id (required): The unique ID of the user.
- Description: Fetches user details by their ID. Results are cached in Redis for faster subsequent access.
- Response:
```
{
    "data": {
        "id": 34452,
        "email": "SantiagoMartin@gmail.org",
        "firstname": "Santiago",
        "lastname": "Martin",
        "created_at": "2015-06-11T20:27:11Z"
    }
}
```

## How to Run

1. cd viswals-backend-test
2. Set up environment variables:
    - Copy the .env.example file to .env.
    - Update the configuration values as needed.
3. Build and run services:
    - docker-compose up --build
4. Access APIs:
- Get All Users: localhost:8080/users?page=0&page_size=100&id:min=500&id:max=2000&sort=id:ASC
- Get User By Id : http://localhost:8080/users/<user id>

## How to Debug Producer & Consumer Manually

### Producer 
1. cd producer
2. Set up environment variables for producer.
3. Ensure that rabbitmq is running, If not running then run it using `docker compose up rabbitmq`
4. Run `go run main.go --filepath=../users.go`

### Consumer
1. cd consumer
2. Set up environment variables for consumer.
3. Ensure that rabbitmq, redis and postgres is running, If not running then run it using `docker compose up rabbitmq redis postgres`
4. Run `go run main.go --consume` ( specify flag consume if you want to consume data from the rabbitmq queue )

## TODO

- Write test cases for remaining services.
- Add HTML page to showcase SSE.
- Optimize consumer service for consuming rabbitmq messages and add proper caching layer.