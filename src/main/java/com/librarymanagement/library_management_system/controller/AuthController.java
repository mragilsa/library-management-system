package com.librarymanagement.library_management_system.controller;

import com.librarymanagement.library_management_system.dto.LoginRequest;
import com.librarymanagement.library_management_system.dto.SignupRequest;
import com.librarymanagement.library_management_system.model.User;
import com.librarymanagement.library_management_system.service.BookService;
import com.librarymanagement.library_management_system.service.FineService;
import com.librarymanagement.library_management_system.service.LoanService;
import com.librarymanagement.library_management_system.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.bcrypt.BCrypt;
import org.springframework.web.bind.annotation.*;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

@RestController
@RequestMapping("/api/auth")
public class AuthController {

    @Autowired
    private UserService userService;

    @Autowired
    private BookService bookService;

    @Autowired
    private LoanService loanService;

    @Autowired
    private FineService fineService;

    // LOGIN
    @PostMapping("/login")
    public ResponseEntity<?> login(@RequestBody LoginRequest request) {
        String username = request.username();
        String password = request.password();

        Optional<User> userOpt = userService.getUserByUsername(username);
        if (userOpt.isPresent() && BCrypt.checkpw(password, userOpt.get().getPassword())) {
            User user = userOpt.get();
            Map<String, Object> response = new HashMap<>();
            response.put("username", user.getUsername());
            response.put("role", user.getRole());

            if (user.getRole() == User.Role.ADMIN) {
                // Admin stats
                response.put("totalUsers", userService.countUsers());
                response.put("totalBooks", bookService.countBooks());
                response.put("totalLoans", loanService.countLoans());
                response.put("totalFines", fineService.countFines());
            }

            return ResponseEntity.ok(response);
        }

        Map<String, String> error = new HashMap<>();
        error.put("error", "Invalid username or password");
        return ResponseEntity.status(401).body(error);
    }

    // SIGNUP
    @PostMapping("/signup")
    public ResponseEntity<?> signup(@RequestBody SignupRequest request) {
        if (userService.getUserByUsername(request.username()).isPresent()) {
            return ResponseEntity.badRequest().body(Map.of("error", "Username already exists"));
        }
        if (userService.getUserByEmail(request.email()).isPresent()) {
            return ResponseEntity.badRequest().body(Map.of("error", "Email already exists"));
        }

        String emailRegex = "^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+$";
        if (!request.email().matches(emailRegex)) {
            return ResponseEntity.badRequest().body(Map.of("error", "Invalid email format"));
        }

        String password = request.password();
        if (password.length() < 8 || !password.matches(".*[A-Z].*") || !password.matches(".*\\d.*")) {
            return ResponseEntity.badRequest().body(Map.of(
                    "error", "Password must be at least 8 characters, include an uppercase letter and a number"
            ));
        }

        String hashedPassword = BCrypt.hashpw(password, BCrypt.gensalt());

        User user = new User();
        user.setUsername(request.username());
        user.setPassword(hashedPassword);
        user.setFullName(request.fullName());
        user.setEmail(request.email());
        user.setPhone(request.phone());
        user.setRole(User.Role.USER);

        userService.saveUser(user);

        return ResponseEntity.ok(Map.of("success", "Account created successfully! Please login."));
    }

}