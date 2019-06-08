package context

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"microx/common/errors"
)

func New(ctx *gin.Context) *MxContext {
	return &MxContext{gctx: ctx}
}

type MxContext struct {
	gctx *gin.Context
}

func (c *MxContext) Context() *gin.Context {
	return c.gctx
}

func (c *MxContext) Param(key string) string {
	return c.gctx.Param(key)
}

func (c *MxContext) Query(key string) string {
	return c.gctx.Query(key)
}

func (c *MxContext) ParseJSON(obj interface{}) error {
	if err := c.gctx.ShouldBindJSON(obj); err != nil {
		return err
	}
	return nil
}

func (c *MxContext) Response(obj interface{}) {
	if obj == nil {
		obj = gin.H{}
	}

	m := make(map[string]interface{})
	m["errno"] = 0
	m["data"] = obj
	m["t"] = time.Now().UnixNano()
	c.response(http.StatusOK, m)
}

func (c *MxContext) ResponseError(err error) {
	ce := errors.Parse(err.Error())

	m := make(map[string]interface{})
	m["errno"] = ce.Errno
	m["errmsg"] = ce.Errmsg
	m["t"] = time.Now().UnixNano()
	c.response(http.StatusOK, m)
}

func (c *MxContext) response(status int, obj interface{}) {
	c.gctx.JSON(status, obj)
	c.gctx.Abort()
}
