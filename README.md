# Form3 - Client

The form3 client is a wrapper library which simplifies api calls to the fake **accountapi** running on a docker container.

## Code Project Structure

 - **./account**: Contains the client wrappers to interact the with fake accountapi for the methods Create, Fetch and Delete.
 - **./client/**: Contains the needed DTO objects for input models and output models.
 - **./config/**: Contains constants used for configuration.
 - **./model/**: Contains models.
 - **./scripts/**: Contains scripts for database startup.
 - **./test/**: Contains integration tests and unit tests.
 - **./test/db.go**: This is a helper file used for the integration tests to reset the database.
 - **./Dockerfile**: Builds a new Docker Image used by docker-compose to run tests. It uses golang:1.16.7-alpine which's the same version used during library development.
 - **./docker-compose.yml**: Provided by Form3 and modified to include the client wrapper container and allow tests to run.

## Tests Execution

Tests will automatically run when the `docker-compose up` command is executed. This is going to download images for Postgres, Fake Account API and build the client docker image. Once the images are available and built it will create the container instances and run the tests, once Fake Account API container is available.

<img width="1114" alt="Screen Shot 2021-08-17 at 02 03 37" src="https://user-images.githubusercontent.com/5897525/129666772-236d5d3a-86ec-4a0a-96e0-3746db7d9cc3.png">

## Assumptions

- **Validations**: client-side validations have not been added for Create wrapper, because it was requested in the interview README file (should not section: don't write client side validations and keep it simple and consice). Also these validations were already handled by the domain accountapi so I return these errors to the client's user to learn about these errors and fix data when needed.
- **Tests**: Since the tests were specified to be run agains the fake account api I consider it to be an integration test.
