package main

//go:generate mockgen -package main -source go_mock_test_demo.go --destination go_mock_test_demo_mock.go
type DB interface {
	Get(key string) (int, error)
}

func GetFromDB(db DB, ket string) int {
	if value, err := db.Get(ket); err == nil {
		return value
	}
	return -1
}
