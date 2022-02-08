package main
 
import (
  "rgb/internal/conf"
  "rgb/internal/server"
)
 
func main() {
  server.Start(conf.NewConfig())
}

//For hot reloading using nodemon, inside root dir run :
//nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/rgb/main.go

// We can feed these values to ENV using source command. So, now we can start our server like this:
// source .env
// go run cmd/rgb/main.go