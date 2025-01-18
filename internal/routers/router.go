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
	// authors := v1.NewAuthor()
	// article := v1.NewArticle()

	apiv1 := r.Group("/api/v1")
	{
		v1AdminsGroup := apiv1.Group("/admins")
		{
			v1AdminsGroup.Group("/auth")
			{
				v1AdminsGroup.POST("/login", admin.Login)
				v1AdminsGroup.POST("/logout", admin.Logout)
			}

			v1AdminsGroup.POST("/", admin.Create)
		}

		// 	v1AuthorsGroup := apiv1.Group("/authors")
		// 	{
		// 		v1AuthorsGroup.Group("/auth")
		// 		{
		// 			v1AuthorsGroup.POST("/login", authors.Login)
		// 			v1AuthorsGroup.POST("/logout", authors.Logout)
		// 			v1AuthorsGroup.POST("/register", authors.Register)
		// 		}

		// 		v1AuthorsGroup.GET("/:id", authors.Get)
		// 		v1AuthorsGroup.GET("/", authors.List)
		// 		v1AuthorsGroup.POST("/", authors.Create)
		// 		v1AuthorsGroup.PUT("/:id", authors.Update)
		// 		v1AuthorsGroup.DELETE("/:id", authors.Delete)
		// 	}

		// 	v1ArticlesGroup := apiv1.Group("/articles")
		// 	{
		// 		v1ArticlesGroup.GET("/:id", article.Get)
		// 		v1ArticlesGroup.GET("/", article.List)
		// 		v1ArticlesGroup.POST("/", article.Create)
		// 		v1ArticlesGroup.DELETE("/:id", article.Delete)
		// 		v1ArticlesGroup.PUT("/:id", article.Update)
		// 		v1ArticlesGroup.PATCH("/:id/state", article.Update)
		// 	}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
