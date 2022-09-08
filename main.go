package main

import (
	"gorm-curd-basic/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "user=postgres password=1 dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	r := gin.Default()
	// create
	r.POST("create", func(ctx *gin.Context) {
		var post entity.Post
		ctx.ShouldBindJSON(&post)
		if err := db.Table("posts").Create(&post).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": post,
			})
		}

	})
	// findAll
	r.GET("/", func(ctx *gin.Context) {

		var post []entity.Post
		if err := db.Table("posts").Find(&post).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": post,
			})
		}
	})
	// findById
	r.GET("/:id", func(ctx *gin.Context) {
		id, err1 := strconv.Atoi(ctx.Param("id"))
		if err1 != nil {
			ctx.JSON(200, gin.H{
				"message": "can't get string to id",
			})
			return
		}

		var post entity.Post
		if err := db.Table("posts").First(&post, id).Error; err != nil {
			ctx.JSON(200, gin.H{
				"err": err.Error(),
			})

		} else {
			ctx.JSON(200, gin.H{
				"data": post,
			})
		}

	})
	//updater
	r.PUT("update/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "can't get string to id",
			})
			return
		}
		var post entity.Post
		if err1 := db.Table("posts").First(&post, id).Error; err1 != nil {
			ctx.JSON(200, gin.H{
				"err": err1.Error(),
			})

		} else {
			ctx.BindJSON(&post)
			db.Save(&post)
			ctx.JSON(200, gin.H{
				"data": post,
			})
		}

	})
	r.DELETE("delete/:id", func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "can't get string to id",
			})
			return
		}
		var post entity.Post
		if err1 := db.Table("posts").Delete(&post, id); err != nil {
			ctx.JSON(200, gin.H{
				"err": err1.Error,
			})
		} else {
			ctx.JSON(200, gin.H{
				"message": "success",
			})
		}
	})

	r.Run(":8080")
}
