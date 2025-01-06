# Currency API

Этот проект предоставляет API для получения данных о курсах валют. Он использует MySQL в качестве базы данных и контейнеры Docker для запуска приложения.

## Настройка переменных окружения

Создайте файл `.env` в корне проекта и добавьте следующие переменные:

```bash
MYSQL_ROOT_PASSWORD=<your_pass>
MYSQL_ROOT_USER=root
MYSQL_DATABASE=<your_db>
MYSQL_PORT=<port1>
MYSQL_HOST=localhost
SERVER_PORT=<port2>
```

## Запуск проекта с использованием Docker Compose

Для запуска проекта используйте команду:

```bash
docker-compose up --build
```

## Запросы API

- Получить все валюты
Метод: GET
Путь: /currencies
Описание: Получить все данные о курсах валют.

- Получить валюты по дате
Метод: GET
Путь: /currencies?date=YYYY-MM-DD
Описание: Получить данные о курсах валют для указанной даты
Параметры:
```bash
date (обязательный) — дата в формате YYYY-MM-DD.
```