# 🤖 Server

The [Critter](https://github.com/ahashim/critter) web server.

- [Architecture](#architecture)
- [Local Development](#local-development)
- [Testing](#testing)
- [Clients](#clients)

## Architecture

#### Backend

- [Echo](https://echo.labstack.com/): High performance, extensible, minimalist Go web framework.
- [Ent](https://entgo.io/): Simple, yet powerful ORM for modeling and querying data.

#### Frontend

- [HTMX](https://htmx.org/): Access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes, so you can build modern user interfaces with the simplicity and power of hypertext.
- [Alpine.js](https://alpinejs.dev/): Rugged, minimal tool for composing behavior directly in your markup. Think of it like jQuery for the modern web. Plop in a script tag and get going.
- [Tailwind CSS](https://tailwindcss.com/): A utility-first CSS framework for rapidly building custom user interfaces.

#### Storage

- [PostgreSQL](https://www.postgresql.org/): The world's most advanced open source relational database.
- [Redis](https://redis.io/): In-memory data structure store, used as a database, cache, and message broker.

## Local Development

### Dependencies

Ensure the following are installed on your system:

- [Go](https://go.dev/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

```zsh
git clone git@github.com:ahashim/server.git
cd server
make up
```

**Note:** You might need to run `make up` as root due to `dockerd` requiring
access to unix domain sockets. More information on that
[in their documentation.](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-socket-option)

### Initialization

From the project root directory:

```zsh
make run
```

Access the application in your browser at `localhost:8000`.

### Reset

```zsh
make reset
```

This will wipe all the existing data & restart all containers.

## Testing

```zsh
make test
```

This command ensures that the tests from each package are not run in parallel.

This is required since many packages contain tests that connect to the test
database which is dropped and recreated automatically for each package.

## Clients

The following _make_ commands are available to make it easy to connect to the database and cache.

- `make db`: Connects to the primary database
- `make db-test`: Connects to the test database
- `make cache`: Connects to the primary cache
- `make cache-test`: Connects to the test cache
