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

# Getting started
- Get .env file from: [VŠE - OneDrive](https://vse-my.sharepoint.com/:u:/g/personal/moum02_vse_cz/ERRL73wvUD5AmbOrxpF0GwABkiaAUCsNFWA8aAD11Larig?e=VQIUwb ".env file stored on OneDrive")
- **execute**: `source environment.env`
- **execute**: `make install_dependencies`
- **execute**: `make run`
    - or set EMAIL_SENDER_IDENTITY and SENDGRID_API_KEY (values can be found in the .env file) in the docker-compose.yml file and execute `make compose`

# Project structure
- **App entrypoint** - Web API's entrypoint.
    - Entrypoint is located in /cmd/main.go.
- Transport layer
- Service layer
- Database layer
- Utility folder
- Config folder
- **GitHub Actions definition** - Folder containing definitions (in .yml files) of GitHub Actions pipelines.

# Deployment
Currently there is no deployed version of this app.

But there is automated docker image build process. Docker image can be found here: [DockerHub - michalmoudry/newsletter-platform](https://hub.docker.com/r/michalmoudry/newsletter-platform).

# Documentation
Additional information regarding this system is stored in our [wiki](https://github.com/MichalMoudry/newsletter-platform/wiki "Newsletter platform - wiki pages").

Postman collection for the API can be found here: [Postman collection](./docs/newsletter_platform.json "JSON file with the Postman collection")

# Used libraries
...

More comprehensive information can be found in the [go.mod](./go.mod "Project's go.mod file") file.
