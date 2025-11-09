package com.librarymanagement.library_management_system.controller;

import com.librarymanagement.library_management_system.dto.FineResponse;
import com.librarymanagement.library_management_system.model.Fine;
import com.librarymanagement.library_management_system.model.Loan;
import com.librarymanagement.library_management_system.service.FineService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/api/fines")
public class FineController {

    @Autowired
    private FineService fineService;

    // LIST all fines
    @GetMapping
    public ResponseEntity<List<FineResponse>> getAllFines() {
        List<FineResponse> response = fineService.getAllFines()
                .stream()
                .map(fine -> new FineResponse(
                        fine.getFineId(),
                        fine.getLoan().getUser().getUsername(),
                        fine.getLoan().getBook().getTitle(),
                        fine.getDaysOverdue(),
                        fine.getFineAmount().intValue(),
                        fine.getStatus().toString()
                )).toList();

        return ResponseEntity.ok(response);
    }

    // LIST all unpaid fines
    @GetMapping("/unpaid")
    public ResponseEntity<List<FineResponse>> getUnpaidFines() {
        List<FineResponse> response = fineService.getUnpaidFines()
                .stream()
                .map(fine -> new FineResponse(
                        fine.getFineId(),
                        fine.getLoan().getUser().getUsername(),
                        fine.getLoan().getBook().getTitle(),
                        fine.getDaysOverdue(),
                        fine.getFineAmount().intValue(),
                        fine.getStatus().toString()
                )).toList();

        return ResponseEntity.ok(response);
    }

    // APPROVE fine by ID
    @PostMapping("/approve/{fineId}")
    public ResponseEntity<?> approveFine(@PathVariable Long fineId) {
        Optional<Fine> fineOpt = fineService.getFineById(fineId);

        if (fineOpt.isEmpty()) {
            return ResponseEntity.status(404).body("Fine not found");
        }

        Fine fine = fineOpt.get();

        if (fine.getLoan().getStatus() != Loan.Status.RETURNED) {
            return ResponseEntity.status(400).body("This fine cannot be approved because the book has not been paid and returned.");
        }

        fineService.updateFineStatus(fineId, Fine.Status.PAID);
        return ResponseEntity.ok("Fine approved successfully");
    }
}