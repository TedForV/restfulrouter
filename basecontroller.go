package RestfulRouter

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

const (
	KEY_SEPERATOR = ":"
)

type IBaseController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
	Patch(c *gin.Context)
	Head(c *gin.Context)
	Options(c *gin.Context)

	//
	Mapping() map[string]GinHandler
}

type BaseController struct {
}

func (t *BaseController) Get(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Post(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Put(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Delete(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Patch(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Head(c *gin.Context) {
	returnNotResource(c)
	return
}
func (t *BaseController) Options(c *gin.Context) {
	returnNotResource(c)
	return
}

func (t *BaseController) Mapping() map[string]GinHandler {
	return nil
}

func returnNotResource(c *gin.Context) {
	c.String(http.StatusNotFound, "")
}

func analyseMappingKey(key string) (method string, pathName string, err error) {
	key = strings.TrimSpace(key)

	if len(key) < 5 {
		return "", "", errors.New("key has error.")
	}

	if i := strings.Index(key, KEY_SEPERATOR); i == -1 {
		return "", "", errors.New("key needs a comma.")
	} else {
		method = key[:i]
		if err != nil {
			return "", "", errors.New("key has error.")
		}

		pathName = strings.ToLower(strings.TrimSpace(key[i+1:]))
	}

	return method, pathName, nil
}

//method is http.MethodXXX
func ComposeCustomMappingKey(method string, path string) string {
	return method + KEY_SEPERATOR + path
}

type GinHandler func(c *gin.Context)

func RegisterAPIRoute(ginEngine *gin.Engine, controllers []IBaseController) {
	routesControllerMapping(ginEngine, controllers)
	//if controllers == nil || len(controllers) == 0 {
	//	return
	//}
	//for _, c := range controllers {
	//	cname, err := getControllerValidName(c)
	//	if err != nil {
	//		panic(err)
	//	}
	//	autoMapping(ginEngine, cname, c)
	//	err = autoCustomMapping(ginEngine, cname, c)
	//	if err != nil {
	//		panic(err)
	//	}
	//}

}

func RegisterGroupAPIRoute(basePath string, ginEngine *gin.Engine, controllers []IBaseController) {
	if !strings.HasPrefix(basePath, "/") {
		basePath = "/" + basePath
	}
	g := ginEngine.Group(basePath)
	{
		routesControllerMapping(g, controllers)
	}
}

func routesControllerMapping(router gin.IRouter, controllers []IBaseController) {
	if controllers == nil || len(controllers) == 0 {
		return
	}
	for _, c := range controllers {
		cname, err := getControllerValidName(c)
		if err != nil {
			panic(err)
		}
		autoMapping(router, cname, c)
		err = autoCustomMapping(router, cname, c)
		if err != nil {
			panic(err)
		}
	}
}

const (
	CONTROLLER_SUFFIX     = "Controller"
	ERROR_CONTROLLER_NAME = "Controller name must be suffix with 'Controller'"
)

func getControllerValidName(controller IBaseController) (string, error) {
	typeInfo := reflect.TypeOf(controller)
	fullName := typeInfo.Elem().String()
	lastDotIndex := strings.LastIndex(fullName, ".")
	fullName = fullName[lastDotIndex+1:]
	if strings.HasSuffix(fullName, CONTROLLER_SUFFIX) && len(fullName) > len(CONTROLLER_SUFFIX) {
		validName := fullName[0 : len(fullName)-len(CONTROLLER_SUFFIX)]
		return strings.ToLower(strings.TrimSpace(validName)), nil
	} else {
		return "", errors.New(ERROR_CONTROLLER_NAME)
	}

}

func autoMapping(router gin.IRouter, controllerName string, controller IBaseController) {
	path := "/" + controllerName
	router.GET(path, func(c *gin.Context) {
		controller.Get(c)
	})
	router.POST(path, func(c *gin.Context) {
		controller.Post(c)
	})
	router.PUT(path, func(c *gin.Context) {
		controller.Post(c)
	})
	router.DELETE(path, func(c *gin.Context) {
		controller.Post(c)
	})
	router.HEAD(path, func(c *gin.Context) {
		controller.Post(c)
	})
	router.OPTIONS(path, func(c *gin.Context) {
		controller.Post(c)
	})
	router.PATCH(path, func(c *gin.Context) {
		controller.Post(c)
	})
}

func autoCustomMapping(router gin.IRouter, controllerName string, controller IBaseController) error {
	route := controller.Mapping()

	for k, v := range route {
		method, path, err := analyseMappingKey(k)
		if err != nil {
			return err
		}
		fullPath := "/" + controllerName + "/" + path
		switch method {
		case http.MethodGet:
			func(handler GinHandler) {
				router.GET(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodPost:
			func(handler GinHandler) {
				router.POST(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodPut:
			func(handler GinHandler) {
				router.PUT(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodDelete:
			func(handler GinHandler) {
				router.DELETE(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodHead:
			func(handler GinHandler) {
				router.HEAD(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodOptions:
			func(handler GinHandler) {
				router.OPTIONS(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		case http.MethodPatch:
			func(handler GinHandler) {
				router.PATCH(fullPath, func(c *gin.Context) {
					handler(c)
				})
			}(v)
		}
	}
	return nil
}
