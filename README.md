# Файлы для итогового задания

В директории `tests` находятся тесты для проверки API, которое должно быть реализовано в веб-сервере.

Директория `web` содержит файлы фронтенда.

# Структура проекта

## .env
PORT - порт для запуска сервера
DATA_SOURCE - путь до sqlite базы

## Domain
Содержит описания типов Task, TaskListResponse, TaskCreateResponse

### Database
Пакет со структурой которая содержит методы для настройки соединения с базой и методами для ее обновления

- DeleteTaskById
- InsertTask
- SelectTaskById
- SelectTasks
- UpdateTask


### Server
Пакет со структрой которая содержит методы для запуска сервера и обработчиков endpoint

- PostTaskDone `POST /api/task/done`
- HandleNextDate `GET /api/nextdate`
- HandleTask `/api/task`
    - GetTask `GET /api/task`
    - CreateTask `POST /api/task`
    - PutTask `PUT /api/task`
    - DeleteTask `DELETE /api/task`
- ListTasks `GET /api/tasks`