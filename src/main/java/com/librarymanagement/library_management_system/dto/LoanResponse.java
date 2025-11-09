package com.librarymanagement.library_management_system.dto;

import java.time.LocalDate;

public record LoanResponse(
        Long loanId,
        Long userId,
        String username,
        Long bookId,
        String title,
        LocalDate loanDate,
        LocalDate dueDate,
        LocalDate returnDate,
        String status
) {}