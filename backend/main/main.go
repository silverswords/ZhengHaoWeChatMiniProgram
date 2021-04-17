package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	integration "github.com/shyvana/integration/controller/gin"
	project "github.com/shyvana/project/controller/gin"
	user "github.com/shyvana/user/controller/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())
	dbConn, err := sql.Open("mysql", "root:ZhenghaoAdmin@tcp(139.199.76.73:3307)/test?parseTime=true")
	if err != nil {
		panic(err)
	}

	u := &User{username: "zhenghao", password: "zhenghao888"}
	router.POST("/api/v1/login", u.Login)

	d := &Display{Display: "false"}
	router.GET("/api/v1/setpopup/:set", d.SetPopup)
	router.GET("/api/v1/getpopup", d.GetPopup)

	userCon := user.New(dbConn, "user")
	userCon.RegisterRouter(router.Group("/api/v1/user"))

	projectCon := project.New(dbConn, "project")
	projectCon.RegisterRouter(router.Group("/api/v1/project"))

	integrationCon := integration.New(dbConn, "integration")
	integrationCon.RegisterRouter(router.Group("/api/v1/integration"))

	log.Fatal(router.Run(":8111"))
}

type User struct {
	username string
	password string
}

func (u *User) Login(c *gin.Context) {
	var (
		req struct {
			UserName string `json:"name"      binding:"required"`
			Password string `json:"pwd"       binding:"required"`
		}
	)

	err := c.ShouldBind(&req)
	if err != nil {
		log.Println(err)
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if req.UserName != u.username || req.Password != u.password {
		log.Println("用户名或密码错误")
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "error": errors.New("用户名或密码错误")})
	}

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

type Display struct {
	Display string
}

func (d *Display) SetPopup(c *gin.Context) {
	set := c.Params.ByName("set")
	d.Display = set

	c.JSON(http.StatusOK, gin.H{"set": set})

}

func (d *Display) GetPopup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"set": d.Display})
}
