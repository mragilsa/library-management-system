package com.librarymanagement.library_management_system.dto;

public record BorrowRequest(
        String username,
        String password
) {}