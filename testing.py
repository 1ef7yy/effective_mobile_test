import unittest
from models import Book, Library
import os
from jsonParser import read_json, write_json


class TestLibrary(unittest.TestCase):
    def test_add_book(self) -> None:
        """Тест функции добавления книги в библиотеку"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        self.assertIn(book, library.books)

    def test_get_book_by_id(self) -> None:
        """Тест функции поиска книги по id"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        result = library.get_book_by_id(book.id)
        self.assertEqual(book, result)

    def test_delete_book(self) -> None:
        """Тест функции удаления книги из библиотеки"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        library.delete_book(book.id)
        self.assertNotIn(book, library.books)

    def test_change_book_status(self) -> None:
        """Тест функции изменения статуса книги"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        library.change_book_status(book.id, "выдана")
        self.assertEqual(book.status, "выдана")

    def test_save_to_json(self) -> None:
        """Тест функции сохранения библиотеки в JSON-файл"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        write_json("test1_library.json", library)
        self.assertTrue(os.path.exists("test1_library.json"))

    def test_load_from_json(self) -> None:
        """Тест функции загрузки библиотеки из JSON-файла"""
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        write_json("test2_library.json", library)
        loaded_library = read_json("test2_library.json")
        self.assertEqual(library, loaded_library)


if __name__ == "__main__":
    unittest.main()
