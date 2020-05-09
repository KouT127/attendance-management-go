package attendance

import (
	"github.com/KouT127/attendance-management/api/handler"
	"github.com/KouT127/attendance-management/api/payloads"
	. "github.com/KouT127/attendance-management/api/responses"
	"github.com/KouT127/attendance-management/application/services"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AttendanceHandler interface {
	ListHandler(c *gin.Context)
	MonthlyHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
}

type attendanceService struct {
	service services.AttendanceService
}

func NewAttendanceHandler(service services.AttendanceService) AttendanceHandler {
	return &attendanceService{
		service: service,
	}
}

func (s *attendanceService) ListHandler(c *gin.Context) {
	var (
		userID string
		res    *models.GetAttendancesResults
		err    error
	)

	p := payloads.NewPaginatorPayload(0, 5)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if userID, err = handler.GetIDByKey(c, auth.AuthorizedUserIDKey); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	params := models.GetAttendancesParameters{
		UserID:     userID,
		Pagination: p.ToPagination(),
	}

	if res, err = s.service.GetAttendances(params); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(res.MaxCnt)
	resps := ToAttendancesResponses(hasNext, res.Attendances)
	c.JSON(http.StatusOK, resps)
}

func (s *attendanceService) MonthlyHandler(c *gin.Context) {
	var (
		userID string
		res    *models.GetAttendancesResults
		err    error
	)

	p := payloads.NewPaginatorPayload(0, 31)
	//param := payloads.NewSearchParams()

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := c.Bind(s); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if userID, err = handler.GetIDByKey(c, auth.AuthorizedUserIDKey); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	params := models.GetAttendancesParameters{
		UserID:     userID,
		Pagination: p.ToPagination(),
	}

	if res, err = s.service.GetAttendances(params); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	hasNext := p.HasNext(res.MaxCnt)
	resps := ToAttendancesResponses(hasNext, res.Attendances)
	c.JSON(http.StatusOK, resps)
}

func (s *attendanceService) CreateHandler(c *gin.Context) {
	input := payloads.AttendancePayload{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, NewValidationError("user", err))
		return
	}

	userID, err := handler.GetIDByKey(c, auth.AuthorizedUserIDKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	attendanceTime := input.ToAttendanceTime()
	attendance, err := s.service.CreateOrUpdateAttendance(c, attendanceTime, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, NewError(BadAccessError))
		return
	}

	res := ToAttendanceResult(attendance)
	c.JSON(http.StatusOK, res)
}
