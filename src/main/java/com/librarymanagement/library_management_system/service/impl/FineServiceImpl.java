package com.librarymanagement.library_management_system.service.impl;

import com.librarymanagement.library_management_system.model.Fine;
import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.repository.FineRepository;
import com.librarymanagement.library_management_system.service.FineService;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.time.LocalDate;
import java.time.temporal.ChronoUnit;
import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Service
public class FineServiceImpl implements FineService {

    private final FineRepository fineRepository;

    public FineServiceImpl(FineRepository fineRepository) {
        this.fineRepository = fineRepository;
    }

    @Override
    public Fine saveFine(Fine fine) {
        return fineRepository.save(fine);
    }

    @Override
    public Optional<Fine> getFineById(Long id) {
        return fineRepository.findById(id);
    }

    @Override
    public List<Fine> getAllFines() {
        List<Fine> fines = fineRepository.findAll();
        return fines != null ? fines : Collections.emptyList();
    }

    @Override
    public List<Fine> getFinesByLoan(Loan loan) {
        List<Fine> fines = fineRepository.findByLoan(loan);
        return fines != null ? fines : Collections.emptyList();
    }

    @Override
    public List<Fine> getUnpaidFines() {
        List<Fine> unpaidFines = fineRepository.findByStatus(Fine.Status.UNPAID);
        return unpaidFines != null ? unpaidFines : Collections.emptyList();
    }

    @Override
    public Fine updateFineStatus(Long id, Fine.Status status) {
        Optional<Fine> fineOpt = fineRepository.findById(id);
        if (fineOpt.isPresent()) {
            Fine fine = fineOpt.get();
            fine.setStatus(status);
            return fineRepository.save(fine);
        }
        return null;
    }

    @Override
    public Fine createFineForLoan(Loan loan, double dailyAmount) {
        long overdueDays = ChronoUnit.DAYS.between(loan.getDueDate(), LocalDate.now());
        if (overdueDays <= 0) return null;

        double totalFine = overdueDays * dailyAmount;

        Fine fine = new Fine();
        fine.setLoan(loan);
        fine.setFineAmount(BigDecimal.valueOf(totalFine));
        fine.setStatus(Fine.Status.UNPAID);

        return fineRepository.save(fine);
    }

    @Override
    public long countFines() {
        return fineRepository.count();
    }
}