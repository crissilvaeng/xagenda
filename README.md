# X-Agenda

Build an HTTP API that's responsible for handling a phone agenda.

## Features

- [X] An endpoint for pushing new contacts
- [X] An endpoint for editing contact information
- [X] An endpoint for deleting a contact
- [X] An endpoint for searching a contact by it's id
- [X] An endpoint for searching contacts by a part of their name
- [X] An endpoint that lists all contacts
- [X] The http service should be configurable through flags on startup (timeouts and port to use)
- [X] Log messages should be written for each endpoint hit, with the response status code and the
time it took to fulfill the request
- [X] If an error occurs, a log message should be written with the response status code and a helpful
error message, to help an engineer troubleshoot the issue
- [X] Service and host metrics should be collected. I suggest using Prometheus
- [X] The application should have a reasonable test coverage, preferably above 70%
- [X] The application should have end-to-end tests (this is a good way to try out the http client)
- [X] The application should contain a buildable Dockerfile
(https://levelup.gitconnected.com/complete-guide-to-create-docker-container-for-your-golan
g-application-80f3fb59a15e) -- care on this, as using plainly the scratch image might hinder
you from making https requests. Not that this will impact our example, but something to
always take care into the future
- [X] It would be nice for the application to have some type of storage to persist the data. I'll leave
this open, feel free to pick any type of storage you want

## Usage

### Running

```shell
docker compose up

curl -X POST -H "Content-Type: application/json" -d '{"name": "Joe Doe", "email": "john@doe.com", "phone": "123456789"}' http://localhost:8080/v1/api/people

curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane Doe", "email": "jane@doe.com", "phone": "123456789"}' http://localhost:8080/v1/api/people
curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane Smith", "email": "jane@smith.com", "phone": "123456789"}' http://localhost:8080/v1/api/people

curl -X GET http://localhost:8080/v1/api/people

curl -X GET http://localhost:8080/v1/api/people?q=Jane

curl -X PUT -H "Content-Type: application/json" -d '{"name": "John Doe", "phone": "123456789"}' http://localhost:8080/v1/api/people/1
curl -X GET http://localhost:8080/v1/api/people/1

curl -X DELETE http://localhost:8080/v1/api/people/1
curl -X GET http://localhost:8080/v1/api/people

docker compose down
```

### Configuration

```shell
Usage of xagenda:
  -dsn string
        DSN for the database
  -host string
        TCP address for the server to listen on (default "localhost:8080")
  -log-level uint
        Log level for the server (default 2)
  -timeout duration
        Timeout for the server to wait for a response (default 10s)
```

Also accepts as environment variables:

- `XAGENDA_DSN`
- `XAGENDA_HOST`
- `XAGENDA_LOG_LEVEL`
- `XAGENDA_TIMEOUT`

Log levels:
- 1: `Info`
- 2: `Error`

## License

The MIT License (MIT). To see the details of this license, see the [license file](LICENSE.md).

:octocat:
