package com.librarymanagement.library_management_system.repository;

import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.model.User;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface LoanRepository extends JpaRepository<Loan, Long> {

    List<Loan> findByUser(User user);

    List<Loan> findByBook(Book book);

    List<Loan> findByStatus(Loan.Status status);

    @Query("SELECT l FROM Loan l WHERE l.book.title LIKE %:keywords% OR l.user.username LIKE %:keyword%")
    List<Loan> searchByBookTitleOrUserUsername(@Param("keyword") String keyword);

}
