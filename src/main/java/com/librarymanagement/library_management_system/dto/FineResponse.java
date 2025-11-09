package com.librarymanagement.library_management_system.dto;

public record FineResponse(
        Long fineId,
        String username,
        String bookTitle,
        Integer daysOverdue,
        Integer fineAmount,
        String status
) {}
