# BookShelf API

**BookShelf API** — простой RESTful API для управления личной библиотекой 


### CRUD
- **POST** `/books` — создать книгу
- **GET** `/books?page=1&limit=10` — получить список книг с пагинацией
- **GET** `/books/{id}` — получить книгу по id
- **PUT** `/books/{id}` — обновить книгу
- **DELETE** `/books/{id}` — удалить книгу

### CRUD с бизнес-логика
- **GET** `/books/recommend` — рекомендации (топ-5 книг по рейтингу)
- **POST** `/books/{id}/mark-out-of-stock` — отметить книгу как закончившуюся (`out_of_stock=true`)

## Запуск

### 1. Создать `.env`
В корне проекта создай файл `.env`:

```env
APP_PORT=8080

DB_HOST=db
DB_PORT=5432
DB_NAME=bookshelf
DB_USER=postgres
DB_PASSWORD=postgres
DB_SSLMODE=disable
```

### 2. Запуск контейнеров

```
docker compose up -d
```

### 3. Проверить, что контейнеры поднялись

```
docker compose ps
```

### 4. Проверить, что миграции применились

```
docker compose exec db psql -U postgres -d bookshelf -c “\dt”
docker compose exec db psql -U postgres -d bookshelf -c “SELECT * FROM schema_migrations;”
```

### 5. Создать книгу (POST /books)

```curl -i -X POST http://localhost:8080/books -H “Content-Type: application/json” -d ‘{“title”:“Dune”,“author”:“Frank Herbert”,“year”:1965,“isbn”:“9780441172719”,“out_of_stock”:false,“read”:true,“rating”:9}’
```
### 6. Получить список книг (GET /books?page=1&limit=10)
```
curl -i “http://localhost:8080/books?page=1&limit=10”
```

### 7. Получить книгу по id (GET /books/{id})
```
curl -i http://localhost:8080/books/1
```

### 8. Обновить книгу (PUT /books/{id})
```
curl -i -X PUT http://localhost:8080/books/1 -H “Content-Type: application/json” -d ‘{“title”:“Dune (updated)”,“author”:“Frank Herbert”,“year”:1965,“isbn”:“9780441172719”,“out_of_stock”:false,“read”:true,“rating”:10}’
```

### 9. Удалить книгу (DELETE /books/{id})
```
curl -i -X DELETE http://localhost:8080/books/1
```

### 10. Отметить “закончилась” (POST /books/{id}/mark-out-of-stock)
```
curl -i -X POST http://localhost:8080/books/1/mark-out-of-stock
```

### 11. Рекомендации (GET /books/recommend)
```
curl -i http://localhost:8080/books/recommend
```


