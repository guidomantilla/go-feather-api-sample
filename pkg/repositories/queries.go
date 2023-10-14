package repositories

import (
	"sync/atomic"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

var singleton atomic.Value

type QueriesMap map[string]string

func BuildQueries() QueriesMap {

	value := singleton.Load()
	if value != nil {
		return value.(QueriesMap)
	}

	driverName := feather_sql.MysqlDriverName
	paramHolder := feather_sql.NamedParamHolder

	statements := QueriesMap{
		"FindPrincipalById": feather_sql.CreateSelectSQL("auth_principals", models.AuthPrincipal{}, driverName, paramHolder, feather_sql.PkColumnFilter),
		"FindResourceById":  feather_sql.CreateSelectSQL("auth_resources", models.AuthResource{}, driverName, paramHolder, feather_sql.PkColumnFilter),
		"SaveResource":      feather_sql.CreateInsertSQL("auth_resources", models.AuthResource{}, driverName, paramHolder) + " ON DUPLICATE KEY UPDATE name = :name, application = :application",
	}

	singleton.Store(statements)
	return statements
}
