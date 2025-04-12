package simple

type Database struct {
	Name string
}

type DatabasePostgreeSql Database
type DatabaseMySql Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreeSql
	DatabaseMySQL *DatabaseMySql
}



func NewDatabasePostgreSQL () *DatabasePostgreeSql {
	return (*DatabasePostgreeSql)(&Database{
		Name: "PostgreSQL",
	})
}

func NewDatabaseMySQL () *DatabaseMySql {
	return (*DatabaseMySql)(&Database{
		Name: "MySQL",
	})
}

func NewDatabase (postgreSQL *DatabasePostgreeSql, mySQL *DatabaseMySql) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMySQL: mySQL,
	}
}