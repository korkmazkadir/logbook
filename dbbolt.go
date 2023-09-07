package logbook

import (
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/bbolt"
)

type DBbolt struct {
	db *bbolt.DB
}

func NewDBbbolt(filePath string) (*DBbolt, error) {
	db, err := bbolt.Open(filePath, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &DBbolt{db: db}, nil
}

func (bolt *DBbolt) CreateBook(bookID string) error {
	err := bolt.db.Update(func(tx *bbolt.Tx) error {
		//_, err := tx.CreateBucketIfNotExists([]byte(bookID))
		_, err := tx.CreateBucket([]byte(bookID))
		return err
	})

	if err == bbolt.ErrBucketExists {
		return fmt.Errorf("logbook with ID %s already exists", bookID)
	}

	return err
}

// AppendLog appends a log to the book specified by bookID
func (bolt *DBbolt) AppendLog(bookID string, log Log) (Log, error) {

	log.Time = time.Now()

	err := bolt.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(bookID))

		if b == nil {
			return fmt.Errorf("logbook with ID %s does not exists", bookID)
		}

		key := []byte(log.Time.Format(time.RFC3339))
		value := b.Get(key)
		if value != nil {
			return fmt.Errorf("too many append log request, try again later")
		}

		id, err := b.NextSequence()
		if err != nil {
			return err
		}

		log.ID = id

		buf, err := json.Marshal(log)
		if err != nil {
			return err
		}

		return b.Put(key, buf)
	})

	return log, err
}

// GetLogs retreives the logs of a book specified by the bookID
func (bolt *DBbolt) GetLogs(bookID string) ([]Log, error) {

	var logs []Log
	err := bolt.db.View(func(tx *bbolt.Tx) error {

		b := tx.Bucket([]byte(bookID))

		if b == nil {
			return fmt.Errorf("book with ID %s is not available", bookID)
		}

		err := b.ForEach(func(_, value []byte) error {

			var log Log
			err := json.Unmarshal(value, &log)
			if err != nil {
				return err
			}

			logs = append(logs, log)
			return nil
		})

		return err
	})

	return logs, err
}
