package restfulrouter

// import (
// 	"github.com/gin-gonic/gin"
// 	"net/http"
// )

// type TestController struct {
// 	restfulrouter.BaseController
// }

// func (tc *TestController) Mapping() map[string]restfulrouter.GinHandler {
// 	m := make(map[string]restfulrouter.GinHandler, 1)
// 	m[restfulrouter.ComposeCustomMappingKey(http.MethodGet, "customTest")] = CustomMethodTest
// 	return m
// }

// func CustomMethodTest(c *gin.Context) {
// 	c.String(http.StatusNotFound, "")
// }

// func ExampleRegisterAPIRoute() {
// 	r := gin.Default()
// 	restfulrouter.RegisterAPIRoute(r, []IBaseController{&TestController{}})
// }

// func ExampleRegisterGroupAPIRoute() {
// 	r := gin.Default()
// 	restfulrouter.RegisterGroupAPIRoute("/api", r, []restfulrouter.IBaseController{&TestController{}})
// }
