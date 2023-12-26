# DataDemise
 Cloud Data Destruction Certification/Verification

## Overview
DataDemise is an application for certifying and verifying the destruction of data stored across various cloud providers. It ensures secure and verifiable destruction of data, providing certificates as proof of destruction.

## Components:
- **Understanding Cloud Providers APIs**: Interaction with APIs of AWS, Google Cloud, and Azure to track and confirm data destruction.
- **Security and Authentication**: Enhanced security handling with custom implementations for credential management, encryption, and secure communication.
- **Data Destruction Verification**: Service implementations for verifying the complete destruction of data across different cloud platforms.
- **Certificate Generation**: Automated generation of credible certificates post data destruction using Go, with details like Cloud Provider, Service, Media Type, etc.
- **User Interface**: A user-friendly interface using HTMX for managing the data destruction process, with real-time updates and interactive elements.
- **Backend Logic**: Robust backend logic in Go (using the Echo framework) to handle requests, process data destruction, and generate certificates.

## Recent Updates
1. **GO-Echo with HTMX**: 
   - Leveraging Go for backend development, with a focus on performance, cloud integration, and security.
   - Using HTMX for dynamic frontend interaction and workflow management.

2. **Containerization with Docker**:
   - Docker setup for building and running the Go application.
   - `docker-compose.yaml` configuration for orchestrating the Go service and the HTMX frontend served via Nginx.

3. **Security Enhancements**:
   - Added a `security` directory in the project structure, containing modules like `auth.go`, `credentials.go`, `encryption.go`, and others for comprehensive security handling.

4. **API Development**:
   - Developed `api/handlers` and `api/services` with controllers and services for cloud provider interactions, data verification, and certificate generation.

5. **Verification Service**:
   - Implemented `VerificationService` interface and specific services like `AWSVerificationService` to check the complete destruction of data.







