package v1

import (
	. "github.com/gin-gonic/gin"
)

func Router(r *Engine) {
	group := r.Group("/v1")
	UsersRouter(group)
	AttendancesRouter(group)
}
