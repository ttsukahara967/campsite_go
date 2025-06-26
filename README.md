# Campsite Go API

This is a sample API server for campsite information, built with Go and Gin.

## Features

- Simple RESTful API for campsites
- MySQL backend (runs in a Docker container)
- Hot reload development with [Air](https://github.com/air-verse/air)
- Swagger UI (OpenAPI spec) available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
- (For development) [Air](https://github.com/air-verse/air) is included in the Docker image

### Quick Start

1. Clone this repository:
    ```sh
    git clone https://github.com/ttsukahara967/campsite_go.git
    cd campsite_go
    ```

2. Start the backend & database:
    ```sh
    docker-compose up
    ```

- The API server will start at [http://localhost:8080](http://localhost:8080)
- The MySQL database is started automatically (default credentials in `docker-compose.yml`)
- Swagger UI is available at [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Development with Hot Reload

You can edit Go source files and see changes immediately thanks to [Air](https://github.com/air-verse/air) – no need to restart the container manually.

---

## License

MIT
