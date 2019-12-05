# go-battlesnake
A [Battlesnake.io](http://battlesnake.io) Snake AI server written in Go

Written for Battlesnake Snake API version 2018.03.beta (found at https://docs.battlesnake.com/snake-api)

### Steps to run the AI locally (on Windows)

1. Download and install Go (version 1.12) from https://golang.org/doc/install

2. Clone repo to your GOPATH src folder:
    ```
    git clone https://github.com/sanderaido/go-battlesnake.git %GOPATH%/src/github.com/sanderaido/go-battlesnake
    ```
3. Run the server with:
    ```
    go run github.com/sanderaido/go-battlesnake
    ```
4. The AI is accessible at http://localhost:8080/
   
   
### Run tests with
```
go test github.com/sanderaido/go-battlesnake/...
```