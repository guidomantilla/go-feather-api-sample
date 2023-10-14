package repositories

import (
	"sync/atomic"

	feather_sql "github.com/guidomantilla/go-feather-sql/pkg/sql"

	"github.com/guidomantilla/go-feather-api-sample/pkg/models"
)

type Query int

const (
	FindPrincipalByIdQuery Query = iota
	FindResourceByIdQuery
	SaveResourceQuery
)

var singleton atomic.Value

type QueriesMap map[Query]string

func BuildQueries() QueriesMap {

	value := singleton.Load()
	if value != nil {
		return value.(QueriesMap)
	}

	driverName := feather_sql.MysqlDriverName
	paramHolder := feather_sql.NamedParamHolder

	statements := QueriesMap{
		FindPrincipalByIdQuery: feather_sql.CreateSelectSQL("auth_principals", models.AuthPrincipal{}, driverName, paramHolder, feather_sql.PkColumnFilter),
		FindResourceByIdQuery:  feather_sql.CreateSelectSQL("auth_resources", models.AuthResource{}, driverName, paramHolder, feather_sql.PkColumnFilter),
		SaveResourceQuery:      feather_sql.CreateInsertSQL("auth_resources", models.AuthResource{}, driverName, paramHolder) + " ON DUPLICATE KEY UPDATE name = :name, application = :application",
	}

	singleton.Store(statements)
	return statements
}
