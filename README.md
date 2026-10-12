# wisp server log injection

How to use:
Modify the following fields in origin.go
```
    target := "wisp server" <- wisp server
    path := "/" = <- wisp path
    useTLS := true = <- whether to use TLS or not
    

    payload := "\n\n\r\n\n\r\n\n\r\n\n\rwhats up g\n\n\n" <- payload to log
```

run by using ```go run origin.go```

Originially discovered on march 26, publicly released today 
