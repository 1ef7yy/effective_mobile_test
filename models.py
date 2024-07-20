# -*- coding: utf-8 -*-

import datetime


class Book:
    def __init__(self, title: str, author: str, year: int, status: str):
        self.title = title
        self.author = author
        self.year = year
        self.status = status

    def __eq__(self, other):
        return [
            self.title,
            self.author,
            self.year,
            self.status,
        ] == [
            other.title,
            other.author,
            other.year,
            other.status,
        ]


class Library:
    def __init__(self):
        self.books: list[Book] = []
        self.curr_id = 1

    def __eq__(self, other):
        return self.books == other.books

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
        table = []
        for book in self.books:
            table.append([book.id, book.title, book.author, book.year, book.status])

        max_cols = [max(len(str(row[i])) for row in table) for i in range(5)]

        print(
            " " * (max_cols[0] + 2)
            + " " * (max_cols[1] + 2)
            + " " * (max_cols[2] + 2)
            + " " * (max_cols[3] + 2)
            + " " * (max_cols[4] + 2)
        )
        print(
            " ID "
            + " " * (max_cols[0] - 2)
            + " Название "
            + " " * (max_cols[1] - 6)
            + " Автор "
            + " " * (max_cols[2] - 4)
            + " Год "
            + " " * (max_cols[3] - 2)
            + " Статус "
        )

        for row in table:
            print(
                f" {row[0]:{max_cols[0]}} | {row[1]:{max_cols[1]}} | {row[2]:{max_cols[2]}} | {row[3]:{max_cols[3]}} | {row[4]:{max_cols[4]}} "
            )

    def get_books(self) -> list[Book]:
        return [book for book in self.books]
