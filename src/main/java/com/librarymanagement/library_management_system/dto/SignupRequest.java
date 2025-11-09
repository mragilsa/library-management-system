package com.librarymanagement.library_management_system.dto;

public record SignupRequest(
        String username,
        String password,
        String fullName,
        String email,
        String phone
) {}