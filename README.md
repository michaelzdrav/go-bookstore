# go-bookstore

Run MYSQL container locally (please change the root password as this was an example):
docker run --name basic-mysql --rm -v /tmp/mysql-data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=ANSKk08aPEDbFjDO -e MYSQL_DATABASE=testing -p 3306:3306 -it mysql:8.0

How to setup GORM connection to db:
gorm db connection: "dbType",  "username:password/nameOfTable?charset=utf8&parseTime=True&loc=Local"

Compile this locally:
Navigate to cmd/main 
Run `go build`
Run `go run main.go`

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