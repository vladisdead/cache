package storage

import (
	"context"
	"fmt"
	"task4/model"
	"time"
)

func (p *Provider) GetAllUsersFromDB() ([]model.Users, error) {
	var (
		q     string
		err   error
		users = make([]model.Users, 0)
	)

	q = `SELECT id, name, birthday, age, is_male FROM ref.tb_table`
	rows, err := p.Conn.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.Users

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Birthday,
			&user.Age,
			&user.IsMale,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	p.Log.Print(q)

	return users, rows.Err()
}

func (p *Provider) GetUsersWithMinAgeFromDB(age int) ([]model.Users, error) {
	var (
		q     string
		err   error
		users = make([]model.Users, 0)
	)

	q = `SELECT id, name, birthday, age, is_male FROM ref.tb_table WHERE age > $1`
	rows, err := p.Conn.Query(q, age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.Users

		if err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Birthday,
			&user.Age,
			&user.IsMale,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	p.Log.Print(q)

	return users, rows.Err()
}

func (p *Provider) GetUserByIdDB(id int) (model.Users, error) {
	var (
		q    string
		err  error
		user model.Users
	)

	q = `SELECT id, name, birthday, age, is_male FROM ref.tb_table WHERE id = $1`

	if err = p.Conn.QueryRow(q, id).Scan(&user.Id,
		&user.Name,
		&user.Birthday,
		&user.Age,
		&user.IsMale); err != nil {
		fmt.Println(err)
		return model.Users{}, fmt.Errorf("Такого пользователся нет. Id = %d", id)
	}

	p.Log.Print(q)

	return user, nil
}

func (p *Provider) CreateUserDB(user model.Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var q string

	t1, _ := time.Parse("2006-01-02", user.Birthday)
	t2 := time.Now()

	user.Age = int(t2.Sub(t1).Hours() / 24 / 365)

	q = `INSERT INTO ref.tb_table (name, birthday, age, is_male) VALUES ($1, $2, $3, $4)`

	_, err = tx.Exec(q, user.Name, user.Birthday, user.Age, user.IsMale)
	if err != nil {
		return err
	}

	p.Log.Print(q)

	return tx.Commit()
}

func (p *Provider) ChangeUserDB(id int, user model.Users) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		q      string
		userId int
	)

	q = `SELECT id FROM ref.tb_table WHERE id = $1`

	if err = tx.QueryRow(q, id).Scan(&userId); err != nil {
		return fmt.Errorf("пользователся не существует. Id пользователя %d", id)
	}

	q = `UPDATE ref.tb_table SET name = $1, birthday = $2, age = $3, is_male = $4 WHERE id = $5`

	t1, _ := time.Parse("2006-01-02", user.Birthday)
	t2 := time.Now()

	user.Age = int(t2.Sub(t1).Hours() / 24 / 365)

	if _, err := tx.Exec(q, user.Name, user.Birthday, user.Age, user.IsMale, id); err != nil {
		return nil
	}

	p.Log.Print(q)

	return tx.Commit()
}

func (p *Provider) DeleteUserDB(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultCtxTimeout)
	defer cancel()

	tx, err := p.Conn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	var (
		q      string
		userId int
	)

	q = `SELECT id FROM ref.tb_table WHERE id = $1`

	if err = tx.QueryRow(q, id).Scan(&userId); err != nil {
		return fmt.Errorf("пользователся не существует. Id пользователя %d", id)
	}

	q = `DELETE FROM ref.tb_table WHERE id = $1`

	_, err = tx.Exec(q, id)
	if err != nil {
		return err
	}

	p.Log.Print(q)

	return tx.Commit()
}

func (p *Provider) GetAllCache() ([]model.LocalCache,error) {
	return nil, nil
}