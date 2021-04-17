/*
 * Revision History:
 *     Initial: 2020/11/23      oiar
 */

package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/shyvana/project/model/mysql"
	"log"
	"net/http"
)

// ProjectController -
type ProjectController struct {
	db        *sql.DB
	tableName string
}

// New -
func New(db *sql.DB, tableName string) *ProjectController {
	return &ProjectController{
		db:        db,
		tableName: tableName,
	}
}

// RegisterRouter -
func (b *ProjectController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	err := mysql.CreateTable(b.db, b.tableName)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/create", b.create)
	r.POST("/delete", b.deleteByID)
	r.POST("/info/id", b.infoByID)
	r.POST("/list", b.listProject)
	r.POST("/update", b.updateByID)
}

func (b *ProjectController) create(c *gin.Context) {
	var (
		req struct {
			ProjectName  string `json:"projectName"      binding:"required"`
			Introduction string `json:"introduction"       binding:"required"`
			Rule         string `json:"rule"         binding:"required"`
			PathOne      string `json:"pathOne"       binding:"required"`
			PathTwo      string `json:"pathTwo"       binding:"required"`
			PathThree    string `json:"pathThree"       binding:"required"`
			PathFour     string `json:"pathFour" binding:"required"`
			PathFive     string `json:"pathFive"       binding:"required"`
			PathSix      string `json:"pathSix"       binding:"required"`
			PathSeven    string `json:"pathSeven"       binding:"required"`
			PathEight    string `json:"pathEight"       binding:"required"`
			PathNine     string `json:"pathNine"       binding:"required"`
			QRCode       string `json:"qrCode"       binding:"required"`
			AddPoints    uint64 `json:"addPoints"          binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.InsertProject(b.db, b.tableName, req.ProjectName, req.Introduction, req.Rule, req.PathOne, req.PathTwo, req.PathThree, req.PathFour, req.PathFive, req.PathSix, req.PathSeven, req.PathEight, req.PathNine, req.QRCode, req.AddPoints)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ID": id})
}

func (b *ProjectController) listProject(c *gin.Context) {
	Projects, err := mysql.ListProject(b.db, b.tableName)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Projects": Projects})
}

func (b *ProjectController) infoByID(c *gin.Context) {
	var (
		req struct {
			ProjectID uint64 `json:"ProjectId"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	Project, err := mysql.InfoByID(b.db, b.tableName, req.ProjectID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ban": Project})
}

func (b *ProjectController) deleteByID(c *gin.Context) {
	var (
		req struct {
			ProjectID int `json:"ProjectId"    binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.DeleteByID(b.db, b.tableName, req.ProjectID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (b *ProjectController) updateByID(c *gin.Context) {
	var (
		req struct {
			ProjectID    int    `json:"ProjectId"    binding:"required"`
			ProjectName  string `json:"projectName"      binding:"required"`
			Introduction string `json:"introduction"       binding:"required"`
			Rule         string `json:"rule"         binding:"required"`
			PathOne      string `json:"pathOne"       binding:"required"`
			PathTwo      string `json:"pathTwo"       binding:"required"`
			PathThree    string `json:"pathThree"       binding:"required"`
			PathFour     string `json:"pathFour" binding:"required"`
			PathFive     string `json:"pathFive"       binding:"required"`
			PathSix      string `json:"pathSix"       binding:"required"`
			PathSeven    string `json:"pathSeven"       binding:"required"`
			PathEight    string `json:"pathEight"       binding:"required"`
			PathNine     string `json:"pathNine"       binding:"required"`
			QRCode       string `json:"qrCode"       binding:"required"`
			AddPoints    uint64 `json:"addPoints"          binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.UpdateByID(b.db, b.tableName, req.ProjectName, req.Introduction, req.Rule, req.PathOne, req.PathTwo, req.PathThree, req.PathFour, req.PathFive, req.PathSix, req.PathSeven, req.PathEight, req.PathNine, req.QRCode, req.AddPoints, req.ProjectID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
