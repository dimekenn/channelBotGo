package repository

import "database/sql"

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) *MessageRepo  {
	return &MessageRepo{db: db}
}

func (mr *MessageRepo) InsertMessage(id int, message string) error {
	tx, err := mr.db.Begin()
	if err != nil{
		return err
	}
	_, err = tx.Query(
		"insert into messages (id, message) values ($1, $2)",
		id, message,
	)
	if err != nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (mr *MessageRepo) GetText(id int) (string, error)  {
	tx, err := mr.db.Begin()
	if err != nil{
		return "", err
	}
	var text string
	res, err := tx.Query(
		"select message from messages where id = $1",
			id,
	)
	if err != nil{
		tx.Rollback()
		return "", err
	}
	for res.Next() {
		err = res.Scan(&text)
		if err != nil {
			return "", err
		}
	}
	tx.Commit()
	return text, nil
}
