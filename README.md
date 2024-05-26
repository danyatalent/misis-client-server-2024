# misis-client-server-2024

Группа БПМ-22-1
Каневский Даниил, Труфманов Михаил, Махров Матвей

## Бэкенд смартапа "Тренировка ЧГК"

### Реализация
Используется стороннее API для получения вопросов в формате ЧГК

Вопросы кэшируются в Redis для того, чтобы последующие запросы происходили быстрее

Реализовано получение запросов по REST и по gRPC

### Endpoints

```
GET /api/v1/question/random - получение случайного вопроса, в формате JSON
```
---
```
POST /api/v1/answer/check - проверка ответа

Request Body
{
    "userAnswer": <ответ-пользователя>
    "correctAnswer": <правильный ответ>
}
```

Также предусмотрены запросы по gRPC, порт для которого указывается в конфиге

### Запуск для локальной разработки с помощью docker-compose

Для начала необходимо собрать образ из [данного репозитория](https://github.com/Mihail20052005/testServer) и назвать образ **ml-service**

Затем
```bash
cd backend/question-service && docker-compose up --build
```

### Демо проекта
Демо проекта можно проверить по адресу
```
GET  https::/4-gk.ru/api/v1/question/random
POST https::/4-gk.ru/api/v1/answer/check
```

### TODO
* ~~[Реализовать сервис обработки ответов](https://github.com/Mihail20052005/testServer)~~
* Добавить Swagger документацию
* Добавить тесты
* Настроить CI/CD
* Реализовать [полноценный фронтенд для смартаппа](https://github.com/MatveyMakhrov/module-for-SBER/)

