package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"todo-golang/pkg/system/database"

	_ "github.com/lib/pq"
)

var (
	db          *sql.DB
	stmp        *sql.Stmt
	err         error
	isConnected bool = false
	queryString string
)

type Postgres struct {
	driverData database.DriverData
}

func NewPostgres(dbData database.DriverData) *Postgres {
	return &Postgres{
		driverData: dbData,
	}
}

func (p *Postgres) Connect() error {
	dsn := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=disable",
		p.driverData.Host, p.driverData.User, p.driverData.DBName, p.driverData.Password,
	)
	postgre, err := sql.Open(p.driverData.Driver, dsn)
	if err != nil {
		return err
	}

	db = postgre
	isConnected = true
	return nil
}

func (p *Postgres) Disconnect() {
	if isConnected {
		db.Close()
		isConnected = false
	}
}

func (p *Postgres) Add(query string) database.ActiveRecord {
	queryString = fmt.Sprintf("%s %s", queryString, query)
	return p
}

func (p *Postgres) GetQueryString() string {
	return queryString
}

func (p *Postgres) RefreshQueryString() {
	queryString = ""
}

func (p *Postgres) Execute(params ...interface{}) database.ActiveRecord {
	return p.executeQuery(p.GetQueryString(), params...)
}

func (p *Postgres) executeQuery(query string, params ...interface{}) database.ActiveRecord {
	if !isConnected {
		if err := p.Connect(); err != nil {
			log.Fatal(err.Error())
		}
	}

	stmpPrepare, errPrepare := db.Prepare(query)
	if errPrepare != nil {
		log.Fatal(errPrepare)
	}

	stmp = stmpPrepare
	p.RefreshQueryString()
	return p
}

func (p *Postgres) FetchAll() [][]interface{} {
	var response [][]interface{}

	rows, err := stmp.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer p.Disconnect()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	values := make([]interface{}, len(columns))
	valuePointers := make([]interface{}, len(columns))

	for i := range columns {
		valuePointers[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(valuePointers...); err != nil {
			log.Fatal(err)
		}

		rowCopy := make([]interface{}, len(values))
		copy(rowCopy, values)
		response = append(response, rowCopy)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return response
}
