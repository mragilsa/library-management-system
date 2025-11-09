package com.librarymanagement.library_management_system.controller;

import com.librarymanagement.library_management_system.dto.BookRequest;
import com.librarymanagement.library_management_system.model.Book;
import com.librarymanagement.library_management_system.model.Category;
import com.librarymanagement.library_management_system.service.BookService;
import com.librarymanagement.library_management_system.service.CategoryService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Map;
import java.util.Optional;

@RestController
@RequestMapping("/api/books")
public class BookController {

    @Autowired
    private BookService bookService;

    @Autowired
    private CategoryService categoryService;

    @GetMapping
    public ResponseEntity<List<Book>> getAllBooks() {
        List<Book> books = bookService.getAllBooks();
        return ResponseEntity.ok(books);
    }

    @PostMapping
    public ResponseEntity<?> addBook(@RequestBody BookRequest request) {
        Optional<Category> categoryOpt = categoryService.getCategoryById(request.categoryId());
        if (categoryOpt.isEmpty()) {
            return ResponseEntity.badRequest().body("Category not found");
        }

        Book book = new Book();
        book.setIsbn(request.isbn());
        book.setTitle(request.title());
        book.setAuthor(request.author());
        book.setPublisher(request.publisher());
        book.setPublishYear(request.publishYear());
        book.setTotalCopies(request.totalCopies());
        book.setAvailableCopies(request.availableCopies());
        book.setCategory(categoryOpt.get());

        if (book.getAvailableCopies() > 0) {
            book.setStatus("AVAILABLE");
        } else {
            book.setStatus("NOT AVAILABLE");
        }

        Book savedBook = bookService.saveBook(book);
        return ResponseEntity.ok(savedBook);
    }

    @GetMapping("/{id}")
    public ResponseEntity<?> getBookById(@PathVariable Long id) {
        Optional<Book> bookOpt = bookService.getBookById(id);
        return bookOpt.map(ResponseEntity::ok)
                .orElse(ResponseEntity.notFound().build());
    }

    @PutMapping("/{id}")
    public ResponseEntity<?> updateBook(@PathVariable Long id, @RequestBody BookRequest request) {
        Optional<Book> existingBookOpt = bookService.getBookById(id);
        if (existingBookOpt.isEmpty()) {
            return ResponseEntity.notFound().build();
        }

        Book existingBook = existingBookOpt.get();

        if (request.isbn() != null) existingBook.setIsbn(request.isbn());
        if (request.title() != null) existingBook.setTitle(request.title());
        if (request.author() != null) existingBook.setAuthor(request.author());
        if (request.publisher() != null) existingBook.setPublisher(request.publisher());
        if (request.publishYear() != null) existingBook.setPublishYear(request.publishYear());
        if (request.totalCopies() != null) existingBook.setTotalCopies(request.totalCopies());
        if (request.availableCopies() != null) existingBook.setAvailableCopies(request.availableCopies());
        if (request.categoryId() != null && request.categoryId() > 0) {
            Optional<Category> categoryOpt = categoryService.getCategoryById(request.categoryId());
            categoryOpt.ifPresent(existingBook::setCategory);
        }

        if (existingBook.getAvailableCopies() > 0) {
            existingBook.setStatus("AVAILABLE");
        } else {
            existingBook.setStatus("NOT AVAILABLE");
        }

        Book updatedBook = bookService.saveBook(existingBook);
        return ResponseEntity.ok(updatedBook);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<?> deleteBook(@PathVariable Long id) {
        Optional<Book> existingBook = bookService.getBookById(id);
        if (existingBook.isPresent()) {
            bookService.deleteBook(id);
            return ResponseEntity.ok(Map.of("message", "Your book was successfully deleted!"));
        }
        return ResponseEntity.notFound().build();
    }

    @GetMapping("/search")
    public ResponseEntity<List<Book>> searchBooks(@RequestParam String keyword) {
        if (keyword == null || keyword.trim().isEmpty()) {
            return ResponseEntity.ok(bookService.getAllBooks());
        }
        List<Book> books = bookService.searchBooks(keyword);
        return ResponseEntity.ok(books);
    }

    @GetMapping("/status/{status}")
    public ResponseEntity<List<Book>> filterByStatus(@PathVariable String status) {
        List<Book> books = bookService.getBooksByStatus(status.toUpperCase());
        return ResponseEntity.ok(books);
    }

    @GetMapping("/category/{categoryId}")
    public ResponseEntity<List<Book>> filterByCategory(@PathVariable Long categoryId) {
        Optional<Category> categoryOpt = categoryService.getCategoryById(categoryId);
        if (categoryOpt.isEmpty()) {
            return ResponseEntity.badRequest().body(List.of());
        }
        List<Book> books = bookService.getBooksByCategory(categoryOpt.get());
        return ResponseEntity.ok(books);
    }
}