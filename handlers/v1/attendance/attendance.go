package attendance

import (
	"github.com/KouT127/attendance-management/application/facades"
	"github.com/KouT127/attendance-management/handlers"
	"github.com/KouT127/attendance-management/infrastructure/sqlstore"
	"github.com/KouT127/attendance-management/models"
	"github.com/KouT127/attendance-management/modules/auth"
	"github.com/KouT127/attendance-management/modules/payloads"
	. "github.com/KouT127/attendance-management/modules/responses"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListHandler(c *gin.Context) {
	var (
		userId string
		res    *models.GetAttendancesResults
		err    error
	)

	store := sqlstore.InitDatabase()
	facade := facades.NewAttendanceFacade(store)

	p := payloads.NewPaginatorPayload(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if userId, err = handlers.GetIdByKey(c, auth.AuthorizedUserIDKey); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	params := models.GetAttendancesParameters{
		UserID:    userId,
		Paginator: p.ToPaginator(),
	}

	if res, err = facade.GetAttendances(params); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(res.MaxCnt)
	resps := ToAttendancesResponses(hasNext, res.Attendances)
	c.JSON(http.StatusOK, resps)
}

func MonthlyHandler(c *gin.Context) {
	var (
		userId string
		res    *models.GetAttendancesResults
		err    error
	)

	store := sqlstore.InitDatabase()
	facade := facades.NewAttendanceFacade(store)

	p := payloads.NewPaginatorPayload(0, 31)
	s := payloads.NewSearchParams()

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if userId, err = handlers.GetIdByKey(c, auth.AuthorizedUserIDKey); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	params := models.GetAttendancesParameters{
		UserID:    userId,
		Paginator: p.ToPaginator(),
	}

	if res, err = facade.GetAttendances(params); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(res.MaxCnt)
	resps := ToAttendancesResponses(hasNext, res.Attendances)
	c.JSON(http.StatusOK, resps)
}

func CreateHandler(c *gin.Context) {
	store := sqlstore.InitDatabase()
	facade := facades.NewAttendanceFacade(store)

	input := payloads.AttendancePayload{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, NewValidationError("user", err))
		return
	}

	userId, err := handlers.GetIdByKey(c, auth.AuthorizedUserIDKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	attendanceTime := input.ToAttendanceTime()
	attendance, err := facade.CreateOrUpdateAttendance(attendanceTime, userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	res := ToAttendanceResult(attendance)
	c.JSON(http.StatusOK, res)
}
