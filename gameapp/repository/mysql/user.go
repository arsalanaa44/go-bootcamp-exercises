package mysql

import (
	"database/sql"
	"fmt"
	"gameapp/entity"
)

func (d *MySQLDB) IsPhoneNumberUnique(phoneNumber string) (bool, error) {
	row := d.db.QueryRow(`select * from users where phone_number = ?`, phoneNumber)
	//u := entity.User{}
	//var createdAt []uint8

	//fmt.Println(row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &createdAt))
	//fmt.Println(row)
	//fmt.Println(u, string(createdAt))
	//if sErr := row.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.Password, &createdAt); sErr != nil {
	//	if sErr == sql.ErrNoRows {
	//
	//		return true, nil
	//	}
	//
	//	return false, sErr
	//}
	if _, sErr := scanUser(row); sErr != nil {
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

	if user, sErr := scanUser(row); sErr != nil {
		if sErr == sql.ErrNoRows {

			return entity.User{}, false, nil
		}

		return entity.User{}, false, fmt.Errorf("can't scan query result: %w", sErr)
	} else {

		return user, true, nil
	}
}

func (d *MySQLDB) GetUserByID(userID int) (entity.User, error) {

	row := d.db.QueryRow("select * from users where id = ?", userID)
	if user, sErr := scanUser(row); sErr != nil {
		if sErr == sql.ErrNoRows {

			return entity.User{}, fmt.Errorf("record not found: %w", sErr)
		}

		return entity.User{}, fmt.Errorf("can't scan query result: %w", sErr)
	} else {

		return user, nil
	}

}

// fewer duplicated code
func scanUser(row *sql.Row) (entity.User, error) {
	var createdAt []uint8
	var user entity.User

	err := row.Scan(&user.ID, &user.Name, &user.PhoneNumber, &user.Password, &createdAt)

	return user, err
}
