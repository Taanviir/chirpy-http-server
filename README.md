# chirpy-http-server 🐤

## Overview
Chirpy is a simple http server built using Go. The aim of Chirpy is to let users register, login, and post chirps🐤!

## Features
Chirpy offers the following features through API endpoints:
- User authentication using JSON Web Tokens
- Protected endpoints using JWT access tokens
- Refresh tokens to extend user sessions
- CRUD operations on a PostgreSQL database for users and chirps
- Hashed passwords using Bcrypt
- Webhook handling to allow users to upgrade to Chirpy Red

## Prerequisites

- **Go**: Version 1.18 or higher.
- **PostgreSQL**: A running PostgreSQL instance for the database.
- **Goose**: To automatically do database migratiosn.
- **sqlc**: To regenerate database code from SQL queries.

## Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/Taanviir/chirpy-http-server.git
    cd chirpy-http-server
    ```

2. **Install Dependencies**:
    ```bash
    go mod tidy
    ```

3. **Set Up Environment Variables**:
    - Create a `.env` file in the current directory. Use `.env.example` to set up the `.env`.

4. **Set Up Database**:
    - Create a PostgreSQL database.
    - Apply schema migrations from `sql/schema/`.
      ```bash
      goose postgres DB_URL down
      ```

5. **Install the CLI**:
    ```bash
    go install github.com/Taanviir/chirpy@latest
    ```
    - This will compile and install the chirpy program, making it available in your $GOPATH/bin directory.
    - Ensure this directory is in your PATH to use the gator command globally.

## Usage

### CLI Commands

To use the application, run the `chirpy` executable to start the server on port `8080`.
Then use another terminal to send requests to the server using `cURL` or use a browser.

## Contributing

1. Fork the repository.
2. Create a new branch for your feature/fix.
3. Submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
