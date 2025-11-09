package com.librarymanagement.library_management_system.controller;

import com.librarymanagement.library_management_system.dto.BorrowRequest;
import com.librarymanagement.library_management_system.dto.LoanResponse;
import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.model.User;
import com.librarymanagement.library_management_system.service.BookService;
import com.librarymanagement.library_management_system.service.LoanService;
import com.librarymanagement.library_management_system.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCrypt;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/loans")
public class LoanController {

    @Autowired
    private LoanService loanService;

    @Autowired
    private BookService bookService;

    @Autowired
    private UserService userService;

    @GetMapping
    public ResponseEntity<List<LoanResponse>> getAllLoans() {
        List<Loan> loans = loanService.getAllLoans();
        List<LoanResponse> response = loans.stream().map(loan -> new LoanResponse(
                loan.getLoanId(),
                loan.getUser().getUserId(),
                loan.getUser().getUsername(),
                loan.getBook().getBookId(),
                loan.getBook().getTitle(),
                loan.getLoanDate(),
                loan.getDueDate(),
                loan.getReturnDate(),
                loan.getStatus().name()
        )).toList();
        return ResponseEntity.ok(response);
    }

    @GetMapping("/user/{userId}")
    public ResponseEntity<?> getLoansByUser(@PathVariable Long userId) {
        Optional<User> userOpt = userService.getUserById(userId);
        if (userOpt.isEmpty()) {
            return ResponseEntity.status(HttpStatus.NOT_FOUND)
                    .body("User dengan ID " + userId + " tidak ditemukan.");
        }
        List<Loan> loans = loanService.getLoansByUser(userOpt.get());
        List<LoanResponse> response = loans.stream().map(loan -> new LoanResponse(
                loan.getLoanId(),
                loan.getUser().getUserId(),
                loan.getUser().getUsername(),
                loan.getBook().getBookId(),
                loan.getBook().getTitle(),
                loan.getLoanDate(),
                loan.getDueDate(),
                loan.getReturnDate(),
                loan.getStatus().name()
        )).toList();
        return ResponseEntity.ok(response);
    }

    @PostMapping("/borrow/{bookId}")
    public ResponseEntity<?> borrowBook(@PathVariable Long bookId, @RequestBody BorrowRequest request) {

        Optional<Book> bookOpt = bookService.getBookById(bookId);
        if (bookOpt.isEmpty()) return ResponseEntity.badRequest().body("Book not found");

        Optional<User> userOpt = userService.getUserByUsername(request.username());
        if (userOpt.isEmpty()) return ResponseEntity.badRequest().body("User not found");

        if (!BCrypt.checkpw(request.password(), userOpt.get().getPassword()))
            return ResponseEntity.status(HttpStatus.UNAUTHORIZED).body("Invalid password");

        Loan loan = new Loan();
        loan.setBook(bookOpt.get());
        loan.setUser(userOpt.get());

        loanService.saveLoan(loan);

        return ResponseEntity.ok("Book borrowed successfully");
    }

    @PostMapping("/return/{loanId}")
    public ResponseEntity<String> returnBook(@PathVariable Long loanId) {

        Optional<Loan> loanOpt = loanService.getLoanById(loanId);
        if (loanOpt.isEmpty()) return ResponseEntity.badRequest().body("Loan not found");

        loanService.updateLoanStatus(loanId, Loan.Status.RETURNED);

        return ResponseEntity.ok("Book returned successfully");
    }
}