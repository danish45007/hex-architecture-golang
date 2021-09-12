package ports

type DBPort interface {
	CloseDbConn()
	AddToHistory(answer int32, operation string) error
}
