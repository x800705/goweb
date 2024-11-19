package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Test struct {
	Id      uint `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	dsn := "pro:trHT4XTrFiJEpTjC@tcp(127.0.0.1:3306)/pro?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			//禁止复数
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
		if err != nil {
			panic("failed to connect database")
		}
		// 自动迁移
	   //db.AutoMigrate(&User{})
    //查询所有
	r.GET("/", func(c *gin.Context) {
		var test []Test

		//db.Find(&test) // 根据整型主键查找
		sql := "select * from test"
		db.Raw(sql).Scan(&test)

		c.JSON(http.StatusOK, test)
	})
	//查询单个
	r.GET("/findOne", func(c *gin.Context) {
	    a := c.Query("a")

	    var test[] Test

		//db.Find(&test) // 根据整型主键查找
		sql := "select * from test where id = ?" 
   
		//db.Where("id = ?",a).Find(&test)
		db.Raw(sql,a).Scan(&test)

		c.JSON(http.StatusOK, test)
	   
	})
	//新增单个
    r.POST("/addOne", func(c *gin.Context) {
        var test Test
   
        // 解析JSON
        if err := c.ShouldBindJSON(&test); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
   
        // 新增
        if err := db.Create(&test).Error; err != nil {
         
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add the element: " + err.Error()})
            return
        }
   
        // 如果没有错误，说明创建成功
        c.JSON(200, gin.H{"message": "Element added successfully"})
    })
   
   
    //删除单个
	r.POST("/deleteOne", func(c *gin.Context) {
	  id := 12354
	  db.Delete(&Test{},id)
	 
	  c.JSON(200,"success")
	   
	})
	//更改单个
	r.POST("/changeOne", func(c *gin.Context) {
	     var test Test
	     
	     
	    // 解析JSON
        if err := c.ShouldBindJSON(&test); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
       
        db.Model(&test).Updates(test)
       
        c.JSON(200,"ok")
       
	   
	})
	r.GET("/add", func(c *gin.Context) {
	    c.JSON(200,"Hello,World")
	   
	})

	r.GET("/test", func(c *gin.Context) {
		a := c.Query("a")

		fmt.Println(a)

		dsn := "pro:trHT4XTrFiJEpTjC@tcp(127.0.0.1:3306)/pro?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
			//禁止复数
			NamingStrategy: schema.NamingStrategy{
				SingularTable: false,
			},
		})
		if err != nil {
			panic("failed to connect database")
		}

		var test []Test

		db.Where("id = ?", a).Find(&test)

		fmt.Println(test)

		c.String(http.StatusOK, "hello World!")

	})

	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}
