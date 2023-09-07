package logbook

type Database interface {
	CreateBook(bookID string) error
	AppendLog(bookID string, log Log) (Log, error)
	GetLogs(bookID string) ([]Log, error)
	//GetLog(bookID string, logID uint64)
}
