# -*- coding: utf-8 -*-

import json
from models import Book, Library
import os


def read_json(filepath: str) -> Library:
    """Чтение данных из JSON-файла"""
    if not os.path.exists(filepath):
        raise FileNotFoundError(f"File not found: {filepath}")
    with open(filepath, "r", encoding="utf-8") as f:
        data = json.load(f)
        books = data.get("books", [])
        library = Library()
        for book in books:
            library.add_book(
                Book(book["title"], book["author"], book["year"], book["status"])
            )
        return library


def write_json(filename: str, library: Library) -> None:
    """Запись данных в JSON-файл"""
    with open(filename, "w+", encoding="utf-8") as f:
        books = [book.__dict__ for book in library.books]
        data = {"books": books}
        json.dump(data, f, ensure_ascii=False)
