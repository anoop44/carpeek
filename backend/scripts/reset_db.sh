#!/bin/bash

# Load environment variables
source ../.env.local 2>/dev/null || source ../.env 2>/dev/null

POSTGRES_USER=${POSTGRES_USER:-autocorrect_user}
POSTGRES_DB=${POSTGRES_DB:-autocorrect_db}
POSTGRES_CONTAINER="autocorrect_postgres"

echo "WARNING: This will drop ALL tables in the '$POSTGRES_DB' database within the '$POSTGRES_CONTAINER' container."
read -p "Are you sure you want to continue? (y/N) " -n 1 -r
echo    # move to a new line

if [[ $REPLY =~ ^[Yy]$ ]]
then
    echo "Dropping public schema and recreating it..."
    
    # This single command drops the public schema (along with all tables, views, etc. inside it) 
    # and then recreates it clean. This is the fastest and most thorough way to reset a database.
    docker exec -i $POSTGRES_CONTAINER psql -U $POSTGRES_USER -d $POSTGRES_DB -c "DROP SCHEMA public CASCADE; CREATE SCHEMA public;"
    
    echo "Database has been completely cleared."
    echo "You can now run your migrations again."
else
    echo "Operation cancelled."
fi
