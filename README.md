# Multi-Tier Web Application (Go, React, Redis)

This repository contains a simple multi-tier web application designed for learning Docker containerization, networking, and deployment workflows with GitHub Actions and Docker Hub.


## Project Structure
![project structure](https://github.com/user-attachments/assets/7c3651f6-1eed-44dd-ac63-a9e1602d029e)
## Project Overview


This application demonstrates a classic 3-tier architecture:

* **Frontend (React):** A single-page React application that serves as the user interface. 
* **Backend (Go):** A Go web server that exposes an API endpoint. It connects to Redis to increment and retrieve a page view counter.
* **Database (Redis):** An in-memory data store used by the Go backend to persist the page view count.

The core purpose of this project is to provide a hands-on learning experience for:

* **Dockerizing applications:** Creating efficient `Dockerfile`s for Go, React, and Redis.
* **Docker Networking:** Understanding how to connect different containers using custom Docker networks.
* **Multi-Container Orchestration:** Setting up services that depend on each other (e.g., Go backend depending on Redis).
* **CI/CD with GitHub Actions:** Automating the build, test, and push of Docker images to Docker Hub.

## Application Flow

1.  The React frontend makes a request to the Go backend.
2.  The Go backend receives the request and interacts with the Redis database.
3.  Redis increments a counter and returns the updated value to the Go backend.
4.  The Go backend sends the counter value back to the React frontend.
5.  The React frontend displays the updated page view count.

## Technologies Used

* **Go:** Backend API server
* **Redis:** In-memory data store
* **React:** Frontend (Conceptual in this repository's scope)
* **Docker:** Containerization platform
* **GitHub Actions:** CI/CD pipeline for automated builds and deployments
* **Docker Hub:** Container image registry

## Getting Started

These instructions will get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop/) (includes Docker Engine and Docker Compose)
* [Go](https://golang.org/doc/install) (for local development of the Go backend, though not strictly required for running via Docker)

### Running Locally with Docker

This project relies on Docker's networking capabilities to allow containers to communicate.

1.  **Clone the repository:**

    ```bash
    git clone this repo
    cd YOUR_REPO_NAME/go-redis-app # Navigate to the Go backend directory
    ```

2.  **Build the Go Backend Docker Image:**
    First, ensure you have the `go-redis` dependency properly added in your `go.mod` file:
    ```bash
    go get [github.com/go-redis/redis/v8](https://github.com/go-redis/redis/v8)
    ```
    Then, build the Docker image for your Go application:

    ```bash
    docker build -t go-redis-app .
    ```

3.  **Create a Docker Network:**
    This network will allow your Go app and Redis to communicate with each other using their service names.

    ```bash
    docker network create full-stack-app-network
    ```

4.  **Run the Redis Database Container:**
    Start the Redis container and connect it to the `full-stack-app-network`.

    ```bash
    docker run -d --name my-full-stack-redis--network full-stack-app-network redis
    ```
    * `-d`: Run in detached mode (background).
    * `--name my-full-stack-redis`: Assigns a readable name `my-full-stack-redis` to the container. This name is used by your Go application to connect to Redis.
    * `--network full-stack-app-network`: Connects the container to the custom network.

5.  **Run the Go Backend Container:**
    Start your Go application container, connecting it to the same network and mapping its internal port 8080 to your host's port 8000.

    ```bash
    docker run -d --name my-full-stack-go-app -p 8000:8080 --network full-stack-app-network go-redis-app
    ```
    * `-p 8000:8080`: Maps host port `8000` to container port `8080`.
    * `--name my-full-stack-go-app`: Assigns a name to your Go application container.

6. **Run the React Frontend Container:**
     ```bash
    
     cd YOUR_REPO_NAME/react-frontend # Navigate to the react frontend directory
    ```
    Start your React application container, connecting it to the same network and mapping its internal web server port (commonly 80 for Nginx/HTTP server) to your host's port 3000 (typical for React dev servers).
   
    build the Docker image for your react application:

    ```bash
    docker build -t react-frontend-app .
    ```
    ```bash
    docker run -d --name my-full-stack-react-app -p 3000:80 --network full-stack-app-network react-frontend-app
    ```
    * `-p 3000:80`: Maps host port `3000` to container port `80` (assuming your React Dockerfile serves on port 80 internally).
    * `--name my-full-stack-react-app`: Assigns a name to your React application container.


7.  **Verify Application Status (Troubleshooting)**

    * Check if containers are running:
        ```bash
        docker ps
        ```
        You should see `my-redis` and `my-full-stack-go-app` with `Up ...` status.
    * Check logs of your Go app if it exits prematurely:
        ```bash
        docker logs my-full-stack-go-app
        ```
        This is crucial for debugging connection issues or application errors.

8.  **Access the Application:**
    Once both containers are running, open your web browser and navigate to:

    ```
    http://localhost:8000
    ```
    You should see the message 

    Simple React App with Go Backend & Redis

    Page views: 3


