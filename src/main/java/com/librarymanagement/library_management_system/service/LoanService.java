package com.librarymanagement.library_management_system.service;

import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.model.User;
import com.librarymanagement.library_management_system.model.Book;

import java.util.List;
import java.util.Optional;

public interface LoanService {

    Loan saveLoan(Loan loan);

    List<Loan> getAllLoans();

    Optional<Loan> getLoanById(Long id);

    void deleteLoan(Long id);

    List<Loan> getLoansByUser(User user);

    List<Loan> getLoansByBook(Book book);

    List<Loan> getLoansByStatus(Loan.Status status);

    List<Loan> searchLoans(String keyword);

    Loan updateLoanStatus(Long id, Loan.Status status);

    long countLoans();

}