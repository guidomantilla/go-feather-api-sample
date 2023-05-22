package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/sijms/go-ora/v2"

	"github.com/guidomantilla/go-feather-api-sample/cmd"
)

func main() {
	cmd.ExecuteAppCmd()
}
