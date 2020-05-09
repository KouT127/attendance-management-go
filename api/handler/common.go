package handler

import (
	"errors"
	. "github.com/gin-gonic/gin"
)

func GetIDByKey(ctx *Context, key string) (string, error) {
	value, exists := ctx.Get(key)
	if !exists {
		return "", errors.New("user not found")
	}
	id := value.(string)
	return id, nil
}
