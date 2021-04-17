/*
 * Revision History:
 *     Initial: 2020/11/24      oiar
 */

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type Integration struct {
	IntegrationID uint64
	UserID        string
	ProjectID     uint64
	HelpUserID    string
	BoostPoints   uint64
}

type Rankings struct {
	Number   uint64
	UserId   string
	UserName string
	Path     string
	Points   uint64
}

const (
	mysqlIntegrationCreateTable = iota
	mysqlIntegrationInsert
	mysqlIntegrationList
	mysqlIntegrationInfoByID
	mysqlIntegrationDeleteByID
	mysqlGetSumPointsByUserAndProjectID
	mysqlDuplicateCheck
	mysqlGetRankings
	mysqlGetPoints
)

var (
	errInvalidInsert = errors.New("insert Integration:insert affected 0 rows")

	IntegrationSQLString = []string{
		`CREATE TABLE IF NOT EXISTS %s (
			integrationId BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			userId      VARCHAR(512) NOT NULL DEFAULT ' ',
		    projectId      BIGINT UNSIGNED NOT NULL DEFAULT 0,
			helpUserId      VARCHAR(512) NOT NULL DEFAULT ' ',
			boostPoints     BIGINT UNSIGNED NOT NULL DEFAULT 0,
			PRIMARY KEY (integrationId)
		)ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4`,
		`INSERT INTO %s (userId,projectId,helpUserId,boostPoints) VALUES (?,?,?,?)`,
		`SELECT * FROM %s`,
		`SELECT * FROM %s WHERE integrationId = ? LIMIT 1 LOCK IN SHARE MODE`,
		`DELETE FROM %s WHERE integrationId = ? LIMIT 1`,
		`SELECT COALESCE(SUM(boostPoints),0) FROM %s WHERE userId = ? AND projectId = ?`,
		`SELECT COUNT(*) FROM %s WHERE userId = ? AND projectId = ? AND helpUserId = ?`,
		`SELECT  (@a :=@a + 1) as a,T1.userId, T1.userName, T1.path, T1.points FROM(SELECT user.userId, user.userName, user.path, SUM(boostPoints) as points FROM integration INNER JOIN user ON integration.userId=user.userId GROUP BY integration.userId ORDER BY SUM(boostPoints) DESC) as T1, (SELECT @a := 0) as R`,
		`SELECT addPoints FROM project WHERE projectId = %d`,
	}
)

// CreateTable -
func CreateTable(db *sql.DB, tableName string) error {
	s := fmt.Sprintf(IntegrationSQLString[mysqlIntegrationCreateTable], tableName)
	_, err := db.Exec(s)
	return err
}

// GetRankings get rankings list.
func GetRankings(db *sql.DB) ([]*Rankings, error) {
	var (
		rankings []*Rankings

		number    uint64
		userId    string
		userName  string
		imagePath string
		points    uint64
	)

	s := fmt.Sprintf(IntegrationSQLString[mysqlGetRankings])
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&number, &userId, &userName, &imagePath, &points); err != nil {
			return nil, err
		}

		ranking := &Rankings{
			Number:   number,
			UserId:   userId,
			UserName: userName,
			Path:     imagePath,
			Points:   points,
		}

		rankings = append(rankings, ranking)
	}

	return rankings, nil
}

// GetSumPointsByUserAndProjectID query sum points.
func GetSumPointsByUserAndProjectID(db *sql.DB, tableName string, userId string, projectId uint64) (*uint64, error) {
	var i uint64

	s := fmt.Sprintf(IntegrationSQLString[mysqlGetSumPointsByUserAndProjectID], tableName)
	rows, err := db.Query(s, userId, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&i); err != nil {
			return nil, err
		}
	}

	return &i, nil
}

// DuplicateCheck check if duplicate. If true, duplicate help.
func DuplicateCheck(db *sql.DB, tableName string, userId string, projectId uint64, helpUserId string) (bool, error) {
	var i uint64

	s := fmt.Sprintf(IntegrationSQLString[mysqlDuplicateCheck], tableName)
	rows, err := db.Query(s, userId, projectId, helpUserId)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&i); err != nil {
			return false, err
		}
	}

	if i > 0 {
		return true, nil
	}

	return false, nil
}

// InsertIntegration return id
func InsertIntegration(db *sql.DB, tableName string, helpUserId, userId string, projectId uint64) (int, error) {
	pointsFmt := fmt.Sprintf(IntegrationSQLString[mysqlGetPoints], projectId)
	var i int
	rows, err := db.Query(pointsFmt)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&i); err != nil {
			return 0, err
		}
	}

	s := fmt.Sprintf(IntegrationSQLString[mysqlIntegrationInsert], tableName)

	result, err := db.Exec(s, userId, projectId, helpUserId, i)
	if err != nil {
		return 0, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return 0, errInvalidInsert
	}

	integrationId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(integrationId), nil
}

// listIntegration return Integration list
func ListIntegration(db *sql.DB, tableName string) ([]*Integration, error) {
	var (
		integrations []*Integration

		integrationId uint64
		userId        string
		projectId     uint64
		helpUserId    string
		boostPoints   uint64
	)

	s := fmt.Sprintf(IntegrationSQLString[mysqlIntegrationList], tableName)
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&integrationId, &userId, &projectId, &helpUserId, &boostPoints); err != nil {
			return nil, err
		}

		integration := &Integration{
			IntegrationID: integrationId,
			UserID:        userId,
			ProjectID:     projectId,
			HelpUserID:    helpUserId,
			BoostPoints:   boostPoints,
		}

		integrations = append(integrations, integration)
	}

	return integrations, nil
}

// InfoByID query by id
func InfoByID(db *sql.DB, tableName string, integrationId uint64) (*Integration, error) {
	var Integration Integration

	s := fmt.Sprintf(IntegrationSQLString[mysqlIntegrationInfoByID], tableName)
	rows, err := db.Query(s, integrationId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&Integration.IntegrationID, &Integration.UserID, &Integration.ProjectID, &Integration.HelpUserID, &Integration.BoostPoints); err != nil {
			return nil, err
		}
	}

	return &Integration, nil
}

// DeleteByID delete by id
func DeleteByID(db *sql.DB, tableName string, IntegrationId int) error {
	s := fmt.Sprintf(IntegrationSQLString[mysqlIntegrationDeleteByID], tableName)
	_, err := db.Exec(s, IntegrationId)
	return err
}
