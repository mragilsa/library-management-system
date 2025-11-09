package com.librarymanagement.library_management_system.repository;

import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Category;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public interface BookRepository extends JpaRepository<Book, Long> {

    Optional<Book> findByTitle(String title);

    List<Book> findByAuthorContaining(String author);

    List<Book> findByCategory(Category category);

    List<Book> findByTitleContainingOrAuthorContaining(String title, String author);

    List<Book> findByStatus(String status);
}