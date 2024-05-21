# Rate-Limited Notification Service

This project implements a rate-limited notification service using Golang and the Echo framework. The service allows sending notifications of various types while enforcing rate limits to prevent spamming.

## Features

- Rate limiting based on notification type and user.
- Flexible configuration for rate limit rules.
- REST API endpoint for sending notifications.
- Integration with Redis for efficient rate limit tracking.

## Requirements

- Go (version 1.22.1)
- Echo framework (version 4.12.0)
- Redis server

## Installation

By default, the server will bbe running on port `8080`

1. Clone the repository: `git clone https://github.com/juanesmendez/rate-limited-notification-service.git`
2. Install dependencies: `go mod tidy`
3. Build the project: `go build`
4. Run the project: `./rate-limited-notification-service`

## Usage
### Sending Notifications
To send a notification, make a POST request to the `/notifications/send` endpoint with the following parameters:

* `notification_type`: The type of notification (e.g., "status", "news", "marketing").
* `user_id`: The ID of the user receiving the notification, of type string.
* `message`: The content of the notification message.

Example request:

```
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"notification_type":"status","user_id":"123","message":"Hello, world!"}' \
  http://localhost:8080/notifications/send
```