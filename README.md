# Form3 - Client

The form3 client is a wrapper library which simplifies api calls to the fake **accountapi** running on a docker container.

## Code Project Structure

 - **./account**: Contains the client wrappers to interact the with fake accountapi for the methods Create, Fetch and Delete.
 - **./client**: Contains the needed DTO objects for input models and output models.
 - **./config**: Contains constants used for configuration.
 - **./model**: Contains models.
 - **./scripts**: Contains scripts for database startup.
 - **./test**: Contains integration tests and unit tests.
 - **./test/db.go**: This is a helper file used for the integration tests to reset the database.
 - **Dockerfile**: Builds a new Docker Image used by docker-compose to run tests.

## Tests Execution

Tests will automatically run when the `docker-compose up` command is executed. This is going to download images for Postgres, Fake Account API and build the client docker image. Once the images are available and built it will create the container instances and run the tests, once Fake Account API container is available.