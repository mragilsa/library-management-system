package com.librarymanagement.library_management_system.service;

import com.librarymanagement.library_management_system.model.Category;

import java.util.List;
import java.util.Optional;

public interface CategoryService {

    Category saveCategory(Category category);

    List<Category> getAllCategories();

    Optional<Category> getCategoryById(Long id);

    void deleteCategory(Long id);

    Optional<Category> getCategoryByName(String name);

    List<Category> searchCategories(String keyword);

    Category updateCategoryName(Long id, String newName);
}