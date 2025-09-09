# Go API Link Service
Минималистичный RESTful API для управления ссылками, написанный на Go.

Функциональность:
- Добавление, получение, обновление и удаление ссылок (CRUD).
- Простой JSON-based интерфейс.

Установка и запуск

1. Клонирование репозитория:
git clone https://github.com/LordVillain/Go_API_Link_Service.git
cd Go_API_Link_Service

2. Установка зависимостей:
go mod tidy

3. Запуск приложения:
go run main.go

4. Работа с API
Метод	    URL	        Описание
GET	    /links	     Получить список всех ссылок
GET	    /links/{id}	 Получить ссылку по её ID
POST	  /links	     Создать новую ссылку
PUT	    /links/{id}	 Обновить существующую ссылку
DELETE	/links/{id}	 Удалить ссылку по ID

Пример POST /links:
{
  "url": "https://example.com",
  "description": "Example website"
}
Ответ:
{
  "id": 1,
  "url": "https://example.com",
  "description": "Example website"
}



s/1
