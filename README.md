# AutoCorrect

AutoCorrect is a daily car identification challenge game where players guess vehicles from their tail lights. The app presents users with close-up images of car tail lights and challenges them to identify the make, model, and year of the vehicle.

## Architecture

The application consists of two main components:

- **Frontend**: A Next.js application located in the `frontend/` directory
- **Backend**: A Go API server located in the `backend/` directory

## Tech Stack

### Frontend
- Next.js 16
- React 19
- TypeScript
- Tailwind CSS (custom theme)

### Backend
- Go (Golang)
- Gorilla Mux
- PostgreSQL
- Docker & Docker Compose

## Setup Instructions

### Prerequisites

- Docker and Docker Compose
- Git

### Development Setup

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd AutoCorrect
   ```

2. Run the entire application with Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Access the application:
   - Frontend: http://localhost
   - Backend API: http://localhost/api/v1

### Production Deployment

1. On your 2CPU/8GB RAM VM, ensure Docker and Docker Compose are installed
2. Clone the repository to your VM
3. Configure environment variables in a `.env` file
4. Run the application in detached mode:
   ```bash
   docker-compose up -d --build
   ```

## API Endpoints

### GET `/api/v1/challenge/today`
Returns the challenge for the current day.

### POST `/api/v1/challenge/submit`
Validates a user's submission for the current challenge.

## Database Schema

The application uses PostgreSQL with the following main tables:

- `makes`: Contains car manufacturers (Toyota, Honda, etc.)
- `models`: Contains car models linked to makes
- `challenges`: Contains daily challenges with solution references

## Project Structure

```
AutoCorrect/
├── backend/          # Go API server
│   ├── api/          # Route definitions
│   ├── handlers/     # Request/response handling
│   ├── models/       # Data structures
│   ├── database/     # Database connections and queries
│   ├── migrations/   # Database schema migrations
│   └── cmd/          # Application entry points
├── frontend/         # Next.js application
├── docker-compose.yml # Container orchestration
├── nginx.conf        # Reverse proxy configuration
└── README.md
```

## Adding New Challenges

To add new daily challenges:

1. Add the car image to the appropriate directory
2. Insert a record into the `challenges` table with the date and image URL
3. Link the challenge to the correct make and model IDs

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Submit a pull request

## License

[Specify license here]