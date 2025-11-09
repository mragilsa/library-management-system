package com.librarymanagement.library_management_system.dto;

public record UserResponse(
        Long userId,
        String username,
        String fullName,
        String email,
        String phone,
        String role
) {}