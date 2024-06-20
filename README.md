# Boilerplate Auth Golang

This project is a boilerplate for implementing authentication in Golang applications using JSON Web Tokens (JWT).

## Project Structure

```
├── db
│   └── db.go
├── handlers
│   └── auth.go
├── main.go
├── middlewares
│   └── jwt
│       └── jwt.go
└── models
    └── user.go
```

### Directory and File Description

**`db/`**: Contains `db.go`, which configures the MongoDB database connection.

**`handlers/`**: Includes `auth.go`, handling authentication requests.

**`main.go`**: The main file starting the application. It includes a `/private` route that can serve as a guide for using JWT authentication middleware.

**`middlewares/jwt/`**: Contains `jwt.go`, middleware for protecting routes and verifying authentication using JWT.

**`models/`**: Defines `user.go`, specifying the user model structure.

## Getting Started

**Environment Setup**:
   - Create a `.env` file in the project root directory.
   - Define the necessary environment variables:

     ```plaintext
     DB_HOST=localhost
     DB_PORT=27017
     DB_NAME=app
     DB_USER=myuser
     DB_PASSWORD=mypassword
     JWT_SECRET=your_secret_key
     ```

     Adjust these values according to your local MongoDB configuration and the JWT secret key required for authentication.

**Running the Application**:
   - Ensure MongoDB is installed and configured.
   - Start the application by running:

     ```bash
     go run main.go
     ```

     This will start the server on `localhost:8080`.

**Authentication Endpoints**:
   - `POST /signup`: Register new users.
   - `POST /login`: Log in and obtain a JWT token.
   - Customize the handlers in `handlers/auth.go` as per your application's requirements.

## Contribution

Contributions are welcome! If you wish to improve this project, please send a pull request.
