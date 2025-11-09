package com.librarymanagement.library_management_system.repository;

import com.librarymanagement.library_management_system.model.Fine;
import com.librarymanagement.library_management_system.model.Loan;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface FineRepository extends JpaRepository<Fine, Long> {

    List<Fine> findByLoan(Loan loan);

    List<Fine> findByStatus(Fine.Status status);

    Optional<Fine> findByLoanAndStatus(Loan loan, Fine.Status status);

}