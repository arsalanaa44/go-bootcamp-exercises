package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	u := entity.User{
		ID:          0,
		Name:        "",
		PhoneNumber: "",
	}
	var createdAt []uint8

	//fmt.Println(row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt))
	//fmt.Println(row)
	//fmt.Println(u, string(createdAt))
	if sErr := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt); sErr != nil {
		if sErr == sql.ErrNoRows {

			return true, nil
		}

		return false, sErr
	}

	return false, nil
}

func (d *MySQLDB) Register(user entity.User) (entity.User, error) {
	res, err := d.db.Exec(`insert into users(name , phone_number) values(?, ?)`, user.Name, user.PhoneNumber)
	if err != nil {
		return entity.User{}, fmt.Errorf("can't execute command: %w", err)
	}

	// error is always nil
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return user, nil
}
