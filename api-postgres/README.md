for getting start with go you need to 
 brew install go



go to link below to be familiar with go syntax:
 https://gobyexample.com/closures

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

Before we begin writing our application, we need to fetch two packages that our application will depend on:

mux - The Gorilla Mux router, and
pq - The PostgreSQL driver.
You can fetch these using the following command:

go get github.com/gorilla/mux github.com/lib/pq





 

MAYBE NEW FOR GO BEGINNER:

