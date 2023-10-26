package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Usuario struct {
	gorm.Model
	Nombre string
	Email  string
}

func main() {
	r := gin.Default()

	// Configuración de la base de datos
	dsn := "root:@tcp(127.0.0.1:3306)/crud_app?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error al conectar a la base de datos")
	}
	db.AutoMigrate(&Usuario{})

	// Rutas para las operaciones CRUD
	r.GET("/users", ObtenerUsuarios(db))
	r.GET("/users/:id", ObtenerUsuario(db))
	r.POST("/users", CrearUsuario(db))
	r.PUT("/users/:id", ActtualizarUsuario(db))
	r.DELETE("/users/:id", BorrarUsuario(db))

	r.Run(":8080")
}

func InfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"nombre":    "Daniel Flores",
			"matricula": "200614",
			"grupo":     "IDGS 10-B",
		})
	}
}

func ErrorMiddleware(c *gin.Context, err error) {
	c.JSON(400, gin.H{"error": err.Error()})
}

// Manejadores para las rutas CRUD
func ObtenerUsuarios(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []Usuario
		db.Find(&users)
		c.JSON(200, gin.H{
			"data": users,
		})
	}
}

func ObtenerUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Usuario
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func CrearUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Usuario
		if err := c.ShouldBindJSON(&user); err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Create(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func ActtualizarUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Usuario
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		if err := c.ShouldBindJSON(&user); err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Save(&user)
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func BorrarUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user Usuario
		if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
			ErrorMiddleware(c, err)
			return
		}
		db.Delete(&user)
		c.JSON(200, gin.H{
			"message": "Usuario eliminado!",
		})
	}
}
