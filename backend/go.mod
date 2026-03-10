module autocorrect-backend

go 1.24.0

toolchain go1.24.12

require (
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/lib/pq v1.10.9
	github.com/rs/cors v1.11.1
	golang.org/x/time v0.14.0
)

require github.com/razorpay/razorpay-go v1.4.0 // indirect
