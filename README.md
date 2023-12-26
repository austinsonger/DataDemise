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


### File Structure

```
├── Dockerfile
├── Makefile
├── README.md
├── _README.md
├── api
│   ├── handlers
│   │   ├── certificate_controller.go
│   │   ├── cloud_providers_controller.go
│   │   └── verification_controller.go
│   └── services
│       ├── aws.go
│       ├── aws_verification_service.go
│       ├── azure.go
│       ├── azure_verification_service.go
│       ├── gcp.go
│       ├── gcp_verification_service.go
│       ├── service_factory.go
│       └── verification_service.go
├── certificates
│   └── generator.go
├── config
│   └── config.go
├── db
│   ├── database.go
│   ├── migration
│   │   └── migration.go
│   └── seed
│       └── seed.go
├── dist
│   ├── favicon.ico
│   ├── htmx.min.js
│   ├── main.css
│   └── tailwind.css
├── docker-compose.yaml
├── docs
│   ├── Building the Application.md
│   └── Frontend with HTMX.md
├── env.sample
├── go.mod
├── go.sum
├── main.go
├── public
│   ├── base.templ
│   ├── base_templ.go
│   ├── index.templ
│   └── index_templ.go
├── tailwind.config.js
├── template
│   └── template.go
└── utils
    ├── constants
    │   └── common.go
    ├── errs
    │   └── errs.go
    ├── markd
    │   └── markdown.go
    ├── resp
    │   ├── errors.go
    │   └── response.go
    ├── routing
    │   └── setup.go
    ├── scopes
    │   └── pagination.go
    ├── security
    │   ├── access_control.go
    │   ├── api_security.go
    │   ├── audit_log.go
    │   ├── auth.go
    │   ├── config.go
    │   ├── credentials.go
    │   ├── encryption.go
    │   ├── sanitizer.go
    │   ├── security_errors.go
    │   ├── security_test.go
    │   ├── tls_config.go
    │   └── utils.go
    ├── tern
    │   └── ternary.go
    └── typeext
        └── jsonb.go
```






