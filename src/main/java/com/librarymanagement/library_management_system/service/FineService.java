package com.librarymanagement.library_management_system.service;

import com.librarymanagement.library_management_system.model.Fine;
import com.librarymanagement.library_management_system.model.Loan;

import java.util.List;
import java.util.Optional;

public interface FineService {

    Fine saveFine(Fine fine);

    Optional<Fine> getFineById(Long id);

    List<Fine> getAllFines();

    List<Fine> getFinesByLoan(Loan loan);

    List<Fine> getUnpaidFines();

    Fine updateFineStatus(Long id, Fine.Status status);

    Fine createFineForLoan(Loan loan, double dailyAmount);

    long countFines();
}