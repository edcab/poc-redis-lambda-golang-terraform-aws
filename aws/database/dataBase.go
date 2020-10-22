package database

// DataBase representation basic actions on data base
type DataBase interface {
	OpenConnection() error
	CloseConnection()
	GetConnection() interface{}
}
