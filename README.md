# Form3 - Client

Wrapper library created by Hugo Marcelo Del Negro for Form3 Technical Interview. This library acts as a wrapper to communicate with Form3 fake accountapi provided in a Docker image.

## Code Project Structure

 - **./account**: Contains the client wrappers to interact the with fake accountapi for the methods Create, Fetch and Delete.
 - **./client/**: Contains the needed DTO objects for input models and output models.
 - **./config/**: Contains the configuration constants and helper functions.
 - **./model/**: Contains DTO/models.
 - **./scripts/**: Contains scripts for database startup (powered by Form3).
 - **./test/**: Contains integration tests and unit tests.
 - **./test/db.go**: This is a helper file used for the integration tests to reset the database.
 - **./Dockerfile**: Builds a new Docker Image used by docker-compose to run tests. It uses golang:1.16.7-alpine which's the same version used during library development.
 - **./docker-compose.yml**: Provided by Form3 and modified to include the client wrapper container and allow tests to run.

## Tests Execution

Tests will automatically run when the `docker-compose up` command is executed. This is going to download images for Postgres, Fake Account API and build the client docker image. Once the images are available and built it will create the container instances and run the tests, once Fake Account API container is available.

<img width="796" alt="Screen Shot 2021-08-18 at 03 01 40" src="https://user-images.githubusercontent.com/5897525/129845911-74e598a8-8a5d-44fd-8f80-23e36f8d87e5.png">

## Integration Tests

Integration tests have been added to test the e2e wrapper functionality. I created db.go file which perform 2 operations:

1. Seed: Produces data to make it available for tests. (it runs at the beggining of the test).
2. Truncate: Cleans up the accounts table after tests were ran. (it runs on test cleanup).

## Unit Tests

Unit tests have been added to test individual component behavior. I created the mocks.go and http_client.go files.

1. mocks.go: This mocks the HttpClient and make the method Do() configurable so I can mock it during unit tests execution.
2. http_client.go: This is an interface needed for polymorphism, so I can mock it later using my own MockClient struct.

## Dockerization

A Dockerfile has been added to containerize the wrapper library and make it work later on docker-compose up for running all the test's batery.

## Docker Composition

The docker-compose.yml file has been extended to support the wrapper library and make it possible to execute tests as requested (during docker-compose up).

## Assumptions

- **Validations**: No client side validations were added as Form3 requested. I leave this responsability to the accountapi business domain. Instead if validations problems are being returned I bubble them up to the client code so the user learn from domain messages.
- **Tests**: Given the instructions from Form3 and the explicit need to run tests agains the fake accountapi I saw the need to add integration tests in addition to the written unit tests.

NOTE: I do not have experience with Golang.

