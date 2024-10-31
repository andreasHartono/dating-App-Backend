# Dating App Backend (Golang)

This repository contains the backend code for a dating mobile app, built with Golang and MySQL. It includes features like user registration, login, daily profile swiping, and premium purchases.

## Features

- **User Registration**: Register a new user with a unique username and password.
- **User Login**: Secure login with JWT-based authentication.
- **Profile Swiping**: Daily limit of 10 profile swipes (like or pass).
- **Premium Purchase**: Unlocks unlimited swipes and a verified user label.

## Tech Stack

- **Golang**: Backend API development
- **MySQL**: Database management
- **GORM**: ORM for Golang and MySQL integration
- **JWT**: For secure, stateless authentication
- **Postman**: API testing

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd dating-app-backend
   ```
2. Update MySQL configuration in database/db.go:
  ```Go
  dsn := "user:password@tcp(127.0.0.1:3306)/datingapp?charset=utf8mb4&parseTime=True&loc=Local"
  ```
3. Install dependencies:
  ```bash
  go mod tidy
  ```
4. Run the server:
  ```bash
  go run main.go
  ````
## Endpoints
  | Method  | Endpoint | Description |
  | ------------- | ------------- | ------------------ |
  | POST  | /signup    | Register a new user | 
  | POST  | /login     | Login with credentials |
  | GET  | /profiles   | View profiles to swipe |
  | POST  | /swipe     | Swipe left/right |
  | POST  | /purchase  | Purchase premium feature |

## Testing
1. Import the Postman collection (provided in the repository) to test each endpoint.
2. For unit testing, run:
   ```bash
   go test ./...
   ```
## Deployment and Linting
1. Docker: A Dockerfile is included to set up the app in a containerized environment.
2. Linting: Use golint for code linting
  ```bash
  golint ./...
  ```
## ERD and Sequence Diagrams
Refer to the docs/ directory for the ERD and Sequence Diagram.

## Licence
MIT License
```mathematica

Pastikan untuk melampirkan dokumentasi yang sesuai ke dalam file PDF serta menyimpan README di dalam repository GitHub agar dapat memenuhi seluruh kebutuhan submission.

```
