for getting start with go you need to 
 brew install go



go to link below to be familiar with go syntax:
 https://gobyexample.com/closures




Before we begin writing our application, we need to fetch two packages that our application will depend on:

mux - The Gorilla Mux router, and
pq - The PostgreSQL driver.
You can fetch these using the following command:

go get github.com/gorilla/mux github.com/lib/pq

DATABASE PART

make a database in psql
connect to your data base
make a products table :

CREATE TABLE products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)

The database/sql package provides a generic interface around SQL (or SQL-like) databases. 

How we handle the databases in Go?

Open :is used to create a database handle

Ping: to verify that a connection can be made before making a query, use the Ping function

Exec: is used for queries where no rows are returned

Query :is used for retrieval,if a function name includes Query, it is designed to "ask a question of the database", and will return a set of rows, even if it’s empty

QueryRow :is used where only a single row is expected

Scan():We will assign results to variables, a row at a time, with rows.Scan()

good link too read before start to query:
https://github.com/golang/go/wiki/SQLInterface
http://go-database-sql.org/retrieving.html


