# -*- coding: utf-8 -*-
# Само консольное приложение
from models import Book, Library
from jsonParser import read_json, write_json

user_library = Library()


def main() -> None:
    """Главное меню"""

    print("\n\n Что вы хотите сделать?")
    print("""
 1. Добавить книгу
 2. Удалить книгу
 3. Найти книгу по id
 4. Вывести список книг
 5. Изменить статус книги
 6. Загрузить данные из JSON
 7. Выгрузить данные в JSON
""")
    prompt = input(" > ")

    if str(int(prompt)) != prompt:
        print(" Введите целое число")
        main()

    if int(prompt) not in range(1, 8):
        print(" Введите число от 1 до 7")
        main()

    match int(prompt):
        case 1:
            add_book_prompt()
        case 2:
            delete_book_prompt()
        case 3:
            get_book_by_id_prompt()
        case 4:
            print_books_prompt()
        case 5:
            change_book_status_prompt()
        case 6:
            load_from_json_prompt()
        case 7:
            save_to_json_prompt()
        case _:
            print(" Произошла ошибка, попробуйте еще раз")
            main()


def add_book_prompt() -> None:
    print(" Добавить книгу?")
    print(" 1. Да\n 2. Нет, назад")
    add_prompt = input(" > ")
    if int(add_prompt) not in (1, 2):
        print(" Введите 1 или 2")
        add_book_prompt()
    if add_prompt == "2":
        main()

    try:
        title = input(" Название: ")
        author = input(" Автор: ")
        year = int(input(" Год: "))
        status = input(" Статус: ")
        book = Book(title, author, year, status)
        user_library.add_book(book)
        print("Книга успешно добавлена!")
        main()
    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        add_book_prompt()


def delete_book_prompt() -> None:
    if len(user_library.books) == 0:
        print(" Библиотека пуста")
        main()

    print(" Удалить книгу?")
    print(" 1. Да\n 2. Нет, назад")
    delete_prompt = input(" > ")
    if int(delete_prompt) not in (1, 2):
        print(" Введите 1 или 2")
        delete_book_prompt()
    if delete_prompt == "2":
        main()

    try:
        book_id = int(input(" ID книги: "))
        user_library.delete_book(book_id)
        print("Книга успешно удалена!")
        main()
    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        delete_book_prompt()


def get_book_by_id_prompt() -> None:
    if len(user_library.books) == 0:
        print(" Библиотека пуста")
        main()

    try:
        book_id = int(input(" ID книги: "))
        book = user_library.get_book_by_id(book_id)
        print(f"{book.id}: {book.title}, {book.author}, {book.year}, {book.status}")
        main()
    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        get_book_by_id_prompt()


def print_books_prompt() -> None:
    if len(user_library.books) == 0:
        print(" Библиотека пуста")
        main()

    user_library.print_books()
    main()


def change_book_status_prompt() -> None:
    if len(user_library.books) == 0:
        print(" Библиотека пуста")
        main()

    try:
        book_id = int(input(" ID книги: "))
        new_status = input(" Новый статус: ")
        user_library.change_book_status(book_id, new_status)
        print("Статус книги успешно изменен!")
        main()
    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        change_book_status_prompt()


def load_from_json_prompt() -> None:
    print(" Загрузить данные из JSON?")
    print(" 1. Да\n 2. Нет, назад")
    load_prompt = input(" > ")
    if int(load_prompt) not in (1, 2):
        print(" Введите 1 или 2")
        load_from_json_prompt()
    if load_prompt == "2":
        main()

    print("Введите путь к JSON-файлу:")
    path = input(" > ")

    try:
        global user_library
        user_library = read_json(path)
        print("Данные успешно загружены!")
        main()

    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        load_from_json_prompt()


def save_to_json_prompt() -> None:
    try:
        print("Введите путь к JSON-файлу:")
        path = input(" > ")

        write_json(path, user_library)

        print("Данные успешно сохранены!")
        main()
    except Exception as e:
        print(f" Ошибка: {repr(e)}\n\n")
        main()


if __name__ == "__main__":
    print("\n Система управления библиотекой\n\n")

    main()
