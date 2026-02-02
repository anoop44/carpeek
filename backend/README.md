# CarPeek Backend

Backend API for the CarPeek application - a daily car identification challenge game.

## Tech Stack

- **Language**: Go (Golang)
- **Web Framework**: Gorilla Mux
- **Database**: PostgreSQL
- **Database Migration**: golang-migrate
- **Containerization**: Docker & Docker Compose

## Features

- Daily challenge API: Provides the same challenge to all users for a given day
- Challenge validation API: Validates user submissions against the solution
- RESTful API design with proper error handling

## API Endpoints

### GET `/api/v1/challenge/today`
Returns the challenge for the current day.

Response:
```json
{
  "id": 1,
  "date": "2024-01-01",
  "image_url": "/images/challenges/2024-01-01.jpg"
}
```

### GET `/api/v1/challenge/{id}`
Returns a specific challenge by ID.

Response:
```json
{
  "id": 1,
  "date": "2024-01-01",
  "image_url": "/images/challenges/2024-01-01.jpg"
}
```

### POST `/api/v1/challenge/submit`
Validates a user's submission for the current challenge.

Request Body:
```json
{
  "make_id": 1,
  "model_id": 5
}
```

Response:
```json
{
  "id": 1,
  "date": "2024-01-01",
  "correct": true,
  "partial": false,
  "message": "Perfect! You correctly identified the car.",
  "image_url": "/images/challenges/2024-01-01.jpg",
  "solution": {
    "make_name": "Toyota",
    "model_name": "Camry"
  }
}
```

## Setup Instructions

### Prerequisites

- Docker and Docker Compose
- Go 1.21+ (optional, for local development)

### Local Development

1. Clone the repository
2. Navigate to the project root directory
3. Use the local environment file: `cp .env.local .env`
4. Modify the environment variables in `.env` as needed
5. Run the application with Docker Compose: `docker-compose up --build`

### Production Deployment

1. Ensure your VM has Docker and Docker Compose installed
2. Clone the repository to your VM
3. Update the `nginx.conf` file with your domain name if applicable
4. Use the production environment file: `cp .env.production .env`
5. Run: `docker-compose up -d --build`

## Database Schema

The application uses three main tables:

- `makes`: Contains car manufacturers (Toyota, Honda, etc.)
- `models`: Contains car models linked to makes
- `challenges`: Contains daily challenges with solution references
- `schema_migrations`: Tracks applied database migrations

## Database Migrations

The application uses an automatic migration system with separate up and down migration files. Migrations are SQL files located in the `backend/migrations` folder.

### How It Works

When the application starts, it automatically:

1. Checks the `schema_migrations` table to see which migrations have been applied
2. Scans the `migrations` folder for all `*_up.sql` files
3. Applies any new up migrations in order (sorted by filename)

### Migration File Format

Migrations use separate files for up (apply) and down (rollback) operations:

```
XXX_description_up.sql    # Applied when migrating forward
XXX_description_down.sql  # Applied when rolling back
```

Where `XXX` is a sequence number (e.g., `001`, `002`, `003`).

**Examples:**
- `001_initial_schema_up.sql` - Creates tables
- `001_initial_schema_down.sql` - Drops tables
- `002_add_users_up.sql` - Adds users table
- `002_add_users_down.sql` - Removes users table

**Note:** Files without `_up` or `_down` suffix are treated as up migrations.

### Adding New Migrations

To add a new migration:

1. Create two new `.sql` files in the `backend/migrations` folder:
   - `XXX_description_up.sql` - Contains the forward migration SQL
   - `XXX_description_down.sql` - Contains the rollback SQL
2. Use the next sequence number (e.g., if `001_*` exists, use `002_*`)
3. Restart the application - the up migration will be applied automatically

**Example - Adding a users table:**

`002_add_users_up.sql`:
```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

`002_add_users_down.sql`:
```sql
DROP TABLE IF EXISTS users;
```

### Running Down Migrations (Rollback)

To rollback migrations, use the migration CLI tool:

```bash
# Build the migration tool
go build -o migrate ./cmd/migrate

# Show current migration status
./migrate -status

# Rollback all migrations
./migrate -down

# Rollback to a specific version (exclusive - stops before this version)
./migrate -down -target=001_initial_schema

# Run up migrations (default behavior)
./migrate
```

**CLI Options:**
| Flag | Description |
|------|-------------|
| `-down` | Run down migrations (rollback) |
| `-target=XXX` | Target version to rollback to (used with `-down`) |
| `-status` | Show applied migrations |
| `-help` | Show help message |

### Environment Variables

- `MIGRATIONS_PATH`: Path to migrations folder (**required**)
  - Docker: Set to `/app/migrations`
  - Local: Set to `./migrations` or absolute path

## Environment Variables

- `DATABASE_URL`: PostgreSQL connection string
- `PORT`: Port for the backend server (default: 8080)
- `DB_PASSWORD`: Password for the database user

## Docker Services

The application consists of four services:

- `postgres`: PostgreSQL database
- `backend`: Go API server
- `frontend`: Next.js frontend
- `nginx`: Reverse proxy and load balancer

## Architecture

The backend follows a clean architecture pattern:

- `cmd/server`: Entry point of the application
- `api`: Route definitions
- `handlers`: Request/response handling
- `models`: Data structures
- `database`: Database connections and queries
- `migrations`: Database schema migrations

## Adding Challenges

To add new daily challenges:

1. Insert records into the `challenges` table with the appropriate date and image URL
2. Ensure the `solution_make_id` and `solution_model_id` reference valid entries in their respective tables

## Troubleshooting

- If the database fails to start, check the `init.sql` file and ensure proper permissions
- If API endpoints return 404, verify that the router is properly configured
- For database connection issues, confirm that the `DATABASE_URL` is correctly set