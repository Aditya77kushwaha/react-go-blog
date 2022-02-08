package database
 
import (
  "rgb/internal/conf"
 
  "github.com/go-pg/pg/v10"
)
 
func NewDBOptions(cfg conf.Config) *pg.Options {
  return &pg.Options{
    Addr:     cfg.DbHost + ":" + cfg.DbPort,
    Database: cfg.DbName,
    User:     cfg.DbUser,
    Password: cfg.DbPassword,
  }
}

// # Create a new PostgreSQL user called testuser, allow user to login, but NOT creating databases
// $ sudo -u postgres createuser --login --pwprompt testuser