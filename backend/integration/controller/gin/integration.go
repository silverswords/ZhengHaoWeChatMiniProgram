/*
 * Revision History:
 *     Initial: 2020/11/24      oiar
 */

package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shyvana/integration/model/mysql"
)

var (
	errDuplicateHelper = errors.New("您已经助力过了,不能再助力了哦")
	errSelfHelp        = errors.New("您不能助力您自己")
	SuccessHelp        = "助力成功"
)

// IntegrationController -
type IntegrationController struct {
	db        *sql.DB
	tableName string
}

// New -
func New(db *sql.DB, tableName string) *IntegrationController {
	return &IntegrationController{
		db:        db,
		tableName: tableName,
	}
}

// RegisterRouter -
func (b *IntegrationController) RegisterRouter(r gin.IRouter) {
	if r == nil {
		log.Fatal("[InitRouter]: server is nil")
	}

	err := mysql.CreateTable(b.db, b.tableName)
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/create", b.create)
	r.POST("/check", b.Check)
	r.POST("/delete", b.deleteByID)
	r.POST("/info/id", b.infoByID)
	r.POST("/list", b.listIntegration)
	r.POST("/points", b.pointsByUserIDAndProjectID)
	r.POST("/rankings", b.rankings)
}

func (b *IntegrationController) Check(c *gin.Context) {
	var (
		req struct {
			UserID      string `json:"userId"      binding:"required"`
			ProjectID   uint64 `json:"projectId"       binding:"required"`
			HelpUserID  string `json:"helpUserId"         binding:"required"`
			BoostPoints uint64 `json:"boostPoints"       binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if req.UserID == req.HelpUserID {
		c.Error(errSelfHelp)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": errSelfHelp.Error(), "check": "true"})
		return
	}

	duplicate, err := mysql.DuplicateCheck(b.db, b.tableName, req.UserID, req.ProjectID, req.HelpUserID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "err": err.Error(), "check": "true"})
		return
	}
	if duplicate {
		c.Error(errDuplicateHelper)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "err": errDuplicateHelper.Error(), "check": "true"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "check": "false"})
}

func (b *IntegrationController) create(c *gin.Context) {
	var (
		req struct {
			UserID     string `json:"userId"      binding:"required"`
			ProjectID  uint64 `json:"projectId"       binding:"required"`
			HelpUserID string `json:"helpUserId"         binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if req.UserID == req.HelpUserID {
		c.Error(errSelfHelp)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "err": errSelfHelp.Error()})
		return
	}

	duplicate, err := mysql.DuplicateCheck(b.db, b.tableName, req.UserID, req.ProjectID, req.HelpUserID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "err": err.Error()})
		return
	}
	if duplicate {
		c.Error(errDuplicateHelper)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "err": errDuplicateHelper.Error()})
		return
	}

	_, err = mysql.InsertIntegration(b.db, b.tableName, req.HelpUserID, req.UserID, req.ProjectID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway, "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "msg": SuccessHelp})
}

func (b *IntegrationController) pointsByUserIDAndProjectID(c *gin.Context) {
	var (
		req struct {
			UserId    string `json:"userId"      binding:"required"`
			ProjectId uint64 `json:"projectId"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)

	if err != nil {
		c.Error(err)
		log.Println("错误内容", c.PostForm("projectId"))
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	points, err := mysql.GetSumPointsByUserAndProjectID(b.db, b.tableName, req.UserId, req.ProjectId)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "points": points})
}

func (b *IntegrationController) rankings(c *gin.Context) {
	rankings, err := mysql.GetRankings(b.db)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "rankings": rankings})
}

func (b *IntegrationController) listIntegration(c *gin.Context) {
	integrations, err := mysql.ListIntegration(b.db, b.tableName)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "Integrations": integrations})
}

func (b *IntegrationController) infoByID(c *gin.Context) {
	var (
		req struct {
			IntegrationID uint64 `json:"integrationId"     binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	integration, err := mysql.InfoByID(b.db, b.tableName, req.IntegrationID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "ban": integration})
}

func (b *IntegrationController) deleteByID(c *gin.Context) {
	var (
		req struct {
			IntegrationID int `json:"integrationId"    binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = mysql.DeleteByID(b.db, b.tableName, req.IntegrationID)
	if err != nil {
		c.Error(err)
		c.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
