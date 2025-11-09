package com.librarymanagement.library_management_system.service.impl;

import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Category;
import com.librarymanagement.library_management_system.repository.BookRepository;
import com.librarymanagement.library_management_system.service.BookService;
import org.springframework.stereotype.Service;

import java.util.Collections;
import java.util.List;
import java.util.Optional;

@Service
public class BookServiceImpl implements BookService {

    private final BookRepository bookRepository;

    public BookServiceImpl(BookRepository bookRepository) {
        this.bookRepository = bookRepository;
    }

    @Override
    public Book saveBook(Book book) {
        return bookRepository.save(book);
    }

    @Override
    public List<Book> getAllBooks() {
        return bookRepository.findAll();
    }

    @Override
    public Optional<Book> getBookById(Long id) {
        return bookRepository.findById(id);
    }

    @Override
    public void deleteBook(Long id) {
        bookRepository.deleteById(id);
    }

    @Override
    public Optional<Book> getBookByTitle(String title) {
        return bookRepository.findByTitle(title);
    }

    @Override
    public List<Book> getBooksByAuthor(String author) {
        return bookRepository.findByAuthorContaining(author);
    }

    @Override
    public List<Book> getBooksByCategory(Category category) {
        List<Book> books = bookRepository.findByCategory(category);
        return books != null ? books : Collections.emptyList();
    }

    @Override
    public List<Book> searchBooks(String keyword) {
        List<Book> books = bookRepository.findByTitleContainingOrAuthorContaining(keyword, keyword);
        return books != null ? books : Collections.emptyList();
    }

    @Override
    public List<Book> getBooksByStatus(String status) {
        List<Book> books = bookRepository.findByStatus(status);
        return books != null ? books : Collections.emptyList();
    }

    @Override
    public long countBooks() {
        return bookRepository.count();
    }
}