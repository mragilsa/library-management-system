package com.librarymanagement.library_management_system.service.impl;

import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.model.User;
import com.librarymanagement.library_management_system.repository.LoanRepository;
import com.librarymanagement.library_management_system.service.LoanService;
import org.springframework.stereotype.Service;

import java.time.LocalDate;
import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Service
public class LoanServiceImpl implements LoanService {

    private final LoanRepository loanRepository;

    public LoanServiceImpl(LoanRepository loanRepository) {
        this.loanRepository = loanRepository;
    }

    @Override
    public Loan saveLoan(Loan loan) {

        Book book = loan.getBook();

        if (book.getAvailableCopies() <= 0) {
            throw new RuntimeException("No copies available to borrow");
        }

        book.setAvailableCopies(book.getAvailableCopies() - 1);

        loan.setLoanDate(LocalDate.now());
        loan.setDueDate(LocalDate.now().plusDays(7));
        loan.setStatus(Loan.Status.BORROWED);

        return loanRepository.save(loan);
    }

    @Override
    public List<Loan> getAllLoans() {
        List<Loan> loans = loanRepository.findAll();
        if (loans == null) return Collections.emptyList();

        for (Loan loan : loans) {
            if (loan.getStatus() == Loan.Status.BORROWED &&
                    loan.getDueDate() != null &&
                    LocalDate.now().isAfter(loan.getDueDate())) {
                loan.setStatus(Loan.Status.OVERDUE);
                loanRepository.save(loan);
            }
        }

        return loans;
    }

    @Override
    public Optional<Loan> getLoanById(Long id) {
        return loanRepository.findById(id);
    }

    @Override
    public void deleteLoan(Long id) {
        loanRepository.deleteById(id);
    }

    @Override
    public List<Loan> getLoansByUser(User user) {
        List<Loan> loans = loanRepository.findByUser(user);
        return loans != null ? loans : Collections.emptyList();
    }

    @Override
    public List<Loan> getLoansByBook(Book book) {
        List<Loan> loans = loanRepository.findByBook(book);
        return loans != null ? loans : Collections.emptyList();
    }

    @Override
    public List<Loan> getLoansByStatus(Loan.Status status) {
        List<Loan> loans = loanRepository.findByStatus(status);
        return loans != null ? loans : Collections.emptyList();
    }

    @Override
    public List<Loan> searchLoans(String keyword) {
        List<Loan> loans = loanRepository.searchByBookTitleOrUserUsername(keyword);
        return loans != null ? loans : Collections.emptyList();
    }

    @Override
    public Loan updateLoanStatus(Long id, Loan.Status status) {

        Optional<Loan> loanOpt = loanRepository.findById(id);

        if (loanOpt.isPresent()) {
            Loan loan = loanOpt.get();

            if (status == Loan.Status.RETURNED && loan.getStatus() == Loan.Status.BORROWED) {

                Book book = loan.getBook();

                int newCopies = book.getAvailableCopies() + 1;
                if (newCopies <= book.getTotalCopies()) {
                    book.setAvailableCopies(newCopies);
                } else {
                    book.setAvailableCopies(book.getTotalCopies());
                }

                loan.setReturnDate(LocalDate.now());
            }

            loan.setStatus(status);
            return loanRepository.save(loan);
        }

        return null;
    }

    @Override
    public long countLoans() {
        return loanRepository.count();
    }
}