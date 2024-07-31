package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend-tugas-reactjs/controllers"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB, r *gin.Engine) {

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization"}

	// To be able to send tokens to the server.
	corsConfig.AllowCredentials = true
	// OPTIONS method for ReactJS
	corsConfig.AddAllowMethods("OPTIONS")

	r.Use(cors.New(corsConfig))

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/books", controllers.GetAllBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.GetBookById)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook)

	// MataKuliah routes
	r.POST("/mata_kuliah", controllers.CreateMataKuliah)
	r.GET("/mata_kuliah", controllers.GetMataKuliah)
	r.PATCH("/mata_kuliah/:id", controllers.UpdateMataKuliah)
	r.DELETE("/mata_kuliah/:id", controllers.DeleteMataKuliah)

	// Dosen routes
	r.POST("/dosen", controllers.CreateDosen)
	r.GET("/dosen", controllers.GetDosen)
	r.PATCH("/dosen/:id", controllers.UpdateDosen)
	r.DELETE("/dosen/:id", controllers.DeleteDosen)

	// Mahasiswa routes
	r.POST("/mahasiswa", controllers.CreateMahasiswa)
	r.GET("/mahasiswa", controllers.GetMahasiswa)
	r.PATCH("/mahasiswa/:id", controllers.UpdateMahasiswa)
	r.DELETE("/mahasiswa/:id", controllers.DeleteMahasiswa)

	// Nilai routes
	r.POST("/nilai", controllers.CreateNilai)
	r.GET("/nilai", controllers.GetNilai)
	r.PATCH("/nilai/:id", controllers.UpdateNilai)
	r.DELETE("/nilai/:id", controllers.DeleteNilai)

	// JadwalKuliah routes
	r.POST("/jadwal_kuliah", controllers.CreateJadwalKuliah)
	r.GET("/jadwal_kuliah", controllers.GetJadwalKuliah)
	r.PATCH("/jadwal_kuliah/:id", controllers.UpdateJadwalKuliah)
	r.DELETE("/jadwal_kuliah/:id", controllers.DeleteJadwalKuliah)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
