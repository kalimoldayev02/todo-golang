package database

type DriverInterface interface {
	Connect(driverData *DriverData)
	Disconnect()
	Execute() DriverInterface
	Add(query string) DriverInterface
	GetQueryString() string
	RefreshQueryString()
	FetchAll() []interface{}
	Fetch(key string) []interface{}
}
