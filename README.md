# serverTesting

```
этот проект представляет собой веб-приложение, для регистрации и авторизации пользователей с возможностью общения друг с другом и построенное с использованием принципов монолитной архитектуры (Monolithic)
```

## Архитектура проекта
```
Проект следует принципам чистой архитектуры, которая разделяет код на несколько слоев:
```
### Слои архитектуры

```
1. Domain - ядро. Содержит сущности, интерфейсы репозиториев и бизнес-правила
2. Application - сценарии использования. Оркестрация процессов
3. Infrastructure - инфраструктура. Содержит репозитории, внешние сервисы и конфигурации(подключение к базе данных)
4. Configs - конфигурация. Содержит env-переменные, настройки. Подключение к БД, API-ключи
5. API  - интерфейс взаимодействия с приложением извне. Маршрутизация, Аутентификация/авторизация, документацию, контроллеры(http-обработчики), роутеры
```
## Основные функции
- Регистрация пользователей
- Авторизация пользователей
- Аутентификация пользователей
## Структура проекта
```
myApp/
├── application/         # сценари использования (Use cases)
├── domain/              # Доменный слой(бизнес-сущности, интерфейсы репозиториев и сервисов, доменные исключения)
├── database/            # БД
├── cmd/                 # исполняемые файлы
├──configs/              # конфигурация
└── api/                 # все для API
    ├── handlers/        # HTTP-обработчики
    └── routers/         # Роутеры

migrations/              # Миграции баз данных
tests/                   # Тесты

```