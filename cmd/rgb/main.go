package main
 
import "rgb/internal/server"
 
func main() {
  server.Start()
}

//For hot reloading using nodemon, inside root dir run :
//nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/rgb/main.go