package com.librarymanagement.library_management_system.service.impl;

import com.librarymanagement.library_management_system.model.User;
import com.librarymanagement.library_management_system.repository.UserRepository;
import com.librarymanagement.library_management_system.service.UserService;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Service
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;

    public UserServiceImpl(UserRepository userRepository) {
        this.userRepository = userRepository;
    }

    @Override
    public User saveUser(User user) {
        return userRepository.save(user);
    }

    @Override
    public List<User> getAllUser() {
        List<User> users = userRepository.findAll();
        return users != null ? users : Collections.emptyList();
    }

    @Override
    public Optional<User> getUserById(Long id) {
        return userRepository.findById(id);
    }

    @Override
    public void deleteUser(Long id) {
        userRepository.deleteById(id);
    }

    @Override
    public Optional<User> getUserByUsername(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public Optional<User> getUserByEmail(String email) {
        return userRepository.findByEmail(email);
    }

    @Override
    public List<User> getUsersByRole(User.Role role) {
        List<User> users = userRepository.findByRole(role);
        return users != null ? users : Collections.emptyList();
    }

    @Override
    public List<User> searchUsers(String keyword) {
        List<User> users = userRepository.findByUsernameContainingOrEmailContaining(keyword, keyword);
        return users != null ? users : Collections.emptyList();
    }

    @Override
    public long countUsers() {
        return userRepository.count();
    }
}