gonuts
======
A go experiment, intended to help me learn go


Features
--------
 - acts as a webserver that displays the current time when you GET /
 - rot13 encodes text when you GET /mytext
 - displays http request headers when you get /
 - generates random ints in a concurrent fashion
 - solves some project euler problems

Running it
----------
~~~
go run gonuts.go
~~~

choose an option from the menu

Dependencies
------------
**go 1.1** and greater is required due to dependency on Martini(https://github.com/go-martini/martini/)):
~~~
go get github.com/go-martini/martini
go get github.com/joshlf13/rot13
~~~
