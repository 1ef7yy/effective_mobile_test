import datetime


class Book:
    def __init__(self, title: str, author: str, year: int, status: str):
        self.title = title
        self.author = author
        self.year = year
        self.status = status


class Library:
    def __init__(self):
        self.books: list[Book] = []
        self.curr_id = 1

    def add_book(self, book: Book) -> None:
        """Функция добавления книги в библиотеку"""
        if book.status not in ("в наличии", "выдана"):
            raise ValueError("Invalid status")

        if book in self.books:
            raise ValueError("Book already exists")

        if book.year > datetime.datetime.now().year:
            raise ValueError("Book is not yet published")
        # присвоение id книге
        book.id = self.curr_id
        self.books.append(book)
        self.curr_id += 1

    def delete_book(self, book_id: int) -> None:
        """Функция удаления книги из библиотеки"""
        for book in self.books:
            if book.id == book_id:
                self.books.remove(book)
                return

        raise ValueError("Book not found")

    def get_book_by_id(self, book_id: int) -> Book:
        """Функция поиска книги по id"""
        for book in self.books:
            if book.id == book_id:
                return book

        raise ValueError("Book not found")

    def change_book_status(self, book_id: int, new_status: str) -> None:
        """Функция изменения статуса книги"""
        if new_status not in ("в наличии", "выдана"):
            raise ValueError("Invalid status")

        for book in self.books:
            if book.id == book_id:
                book.status = new_status
                return

        raise ValueError("Book not found")

    def print_books(self) -> None:
        """Функция вывода списка книг"""
        for book in self.books:
            print(f"{book.id}: {book.title}, {book.author}, {book.year}, {book.status}")
