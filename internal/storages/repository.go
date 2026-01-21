package storages

func NewRepository(dbType string) ExchangeStorage {
	switch dbType {
	case "mock":
		return NewMockStorage()
	// case "postgres":
	//     return postgres.NewPostgresStorage(...)
	default:
		return NewMockStorage() // fallback
	}
}
