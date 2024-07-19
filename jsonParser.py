import json
from models import Book, Library
import os


def read_json(filepath: str) -> Library:
    if not os.path.exists(filepath):
        raise FileNotFoundError(f"File not found: {filepath}")
    with open(filepath, "r") as f:
        data = json.load(f)
        books = data.get("books", [])
        library = Library()
        for book in books:
            library.add_book(
                Book(book["title"], book["author"], book["year"], book["status"])
            )
        return library


def write_json(filename: str, library: Library) -> None:
    with open(filename, "w+") as f:
        books = [book.__dict__ for book in library.books]
        data = {"books": books}
        json.dump(data, f)
