/*
 * Revision History:
 *     Initial: 2020/11/23      oiar
 */

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
)

type Project struct {
	ProjectID    uint64
	ProjectName  string
	Introduction string
	Rule         string
	PathOne      string
	PathTwo      string
	PathThree    string
	PathFour     string
	PathFive     string
	PathSix      string
	PathSeven    string
	PathEight    string
	PathNine     string
	QRCode       string
	AddPoints    uint64
}

const (
	mysqlProjectCreateTable = iota
	mysqlProjectInsert
	mysqlProjectList
	mysqlProjectInfoByID
	mysqlProjectDeleteByID
	mysqlProjectUpdateByID
)

var (
	errInvalidInsert = errors.New("insert project:insert affected 0 rows")

	projectSQLString = []string{
		`CREATE TABLE IF NOT EXISTS %s (
			projectId      BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
			projectName    VARCHAR(512) UNIQUE DEFAULT NULL DEFAULT ' ',
			introduction   VARCHAR(10240) NOT NULL DEFAULT ' ',
			rule           VARCHAR(512) NOT NULL DEFAULT ' ',
			pathOne     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathTwo     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathThree     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathFour     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathFive     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathSix     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathSeven     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathEight     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			pathNine     	   VARCHAR(256) NOT NULL DEFAULT ' ',
			qrCode     	   VARCHAR(256) NOT NULL DEFAULT ' ',
		    addPoints      BIGINT UNSIGNED NOT NULL DEFAULT 0,
			PRIMARY KEY (projectId)
		)ENGINE=InnoDB AUTO_INCREMENT=1000000 DEFAULT CHARSET=utf8mb4`,
		`INSERT INTO %s (projectName,introduction,rule,pathOne,pathTwo,pathThree,pathFour,pathFive,pathSix,pathSeven,pathEight,pathNine,qrCode,addPoints) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		`SELECT * FROM %s`,
		`SELECT * FROM %s WHERE projectId = ? LIMIT 1 LOCK IN SHARE MODE`,
		`DELETE FROM %s WHERE projectId = ? LIMIT 1`,
		`UPDATE %s SET projectName = ?,introduction = ?,rule = ?,pathOne = ?,pathTwo = ?,pathThree = ?,pathFour = ?,pathFive = ?,pathSix = ?,pathSeven = ?,pathEight = ?,pathNine = ?,qrCode = ?,addPoints = ? WHERE projectId = ? LIMIT 1`,
	}
)

// CreateTable -
func CreateTable(db *sql.DB, tableName string) error {
	s := fmt.Sprintf(projectSQLString[mysqlProjectCreateTable], tableName)
	_, err := db.Exec(s)
	return err
}

// InsertProject return id
func InsertProject(db *sql.DB, tableName, projectName, introduction, rule, pathOne, pathTwo, pathThree, pathFour, pathFive, pathSix, pathSeven, pathEight, pathNine, qrCode string, addPoints uint64) (int, error) {
	s := fmt.Sprintf(projectSQLString[mysqlProjectInsert], tableName)
	result, err := db.Exec(s, projectName, introduction, rule, pathOne, pathTwo, pathThree, pathFour, pathFive, pathSix, pathSeven, pathEight, pathNine, qrCode, addPoints)
	if err != nil {
		return 0, err
	}

	if rows, _ := result.RowsAffected(); rows == 0 {
		return 0, errInvalidInsert
	}

	projectID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(projectID), nil
}

// listProject return project list
func ListProject(db *sql.DB, tableName string) ([]*Project, error) {
	var (
		projects []*Project

		projectID    uint64
		projectName  string
		introduction string
		rule         string
		pathOne      string
		pathTwo      string
		pathThree    string
		pathFour     string
		pathFive     string
		pathSix      string
		pathSeven    string
		pathEight    string
		pathNine     string
		qrCode       string
		addPoints    uint64
	)

	s := fmt.Sprintf(projectSQLString[mysqlProjectList], tableName)
	rows, err := db.Query(s)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&projectID, &projectName, &introduction, &rule, &pathOne, &pathTwo, &pathThree, &pathFour, &pathFive, &pathSix, &pathSeven, &pathEight, &pathNine, &qrCode, &addPoints); err != nil {
			return nil, err
		}

		project := &Project{
			ProjectID:    projectID,
			ProjectName:  projectName,
			Introduction: introduction,
			Rule:         rule,
			PathOne:      pathOne,
			PathTwo:      pathTwo,
			PathThree:    pathThree,
			PathFour:     pathFour,
			PathFive:     pathFive,
			PathSix:      pathSix,
			PathSeven:    pathSeven,
			PathEight:    pathEight,
			PathNine:     pathNine,
			QRCode:       qrCode,
			AddPoints:    addPoints,
		}

		projects = append(projects, project)
	}

	return projects, nil
}

// InfoByID query by id
func InfoByID(db *sql.DB, tableName string, projectId uint64) (*Project, error) {
	var project Project

	s := fmt.Sprintf(projectSQLString[mysqlProjectInfoByID], tableName)
	rows, err := db.Query(s, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&project.ProjectID, &project.ProjectName, &project.Introduction, &project.Rule, &project.PathOne, &project.PathTwo, &project.PathThree, &project.PathFour, &project.PathFive, &project.PathSix, &project.PathSeven, &project.PathEight, &project.PathNine, &project.QRCode, &project.AddPoints); err != nil {
			return nil, err
		}
	}

	return &project, nil
}

// DeleteByID delete by id
func DeleteByID(db *sql.DB, tableName string, ProjectId int) error {
	s := fmt.Sprintf(projectSQLString[mysqlProjectDeleteByID], tableName)
	_, err := db.Exec(s, ProjectId)
	return err
}

// UpdateByID update by id
func UpdateByID(db *sql.DB, tableName, projectName, introduction, rule, pathOne, pathTwo, pathThree, pathFour, pathFive, pathSix, pathSeven, pathEight, pathNine, qrCode string, addPoints uint64, ProjectId int) error {
	s := fmt.Sprintf(projectSQLString[mysqlProjectUpdateByID], tableName)
	_, err := db.Exec(s, projectName, introduction, rule, pathOne, pathTwo, pathThree, pathFour, pathFive, pathSix, pathSeven, pathEight, pathNine, qrCode, addPoints, ProjectId)
	return err
}
