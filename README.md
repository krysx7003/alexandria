
## Nodes

- /api/books/
- /api/series/
- /api/authors/
- /api/genres/

- /media/book/{format}/

## Json

    {
        "id":"1",
        "title":"Book 10",
        "series":{
            "name":"He Who Fights with Monsters",
            "book":10,
            "ongoing":true
            "books":nil
        },
        "genres":[
            {"name":"LitRPG"},
            {"name":"Fantasy"},
            {"name":"Adventure"}
        ],
        "author":[
            {"name":"Shirtaloon"},
            {"name":"Travis Deverell"}
        ],
        "publisher":"Aethon Books",
        "publish_date":"",
        "cover_url":"",
        "url":"",
        "isbn":{
            "isbn10":""
            "isbn13":""
        },
        "created_at":"0001-01-01T00:00:00Z",
        "updated_at":"0001-01-01T00:00:00Z"
    }

## DataBase

GENRES

| ID       | NAME   | CREATED | UPDATED |
| -------- | ------ | ------- | ------- |
| Integer  | String | Date    | Date    |

BOOKS_GENRES

| BOOK_ID  | GENRE_ID |
| -------- | -------- |
| Integer  | Integer  |

AUTHORS

| ID       | NAME   | CREATED | UPDATED |
| -------- | ------ | ------- | ------- |
| Integer  | String | Date    | Date    |

BOOKS_AUTHORS

| BOOK_ID  | AUTHOR_ID |
| -------- | --------- |
| Integer  | Integer   |

ISBNS

| ID       | v10    | v13    | CREATED | UPDATED |
| -------- | ------ | ------ | ------- | ------- |
| Integer  | String | String | Date    | Date    |

SERIES

| ID       | NAME   | ONGOING | TOTAL_BOOKS | CREATED | UPDATED |
| -------- | ------ | ------- | ----------- | ------- | ------- |
| Integer  | String | Bool    | Integer     | Date    | Date    |

BOOKS

| ID       | TITLE  | SERIES_ID | BOOK_IN_SERIES | PUBLISHER | PUBLISH_DATE | COVER_URL | MEDIA_URL | ISBN_ID | CREATED | UPDATED |
| -------- | ------ | --------- | -------------- | --------- | ------------ | --------- | --------- | ------- | ------- | ------- |
| Integer  | String | Integer   | Integer        | String    | Date         | String    | String    | Integer | Date    | Date    |
