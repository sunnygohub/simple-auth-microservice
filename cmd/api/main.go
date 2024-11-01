package main

import (
	"database/sql"

	"github.com/sunnygohub/simple-auth-microservice/data"
)

const webPort = "80"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

}
