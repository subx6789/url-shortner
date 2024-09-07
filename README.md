# URL Shortener

## Overview

The URL Shortener is a simple web application that allows users to shorten long URLs and redirect short URLs to their original destinations. This project demonstrates a URL shortening service using Go with a basic in-memory database.

## Features

- **Shorten URLs**: Convert long URLs into shorter, more manageable URLs.
- **Redirect**: Redirect short URLs to their original destinations.
- **Unique ID Generation**: Uses UUIDs for unique identifier generation.
- **Short URL Generation**: Uses MD5 hashing to generate short URLs.

## Technologies Used

- Go (Golang)
- MD5 Hashing for short URL generation
- UUIDs for unique identification
- HTTP server for handling requests

## Installation

### Prerequisites

- Go (version 1.16 or higher)

### Clone the Repository

```bash
 git clone https://github.com/subx6789/url-shortner.git
 cd url-shortner
```

### Build and Run

#### Build the Application

```bash
 go build -o url-shortner
```

### Run the Application

```bash
 ./url-shortner
```

The server will start on port 8080. You should see a message indicating that the server is running.

## API Endpoints

1. **Root Endpoint**

GET `/`

- **Description:** Returns a welcome message.
- **Response:** `Welcome to the URL Shortener`

2. **Shorten URL**

POST `/shorten`

- **Description:** Shortens a given URL.
- **Request Body:**

  ```json
  {
    "url": "http://example.com/very-long-url"
  }
  ```

- **Response:**

  ```json
  {
    "short_url": "abcd1234"
  }
  ```

- **Status Codes:**

- `200 OK` if the URL was successfully shortened.
- `400 Bad Request` if the request body is invalid.
- `405 Method Not Allowed` if the request method is not POST.

3. **Redirect URL**

GET `/redirect/{short_url}`

- **Description:** Redirects to the original URL corresponding to the given short URL.
- **Example Request:** `/redirect/abcd1234`
- **Response:** Redirects to the original URL.
- **Status Codes:**

- `302 Found` if the redirection is successful.
- `404 Not Found` if the short URL does not exist.
- `405 Method Not Allowed` if the request method is not GET.

## Contributing

If you want to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgements

- [Go Programming Language](https://go.dev/)
- [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier)
- [MD5 Hashing](https://en.wikipedia.org/wiki/MD5)
