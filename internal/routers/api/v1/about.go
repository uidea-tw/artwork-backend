package v1

import (
	"github.com/gin-gonic/gin"
)

type About struct{}

func NewAbout() About {
	return About{}
}

func (t About) Get(c *gin.Context) {}

func (t About) List(c *gin.Context) {}

func (t About) Create(c *gin.Context) {}

func (t About) Update(c *gin.Context) {}

func (t About) Delete(c *gin.Context) {}
