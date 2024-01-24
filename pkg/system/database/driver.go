package database

// ActiveRecord - интерфейс, представляющий базовый функционал для работы с БД
type ActiveRecord interface {
	Add(query string) ActiveRecord
	Connect() error
	Disconnect()
	GetQueryString() string
	RefreshQueryString()
	Execute(params ...interface{}) ActiveRecord
	FetchAll() [][]interface{}
}

// DriverData - структура для хранения данных драйвера
type DriverData struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
	Driver   string
}
