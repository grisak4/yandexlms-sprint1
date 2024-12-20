```markdown
# Yandex LMS Sprint 1

Этот проект является результатом выполнения задания первого спринта на платформе Yandex LMS. Он реализует базовую функциональность калькулятора, предоставляя возможность выполнять операции над математическими выражениями через REST API.

## Описание

Приложение предоставляет HTTP API для выполнения арифметических операций, поддерживает скобки, и работает с числами с плавающей точкой. Основной функционал:
- Обработка запросов с математическими выражениями в формате JSON.
- Возвращение результата вычисления.
- Обработка ошибок, таких как деление на ноль или некорректный ввод.

## Технологии

- **Go**: основной язык разработки.
- **Gin**: веб-фреймворк для работы с HTTP.
- **Testify**: библиотека для тестирования.

## Установка и запуск

1. Клонируйте репозиторий:
   ```bash
   git clone https://github.com/grisak4/yandexlms-sprint1.git
   cd yandexlms-sprint1
   ```

2. Установите зависимости:
   ```bash
   go mod tidy
   ```

3. Запустите приложение:
   ```bash
   go run main.go
   ```

Приложение будет доступно на `http://localhost:8081`.

## API

### POST `/calculator`

#### Тело запроса
```json
{
  "expression": "2+2*2"
}
```

#### Ответы

- **200 OK**:
  ```json
  {
    "result": 6.0
  }
  ```

- **422 Unprocessable Entity**: если выражение отсутствует или некорректно.
  ```json
  {
    "error": "Expression is not valid"
  }
  ```

- **500 Internal Server Error**: если произошла внутренняя ошибка вычислений (например, деление на ноль).
  ```json
  {
    "error": "Internal server error"
  }
  ```

## Тестирование

Для запуска тестов выполните:
```bash
go test ./... -v
```

Тесты проверяют:
- Корректные вычисления.
- Обработку ошибок (деление на ноль, некорректные выражения).
- Корректный HTTP-ответ при отсутствии поля `expression`.
```