# api-go

for getting start with go you need to 
 brew install go




go to link below to be familiar with go syntax:
 https://gobyexample.com/closures

MAYBE NEW FOR GO BEGINNER:

for using mux you should:
 go get github.com/gorilla/mux
 Package mux implements a request router and dispatcher

"Content-Type: application/json" :
 The header just denotes what the content is encoded in. It is not necessarily possible to deduce the type of the content from the content itself, i.e. you can't necessarily just look at the content and know what to do with it. That's what HTTP headers are for, they tell the recipient what kind of content they're (supposedly) dealing with.
 Content-type: application/json; charset=utf-8 designates the content to be in JSON format, encoded in the UTF-8 character encoding. Designating the encoding is somewhat redundant for JSON, since the default (only?) encoding for JSON is UTF-8. So in this case the receiving server apparently is happy knowing that it's dealing with JSON and assumes that the encoding is UTF-8 by default, that's why it works with or without the header. if you want more got to https://stackoverflow.com/questions/9254891/what-does-content-type-application-json-charset-utf-8-really-mean

mux.Vars(r):
 makes mux variable 

panic:
 A panic typically means something went unexpectedly wrong. Mostly we use it to fail fast on errors that shouldn’t occur during normal operation, or that we aren’t prepared to handle gracefully.

NewDecoder and NewEncoder:
 Package json implements encoding and decoding of JSON as defined in RFC 4627. The mapping between JSON and Go values is described in the documentation for the Marshal and Unmarshal functions.
 The json package provides Decoder and Encoder types to support the common operation of reading and writing streams of JSON data. The NewDecoder and NewEncoder functions wrap the io.Reader and io.Writer interface types.
 //func NewDecoder(r io.Reader) *Decoder 
 //func NewEncoder(w io.Writer) *Encoder 

"io" :
The io package specifies the io.Reader interface, which represents the read end of a stream of data.
 more information:https://tour.golang.org/methods/21

 "net/http":
 Package http provides HTTP client and server implementations.

"strconv":
 Package strconv implements conversions to and from string representations of basic data types.
 The most common numeric conversions are Atoi (string to int) and Itoa (int to string).

"fmt":
 Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

"log":
  Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output. It also has a predefined 'standard' Logger accessible through helper functions Print[f|ln], Fatal[f|ln], and Panic[f|ln], which are easier to use than creating a Logger manually. That logger writes to standard error and prints the date and time of each logged message


for posting data use curl 
example:
curl -H "Content-Type: application/json" -d '{}' http://localhost:8080/todos

for deleting data use curl too 
example:
curl -X DELETE localhost:8080/todos/3



