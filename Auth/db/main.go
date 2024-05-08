package db

type DB struct {
	postgresString     string
	PostgresConnection string
}

func NewDBConnection(db string) *DB {
	mydb := &DB{
		postgresString: db,
	}
	mydb.Connect()

	return mydb
}

func (db *DB) Connect() error {
	db.PostgresConnection = "connected"
	return nil
}
