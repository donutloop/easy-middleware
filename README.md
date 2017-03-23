[![Build Status](https://travis-ci.org/donutloop/easy-middleware.svg?branch=master)](https://travis-ci.org/donutloop/easy-middleware)

# easy-middleware

# What is easy-middleware ?

easy-middleware is a lightweight json middleware stack for Go >= 1.7.

## Features:

* Dump incoming request middleware
* Set body limit middleware
* Json header check middleware
* Validator middleware
* Sql db middleware 
* Recovery middleware
* With value middleware
* Json response writer middleware 

## Feature request are welcome

## Example :

```go
    package main

    import (
        "net/http"
        "fmt"
        "os"

        "github.com/donutloop/easy-middleware"
    )

    func main() {

        stack := easy_middleware.New()
    	stack.Add(easy_middleware.Json())
    	stack.Add(easy_middleware.SqlDb("postgres", "postgres://postgres:postgres@db/postgres?sslmode=disable"))
    
        http.Handle("/user", stack.Then(http.HandlerFunc(userHandler)))
    }

    func userHandler(rw http.ResponseWriter, req *http.Request) {
         jrw := rw.(*easy_middleware.JsonResponseWriter)          
          
         var db *sql.DB
         if rv := r.Context().Value("db"); rv != nil {
         		if v, ok := rv.(*sql.DB); ok {
         			 db = v 
         		} else if err := rv.(*easy_middleware.DatabaseError); err != nil {
         			 jrw.WriteError(err, http.StatusInternalServerError)
         			 return 
         		}
         }  
        //...
        
    }
```

## More documentation comming soon
