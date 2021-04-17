/*
 * Revision History:
 *     Initial: 2020/11/24      oiar
 */

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	UserID   string
	UserName string
	Path     string
}

const (
	mysqlUserCreateTable = iota
	mysqlUserInsert
	mysqlUserList
	mysqlUserInfoByID
	mysqlUserDeleteByID
	mysqlUpdateUserInfoByID
)

var (
	errInvalidInsert = errors.New("insert User:insert affected 0 rows")

	UserSQLString = []string{
		`CREATE TABLE IF NOT EXISTS %s (
			userId      VARCHAR(512) NOT NULL DEFAULT ' ',
			userName    VARCHAR(512) NOT NULL DEFAULT ' ',
			path     	VARCHAR(512) NOT NULL DEFAULT ' ',
			PRIMARY KEY (userId)
		)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4`,
		`INSERT INTO %s (userId, userName,path) VALUES (?,?,?)`,
		`SELECT * FROM %s`,
		`SELECT * FROM %s WHERE userId = ? LIMIT 1 LOCK IN SHARE MODE`,
		`DELETE FROM %s WHERE userId = ? LIMIT 1`,
		`UPDATE %s SET userName = ?, path = ? WHERE userId = ?`,
	}
)

// CreateTable -
func CreateTable(db *sql.DB, tableName string) error {
	s := fmt.Sprintf(UserSQLString[mysqlUserCreateTable], tableName)
	_, err := db.Exec(s)
	return err
}

// InsertUser return id
func InsertUser(db *sql.DB, tableName, userid, userName, path string) (int, error) {
	s := fmt.Sprintf(UserSQLString[mysqlUserInsert], tableName)
	result, err := db.Exec(s, userid, userName, path)
	if err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			// 1146 means duplicate key. Here is userid.
			if mysqlError.Number == 1146 {
				// update user info
				log.Println("update user info by", userid, userName)
				s = fmt.Sprintf(UserSQLString[mysqlUpdateUserInfoByID], tableName)
				result, err = db.Exec(s, userName, path, userid)
				if err != nil {
					return 0, err
				}

				return 0, nil
			}
		}
		return 0, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return 0, errInvalidInsert
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(userID), nil
}

// listUser return User list
func ListUser(db *sql.DB, tableName string) ([]*User, error) {
	var (
		users []*User

		userID   string
		userName string
		path     string
	)

	s := fmt.Sprintf(UserSQLString[mysqlUserList], tableName)
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&userID, &userName, &path); err != nil {
			return nil, err
		}

		user := &User{
			UserID:   userID,
			UserName: userName,
			Path:     path,
		}

		users = append(users, user)
	}

	return users, nil
}

// InfoByID query by id
func InfoByID(db *sql.DB, tableName string, userId string) (*User, error) {
	var User User

	s := fmt.Sprintf(UserSQLString[mysqlUserInfoByID], tableName)
	rows, err := db.Query(s, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&User.UserID, &User.UserName, &User.Path); err != nil {
			return nil, err
		}
	}

	return &User, nil
}

// DeleteByID delete by id
func DeleteByID(db *sql.DB, tableName string, userId string) error {
	s := fmt.Sprintf(UserSQLString[mysqlUserDeleteByID], tableName)
	_, err := db.Exec(s, userId)
	return err
}
