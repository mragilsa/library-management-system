package com.librarymanagement.library_management_system.service;

import com.librarymanagement.library_management_system.model.User;
import java.util.List;
import java.util.Optional;

public interface UserService {

    User saveUser(User user);

    List<User> getAllUser();

    Optional<User> getUserById(Long id);

    void deleteUser(Long id);

    Optional<User> getUserByUsername(String username);

    Optional<User> getUserByEmail(String email);

    List<User> getUsersByRole(User.Role role);

    List<User> searchUsers(String keyword);

    long countUsers();

}
