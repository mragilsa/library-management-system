package com.librarymanagement.library_management_system.dto;

public record LoginRequest (
    String username,
    String password
) {}
