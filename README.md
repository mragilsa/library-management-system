# Library Management System

This project is a **Library Management System** developed using the Go programming language. It provides a simple yet powerful interface to manage books and loan records for libraries. The system supports multiple languages and allows users to perform a variety of tasks such as adding, editing, deleting, searching, and viewing books, as well as borrowing and returning books.

The system supports two main user roles:
- **Admin**:
  - Manages the library's book collection and loan records.
- **Loaner**:
  - Borrows books
  - edits loan information
  - returns books.

Additionally, the program is multilingual, allowing the user to choose between several languages, including **English**, **Indonesian**, **Arabic**, and **Chinese**.

## Features

### Admin Role:
- **Add Books**: Admins can add new books to the system with details like ID, title, author, genre, publication year, and available stock.
- **Edit Books**: Admins can edit various book details, including the title, author, genre, publication year, and stock availability.
- **Delete Books**: Admins can remove books from the library.
- **Search Books**: Admins can search for books by:
  - Book ID
  - Title
  - Author
  - Genre
  - Publication year
- **Manage Loans**: Admins can view and manage loans, including seeing which books are currently borrowed and by whom.
- **View Favorite Books**: Admins can view a list of favorite books in the library.

### Loaner Role:
- **Borrow Books**: Loaners can borrow books from the library, provided there is stock available.
- **Edit Loan Information**: Loaners can modify their loan information if needed.
- **Return Books**: Loaners can return books they have borrowed.
- **View Current Loans**: Loaners can view the list of books they have currently borrowed.

### Multi-language Support:
- The system supports four languages:
  - **Indonesian** (Bahasa Indonesia)
  - **English**
  - **Arabic**
  - **Chinese**
- The user can choose their preferred language at the start of the program.

## Setup and Installation

### Prerequisites:
- Ensure that you have Go installed on your machine. You can download it from the [official Go website](https://go.dev/dl/).

### Instructions:
1. **Clone the repository**:
   ```bash
   git clone https://github.com/mragilsa/library-management-system.git

2. **Make sure you’re in the correct directory**:
   ```bash
   cd library-management-system
3. **Run the program**:
   ```bash
   go run main.go
   
Once the program is running, you will be prompted to choose a language and your role (Admin or Loaner). Based on your role, you will be presented with different options to manage books or loans.

## System Overview and Code Structure

- **Core Components**:
	-	Book Struct: This is the central structure representing books in the system. It holds important details about each book such as ID, title, author, genre, publication year, available stock, and the number of times it has been borrowed.
	-	Loan Struct: This structure is used to store information about each loan, such as the loaner’s ID, the book being borrowed, and the loan dates.
	-	Book List: An array or list that holds all the books available in the library. It’s used by admins to perform various actions such as adding, editing, and searching for books.
	-	Loan List: An array or list that holds all the loans in the system, used for tracking the status of borrowed books and managing the loan process.

- **Key Functions**:
  -	Language Selection: The program will prompt the user to select a language at the beginning of the program. This ensures that the system is user-friendly and accessible to people who speak different languages.
	-	Admin and Loaner Menus: Based on the user’s role (Admin or Loaner), the system will display a corresponding menu with options specific to that role. Admins will have options to manage books and loans, while loaners can borrow and return books.
	-	Book Management: Admins can perform various actions related to book management, including adding new books, editing existing books, deleting books, and searching for books using multiple search criteria.
	-	Loan Management: Loaners can borrow and return books. The system tracks the loan status and ensures that borrowed books are returned on time.

### Example Interaction:

When you run the program, you will be prompted to select a language. After selecting the language, you will choose whether you are an Admin or a Loaner. Based on your selection, you will be presented with different menus that allow you to interact with the system:
	•	Admin Menu: You can add books, edit their details, delete books, and manage loans.
	•	Loaner Menu: You can borrow books, return books, and view the current status of your loans.

## License

This project is licensed under the MIT License. Feel free to use, modify, or distribute the code as per the terms of the license.

## How to Contribute

- **If you’d like to contribute to this project, here are a few steps to follow**:
	- 1.	Fork the repository: This will create a copy of the project under your own GitHub account.
	- 2.	Make changes: You can now make changes to the project. This could involve fixing bugs, adding new features, or improving documentation.
	- 3.	Create a pull request: Once you’re happy with your changes, submit a pull request to merge your changes into the main project. Be sure to explain the changes you’ve made in the pull request description.

## We welcome contributions and look forward to seeing how you can improve this project!

- **FAQ**
	- 1.	How can I change the language of the program?
At the start of the program, you will be prompted to select a language. You can choose from Indonesian, English, Arabic, or Chinese.
	- 2.	Can I add my own books to the library?
Yes, if you are an Admin, you can add new books to the library by entering details such as the book title, author, genre, and available stock.
	- 3.	Can Loaners return books they borrowed?
Yes, Loaners can return books they have borrowed by selecting the “Return loan” option in the Loaner menu.
	- 4.	How many books can I add to the system?
The system currently supports a maximum of 1000 books, which is defined by a constant NMAX in the code.
	- 5.	What happens if I try to borrow a book that is not available?
If there is no stock of a book, the system will notify you that the book is unavailable for borrowing.

This project is a great example of how Go can be used to build simple yet robust systems with multi-language support and role-based access. You can extend this project further by adding more features such as user authentication, advanced search options, or a web-based interface.
