# Go-Based Web Server for Book Management

This project is a web server built with Go and the Gin framework, focusing on user and book management. It features user authentication, book operations, user profile management, and utilizes JWT for secure access control.

## Features

- **User Authentication**: Register and login functionality with JWT token generation for session management.
- **Book Management**: Add, update, delete, and fetch books. Books can be managed by admin users.
- **User Profile Management**: Fetch user profiles and their associated books, add books to user profiles with ratings and feedback, and update feedback for books.
- **Secure Access Control**: Middleware ensures endpoints are accessed only by authenticated users, with additional checks for admin role for specific operations.
- **Dynamic Routing**: Organized and secured routes for handling different functionalities within the application.

## Getting Started

### Prerequisites

- Go installed on your machine.
- Gin web framework.
- A compatible database setup (refer to `models` package for schema).

### Installation

1. Clone the repository to your local machine.
2. Ensure Go is installed and set up correctly.
3. Install the Gin framework and any other dependencies.
4. Set up your database and ensure it is running.
5. Update database connection settings in the application.

### Running the Server

Execute the following command from the root directory of the project:
This will start the server on the default port. Access the API endpoints through your preferred API client or browser.

## API Endpoints

- `POST /login`: Authenticate user and receive JWT token.
- `POST /register`: Register a new user.
- `GET /profile/:id`: Get user profile and books (Authenticated).
- `POST /add/:id`: Add a book to user profile with rating and feedback (Authenticated).
- `PUT /profile/:id`: Update feedback for a book (Authenticated).
- `GET /reviews/:id`: Get reviews for a book (Authenticated).
- `GET /`: Fetch all books (Authenticated).
- `GET /:id`: Get a book by ID (Authenticated).
- `POST /create`: Add a new book (Admin only).
- `DELETE /:id`: Delete a book by ID (Admin only).
- `GET /category/:category`: Get books by category (Authenticated).

## Contributing

Contributions are welcome! For major changes, please open an issue first to discuss what you would like to change.

