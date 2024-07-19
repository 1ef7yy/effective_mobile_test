# -*- coding: utf-8 -*-
# Само консольное приложение
from models import Book, Library


def main() -> None:
    """Главное меню"""
    user_library = Library()
    print(" Что вы хотите сделать?")
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


def add_book_prompt() -> None: ...


def delete_book_prompt() -> None: ...


def get_book_by_id_prompt() -> None: ...


def print_books_prompt() -> None: ...


def change_book_status_prompt() -> None: ...


def load_from_json_prompt() -> None: ...


def save_to_json_prompt() -> None: ...


if __name__ == "__main__":
    print("\n Система управления библиотекой\n\n")

    main()
