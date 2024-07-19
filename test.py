import unittest
from models import Book, Library


class TestLibrary(unittest.TestCase):
    def test_add_book(self) -> None:
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        self.assertIn(book, library.books)

    def test_get_book_by_id(self) -> None:
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        result = library.get_book_by_id(book.id)
        self.assertEqual(book, result)

    def test_delete_book(self) -> None:
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        library.delete_book(book.id)
        self.assertNotIn(book, library.books)

    def test_change_book_status(self) -> None:
        library = Library()
        book = Book("title", "author", 2000, "в наличии")
        library.add_book(book)
        library.change_book_status(book.id, "выдана")
        self.assertEqual(book.status, "выдана")


if __name__ == "__main__":
    unittest.main()
