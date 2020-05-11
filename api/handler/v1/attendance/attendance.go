package attendance

import (
	"github.com/KouT127/attendance-management/api/handler"
	"github.com/KouT127/attendance-management/api/payloads"
	"github.com/KouT127/attendance-management/api/responses"
	"github.com/KouT127/attendance-management/application/services"
	"github.com/KouT127/attendance-management/domain/models"
	"github.com/KouT127/attendance-management/infrastructure/auth"
	"github.com/KouT127/attendance-management/utilities/logger"
	"github.com/KouT127/attendance-management/utilities/timeutil"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler interface {
	ListHandler(c *gin.Context)
	CreateHandler(c *gin.Context)
	SummaryHandler(c *gin.Context)
}

type attendanceService struct {
	service services.AttendanceService
}

func NewAttendanceHandler(service services.AttendanceService) Handler {
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

	month, err := timeutil.GetDefaultMonth()
	if err != nil {
		logger.NewWarn(logrus.Fields{}, err.Error())
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	query := payloads.NewAttendancesQueryParam(month)
	if err = c.BindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	if userID, err = handler.GetIDByKey(c, auth.AuthorizedUserIDKey); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	params := models.GetAttendancesParameters{
		UserID: userID,
		Month:  query.Month,
	}

	if res, err = s.service.GetAttendances(c, params); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	resps := responses.ToAttendancesResponses(res.Attendances)
	c.JSON(http.StatusOK, resps)
}

func (s *attendanceService) CreateHandler(c *gin.Context) {
	input := payloads.AttendancePayload{}
	if err := c.Bind(&input); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewValidationError("user", err))
		return
	}

	userID, err := handler.GetIDByKey(c, auth.AuthorizedUserIDKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	if err = input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	attendanceTime := input.ToAttendanceTime()
	attendance, err := s.service.CreateOrUpdateAttendance(c, attendanceTime, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	res := responses.ToAttendanceCreatedResponse(attendance)
	c.JSON(http.StatusOK, res)
}

func (s *attendanceService) SummaryHandler(c *gin.Context) {
	userID, err := handler.GetIDByKey(c, auth.AuthorizedUserIDKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}

	results, err := s.service.GetAttendanceSummary(c, models.GetAttendanceSummaryParameters{UserID: userID})
	if err != nil {
		logger.NewWarn(map[string]interface{}{}, err.Error())
		c.JSON(http.StatusBadRequest, responses.NewError(responses.BadAccessError))
		return
	}
	resp := responses.ToAttendanceSummaryResponse(results)
	c.JSON(http.StatusOK, resp)
}
