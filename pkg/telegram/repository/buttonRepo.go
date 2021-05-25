package repository

import "database/sql"

type ButtonRepo struct {
	db *sql.DB
}

func NewButtonRepo(db *sql.DB) *ButtonRepo  {
	return &ButtonRepo{db: db}
}

func (br *ButtonRepo) InsertButton(id , userId int) error  {
	tx, err := br.db.Begin()
	if err !=nil{
		return err
	}
	_, err = tx.Query(
		"insert into buttons (id, user_id) values ($1, $2)",
		id, userId,
	)
	if err != nil{
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (br *ButtonRepo) GetUserId(id int) (int, error)  {
	tx, err := br.db.Begin()
	if err != nil{
		return 0, err
	}
	var userId int
	res, err := tx.Query(
		"select user_id from buttons where id = $1",
		id,
	)
	if err != nil{
		tx.Rollback()
		return 0, err
	}
	for res.Next() {
		err = res.Scan(&userId)
		if err != nil {
			return 0, err
		}
	}
	tx.Commit()
	return userId, nil
}
