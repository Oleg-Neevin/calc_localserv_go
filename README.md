# calc_localserv_go

## Описание проекта
Этот проект представляет собой **HTTP-сервер** для вычисления арифметических выражений. 
Сервер принимает POST-запросы с арифметическим выражением в формате JSON и возвращает результат вычисления.

### Основные функции:
- Обработка арифметических выражений.
- Обработка ошибок:
  - 422 (Unprocessable Entity)
  - 500 (Internal Server Error).

---

## Примеры использования
### Успешное выполнение:
```bash
curl --location 'http://127.0.0.1:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*2"
}'
```
Ответ:
```
{
  result: "6.000000"
}
```

### Ошибка 422 (invalid expression):
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2*-"
}'
```
Ответ:
```
{
  error: "invalid expression"
}
```
### Ошибка 500 (Internal Server Error):
```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "Hello world!"
}'
```
Ответ:
```
{
error: "It's not a bug. It's a feature"
}
```

---

## Инструкция по запуску
1. Убедитесь, что у вас установлен Go:
```
go version
```
2. Склонируйте репозиторий проекта с github
3. Перейдите в корневую директорию проекта
4. Запустите сервер командой:
```bash
go run .\cmd\main.go
```
