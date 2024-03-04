create database bookstore;
use bookstore;




CREATE TABLE users (
                       id INT AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       email VARCHAR(255) NOT NULL UNIQUE,
                       hashed_password BINARY(64) NOT NULL,
                       role ENUM('user', 'admin') NOT NULL
);
CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       author VARCHAR(255) NOT NULL,
                       description TEXT,
                       category VARCHAR(100),
                       image TEXT
);

CREATE TABLE user_books (
                            user_id INT NOT NULL,
                            book_id BIGINT UNSIGNED NOT NULL,
                            PRIMARY KEY (user_id, book_id),
                            FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
                            FOREIGN KEY (book_id) REFERENCES books (id) ON DELETE CASCADE
);


INSERT INTO books (name, author, description, category, image) VALUES
                                                                   ('Book 1', 'Author 1', 'Description for Book 1', 'Fiction', 'https://example.com/image1.jpg'),
                                                                   ('Book 2', 'Author 2', 'Description for Book 2', 'Non-Fiction', 'https://example.com/image2.jpg'),
                                                                   ('Book 3', 'Author 3', 'Description for Book 3', 'Science Fiction', 'https://example.com/image3.jpg');
