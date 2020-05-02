package sqlstore

import (
	"context"
	"fmt"
	"github.com/KouT127/attendance-management/domain/models"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
	"log"
	"os"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
)

const (
	UserTable           = "users"
	AttendanceTable     = "attendances"
	AttendanceTimeTable = "attendances_time"
)

type SqlStore interface {
	InTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	GetAttendancesCount(ctx context.Context, userId string) (int64, error)
	GetLatestAttendance(ctx context.Context, userId string) (*models.Attendance, error)
	GetAttendances(ctx context.Context, query *models.GetAttendancesParameters) ([]*models.Attendance, error)
	UpdateOldAttendanceTime(ctx context.Context, id int64, kindId uint8) error
	CreateAttendance(ctx context.Context, attendance *models.Attendance) error
	CreateAttendanceTime(ctx context.Context, attendanceTime *models.AttendanceTime) error
	GetUser(ctx context.Context, userId string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, user *models.User) error
}

type sqlStore struct {
	engine *xorm.Engine
}

var (
	eng *xorm.Engine
)

func mustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Printf("Warning: %s environment variable not set.\n", key)
	}
	return v
}

func configureConnectionPool(engine *xorm.Engine) {
	engine.SetMaxIdleConns(5)
	engine.SetMaxOpenConns(7)
	engine.SetConnMaxLifetime(1800)
}

func configureLogger(engine *xorm.Engine) {
	logger := engine.Logger()
	logger.ShowSQL(true)
	logger.SetLevel(xlog.LOG_DEBUG)
}

func configureTimezone(engine *xorm.Engine) {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}
	engine.SetTZLocation(loc)
	engine.SetTZDatabase(loc)
}

func InitDatabase() SqlStore {
	var (
		ss  sqlStore
		err error
	)
	dbHost := os.Getenv("DB_TCP_HOST")
	if dbHost == "" {
		eng, err = initSocketConnectionPool()
		if err != nil {
			log.Fatalf("Socket connection is unavailable")
		}
	} else {
		eng, err = initTCPConnectionPool()
		if err != nil {
			log.Fatalf("Tcp connection is unavailable")
		}
	}
	ss.engine = eng
	return &ss
}

func initSocketConnectionPool() (*xorm.Engine, error) {
	var (
		err                    error
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", dbUser, dbPwd, instanceConnectionName, dbName)
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		return nil, xerrors.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(engine)
	configureTimezone(engine)
	return engine, nil
}

func initTCPConnectionPool() (*xorm.Engine, error) {
	var (
		err       error
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTCPHost, dbName)
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		return nil, xerrors.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureConnectionPool(engine)
	configureLogger(engine)
	configureTimezone(engine)
	return engine, nil
}

func InitTestDatabase() SqlStore {
	var (
		ss  sqlStore
		err error
	)
	eng, err = initTestTCPConnectionPool()
	if err != nil {
		log.Fatalf("Socket connection is unavailable")
	}
	if err = DeleteTestData(); err != nil {
		log.Fatalf("Failed delete data %s", err)
	}
	ss.engine = eng
	return &ss
}

func initTestTCPConnectionPool() (*xorm.Engine, error) {
	var (
		err       error
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_TCP_HOST")
		dbName    = mustGetenv("TEST_DB_NAME")
	)

	uri := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPwd, dbTCPHost, dbName)
	engine, err := xorm.NewEngine("mysql", uri)
	if err != nil {
		return nil, xerrors.Errorf("xorm.NewEngine: %v", err)
	}

	// configure settings
	configureTimezone(engine)
	return engine, nil
}
