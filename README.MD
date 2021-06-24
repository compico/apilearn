# Apilearn

## Интерфейсы

Метод|Интефейсы|Описание
-|-|-
GET     |/api/v1/getbookslist        |Получение список всех книг
GET     |/api/v1/getbook/{book_id}   |Получение книги по id
DELETE  |/api/v1/deletebook/{book_id}|Удалить книгу из списка
POST    |/api/v1/updatebook/{book_id}|Обновить информацию книги
PUT     |/api/v1/addbook             |Добавить книгу
GET     |/api/v1/getlengthlist       |Получить количество книг
PUT     |/api/v1/addtestbooks        |Добавить тестовые книги

## Примеры запросов-ответов
---
`Запрос:`
>GET /api/v1/getbookslist

`Ответ:`
>Status code 200
```json
    {
        "books":[{
            "id": 0,               //number
            "name":"bookname",     //string
            "author":"bookauthor", //string
            "pages": 15            //number
            },
            {
            "id": 2,
            "name":"anotherbookname",
            "author":"anotherbookauthor",
            "pages": 15
            },
        ]
    }
```
---
`Запрос:`
>GET /api/v1/getbook/0

`Ответ:`
>Status code 200
```json
    {
            "id": 0,
            "name":"bookname",
            "author":"bookauthor",
            "pages": 15
            }
    }
```
---
`Запрос:`
>DELETE /api/v1/deletebook/0

`Ответ:`
>Status code 200
---
`Запрос:`
>POST /api/v1/updatebook/{book_id}
```json
{
    //Не обязательно писать все поля
    "name":"bookname",
    "author":"bookauthor",
    "pages": 25
}
```

`Ответ:`
>Status code 200
---
`Запрос:`
>PUT /api/v1/addbook
```json
{
    "name":"bookname",
    "author":"bookauthor",
    "pages": 25
}
```

`Ответ:`
>Status code 202
---
`Запрос:`
>GET /api/v1/getlengthlist

`Ответ:`
```json
 {"length":10}
```
>Status code 200
---
`Запрос:`
>PUT /api/v1/addtestbooks

`Ответ:`
>Status code 200