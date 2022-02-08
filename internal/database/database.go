package database
 
import (
  "github.com/go-pg/pg/v10"
)
 
func NewDBOptions() *pg.Options {
  return &pg.Options{
    Addr:     "localhost:5432",
    Database: "rgb",
    User:     "testuser",
    Password: "Aditya123@",
  }
}

// # Create a new PostgreSQL user called testuser, allow user to login, but NOT creating databases
// $ sudo -u postgres createuser --login --pwprompt testuser