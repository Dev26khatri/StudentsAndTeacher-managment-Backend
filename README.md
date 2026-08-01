# GOGIN

GOGIN is a Go REST API for managing users and students. It uses Gin for HTTP routing, GORM for persistence, PostgreSQL as the database, JWT bearer tokens for authentication, and bcrypt for password hashing.

> Project status: the user and student APIs are in progress. RBAC, courses, teachers, and student-fee management are planned but are **not implemented yet**.

## Tech stack

- Go
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/) with PostgreSQL
- JWT (`github.com/golang-jwt/jwt/v5`)
- bcrypt (`golang.org/x/crypto/bcrypt`)

## Getting started

### Prerequisites

- Go installed
- PostgreSQL running
- A PostgreSQL database for the application

### Configuration

Create a `.env` file in the project root. Do not commit credentials or JWT secrets.

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_database_password
DB_NAME=studentDB
DB_SSLMODE=disable
SERVER_PORT=8080
JWT_SECRET=use_a_long_random_secret_here
```

`SERVER_PORT` is currently informational: the application starts on port `8080` in `cmd/main.go`.

### Run the API

```bash
go run ./cmd
```

At startup, GORM runs migrations for the `users` and `students` tables. The API is available at `http://localhost:8080`.

### Verify the build

```bash
go test ./...
```

## Authentication and middleware

`POST /users/register` and `POST /users/login` are public. Every `/students` endpoint is protected by `JWTAuthMiddleware`.

After logging in, send the returned token on every student request:

```http
Authorization: Bearer <jwt-token>
```

The middleware rejects the request with `401 Unauthorized` when:

- the `Authorization` header is missing;
- the value does not begin with exactly `Bearer `;
- the JWT is invalid, expired, or signed with an unexpected algorithm.

Valid tokens contain the user ID and email, expire after 24 hours, and make `userID` and `emails` available in the Gin request context. Authentication only proves that a user is logged in; there is no role or ownership authorization yet.

## API endpoints

### Users

| Method | Endpoint | Auth | Description |
| --- | --- | --- | --- |
| POST | `/users/register` | Public | Register a user. Duplicate emails return `409`. |
| POST | `/users/login` | Public | Verify credentials and return a JWT. |

#### Register a user

```http
POST /users/register
Content-Type: application/json

{
  "name": "Asha Sharma",
  "email": "asha@example.com",
  "password": "a-strong-password"
}
```

Successful response: `201 Created`

```json
{
  "message": "User Created Succesfully"
}
```

#### Log in

```http
POST /users/login
Content-Type: application/json

{
  "email": "asha@example.com",
  "password": "a-strong-password"
}
```

Successful response: `200 OK`

```json
{
  "message": "Login Successfull",
  "token": "<jwt-token>"
}
```

### Students

All routes below require `Authorization: Bearer <jwt-token>`.

| Method | Endpoint | Description |
| --- | --- | --- |
| POST | `/students/` | Create a student. |
| GET | `/students/` | List students. |
| GET | `/students/:studentId` | Get one student. |
| PUT | `/students/:studentId` | Update a student. |
| DELETE | `/students/:studentId` | Soft-delete a student. |

#### Create a student

```http
POST /students/
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "name": "Rahul Kumar",
  "email": "rahul@example.com",
  "age": 20
}
```

The current service requires a non-empty name and an age of at least 18. Its current student-creation flow does not yet attach the authenticated user, so this endpoint needs completion before it can reliably satisfy the required `UserID` database field. The supplied student `email` is also not persisted; email belongs to the associated user model.

#### Update a student

```http
PUT /students/1
Authorization: Bearer <jwt-token>
Content-Type: application/json

{
  "name": "Rahul Kumar",
  "email": "rahul@example.com",
  "age": 21
}
```

The update flow requires a positive ID, non-empty name and email, and age at least 18.

## Current data model

### User

- GORM fields: `id`, `created_at`, `updated_at`, `deleted_at`
- `email` — unique and required
- `password` — bcrypt hash and required
- `role` — required in the current model

### Student

- GORM fields: `id`, `created_at`, `updated_at`, `deleted_at`
- `name`
- `age`
- `user_id` — required and unique; one student is associated with one user

Records are soft-deleted because both models embed `gorm.Model`.

## Important current limitations

- Registration currently does not set the required `role` field. If the database has no default role, PostgreSQL will reject the insert and the endpoint returns `500`. Define a safe default role (for example `student`) or require/select a role through controlled server-side logic before relying on registration.
- The register request includes `name`, but the current `User` model does not persist it.
- Student creation must connect the request to the authenticated user and set `Student.UserID`.
- RBAC is not implemented. All authenticated users can call every student endpoint.
- The student repository preloads `User`. Do not return password hashes in API responses; add response DTOs before production use.
- Request DTOs do not currently use Gin validation tags for required fields, email format, or password strength.
- Error response keys and messages are not yet consistent across endpoints.

## Roadmap

### 1. Authentication and RBAC

- Add role constants such as `admin`, `teacher`, `student`, and `accountant`.
- Set a safe default role during registration; never trust a public role value without authorization.
- Add role-based middleware, for example: only admins manage users, teachers manage their assigned courses, and accountants manage payments.
- Add ownership checks so students can only view their own permitted records.
- Add refresh-token or session revocation support if required.

### 2. Student module

- Complete the user-to-student creation transaction.
- Add profile fields, guardian/contact details, enrollment status, and academic history.
- Add pagination, filtering, and safe response DTOs.

### 3. Courses and enrollment

- Add `Course` with code, name, description, capacity, schedule, and status.
- Add an `Enrollment` join table between students and courses.
- Prevent duplicate enrollment and enforce capacity/business rules.

### 4. Teacher module

- Add teacher profiles linked to users.
- Assign teachers to courses and classes.
- Add teacher-scoped access controls and student/course views.

### 5. Fees and payments

- Add fee structures by course, term, or student category.
- Track invoices, due dates, discounts, scholarships, payments, refunds, and balances.
- Use transaction-safe payment updates and keep an immutable payment/audit history.
- Restrict finance actions to authorized roles.

### 6. Quality and security

- Add unit and integration tests, including authorization tests.
- Add request validation, consistent error DTOs, structured logging, and API documentation (OpenAPI/Swagger).
- Keep `.env` out of version control; rotate any secret that has been exposed.
- Add rate limiting, CORS configuration, HTTPS deployment, and database backups before production.

## Project structure

```text
cmd/         Application entry point
config/      Environment and PostgreSQL configuration
DTO/         Request/response data-transfer objects
middleware/  JWT authentication middleware
Routes/      Route registration
Students/    Student model, handler, service, and repository
Users/       User model, handler, service, and repository
utils/       JWT creation and validation
```
