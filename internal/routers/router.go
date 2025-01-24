package routers

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uidea/artwork-backend/docs"
	"github.com/uidea/artwork-backend/internal/middleware"
	v1 "github.com/uidea/artwork-backend/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	r.Use(cors.New(config))

	docs.SwaggerInfo.BasePath = "/api/v1"

	admin := v1.NewAdmin()
	user := v1.NewUser()
	article := v1.NewArticle()
	artwork := v1.NewArtwork()

	apiv1 := r.Group("/api/v1")
	{
		v1AdminsGroup := apiv1.Group("/admins")
		v1AdminsGroup.Use(middleware.JWT())
		{
			v1AuthGroup := v1AdminsGroup.Group("/auth")
			{
				v1AuthGroup.POST("/login", admin.Login)
				v1AuthGroup.POST("/logout", admin.Logout)
			}

			v1AdminsGroup.POST("/", admin.Create)
			v1AdminsGroup.GET("/profile", admin.Get)

			v1UsersGroup := v1AdminsGroup.Group("/users")
			{
				v1UsersGroup.POST("/", user.Create)
				v1UsersGroup.GET("/:id", user.Get)
				v1UsersGroup.GET("/", user.List)
				v1UsersGroup.PUT("/:id", user.Update)
				v1UsersGroup.DELETE("/:id", user.Delete)
			}

			v1ArticlesGroup := v1AdminsGroup.Group("/articles")
			{
				v1ArticlesGroup.GET("/:id", article.Get)
				v1ArticlesGroup.GET("/", article.List)
				v1ArticlesGroup.POST("/", article.Create)
				v1ArticlesGroup.DELETE("/:id", article.Delete)
				v1ArticlesGroup.PUT("/:id", article.Update)
				v1ArticlesGroup.PATCH("/:id/state", article.Update)
			}

			v1ArtworksGroup := v1AdminsGroup.Group("/artworks")
			{
				v1ArtworksGroup.GET("/:id", artwork.Get)
				v1ArtworksGroup.GET("/", artwork.List)
				v1ArtworksGroup.POST("/", artwork.Create)
				v1ArtworksGroup.DELETE("/:id", artwork.Delete)
				v1ArtworksGroup.PUT("/:id", artwork.Update)
			}

			v1AboutGroup := v1AdminsGroup.Group("/abouts")
			{
				v1AboutGroup.GET("/:id", artwork.Get)
				v1AboutGroup.GET("/", artwork.List)
				v1AboutGroup.POST("/", artwork.Create)
				v1AboutGroup.DELETE("/:id", artwork.Delete)
				v1AboutGroup.PUT("/:id", artwork.Update)
			}
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
