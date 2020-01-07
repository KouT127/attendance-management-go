package controllers

import (
	"errors"
	. "github.com/gin-gonic/gin"
)

func GetIdByKey(ctx *Context, key string) (string, error) {
	value, exists := ctx.Get(key)
	if !exists {
		return "", errors.New("user not found")
	}
	id := value.(string)
	return id, nil
}
