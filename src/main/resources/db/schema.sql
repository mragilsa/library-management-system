DROP DATABASE IF EXISTS libraryhub_db;
CREATE DATABASE libraryhub_db;
USE libraryhub_db;

CREATE TABLE categories (
  category_id INT(11) NOT NULL AUTO_INCREMENT,
  category_name VARCHAR(50) NOT NULL,
  description TEXT DEFAULT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  PRIMARY KEY (category_id),
  UNIQUE KEY category_name (category_name)
);

INSERT INTO categories(category_id,category_name,description) VALUES
(1,'Fiction','Novels and fictional stories'),
(2,'Non-Fiction','Factual and educational books'),
(3,'Science','Science and technology books'),
(4,'History','Historical books'),
(5,'Biography','Biographical works'),
(6,'Technology','Technology and programming books'),
(7,'Business','Business and management books'),
(8,'Self-Help','Self-improvement books');

CREATE TABLE books (
  book_id INT(11) NOT NULL AUTO_INCREMENT,
  isbn VARCHAR(20) NOT NULL,
  title VARCHAR(150) NOT NULL,
  author VARCHAR(100) NOT NULL,
  publisher VARCHAR(100),
  publish_year INT(11),
  category_id INT(11),
  total_copies INT(11) DEFAULT 1,
  available_copies INT(11) DEFAULT 1,
  status ENUM('AVAILABLE','NOT_AVAILABLE') DEFAULT 'AVAILABLE',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
  PRIMARY KEY(book_id),
  UNIQUE KEY isbn(isbn),
  KEY category_id(category_id),
  CONSTRAINT books_ibfk_1 FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE SET NULL
);

INSERT INTO books(isbn,title,author,publisher,publish_year,category_id,total_copies,available_copies,status)
VALUES
('9780439064873','Harry Potter and the Chamber of Secrets','J.K. Rowling','Bloomsbury',1998,1,5,5,'AVAILABLE'),
('9780062316097','Sapiens','Yuval Noah Harari','Harper',2011,4,5,5,'AVAILABLE'),
('9780134685991','The Pragmatic Programmer','Andrew Hunt','Addison-Wesley',1999,6,4,4,'AVAILABLE'),
('9780743273565','The Great Gatsby','F. Scott Fitzgerald','Scribner',1925,1,3,3,'AVAILABLE'),
('9780596009205','Head First Java','Kathy Sierra','OReilly',2003,6,4,4,'AVAILABLE');

CREATE TABLE users (
  user_id INT(11) NOT NULL AUTO_INCREMENT,
  username VARCHAR(50) NOT NULL,
  password VARCHAR(255) NOT NULL,
  full_name VARCHAR(100) NOT NULL,
  email VARCHAR(100) NOT NULL,
  phone VARCHAR(20),
  role ENUM('ADMIN','USER') DEFAULT 'USER',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
  PRIMARY KEY(user_id),
  UNIQUE KEY username(username),
  UNIQUE KEY email(email)
);

INSERT INTO users(username,password,full_name,email,phone,role)
VALUES ('admin','$2a$10$VaKdOYb2I5H6M280vNJYWeBHZBfKCkGtfhnJyQXQRH3IWY/dpcM6K','System Administrator','admin@bookverse.com','081234567890','ADMIN');

CREATE TABLE loans (
  loan_id INT(11) NOT NULL AUTO_INCREMENT,
  user_id INT(11) NOT NULL,
  book_id INT(11) NOT NULL,
  loan_date DATE NOT NULL,
  due_date DATE NOT NULL,
  return_date DATE DEFAULT NULL,
  status ENUM('BORROWED','RETURNED','OVERDUE') DEFAULT 'BORROWED',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
  PRIMARY KEY(loan_id),
  KEY idx_loan_user(user_id),
  KEY idx_loan_book(book_id),
  CONSTRAINT loans_ibfk_1 FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
  CONSTRAINT loans_ibfk_2 FOREIGN KEY (book_id) REFERENCES books(book_id) ON DELETE CASCADE
);

CREATE TABLE fines (
  fine_id BIGINT(20) NOT NULL AUTO_INCREMENT,
  loan_id INT(11) NOT NULL,
  fine_amount DECIMAL(38,2) NOT NULL,
  days_overdue INT(11) DEFAULT 0,
  fine_per_day DECIMAL(38,2),
  status ENUM('UNPAID','PAID','WAIVED') DEFAULT 'UNPAID',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP(),
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP(),
  PRIMARY KEY(fine_id),
  KEY loan_id(loan_id),
  CONSTRAINT fines_ibfk_1 FOREIGN KEY (loan_id) REFERENCES loans(loan_id) ON DELETE CASCADE
);