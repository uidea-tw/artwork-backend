package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/uidea/artwork-backend/docs"
	v1 "github.com/uidea/artwork-backend/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	docs.SwaggerInfo.BasePath = "/api/v1"

	admin := v1.NewAdmin()
	user := v1.NewUser()
	article := v1.NewArticle()
	artwork := v1.NewArtwork()

	apiv1 := r.Group("/api/v1")
	{
		v1AdminsGroup := apiv1.Group("/admins")
		{
			v1AuthGroup := v1AdminsGroup.Group("/auth")
			{
				v1AuthGroup.POST("/login", admin.Login)
				v1AuthGroup.POST("/logout", admin.Logout)
			}

			v1AdminsGroup.POST("/", admin.Create)
		}

		v1UsersGroup := apiv1.Group("/users")
		{
			v1AuthGroup := v1UsersGroup.Group("/auth")
			{
				v1AuthGroup.POST("/login", user.Login)
				v1AuthGroup.POST("/logout", user.Logout)
				v1AuthGroup.POST("/signup", user.Signup)
			}

			v1UsersGroup.GET("/:id", user.Get)
			v1UsersGroup.GET("/", user.List)
			v1UsersGroup.PUT("/:id", user.Update)
			v1UsersGroup.DELETE("/:id", user.Delete)
		}

		v1ArticlesGroup := apiv1.Group("/articles")
		{
			v1ArticlesGroup.GET("/:id", article.Get)
			v1ArticlesGroup.GET("/", article.List)
			v1ArticlesGroup.POST("/", article.Create)
			v1ArticlesGroup.DELETE("/:id", article.Delete)
			v1ArticlesGroup.PUT("/:id", article.Update)
			v1ArticlesGroup.PATCH("/:id/state", article.Update)
		}

		v1ArtworksGroup := apiv1.Group("/artworks")
		{
			v1ArtworksGroup.GET("/:id", artwork.Get)
			v1ArtworksGroup.GET("/", artwork.List)
			v1ArtworksGroup.POST("/", artwork.Create)
			v1ArtworksGroup.DELETE("/:id", artwork.Delete)
			v1ArtworksGroup.PUT("/:id", artwork.Update)
			v1ArtworksGroup.PATCH("/:id/state", artwork.Update)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
