
# File Structure 

## Overview
This file aims to give a clear understanding of the role and functionality of each component within the project.

## Backend Files

### Go Files in the Backend

- **main.go**
  - The entry point of the Go application. It sets up the web server and ties together all the components of the backend.

- **/config/config.go**
  - Contains configuration-related code, such as reading environment variables and setting up global configuration parameters.

- **/db/database.go**
  - Manages database connections and provides an interface for database operations.

- **/db/migration/migration.go**
  - Handles database migrations, ensuring the database schema is up-to-date.

- **/db/seed/seed.go**
  - Used for seeding the database with initial data for development or testing.

- **/utils/**
  - This directory contains utility functions and helpers that are used across the application.
  - Includes files like `errs.go` for error handling, `response.go` for standardized API responses, etc.

### API Handlers

- **/api/handlers/**
  - Contains handler functions for API endpoints.
  - Files like `cloud_providers_controller.go`, `verification_controller.go`, and `certificate_controller.go` define the logic for respective API endpoints.

### Services

- **/api/services/**
  - Implements business logic and interacts with external services and databases.
  - Includes service implementations like `AWSVerificationService`, and interfaces like `VerificationService`.

## Frontend Files

- **/public/**
  - Contains HTML templates and HTMX files for the frontend.
  - `index.templ` is the main entry point for the frontend interface.

- **/dist/**
  - Contains compiled and minified CSS and JavaScript files, along with other static assets like images and icons.

## Security Files

- **/utils/security/**
  - Contains security-related modules like `auth.go`, `credentials.go`, and `encryption.go`, which handle various aspects of application security.

## Docker and Deployment

- **Dockerfile**
  - Contains instructions for building the Docker image of the Go application.

- **docker-compose.yaml**
  - Defines the multi-container Docker setup, including the Go application and the Nginx server for the frontend.

## Documentation and Configuration

- **README.md**
  - Provides an overview of the project, setup instructions, and general information.

- **.env.sample**
  - A sample environment file with necessary environment variables for running the application.

- **Makefile**
  - Contains scripts to automate common development tasks like building, running, and testing the application.

---

This file should be updated as new components are added or existing ones are modified.
