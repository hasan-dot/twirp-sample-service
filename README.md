# Twirp Sample Service

## Introduction

This is a sample Twirp service that demonstrates how to use Twirp to build a simple RPC (Remote Procedure Call) service using [twitchtv/twirp](https://github.com/twitchtv/twirp).

## Installation and Setup

### Development Environment


To install and run the Twirp service, follow these steps:

1. Clone the repository: `git clone https://github.com/hasan-dot/twirp-sample-service.git`
2. Open the project in VScode dev containers using the following steps:
   1. Click on Remote Explorer VS code extention.
   2. Select "Dev Containers" from the top dropdown list.
   3. Once you have the project open in VS Code, click on the "+" icon at the top-right corner of the window, and select "Open Currrent Folder in Container" from the menu.
   4. Once the dev container is built and running, you'll be working in a fully configured development environment. You can start coding, running, and debugging your project as you normally would.
> [!TIP]
> Remember, any changes you make inside the dev container will also be reflected in your local project directory, since the dev container shares its filesystem with your local project directory.
3. Build the project and install tools using `make` command. This will install all tools, dependencies and build the project.
4. Run the server using `./bin/twirp-server` or using the 'Run and Debug' launch menu in VS code.
> [!IMPORTANT]
> You need to re-build the files and restart the server if you have done any code changes in the project.

### Example Usage

#### Go Client

```bash
./bin/twirp-client
```

#### cURL
Once the service is running, you can make RPC requests to it using a Twirp client. Here's an example using cURL:

```bash
curl -X POST \
 -H 'Content-Type: application/json' \
 -d '{"inches": 10}' \
 http://localhost:8080/twirp/service.Haberdasher/MakeHat | jq
```
