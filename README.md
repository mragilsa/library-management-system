# Library Management System

A RESTful Library Management System built with **Spring Boot** and **MySQL**. The application handles users, books, loans, and fines, including authentication and admin functionalities for managing the library effectively.

---

## Features
- User signup & login
- Admin dashboard with basic stats
- Add, update, delete, search books
- Borrow & return books
- Track and approve fines
- Filter books by category or status

---

## Setup

### 1. Clone
```bash
git clone https://github.com/mragilsa/library-management-system.git
```
```
cd library-management-system
```
### 2. Database
- Schema provided in /resources/db/schema.sql
- Create database libraryhub_db in MySQL
- Spring Boot will initialize tables automatically

### 3. Configure application.properties
Check and update database connection in src/main/resources/application.properties:
```
spring.datasource.url=jdbc:mysql://localhost:3306/libraryhub_db?useSSL=false&serverTimezone=Asia/Jakarta&allowPublicKeyRetrieval=true
```
```
spring.datasource.username=YOUR_DB_USERNAME
```
```
spring.datasource.password=YOUR_DB_PASSWORD
```

### 4. Build & Run
```
mvn clean install
```
```
mvn spring-boot:run
```

Access the app at: http://localhost:8080

---

## Example API Endpoints
- **Signup:** `POST /api/auth/signup` – create user  
- **Login:** `POST /api/auth/login` – authenticate user; admin sees stats  
- **Borrow Book:** `POST /api/loans/borrow/{bookId}` – borrow a book using book ID + user credentials  
- **Return Book:** `POST /api/loans/return/{loanId}` – mark borrowed book as returned  
- **Fines:** `GET /api/fines/unpaid` – list unpaid fines; applicable for overdue books  
- **Approve Fine:** `POST /api/fines/approve/{fineId}` – mark fine as paid (only if book is returned)  

For full API documentation, check Swagger UI:
http://localhost:8080/swagger-ui/index.html

