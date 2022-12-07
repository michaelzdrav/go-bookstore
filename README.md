# go-bookstore


How to setup GORM connection to db:
gorm db connection: "dbType",  "username:password/nameOfTable?charset=utf8&parseTime=True&loc=Local"

Routes:
1. Create Book
GET - http://localhost:9010/book/

JSON Body
{
    "name": "Michael Goes To Hollywood",
    "author": "Michael Zdravkovski",
    "publication": "Nad's publication company"
}

2. Get All Books
GET - http://localhost:9010/book/

3. Get Book By ID
GET - http://localhost:9010/book/{id}

4. Update Book By ID
PUT - http://localhost:9010/book/{id}

JSON Body
{
    "name": "Daniel",
    "author": "Stevenson",
    "publication": "Jackson's Publication"
}

5. Delete Book By ID
DELETE - http://localhost:9010/book/{id}