package com.librarymanagement.library_management_system.service.impl;

import com.librarymanagement.library_management_system.model.Category;
import com.librarymanagement.library_management_system.repository.CategoryRepository;
import com.librarymanagement.library_management_system.service.CategoryService;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Service
public class CategoryServiceImpl implements CategoryService {

    private final CategoryRepository categoryRepository;

    public CategoryServiceImpl(CategoryRepository categoryRepository) {
        this.categoryRepository = categoryRepository;
    }

    @Override
    public Category saveCategory(Category category) {
        return categoryRepository.save(category);
    }

    @Override
    public List<Category> getAllCategories() {
        List<Category> categories = categoryRepository.findAll();
        return categories != null ? categories : Collections.emptyList();
    }

    @Override
    public Optional<Category> getCategoryById(Long id) {
        return categoryRepository.findById(id);
    }

    @Override
    public void deleteCategory(Long id) {
        categoryRepository.deleteById(id);
    }

    @Override
    public Optional<Category> getCategoryByName(String name) {
        return categoryRepository.findByCategoryName(name);
    }

    @Override
    public List<Category> searchCategories(String keyword) {
        List<Category> categories = categoryRepository.findByCategoryNameContaining(keyword);
        return categories != null ? categories : Collections.emptyList();
    }

    @Override
    public Category updateCategoryName(Long id, String newName) {
        Optional<Category> catOpt = categoryRepository.findById(id);
        if (catOpt.isPresent()) {
            Category category = catOpt.get();
            category.setCategoryName(newName);
            return categoryRepository.save(category);
        }
        return null;
    }
}