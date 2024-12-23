package main

import (
	"fmt"
	"time"
)

const NMAX = 1000

type book struct {
	bookId         string
	title          string
	author         string
	genre          string
	pubYear        int
	stockAvailable int
	loanCount      int
}

type peminjam struct {
	peminjamId    string
	bookId        string
	namaPeminjam  string
	tanggalPinjam string
}

type tabInt [NMAX]book
type tabPeminjam [NMAX]peminjam

func main() {
    var data tabInt
    var peminjam tabPeminjam
    var nData, nPeminjam, person, number, edit, find, lang int
    var running bool

    fmt.Println("Welcome to Our Library")
    isValid := false
	for !isValid {
    	language()
    	fmt.Print("Select your language: ")
    	fmt.Scan(&lang)
    	if lang >= 1 && lang <= 4 {
      	  isValid = true
   	 	} else {
      	  fmt.Println("Your select is invalid, please try again.")
   		}
	}

    for {
        role(lang)
        fmt.Print("Select your role: ")
        fmt.Scan(&person)
        if person == 1 || person == 2 || person == 0 {
            running = true
            if person == 1 {
                for running {
                    adminMenu(lang)
                    fmt.Print("Pilih nomor: ")
                    fmt.Scan(&number)

                    switch number {
                    case 1:
                        addBook(&data, &nData, lang)
                    case 2:
                        editMenu(lang)
                        fmt.Scan(&edit)
                        switch edit {
                        case 1:
                            editBookTitle(&data, nData, lang)
                        case 2:
                            editBookAuthor(&data, nData, lang)
                        case 3:
                            editBookGenre(&data, nData, lang)
                        case 4:
                            editBookYear(&data, nData, lang)
                        case 5:
                            editBookStock(&data, nData, lang)
                        default:
                            fmt.Println("Your select is invalid, please try again.")
                        }
                    case 3:
                        deleteBook(&data, &nData, lang)
                    case 4:
                        findMenu(lang)
                        fmt.Scan(&find)
                        switch find {
                        case 1:
                            findBookId(&data, nData, lang)
                        case 2:
                            findBookTitle(&data, nData, lang)
                        case 3:
                            findBookAuthor(&data, nData, lang)
                        case 4:
                            findBookGenre(&data, nData, lang)
                        case 5:
                            findBookYear(&data, nData, lang)
                        default:
                            fmt.Println("Your select is invalid, please try again.")
                        }
                    case 5:
                        favorite(&data, nData, lang)
                    case 6:
                        displayLibrary(&data, nData, lang)
                    case 7:
                        displayLoaner(&data, &peminjam, nData, nPeminjam, lang)
                    case 0:
                        running = false
                    default:
                        fmt.Println("Your select is invalid, please try again.")
                    }
                }
            } else if person == 2 {
                for running {
                    LoanerMenu(lang)
                    fmt.Scan(&number)

                    switch number {
                    case 1:
                        addLoan(&data, &peminjam, nData, &nPeminjam, lang)
                    case 2:
                        editLoan(&peminjam, nPeminjam, lang)
                    case 3:
                        deleteLoan(&data, &peminjam, &nData, &nPeminjam, lang)
                    case 4:
                        if nPeminjam > 0 {
                            returnBook(&data, &peminjam, nData, &nPeminjam, lang)
                        } else {
                            fmt.Println("You have not borrowed any books.")
                        }
                    case 0:
                        running = false
                    default:
                        fmt.Println("Your select is invalid, please try again.")
                    }
                }
            } else if person == 0 {
                return
            }
        } else {
            fmt.Println("Your select is invalid, please try again.")
        }
    }
}

func language() {
	fmt.Println("1. Indonesian")
	fmt.Println("2. English")
	fmt.Println("3. Arabic")
	fmt.Println("4. Chinese")
}

func role(lang int) {
	if lang == 1 {
		fmt.Println("1. Admin")
		fmt.Println("2. Peminjam")
		fmt.Println("0. Keluar")
	} else if lang == 2 {
		fmt.Println("1. Admin")
		fmt.Println("2. Loaner")
		fmt.Println("0. Exit")
	} else if lang == 3 {
		fmt.Println("1. مسؤل")
		fmt.Println("2. مستعير")
		fmt.Println("0. خروج")
	} else if lang == 4 {
		fmt.Println("1. 行政")
		fmt.Println("2. 借款人")
		fmt.Println("0. 退出")
	} else {
		fmt.Println("Your select is invalid, please try again.")
	}
}

func editMenu(lang int) {
	if lang == 1 {
		fmt.Println("1. Judul buku")
		fmt.Println("2. Penulis buku")
		fmt.Println("3. Genre buku")
		fmt.Println("4. Tahun terbit buku")
		fmt.Println("5. Stok buku yang tersedia")
		fmt.Print("Pilih tipe buku apa yang ingin anda edit: ")
	} else if lang == 2 {
		fmt.Println("1. Book title")
		fmt.Println("2. Book author")
		fmt.Println("3. Book genre")
		fmt.Println("4. Publication year of book")
		fmt.Println("5. Stock available of book")
		fmt.Print("Select what type of book you want to edit: ")
	} else if lang == 3 {
		fmt.Println("1. عنوان الكتاب")
		fmt.Println("2. مؤلف الكتاب")
		fmt.Println("3. نوع الكتاب")
		fmt.Println("4. سنة نشر الكتاب")
		fmt.Println("5. الكمية المتاحة من الكتاب")
		fmt.Print("اختر نوع الكتاب الذي ترغب في تحريره: ")
	} else if lang == 4 {
		fmt.Println("1. 书名")
		fmt.Println("2. 作者")
		fmt.Println("3. 类别")
		fmt.Println("4. 出版年份")
		fmt.Println("5. 库存量")
		fmt.Print("请选择您想编辑的书籍类型：")
	}
}

func findMenu(lang int) {
	if lang == 1 {
		fmt.Println("1. Id Buku")
		fmt.Println("2. Judul Buku")
		fmt.Println("3. Penulis Buku")
		fmt.Println("4. Genre Buku")
		fmt.Println("5. Tahun Terbit Buku")
		fmt.Print("Pilih jenis buku yang ingin Anda cari: ")
	} else if lang == 2 {
		fmt.Println("1. Book Id")
		fmt.Println("2. Book title")
		fmt.Println("3. Book author")
		fmt.Println("4. Book genre")
		fmt.Println("5. Publication year of book")
		fmt.Print("Select what type of book you want to find: ")
	} else if lang == 3 {
		fmt.Println("1. معرف الكتاب")
		fmt.Println("2. عنوان الكتاب")
		fmt.Println("3. مؤلف الكتاب")
		fmt.Println("4. نوع الكتاب")
		fmt.Println("5. سنة نشر الكتاب")
		fmt.Print("اختر نوع الكتاب الذي ترغب في العثور عليه: ")
	} else if lang == 4 {
		fmt.Println("1. 书籍编号")
		fmt.Println("2. 书名")
		fmt.Println("3. 作者")
		fmt.Println("4. 类别")
		fmt.Println("5. 出版年份")
		fmt.Print("请选择您想查找的书籍类型：")
	}

}

func adminMenu(lang int) {
	if lang == 1 {
		fmt.Println("1. Tambah buku")
		fmt.Println("2. Edit buku")
		fmt.Println("3. Hapus buku")
		fmt.Println("4. Pencarian buku")
		fmt.Println("5. Lihat buku favorit")
		fmt.Println("6. Menampilkan semua daftar buku di perpustakaan")
		fmt.Println("7. Menampilkan semua daftar buku di peminjam")
		fmt.Println("0. Kembali ke menu")
	} else if lang == 2 {
		fmt.Println("1. Add book")
		fmt.Println("2. Edit book")
		fmt.Println("3. Delete book")
		fmt.Println("4. Find book")
		fmt.Println("5. View favorite books")
		fmt.Println("6. Displays all book list in the library")
		fmt.Println("7. Displays all book list in the loaner")
		fmt.Println("0. Back to menu")
	} else if lang == 3 {
		fmt.Println("1. إضافة كتاب")
		fmt.Println("2. تعديل كتاب")
		fmt.Println("3. حذف كتاب")
		fmt.Println("4. البحث عن كتاب")
		fmt.Println("5. عرض الكتب المفضلة")
		fmt.Println("6. عرض قائمة جميع الكتب في المكتبة")
		fmt.Println("7. عرض قائمة جميع الكتب في المستعير")
		fmt.Println("0. العودة إلى القائمة")
	} else if lang == 4 {
		fmt.Println("1. 添加书籍")
		fmt.Println("2. 编辑书籍")
		fmt.Println("3. 删除书籍")
		fmt.Println("4. 查找书籍")
		fmt.Println("5. 查看喜欢的书籍")
		fmt.Println("6. 显示图书馆所有书籍列表")
		fmt.Println("7. 显示借书者所有书籍列表")
		fmt.Println("0. 返回菜单")
	}
}

func LoanerMenu(lang int) {
	if lang == 1 {
		fmt.Println("1. Tambah pinjaman")
		fmt.Println("2. Edit pinjaman")
		fmt.Println("3. Hapus pinjaman")
		fmt.Println("4. Kembalikan pinjaman")
		fmt.Println("0. Kembali ke menu")
		fmt.Print("Pilih nomor: ")
	} else if lang == 2 {
		fmt.Println("1. Add loan")
		fmt.Println("2. Edit loan")
		fmt.Println("3. Delete loan")
		fmt.Println("4. Return loan")
		fmt.Println("0. Back to menu")
		fmt.Print("Select number: ")
	} else if lang == 3 {
		fmt.Println("1. إضافة قرض")
		fmt.Println("2. تعديل قرض")
		fmt.Println("3. حذف قرض")
		fmt.Println("4. إرجاع قرض")
		fmt.Println("0. العودة إلى القائمة")
		fmt.Print("اختر رقمًا: ")
	} else if lang == 4 {
		fmt.Println("1. 添加贷款")
		fmt.Println("2. 编辑贷款")
		fmt.Println("3. 删除贷款")
		fmt.Println("4. 归还贷款")
		fmt.Println("0. 返回菜单")
		fmt.Print("请选择数字：")
	}

}

func addBook(A *tabInt, n *int, lang int) {
	var i int
	if *n > 0 {
		i = *n
	}
	if lang == 1 {
		fmt.Print("Masukkan Id buku: ")
		fmt.Scan(&(*A)[i].bookId)
		fmt.Print("Masukkan judul buku: ")
		fmt.Scan(&(*A)[i].title)
		fmt.Print("Masukkan penulis buku: ")
		fmt.Scan(&(*A)[i].author)
		fmt.Print("Masukkan genre buku: ")
		fmt.Scan(&(*A)[i].genre)
		fmt.Print("Masukkan tahun terbit buku: ")
		fmt.Scan(&(*A)[i].pubYear)
		fmt.Print("Masukkan stok: ")
		fmt.Scan(&(*A)[i].stockAvailable)
		fmt.Println("Buku berhasil ditambahkan")
	} else if lang == 2 {
		if *n > NMAX {
			*n = NMAX
		}
		fmt.Print("Enter book Id: ")
		fmt.Scan(&(*A)[i].bookId)
		fmt.Print("Enter book title: ")
		fmt.Scan(&(*A)[i].title)
		fmt.Print("Enter book author: ")
		fmt.Scan(&(*A)[i].author)
		fmt.Print("Enter book genre: ")
		fmt.Scan(&(*A)[i].genre)
		fmt.Print("Enter publication year of book: ")
		fmt.Scan(&(*A)[i].pubYear)
		fmt.Print("Enter stock: ")
		fmt.Scan(&(*A)[i].stockAvailable)
		fmt.Println("The book has been successfully added")
	} else if lang == 3 {
		if *n > NMAX {
			*n = NMAX
		}
		fmt.Print("أدخل معرف الكتاب: ")
		fmt.Scan(&(*A)[i].bookId)
		fmt.Print("أدخل عنوان الكتاب: ")
		fmt.Scan(&(*A)[i].title)
		fmt.Print("أدخل مؤلف الكتاب: ")
		fmt.Scan(&(*A)[i].author)
		fmt.Print("أدخل نوع الكتاب: ")
		fmt.Scan(&(*A)[i].genre)
		fmt.Print("أدخل سنة نشر الكتاب: ")
		fmt.Scan(&(*A)[i].pubYear)
		fmt.Print("أدخل المخزون: ")
		fmt.Scan(&(*A)[i].stockAvailable)
		fmt.Println("تمت إضافة الكتاب بنجاح")
	} else if lang == 4 {
		if *n > NMAX {
			*n = NMAX
		}
		fmt.Print("输入书籍编号: ")
		fmt.Scan(&(*A)[i].bookId)
		fmt.Print("输入书名: ")
		fmt.Scan(&(*A)[i].title)
		fmt.Print("输入作者: ")
		fmt.Scan(&(*A)[i].author)
		fmt.Print("输入书籍类别: ")
		fmt.Scan(&(*A)[i].genre)
		fmt.Print("输入出版年份: ")
		fmt.Scan(&(*A)[i].pubYear)
		fmt.Print("输入库存数量: ")
		fmt.Scan(&(*A)[i].stockAvailable)
		fmt.Println("图书已成功添加")
	}
	*n++
}


func editBookTitle(A *tabInt, n int, lang int) {
	var Id string
	var title string
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda edit: ")
		fmt.Scan(&Id)
		fmt.Print("Masukkan judul baru: ")
		fmt.Scan(&title)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].title = title
			}
		}
		fmt.Println("Buku berhasil diedit")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to edit: ")
		fmt.Scan(&Id)
		fmt.Print("Enter new title: ")
		fmt.Scan(&title)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].title = title
			}
		}
		fmt.Println("Book successfully edited")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد تحريره: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل العنوان الجديد: ")
		fmt.Scan(&title)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].title = title
			}
		}
		fmt.Println("تم تعديل الكتاب بنجاح")
	} else if lang == 4 {
		fmt.Print("输入您要编辑的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入新标题：")
		fmt.Scan(&title)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].title = title
			}
		}
		fmt.Println("书籍成功编辑")
	}
}

func editBookAuthor(A *tabInt, n int, lang int) {
	var Id string
	var author string
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda edit: ")
		fmt.Scan(&Id)
		fmt.Print("Masukkan penulis baru: ")
		fmt.Scan(&author)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].author = author
			}
		}
		fmt.Println("Buku berhasil diedit")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to edit: ")
		fmt.Scan(&Id)
		fmt.Print("Enter new author: ")
		fmt.Scan(&author)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].author = author
			}
		}
		fmt.Println("Book successfully edited")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد تحريره: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل المؤلف الجديد: ")
		fmt.Scan(&author)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].author = author
			}
		}
		fmt.Println("تم تعديل الكتاب بنجاح")
	} else if lang == 4 {
		fmt.Print("输入您要编辑的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入新作者：")
		fmt.Scan(&author)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].author = author
			}
		}
		fmt.Println("书籍成功编辑")
	}
}

func editBookGenre(A *tabInt, n int, lang int) {
	var Id string
	var genre string
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda edit: ")
		fmt.Scan(&Id)
		fmt.Print("Masukkan genre baru: ")
		fmt.Scan(&genre) 
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].genre = genre
			}
		}
		fmt.Println("Buku berhasil diedit")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to edit: ")
		fmt.Scan(&Id)
		fmt.Print("Enter new genre: ")
		fmt.Scan(&genre)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].genre = genre
			}
		}
		fmt.Println("Book successfully edited")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد تحريره: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل النوع الجديد: ")
		fmt.Scan(&genre)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].genre = genre
			}
		}
		fmt.Println("تم تعديل الكتاب بنجاح")
	} else if lang == 4 {
		fmt.Print("输入您要编辑的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入新类型：")
		fmt.Scan(&genre)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].genre = genre
			}
		}
		fmt.Println("书籍成功编辑")
	}
}

func editBookYear(A *tabInt, n int, lang int) {
	var Id string
	var year, newestYear int

	if lang == 1 {
        fmt.Print("Masukkan Id Buku yang ingin Anda edit: ")
        fmt.Scan(&Id)
        fmt.Print("Masukkan tahun baru: ")
        fmt.Scan(&year)
        for i := 0; i < n; i++ {
            if A[i].bookId == Id {
                A[i].pubYear = year
            }
        }
		fmt.Println("Buku berhasil diedit")
        newestYear = A[0].pubYear
        for i := 1; i < n; i++ {
            if A[i].pubYear > newestYear {
                newestYear = A[i].pubYear
            }
        }
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to edit: ")
		fmt.Scan(&Id)
		fmt.Print("Enter new year: ")
		fmt.Scan(&year)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].pubYear = year
			}
		}
		fmt.Println("Book successfully edited")
		newestYear = A[0].pubYear
        for i := 1; i < n; i++ {
            if A[i].pubYear > newestYear {
                newestYear = A[i].pubYear
            }
        }
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد تحريره: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل السنة الجديدة: ")
		fmt.Scan(&year)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].pubYear = year
			}
		}
		fmt.Println("تم تعديل الكتاب بنجاح")
		newestYear = A[0].pubYear
        for i := 1; i < n; i++ {
            if A[i].pubYear > newestYear {
                newestYear = A[i].pubYear
            }
        }
	} else if lang == 4 {
		fmt.Print("输入您要编辑的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入新年份：")
		fmt.Scan(&year)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].pubYear = year
			}
		}
		fmt.Println("书籍成功编辑")
		newestYear = A[0].pubYear
        for i := 1; i < n; i++ {
            if A[i].pubYear > newestYear {
                newestYear = A[i].pubYear
            }
        }
	}
}

func editBookStock(A *tabInt, n int, lang int) {
	var Id string
	var stock int
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda edit: ")
		fmt.Scan(&Id)
		fmt.Print("Masukkan stok baru: ")
		fmt.Scan(&stock)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].stockAvailable = stock
			}
		}
		fmt.Println("Buku berhasil diedit")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to edit: ")
		fmt.Scan(&Id)
		fmt.Print("Enter new stock: ")
		fmt.Scan(&stock)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].stockAvailable = stock
			}
		}
		fmt.Println("Book successfully edited")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد تحريره: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل المخزون الجديد: ")
		fmt.Scan(&stock)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].stockAvailable = stock
			}
		}
		fmt.Println("تم تعديل الكتاب بنجاح")
	} else if lang == 4 {
		fmt.Print("输入您要编辑的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入新库存：")
		fmt.Scan(&stock)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				A[i].stockAvailable = stock
			}
		}
		fmt.Println("书籍成功编辑")
	}
}

func deleteBook(A *tabInt, n *int, lang int) {
	var Id string
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda hapus: ")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to delete: ")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد حذفه: ")
	} else if lang == 4 {
		fmt.Print("输入您要删除的书籍编号：")
	}
	fmt.Scan(&Id)

	found := false
	for i := 0; i < *n; i++ {
		if A[i].bookId == Id {
			found = true
			for j := i; j < *n-1; j++ {
				A[j] = A[j+1]
			}
			*n--
			i--
		}
	}

	if found {
		if lang == 1 {
			fmt.Println("Buku berhasil dihapus")
		} else if lang == 2 {
			fmt.Println("Book successfully deleted")
		} else if lang == 3 {
			fmt.Println("تم حذف الكتاب بنجاح")
		} else if lang == 4 {
			fmt.Println("书籍成功删除")
		}
	} else {
		if lang == 1 {
			fmt.Println("Buku tidak ditemukan")
		} else if lang == 2 {
			fmt.Println("Book not found")
		} else if lang == 3 {
			fmt.Println("الكتاب غير موجود")
		} else if lang == 4 {
			fmt.Println("找不到书籍")
		}
	}
}

func findBookId(A *tabInt, n int, lang int) {
	//Sequential Search
    var Id string
    var found bool = false

    if lang == 1 {
        fmt.Print("Masukkan Id Buku yang ingin Anda temukan: ")
        fmt.Scan(&Id)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit Buku", "Stok Tersedia")
        for i := 0; i < n; i++ {
            if A[i].bookId == Id {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Pencarian Anda tidak ditemukan")
        }
    } else if lang == 2 {
        fmt.Print("Enter Book Id that you want to find: ")
        fmt.Scan(&Id)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Available Stock")
        for i := 0; i < n; i++ {
            if A[i].bookId == Id {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Your search is not found")
        }
    } else if lang == 3 {
        fmt.Print("أدخل معرف الكتاب الذي تريد العثور عليه: ")
        fmt.Scan(&Id)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة النشر", "المخزون المتاح")
        for i := 0; i < n; i++ {
            if A[i].bookId == Id {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("لم يتم العثور على بحثك")
        }
    } else if lang == 4 {
        fmt.Print("输入您要查找的书籍编号：")
        fmt.Scan(&Id)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "可用库存")
        for i := 0; i < n; i++ {
            if A[i].bookId == Id {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("未找到您的搜索")
        }
    }
}

func findBookTitle(A *tabInt, n int, lang int) {
	//Sequential Search
    var title string
    var found bool = false

    if lang == 1 {
        fmt.Print("Masukkan judul buku yang ingin Anda temukan: ")
        fmt.Scan(&title)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit Buku", "Stok Tersedia")
        for i := 0; i < n; i++ {
            if A[i].title == title {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Pencarian Anda tidak ditemukan")
        }
    } else if lang == 2 {
        fmt.Print("Enter Book title that you want to find: ")
        fmt.Scan(&title)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Available Stock")
        for i := 0; i < n; i++ {
            if A[i].title == title {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Your search is not found")
        }
    } else if lang == 3 {
        fmt.Print("أدخل عنوان الكتاب الذي تريد العثور عليه: ")
        fmt.Scan(&title)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة النشر", "المخزون المتاح")
        for i := 0; i < n; i++ {
            if A[i].title == title {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("لم يتم العثور على بحثك")
        }
    } else if lang == 4 {
        fmt.Print("输入您要查找的书籍标题：")
        fmt.Scan(&title)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "可用库存")
        for i := 0; i < n; i++ {
            if A[i].title == title {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("未找到您的搜索")
        }
    }
}

func findBookAuthor(A *tabInt, n int, lang int) { 
	//Sequential Search
    var author string
    var found bool = false

    if lang == 1 {
        fmt.Print("Masukkan penulis buku yang ingin Anda temukan: ")
        fmt.Scan(&author)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit Buku", "Stok Tersedia")
        for i := 0; i < n; i++ {
            if A[i].author == author {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Pencarian Anda tidak ditemukan")
        }
    } else if lang == 2 {
        fmt.Print("Enter Book author that you want to find: ")
        fmt.Scan(&author)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Available Stock")
        for i := 0; i < n; i++ {
            if A[i].author == author {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Your search is not found")
        }
    } else if lang == 3 {
        fmt.Print("أدخل مؤلف الكتاب الذي تريد العثور عليه: ")
        fmt.Scan(&author)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة النشر", "المخزون المتاح")
        for i := 0; i < n; i++ {
            if A[i].author == author {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("لم يتم العثور على بحثك")
        }
    } else if lang == 4 {
        fmt.Print("输入您要查找的书籍作者：")
        fmt.Scan(&author)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "可用库存")
        for i := 0; i < n; i++ {
            if A[i].author == author {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("未找到您的搜索")
        }
    }
}

func findBookGenre(A *tabInt, n int, lang int) { 
	//Sequential Search
    var genre string
    var found bool = false 

    if lang == 1 {
        fmt.Print("Masukkan genre buku yang ingin Anda temukan: ")
        fmt.Scan(&genre)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit Buku", "Stok Tersedia")
        for i := 0; i < n; i++ {
            if A[i].genre == genre {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Pencarian Anda tidak ditemukan")
        }
    } else if lang == 2 {
        fmt.Print("Enter Book genre that you want to find: ")
        fmt.Scan(&genre)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Available Stock")
        for i := 0; i < n; i++ {
            if A[i].genre == genre {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Your search is not found")
        }
    } else if lang == 3 {
        fmt.Print("أدخل نوع الكتاب الذي تريد البحث عنه: ")
        fmt.Scan(&genre)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة النشر", "المخزون المتاح")
        for i := 0; i < n; i++ {
            if A[i].genre == genre {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("لم يتم العثور على بحثك")
        }
    } else if lang == 4 {
        fmt.Print("输入您要查找的书籍类别: ")
        fmt.Scan(&genre)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "可用库存")
        for i := 0; i < n; i++ {
            if A[i].genre == genre {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("未找到您的搜索")
        }
    }
}

func findBookYear(A *tabInt, n int, lang int) {
	//Sequential Search
    var year int
    var found bool = false

    if lang == 1 {
        fmt.Print("Masukkan tahun terbit buku yang ingin Anda temukan: ")
        fmt.Scan(&year)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit Buku", "Stok Tersedia" )
        for i := 0; i < n; i++ {
            if A[i].pubYear == year {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Pencarian Anda tidak ditemukan")
        }
    } else if lang == 2 {
        fmt.Print("Enter Book year that you want to find: ")
        fmt.Scan(&year)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Available Stock")
        for i := 0; i < n; i++ {
            if A[i].pubYear == year {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("Your search is not found")
        }
    } else if lang == 3 {
        fmt.Print("أدخل سنة نشر الكتاب التي تريد العثور عليها: ")
        fmt.Scan(&year)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة النشر", "المخزون المتاح")
        for i := 0; i < n; i++ {
            if A[i].pubYear == year {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if !found {
            fmt.Println("لم يتم العثور على بحثك")
        }
    } else if lang == 4 {
        fmt.Print("输入您要查找的书籍出版年份：")
        fmt.Scan(&year)
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "可用库存")
        for i := 0; i < n; i++ {
            if A[i].pubYear == year {
                found = true
                fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d\n", A[i].bookId, A[i].title, A[i].author, A[i].genre, A[i].pubYear, A[i].stockAvailable)
            }
        }
        if (!found) {
            fmt.Println("未找到您的搜索")
        }
    }
}

func addLoan(A *tabInt, B *tabPeminjam, n int, nPeminjam *int, lang int) {
	var Id, peminjamId, namaPeminjam, date1 string
	if lang == 1 {
		fmt.Print("Masukkan Id Buku yang ingin Anda pinjam: ")
		fmt.Scan(&Id)
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Masukkan kata sandi Anda: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Masukkan tanggal peminjaman (dd-mm-yyyy): ")
		fmt.Scan(&date1)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				if A[i].stockAvailable > 0 {
					A[i].stockAvailable--
					B[*nPeminjam].peminjamId = peminjamId
					B[*nPeminjam].bookId = Id
					A[i].loanCount++
					B[*nPeminjam].namaPeminjam = namaPeminjam
					B[*nPeminjam].tanggalPinjam = date1
					*nPeminjam++
					fmt.Println("Peminjaman berhasil.")
					return
				} else {
					fmt.Println("Peminjaman gagal. Stok tidak tersedia.")
					return
				}
			}
		}
		fmt.Println("Id Buku tidak ditemukan.")
	} else if lang == 2 {
		fmt.Print("Enter Book Id that you want to loan: ")
		fmt.Scan(&Id)
		fmt.Print("Enter your name: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Enter your password: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Enter the date you loaned (dd-mm-yyyy): ")
		fmt.Scan(&date1)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				if A[i].stockAvailable > 0 {
					A[i].stockAvailable--
					B[*nPeminjam].peminjamId = peminjamId
					B[*nPeminjam].bookId = Id
					A[i].loanCount++
					B[*nPeminjam].namaPeminjam = namaPeminjam
					B[*nPeminjam].tanggalPinjam = date1
					*nPeminjam++
					fmt.Println("The loan has been successful.")
					return
				} else {
					fmt.Println("The loan has been failed. No stock available.")
					return
				}
			}
		}
		fmt.Println("Book Id is not found.")
	} else if lang == 3 {
		fmt.Print("أدخل معرف الكتاب الذي تريد استعارته: ")
		fmt.Scan(&Id)
		fmt.Print("أدخل اسمك: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("أدخل كلمة المرور الخاصة بك: ")
		fmt.Scan(&peminjamId)
		fmt.Print("أدخل تاريخ الاستعارة (dd-mm-yyyy): ")
		fmt.Scan(&date1)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				if A[i].stockAvailable > 0 {
					A[i].stockAvailable--
					B[*nPeminjam].peminjamId = peminjamId
					B[*nPeminjam].bookId = Id
					A[i].loanCount++
					B[*nPeminjam].namaPeminjam = namaPeminjam
					B[*nPeminjam].tanggalPinjam = date1
					*nPeminjam++
					fmt.Println("تمت الاستعارة بنجاح.")
					return
				} else {
					fmt.Println("فشلت الاستعارة. لا يوجد مخزون متاح.")
					return
				}
			}
		}
		fmt.Println("لم يتم العثور على معرف الكتاب.")
	} else if lang == 4 {
		fmt.Print("输入您要借的书籍编号：")
		fmt.Scan(&Id)
		fmt.Print("输入您的姓名：")
		fmt.Scan(&namaPeminjam)
		fmt.Print("输入您的密码：")
		fmt.Scan(&peminjamId)
		fmt.Print("输入借书日期 (dd-mm-yyyy)：")
		fmt.Scan(&date1)
		for i := 0; i < n; i++ {
			if A[i].bookId == Id {
				if A[i].stockAvailable > 0 {
					A[i].stockAvailable--
					B[*nPeminjam].peminjamId = peminjamId
					B[*nPeminjam].bookId = Id
					A[i].loanCount++
					B[*nPeminjam].namaPeminjam = namaPeminjam
					B[*nPeminjam].tanggalPinjam = date1
					*nPeminjam++
					fmt.Println("借书成功。")
					return
				} else {
					fmt.Println("借书失败。没有库存。")
					return
				}
			}
		}
		fmt.Println("未找到书籍编号。")
	}
}

func editLoan(B *tabPeminjam, nPeminjam int, lang int) {
	var peminjamId string
	var namaPeminjam1, namaPeminjam2 string
	if lang == 1 {
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&namaPeminjam1)
		fmt.Print("Masukkan kata sandi Anda: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Masukkan nama baru: ")
		fmt.Scan(&namaPeminjam2)
		for i := 0; i < nPeminjam; i++ {
			if B[i].peminjamId == peminjamId && B[i].namaPeminjam == namaPeminjam1 {
				B[i].namaPeminjam = namaPeminjam2
				fmt.Println("Nama Anda telah berhasil diubah.")
				return
			}
		}
	} else if lang == 2 {
		fmt.Print("Enter your name: ")
		fmt.Scan(&namaPeminjam1)
		fmt.Print("Enter your password: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Enter new name: ")
		fmt.Scan(&namaPeminjam2)
		for i := 0; i < nPeminjam; i++ {
			if B[i].peminjamId == peminjamId && B[i].namaPeminjam == namaPeminjam1 {
				B[i].namaPeminjam = namaPeminjam2
				fmt.Println("Your name has been succesfully changed.")
				return
			}
		}
	} else if lang == 3 {
		fmt.Print("أدخل اسمك: ")
		fmt.Scan(&namaPeminjam1)
		fmt.Print("أدخل كلمة المرور الخاصة بك: ")
		fmt.Scan(&peminjamId)
		fmt.Print("أدخل الاسم الجديد: ")
		fmt.Scan(&namaPeminjam2)
		for i := 0; i < nPeminjam; i++ {
			if B[i].peminjamId == peminjamId && B[i].namaPeminjam == namaPeminjam1 {
				B[i].namaPeminjam = namaPeminjam2
				fmt.Println("تم تغيير اسمك بنجاح.")
				return
			}
		}
	} else if lang == 4 {
		fmt.Print("输入您的姓名：")
		fmt.Scan(&namaPeminjam1)
		fmt.Print("输入您的密码：")
		fmt.Scan(&peminjamId)
		fmt.Print("输入新姓名：")
		fmt.Scan(&namaPeminjam2)
		for i := 0; i < nPeminjam; i++ {
			if B[i].peminjamId == peminjamId && B[i].namaPeminjam == namaPeminjam1 {
				B[i].namaPeminjam = namaPeminjam2
				fmt.Println("您的姓名已成功更改。")
				return
			}
		}
	}
}

func deleteLoan(A *tabInt, B *tabPeminjam, n *int, nPeminjam *int, lang int) {
	var peminjamId, namaPeminjam, bookId string
	switch lang {
	case 1:
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Masukkan kata sandi anda: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Masukkan Id buku yang ingin dihapus: ")
		fmt.Scan(&bookId)
	case 2:
		fmt.Print("Enter your name: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Enter your password: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Enter the book Id you want to delete: ")
		fmt.Scan(&bookId)
	case 3:
		fmt.Print("أدخل اسمك: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("أدخل كلمة المرور: ")
		fmt.Scan(&peminjamId)
		fmt.Print("أدخل معرف الكتاب الذي تريد حذفه: ")
		fmt.Scan(&bookId)
	case 4:
		fmt.Print("输入您的姓名：")
		fmt.Scan(&namaPeminjam)
		fmt.Print("输入您的密码：")
		fmt.Scan(&peminjamId)
		fmt.Print("输入您要删除的书籍编号：")
		fmt.Scan(&bookId)
	}

	found := false
	for i := 0; i < *n; i++ {
		if (*A)[i].bookId == bookId {
			found = true
			(*A)[i].loanCount--
			(*A)[i].stockAvailable++
			break
		}
	}

	if found {
		for i := 0; i < *nPeminjam; i++ {
			if (*B)[i].bookId == bookId && (*B)[i].peminjamId == peminjamId {
				for j := i; j < *nPeminjam-1; j++ {
					(*B)[j] = (*B)[j+1]
				}
				*nPeminjam--
				break
			}
		}
		switch lang {
		case 1:
			fmt.Println("Data pinjaman telah berhasil dihapus.")
		case 2:
			fmt.Println("Loan data has been successfully deleted.")
		case 3:
			fmt.Println("تم حذف بيانات الاستعارة بنجاح.")
		case 4:
			fmt.Println("借书数据已成功删除。")
		}
	} else {
		switch lang {
		case 1:
			fmt.Println("Data pinjaman tidak berhasil dihapus.")
		case 2:
			fmt.Println("Loan data was not successfully deleted.")
		case 3:
			fmt.Println("لم يتم حذف بيانات الاستعارة بنجاح.")
		case 4:
			fmt.Println("借书数据未成功删除。")
		}
	}
}

func returnBook(A *tabInt, B *tabPeminjam, n int, nPeminjam *int, lang int) {
	var peminjamId, namaPeminjam, date2, bookId string
	var messageNotFound string

	switch lang {
	case 1:
		messageNotFound = "Pencarian Anda tidak ditemukan."
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Masukkan kata sandi Anda: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Masukkan Id buku yang ingin dikembalikan: ")
		fmt.Scan(&bookId)
		fmt.Print("Masukkan tanggal pengembalian (dd-mm-yyyy): ")
		fmt.Scan(&date2)
	case 2:
		messageNotFound = "Your search is not found."
		fmt.Print("Enter your name: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("Enter your password: ")
		fmt.Scan(&peminjamId)
		fmt.Print("Enter the book ID to be returned: ")
		fmt.Scan(&bookId)
		fmt.Print("Enter the date you returned (dd-mm-yyyy): ")
		fmt.Scan(&date2)
	case 3:
		messageNotFound = "بحثك غير موجود."
		fmt.Print("أدخل اسمك: ")
		fmt.Scan(&namaPeminjam)
		fmt.Print("أدخل كلمة المرور الخاصة بك: ")
		fmt.Scan(&peminjamId)
		fmt.Print("أدخل معرف الكتاب الذي تريد إرجاعه: ")
		fmt.Scan(&bookId)
		fmt.Print("أدخل تاريخ الإرجاع (dd-mm-yyyy): ")
		fmt.Scan(&date2)
	case 4:
		messageNotFound = "未找到您的搜索."
		fmt.Print("输入您的姓名：")
		fmt.Scan(&namaPeminjam)
		fmt.Print("输入您的密码：")
		fmt.Scan(&peminjamId)
		fmt.Print("请输入要归还的书籍 ID: ")
		fmt.Scan(&bookId)
		fmt.Print("输入您的归还日期 (dd-mm-yyyy):")
		fmt.Scan(&date2)
	default:
		fmt.Println("Invalid language selection.")
		return
	}

	var foundIndex = -1
	for i := 0; i < *nPeminjam; i++ {
		if B[i].peminjamId == peminjamId && B[i].namaPeminjam == namaPeminjam && B[i].bookId == bookId {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		fmt.Println(messageNotFound)
		return
	}

	tanggalPinjam := B[foundIndex].tanggalPinjam
	fine := calculateFine(tanggalPinjam, date2)
	if fine > 0 {
		switch lang {
		case 1:
			fmt.Printf("Denda yang harus dibayar: $%d\n", fine)
			bayarIndo()
		case 2:
			fmt.Printf("Fine must be paid: $%d\n", fine)
			bayarEng()
		case 3:
			fmt.Printf("يجب دفع الغرامة: $%d\n", fine)
			bayarArab()
		case 4:
			fmt.Printf("应付罚款：$%d\n", fine)
			bayarChinese()
		}
	} else {
		switch lang {
		case 1:
			fmt.Println("Tidak ada denda yang harus dibayar.")
		case 2:
			fmt.Println("No fine to be paid.")
		case 3:
			fmt.Println("لا يوجد غرامة يجب دفعها.")
		case 4:
			fmt.Println("没有罚款需要支付。")
		}
	}

	for j := 0; j < n; j++ {
		if A[j].bookId == bookId {
			A[j].stockAvailable++
			A[j].loanCount--
			break
		}
	}

	for k := foundIndex; k < *nPeminjam-1; k++ {
		B[k] = B[k+1]
	}
	*nPeminjam--
}

func bayarIndo() {
	fmt.Println("Pilih metode pembayaran:")
	fmt.Println("1. MasterCard")
	fmt.Println("2. VISA")
	fmt.Println("3. American Express")
	fmt.Println("4. Discover")
	fmt.Println("5. JCB (Japan Credit Bureau)")
	fmt.Println("6. UnionPay")
	fmt.Println("7. Diners Club")
	var metode int
	fmt.Print("Pilih nomor: ")
	fmt.Scan(&metode)
	bayarIndo2(metode)
}

func bayarEng() {
	fmt.Println("Choose payment method:")
	fmt.Println("1. MasterCard")
	fmt.Println("2. VISA")
	fmt.Println("3. American Express")
	fmt.Println("4. Discover")
	fmt.Println("5. JCB (Japan Credit Bureau)")
	fmt.Println("6. UnionPay")
	fmt.Println("7. Diners Club")
	var metode int
	fmt.Print("Select number: ")
	fmt.Scan(&metode)
	bayarEng2(metode)
}

func bayarArab() {
	fmt.Println("اختر طريقة الدفع:")
	fmt.Println("١. ماستر كارد")
	fmt.Println("٢. فيزا")
	fmt.Println("٣. أمريكان إكسبريس")
	fmt.Println("٤. ديسكوفر")
	fmt.Println("٥. جي سي بي (مكتب الائتمان الياباني)")
	fmt.Println("٦. يونيون باي")
	fmt.Println("٧. دينرز كلوب")
	var metode int
	fmt.Print("اختر رقمًا: ")
	fmt.Scan(&metode)
	bayarArab2(metode)
}

func bayarChinese() {
	fmt.Println("选择支付方式:")
	fmt.Println("1. 万事达卡")
	fmt.Println("2. VISA")
	fmt.Println("3. 美国运通 (American Express)")
	fmt.Println("4. 发现卡 (Discover)")
	fmt.Println("5. 日本信用局 (JCB)")
	fmt.Println("6. 银联 (UnionPay)")
	fmt.Println("7. 大来卡 (Diners Club)")
	var metode int
	fmt.Print("请选择数字：")
	fmt.Scan(&metode)
	bayarChinese2(metode)
}

func bayarIndo2(metode int) {
	var nomorRekening, masaBerlaku, cvc string
	fmt.Print("Masukkan nomor rekening anda: ")
	fmt.Scan(&nomorRekening)
	fmt.Print("Masukkan masa berlaku kartu (MM/YY): ")
	fmt.Scan(&masaBerlaku)
	fmt.Print("Masukkan CVC/CVV: ")
	fmt.Scan(&cvc)
	fmt.Println("Mohon menunggu...")
	time.Sleep(5 * time.Second)
	fmt.Println("Pembayaran berhasil")
}

func bayarEng2(metode int) {
	var cardNumber, expiryDate, cvc string
	fmt.Print("Enter your card number: ")
	fmt.Scan(&cardNumber)
	fmt.Print("Enter card expiry date (MM/YY): ")
	fmt.Scan(&expiryDate)
	fmt.Print("Enter CVC/CVV: ")
	fmt.Scan(&cvc)
	fmt.Println("Please wait...")
	time.Sleep(5 * time.Second)
	fmt.Println("Payment successful")
}

func bayarArab2(metode int) {
	var cardNumber, expiryDate, cvc string
	fmt.Print("أدخل رقم البطاقة: ")
	fmt.Scan(&cardNumber)
	fmt.Print("أدخل تاريخ انتهاء البطاقة (MM/YY): ")
	fmt.Scan(&expiryDate)
	fmt.Print("أدخل CVC/CVV: ")
	fmt.Scan(&cvc)
	fmt.Println("أرجو الانتظار...")
	time.Sleep(5 * time.Second)
	fmt.Println("تم الدفع بنجاح")
}

func bayarChinese2(metode int) {
	var cardNumber, expiryDate, cvc string
	fmt.Print("输入卡号: ")
	fmt.Scan(&cardNumber)
	fmt.Print("输入卡的有效期 (MM/YY): ")
	fmt.Scan(&expiryDate)
	fmt.Print("输入 CVC/CVV: ")
	fmt.Scan(&cvc)
	fmt.Println("请稍等...")
	time.Sleep(5 * time.Second)
	fmt.Println("支付成功")
}

func calculateFine(tanggalPinjam, date2 string) int {
	data := diffDays(tanggalPinjam, date2)
	if data > 14 {
		return (data - 14) * 2
	}
	return 0
}

func diffDays(date1Str, date2Str string) int {
	layout := "02-01-2006"
	date1, _ := time.Parse(layout, date1Str)
	date2, _ := time.Parse(layout, date2Str)
	duration := date2.Sub(date1)
	return int(duration.Hours() / 24)
}

func favorite(A *tabInt, n int, lang int) {
	// Selection sort untuk mengurutkan berdasarkan loanCount dalam urutan menurun
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if (*A)[j].loanCount > (*A)[maxIdx].loanCount {
				maxIdx = j
			}
		}
		(*A)[i], (*A)[maxIdx] = (*A)[maxIdx], (*A)[i]
	}

	limit := 5
	if n < 5 {
		limit = n
	}

	if lang == 1 {
		fmt.Println("Buku favorit teratas: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit", "Stok Tersedia", "Jumlah Peminjaman")
		for i := 0; i < limit; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		fmt.Println("Ringkasan:")
		mostFrequentBorrower := binarySearch(A, n)
		if mostFrequentBorrower != "" {
			fmt.Printf("Nama penulis buku yang paling laris: %s\n", mostFrequentBorrower)
		} else {
			fmt.Println("Tidak ada nama yang meminjam lebih dari 1 kali.")
		}
	} else if lang == 2 {
		fmt.Println("Top favorite books: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Stock Available", "Loan Count")
		for i := 0; i < limit; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		fmt.Println("Summary:")
		mostFrequentBorrower := binarySearch(A, n)
		if mostFrequentBorrower != "" {
			fmt.Printf("The name of the author of the best-selling book: %s\n", mostFrequentBorrower)
		} else {
			fmt.Println("No name has borrowed more than once.")
		}
	} else if lang == 3 {
		fmt.Println("أعلى الكتب المفضلة: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة نشر الكتاب", "النسخ المتاحة", "عدد الإعارات")
		for i := 0; i < limit; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		fmt.Println("ملخص:")
		mostFrequentBorrower := binarySearch(A, n)
		if mostFrequentBorrower != "" {
			fmt.Printf("اسم مؤلف الكتاب الأكثر مبيعا: %s\n", mostFrequentBorrower)
		} else {
			fmt.Println("لا يوجد اسم استعار أكثر من مرة واحدة.")
		}
	} else if lang == 4 {
		fmt.Println("热门收藏图书： ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "库存", "借阅次数")
		for i := 0; i < limit; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		fmt.Println("总结:")
		mostFrequentBorrower := binarySearch(A, n)
		if mostFrequentBorrower != "" {
			fmt.Printf("畅销书作者的名字：%s\n", mostFrequentBorrower)
		} else {
			fmt.Println("没有名字借阅超过一次。")
		}
	} else {
		fmt.Println("Language not supported.")
	}
}

func displayLibrary(A *tabInt, n int, lang int) {
	// Insertion sort untuk mengurutkan berdasarkan tahun terbit buku dalam urutan menurun
	for i := 1; i < n; i++ {
		key := (*A)[i]
		j := i - 1
		for j >= 0 && (*A)[j].title > key.title {
			(*A)[j+1] = (*A)[j]
			j = j - 1
		}
		(*A)[j+1] = key
	}

	if lang == 1 {
		fmt.Println("Daftar buku di perpustakaan: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit", "Stok Tersedia", "Jumlah Peminjaman")
		for i := 0; i < n; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		var newestYear int
   		for i := 0; i < n; i++ {
        if A[i].pubYear > newestYear {
            newestYear = A[i].pubYear
        	}
    	}
    	fmt.Printf("Ringkasan:\nTahun penerbitan buku terbaru yang ada dalam database perpustakaan: %d\n", newestYear)
	} else if lang == 2 {
		fmt.Println("List of books in the library: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Stock Available", "Loan Count")
		for i := 0; i < n; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		var newestYear int
   		for i := 0; i < n; i++ {
        if A[i].pubYear > newestYear {
            newestYear = A[i].pubYear
        	}
    	}
    	fmt.Printf("Ringkasan:\nTahun penerbitan buku terbaru yang ada dalam database perpustakaan: %d\n", newestYear)
	} else if lang == 3 {
		fmt.Println("قائمة الكتب في المكتبة: ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة نشر الكتاب", "النسخ المتاحة", "عدد الإعارات")
		for i := 0; i < n; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		var newestYear int
   		for i := 0; i < n; i++ {
        if A[i].pubYear > newestYear {
            newestYear = A[i].pubYear
        	}
    	}
    	fmt.Printf("Ringkasan:\nTahun penerbitan buku terbaru yang ada dalam database perpustakaan: %d\n", newestYear)
	} else if lang == 4 {
		fmt.Println("图书馆书目清单： ")
		fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "库存", "借阅次数")
		for i := 0; i < n; i++ {
			fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20d %-20d\n",
				(*A)[i].bookId, (*A)[i].title, (*A)[i].author, (*A)[i].genre, (*A)[i].pubYear, (*A)[i].stockAvailable, (*A)[i].loanCount)
		}
		var newestYear int
   		for i := 0; i < n; i++ {
        if A[i].pubYear > newestYear {
            newestYear = A[i].pubYear
        	}
    	}
    	fmt.Printf("摘要：\n图书馆数据库中最新出版的图书年份：%d\n", newestYear)
	} else {
		fmt.Println("Language not supported.")
	}
	
	
}

func displayLoaner(A *tabInt, B *tabPeminjam, n int, nPeminjam int, lang int) {
    if lang == 1 {
        fmt.Println("Daftar buku yang sedang dipinjam: ")
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "ID Buku", "Judul Buku", "Penulis Buku", "Genre Buku", "Tahun Terbit", "Nama Peminjam", "Tanggal Peminjaman")
        if nPeminjam > 0 {
            for i := 0; i < nPeminjam; i++ {
                for j := 0; j < n; j++ {
                    if B[i].bookId == A[j].bookId {
                        fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20s %-20s\n", A[j].bookId, A[j].title, A[j].author, A[j].genre, A[j].pubYear, B[i].namaPeminjam, B[i].tanggalPinjam)
                    }
                }
            }
        } else {
            fmt.Println("Maaf, tidak ada buku yang sedang dipinjam")
        }
    } else if lang == 2 {
        fmt.Println("List of books being borrowed: ")
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "Book ID", "Book Title", "Book Author", "Book Genre", "Publication Year", "Borrower's Name", "Date of Borrowing")
        if nPeminjam > 0 {
            for i := 0; i < nPeminjam; i++ {
                for j := 0; j < n; j++ {
                    if B[i].bookId == A[j].bookId {
                        fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20s %-20s\n", A[j].bookId, A[j].title, A[j].author, A[j].genre, A[j].pubYear, B[i].namaPeminjam, B[i].tanggalPinjam)
                    }
                }
            }
        } else {
            fmt.Println("Sorry, no books are being borrowed")
        }
    } else if lang == 3 {
        fmt.Println("قائمة الكتب المعارة:")
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "معرف الكتاب", "عنوان الكتاب", "مؤلف الكتاب", "نوع الكتاب", "سنة نشر الكتاب", "اسم المستعير", "تاريخ الاقتراض")
        if nPeminjam > 0 {
            for i := 0; i < nPeminjam; i++ {
                for j := 0; j < n; j++ {
                    if B[i].bookId == A[j].bookId {
                        fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20s %-20s\n", A[j].bookId, A[j].title, A[j].author, A[j].genre, A[j].pubYear, B[i].namaPeminjam, B[i].tanggalPinjam)
                    }
                }
            }
        } else {
            fmt.Println("آسف، لا توجد كتب تُعار حاليًا")
        }
    } else if lang == 4 {
        fmt.Println("借阅中的图书清单：")
        fmt.Printf("%-20s %-20s %-20s %-20s %-20s %-20s %-20s\n", "书籍编号", "书名", "作者", "书籍类型", "出版年份", "借阅人姓名", "借阅日期")
        if nPeminjam > 0 {
            for i := 0; i < nPeminjam; i++ {
                for j := 0; j < n; j++ {
                    if B[i].bookId == A[j].bookId {
                        fmt.Printf("%-20s %-20s %-20s %-20s %-20d %-20s %-20s\n", A[j].bookId, A[j].title, A[j].author, A[j].genre, A[j].pubYear, B[i].namaPeminjam, B[i].tanggalPinjam)
                    }
                }
            }
        } else {
            fmt.Println("抱歉，当前没有借阅的书籍")
        }
    }
}

func binarySearch(A *tabInt, n int) string {
	//Binary Search
	left, right := 0, n-1
	for left <= right {
		mid := left + (right-left)/2
		if A[mid].loanCount > 1 {
			return A[mid].author
		}
		if A[mid].loanCount < 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ""
}