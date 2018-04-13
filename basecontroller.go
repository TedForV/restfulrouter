package RestfulRouter

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
	"errors"
)

type BaseController interface {
	Get(c *gin.Context)
	Post(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
	Patch(c *gin.Context)
	Head(c *gin.Context)
	Options(c *gin.Context)

	Mapping(ginEngine *gin.Engine)
}

type GinHandler func(c *gin.Context)

func RegisterAPIRoute(ginEngine *gin.Engine, controllers []BaseController) {
	if controllers == nil || len(controllers) == 0 {
		return
	}
	for _, c := range controllers {
		cname, err := getControllerValidName(c)
		if err != nil {
			panic(err)
		}
		autoMapping(ginEngine, cname, &c)
	}

}

const (
	CONTROLLER_SUFFIX     = "Controller"
	ERROR_CONTROLLER_NAME = "Controller name must be suffix with 'Controller'"
)

func getControllerValidName(controller BaseController) (string, error) {
	typeInfo := reflect.TypeOf(controller)
	fullName := typeInfo.Elem().String()
	lastDotIndex := strings.LastIndex(fullName, ".")
	fullName = fullName[lastDotIndex+1:]
	if strings.HasSuffix(fullName, CONTROLLER_SUFFIX) && len(fullName) > len(CONTROLLER_SUFFIX) {
		validName := fullName[0 : len(fullName)-len(CONTROLLER_SUFFIX)]
		return validName, nil
	} else {
		return "", errors.New(ERROR_CONTROLLER_NAME)
	}

}

func autoMapping(ginEngine *gin.Engine, controllerName string, controller BaseController) {
}