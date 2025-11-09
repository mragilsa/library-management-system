package com.librarymanagement.library_management_system.dto;

public record BookRequest(
        String isbn,
        String title,
        String author,
        String publisher,
        Integer publishYear,
        Integer totalCopies,
        Integer availableCopies,
        Long categoryId
) {}