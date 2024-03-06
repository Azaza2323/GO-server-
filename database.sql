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

INSERT INTO books (name, author, description, category, image) VALUES
                                                                   ('1984', 'George Orwell', 'Dystopian political novel.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/en/5/51/1984_first_edition_cover.jpg'),
                                                                   ('To Kill a Mockingbird', 'Harper Lee', 'A novel about racism and injustice.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/4/4f/To_Kill_a_Mockingbird_%28first_edition_cover%29.jpg/330px-To_Kill_a_Mockingbird_%28first_edition_cover%29.jpg'),
                                                                   ('The Great Gatsby', 'F. Scott Fitzgerald', 'A novel on themes of decadence and excess.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/7/7a/The_Great_Gatsby_Cover_1925_Retouched.jpg/800px-The_Great_Gatsby_Cover_1925_Retouched.jpg'),
                                                                   ('Pride and Prejudice', 'Jane Austen', 'A romantic novel of manners.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/17/PrideAndPrejudiceTitlePage.jpg/330px-PrideAndPrejudiceTitlePage.jpg'),
                                                                   ('Brave New World', 'Aldous Huxley', 'A novel anticipating future scientific advancements.', 'Science Fiction', 'https://upload.wikimedia.org/wikipedia/en/6/62/BraveNewWorld_FirstEdition.jpg'),
                                                                   ('The Catcher in the Rye', 'J.D. Salinger', 'A novel about teenage angst and alienation.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/8/89/The_Catcher_in_the_Rye_%281951%2C_first_edition_cover%29.jpg/800px-The_Catcher_in_the_Rye_%281951%2C_first_edition_cover%29.jpg'),
                                                                   ('The Hobbit', 'J.R.R. Tolkien', 'A fantasy novel about the journey of Bilbo Baggins.', 'Fantasy', 'https://upload.wikimedia.org/wikipedia/en/4/4a/TheHobbit_FirstEdition.jpg'),
                                                                   ('Moby Dick', 'Herman Melville', 'A novel about Captain Ahab’s pursuit of the giant whale Moby Dick.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/3/36/Moby-Dick_FE_title_page.jpg/330px-Moby-Dick_FE_title_page.jpg'),
                                                                   ('War and Peace', 'Leo Tolstoy', 'An epic novel about the French invasion of Russia.', 'Historical Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/a/af/Tolstoy_-_War_and_Peace_-_first_edition%2C_1869.jpg/330px-Tolstoy_-_War_and_Peace_-_first_edition%2C_1869.jpg'),
                                                                   ('Hamlet', 'William Shakespeare', 'A tragedy about the Prince of Denmark.', 'Drama', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6a/Edwin_Booth_Hamlet_1870.jpg/330px-Edwin_Booth_Hamlet_1870.jpg');



INSERT INTO books (name, author, description, category, image) VALUES
                                                                   ('Crime and Punishment', 'Fyodor Dostoevsky', 'A psychological novel about the moral dilemmas of a poor ex-student.', 'Classic Literature', 'https://upload.wikimedia.org/wikipedia/en/4/4b/Crimeandpunishmentcover.png'),
                                                                   ('Anna Karenina', 'Leo Tolstoy', 'A novel about love, betrayal and death in Russian society.', 'Classic Literature', 'https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/AnnaKareninaTitle.jpg/330px-AnnaKareninaTitle.jpg'),
                                                                   ('The Adventures of Huckleberry Finn', 'Mark Twain', 'A novel about the adventures of a young boy and an escaped slave.', 'Classic Literature', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/61/Huckleberry_Finn_book.JPG/330px-Huckleberry_Finn_book.JPG'),
                                                                   ('The Odyssey', 'Homer', 'An epic poem about Odysseus journey home after the fall of Troy.', 'Epic Poetry', 'https://upload.wikimedia.org/wikipedia/commons/thumb/1/1a/Odyssey-crop.jpg/330px-Odyssey-crop.jpg'),
                                                                   ('Invisible Man', 'Ralph Ellison', 'A novel on the social and intellectual issues facing African Americans.', 'Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/e/ee/Invisible_Man_%281952_1st_ed_jacket_cover%29.jpg/330px-Invisible_Man_%281952_1st_ed_jacket_cover%29.jpg'),
                                                                   ('Beloved', 'Toni Morrison', 'A novel about an African American family post-Civil War.', 'Historical Fiction', 'https://upload.wikimedia.org/wikipedia/commons/thumb/6/6f/Beloved_%281987_1st_ed_dust_jacket_cover%29.jpg/330px-Beloved_%281987_1st_ed_dust_jacket_cover%29.jpg'),
                                                                   ('Mrs. Dalloway', 'Virginia Woolf', 'A novel detailing a day in the life of Clarissa Dalloway in post-World War I England.', 'Modernist Literature', 'https://upload.wikimedia.org/wikipedia/en/thumb/6/67/Mrs._Dalloway_cover.jpg/330px-Mrs._Dalloway_cover.jpg'),
                                                                   ('Lolita', 'Vladimir Nabokov', 'A novel about a literature professor’s obsession with a twelve-year-old girl.', 'Modern Literature', 'https://upload.wikimedia.org/wikipedia/commons/thumb/5/57/Lolita_1955.JPG/330px-Lolita_1955.JPG'),
                                                                   ('The Catch-22', 'Joseph Heller', 'A satirical novel about World War II soldiers.', 'Satire', 'https://upload.wikimedia.org/wikipedia/en/9/99/Catch22.jpg'),
                                                                   ('Don Quixote', 'Miguel de Cervantes', 'A novel about the adventures of a nobleman who believes he is a knight.', 'Classic Literature', 'https://upload.wikimedia.org/wikipedia/commons/thumb/b/ba/Title_page_first_edition_Don_Quijote.jpg/375px-Title_page_first_edition_Don_Quijote.jpg');
