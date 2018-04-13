package RestfulRouter

import (
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestGetControllerValidName(t *testing.T) {
	tc := TestController{}
	name, err := getControllerValidName(&tc)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, "Test", name)
}

type TestController struct {
}

func (t *TestController) Get(c *gin.Context) {

}
func (t *TestController) Post(c *gin.Context) {

}
func (t *TestController) Put(c *gin.Context) {

}
func (t *TestController) Delete(c *gin.Context) {

}
func (t *TestController) Patch(c *gin.Context) {

}
func (t *TestController) Head(c *gin.Context) {

}
func (t *TestController) Options(c *gin.Context) {

}

func (t *TestController) Mapping(ginEngine *gin.Engine) {

}