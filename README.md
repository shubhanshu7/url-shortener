URL Shortener

This is a simple URL shortener service implemented in Go. It provides a REST API for shortening URLs and redirecting users to the original URLs.

## Features

- Accepts a long URL as input and returns a shortened URL..
- Redirects users to the original URL
- Stores URL mappings and metrics in-memory.
- Provides an API endpoint to retrieve the top 3 domains with the most shortened URLs.

## Getting Started

### Prerequisites

- Go (version 1.13 or higher)
- Git

### Installation

1. Clone the repository:
   git clone https://github.com/your-username/url-shortener.git

2. Navigate to the project repository:
  cd url-shortener.

3. RUn the Application
   1. go run main.go
   2. Docker-
       - install docker desktop
       - docker build -t url-shortener .
       - check status of container(docker ps -a)
       - docker run -p 8080:8080 url-shortener
     

## Usage

- curl -X POST -H "Content-Type: application/json" -d '{"url":"https://example.com"}' http://localhost:8080/shorten
- curl http://localhost:8080/redirect?short_url=<shortened-url>
- curl http://localhost:8080/metrics


## Testing

go test -v ./...






