gonuts
======
A go experiment.

features
--------
 - tells the user the current time 
 - rot13 encodes text
 - displays http request headers


instructions
------------
go run gonuts.go

Features
--------
 - acts as a webserver that displays the current time when you GET /


Dependencies
------------
**go 1.1** and greater is required due to dependency on Martini(https://github.com/go-martini/martini/)):
~~~
go get github.com/go-martini/martini
go get github.com/joshlf13/rot13
~~~
