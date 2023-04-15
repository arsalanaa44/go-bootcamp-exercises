package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	u := entity.User{}
	var createdAt []uint8

	//fmt.Println(row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt))
	//fmt.Println(row)
	//fmt.Println(u, string(createdAt))
	if sErr := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt); sErr != nil {
		if sErr == sql.ErrNoRows {

			return true, nil
		}

		return false, sErr
	}

	return false, nil
}

func (d *MySQLDB) Register(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name , phone_number, password) values(?, ?, ?)`,
		user.Name, user.PhoneNumber, user.Password)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (d *MySQLDB) GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error) {

	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	u := entity.User{}
	var createdAt []uint8

	if sErr := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt); sErr != nil {
		if sErr == sql.ErrNoRows {

			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("can't scan query result: %w", sErr)
	}

	return u, true, nil
}
