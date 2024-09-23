package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	//"github.com/gin-gonic/gin"
	"log"
	//"os"
	"fmt"
	//"hanoi/rabbitmq"
	"hanoi/route"
	"hanoi/database"
	"hanoi/models"
	"gorm.io/gorm"
	"time"
	//"github.com/swaggo/gin-swagger"
	"github.com/swaggo/fiber-swagger"
	//"github.com/swaggo/files"
	 _ "hanoi/docs" // สำหรับเอกสาร Swagger

	
)


// func loadDatabase() {
// 	if err := database.Connect(); err != nil {
// 		handleError(err)
// 	}

// }

// func DropTable () {

// 	database.Database.Migrator().DropTable(&models.TransactionSub{})
// 	database.Database.Migrator().DropTable(&models.BuyInOut{})

// }

func migrateNormal(db *gorm.DB) {

	if err := db.AutoMigrate(&models.Product{},&models.BanksAccount{},&models.Users{},&models.TransactionSub{},&models.BankStatement{},&models.BuyInOut{}); err != nil {
		handleError(err)
	}
	 
	fmt.Println("Migrations Normal Tables executed successfully")
}
// func migrateAdmin() {
 
// 	if err := database.Database.AutoMigrate(&models.TsxAdmin{},&models.Provider{}); err != nil {
// 		handleError(err)
// 	}
// 	fmt.Println("Migrations Admin Tables executed successfully")
// }



// @title Swagger Example API
// @version 1.0
// @description This is a sample server for Swagger with GORM and MySQL in Go.
// @host localhost:3003
// @BasePath /api/v1


// @Summary Get all users
// @Description Get a list of all users in the database.
// @Tags users
// @Produce json
// @Success 200 {array} models.SwaggerUser
// @Router /users [get]
func GetUsers(c *fiber.Ctx) error {
    var users []models.Users
	db, _ := database.ConnectToDB("hanoi")
    db.Find(&users)
    return c.JSON(users)
}


func handleError(err error) {
	log.Fatal(err)
}

func main() {

	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		loc, _ := time.LoadLocation("Asia/Bangkok")
		c.Locals("location", loc)
		return c.Next()
	})
	db, _ := database.ConnectToDB("hanoi")

	//rabbitmq.Init()
	
	// loadDatabase()
	 //DropTable()
	 migrateNormal(db)
	 //migrateAdmin()

	app.Use(logger.New())

	route.SetupRoutes(app)
	
	// เส้นทาง Swagger
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

	// เส้นทาง API สำหรับดึงข้อมูลผู้ใช้งาน
	app.Get("/api/users", GetUsers)

    // เรียกใช้ฟังก์ชันจาก efinity.go
	log.Fatal(app.Listen(":8030"))
	 
	
}