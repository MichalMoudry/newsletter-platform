# Newsletter platform
A repository for a web API for the newsletter platform.

[![Build and test project](https://github.com/MichalMoudry/newsletter-platform/actions/workflows/go.yml/badge.svg)](https://github.com/MichalMoudry/newsletter-platform/actions/workflows/go.yml)

# Application description
The newsletter app enables newsletter editors to register and manage their own newsletter that consumers can subscribe to. Editors can then publish news in the newsletter and send them to all subscribers.

API should be able to serve both mobile apps and websites using REST API. The API backend should use two different storage services to maintain the data. User accounts are stored in the Postgre SQL database, whereas subscribers are stored in Firebase.
## Requirements
- *Design & Implement an API for a Go Newsletter platform*
- **API architecture**: REST
- **Language**: Go
- **Database**: PostgreSQL
    - ORMs are not allowed
- For deployment, you can use any cloud platform
## Description source
MICROSERVICE DEVELOPMENT W/ GO, SEMESTRAL WORK - Version 0.1.0

# Project structure
- **App entrypoint** - Web API's entrypoint.
    - Entrypoint is located in /cmd/main.go.
- Transport layer
- Service layer
- Database layer
- Utility folder
- Config folder
- GitHub Actions definition

# Documentation
Additional information regarding this system is stored in our [wiki](https://github.com/MichalMoudry/newsletter-platform/wiki "Newsletter platform - wiki pages").

# Security


# Used libraries
...

More comprehensive information can be found in the [go.mod](./go.mod "Project's go.mod file") file.
