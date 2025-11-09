package com.librarymanagement.library_management_system.service;

import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Category;

import java.util.List;
import java.util.Optional;

public interface BookService {

    Book saveBook(Book book);

    List<Book> getAllBooks();

    Optional<Book> getBookById(Long id);

    void deleteBook(Long id);

    Optional<Book> getBookByTitle(String title);

    List<Book> getBooksByAuthor(String author);

    List<Book> getBooksByCategory(Category category);

    List<Book> searchBooks(String keyword);

    List<Book> getBooksByStatus(String status);

    long countBooks();

}