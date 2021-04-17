/*
 * Revision History:
 *     Initial: 2020/11/23      oiar
 */

package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/shyvana/user/model/mysql"
	"log"
	"net/http"
)

// userController -
type userController struct {
	db        *sql.DB
	tableName string
}

// New -
func New(db *sql.DB, tableName string) *userController {
	return &userController{
		db:        db,
		tableName: tableName,
	}
}

// RegisterRouter -
func (b *userController) RegisterRouter(r gin.IRouter) {
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
	r.POST("/list", b.listUser)
}

func (b *userController) create(c *gin.Context) {
	var (
		req struct {
			UserID   string `json:"userId" binding:"required"`
			UserName string `json:"userName"      binding:"required"`
			Path     string `json:"path"       binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	id, err := mysql.InsertUser(b.db, b.tableName, req.UserID, req.UserName, req.Path)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ID": id})
}

func (b *userController) listUser(c *gin.Context) {
	users, err := mysql.ListUser(b.db, b.tableName)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "users": users})
}

func (b *userController) infoByID(c *gin.Context) {
	var (
		req struct {
			UserID string `json:"userId"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	user, err := mysql.InfoByID(b.db, b.tableName, req.UserID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ban": user})
}

func (b *userController) deleteByID(c *gin.Context) {
	var (
		req struct {
			UserID string `json:"userId"    binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.DeleteByID(b.db, b.tableName, req.UserID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
