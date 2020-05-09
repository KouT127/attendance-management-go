package sqlstore

//go:generate mockgen -source=sqlstore.go -destination=mock/mock_sqlstore.go -package=sqlstore -aux_files github.com/KouT127/attendance-management/infrastructure/sqlstore=user.go,github.com/KouT127/attendance-management/infrastructure/sqlstore=attendance.go,github.com/KouT127/attendance-management/infrastructure/sqlstore=transaction.go
import (
	"fmt"
	"golang.org/x/xerrors"
	"log"
	"os"
	"time"
	"xorm.io/xorm"
	xlog "xorm.io/xorm/log"
	"xorm.io/xorm/names"
)

const (
	UserTable           = "users"
	AttendanceTable     = "attendances"
	AttendanceTimeTable = "attendances_time"
)

type SQLStore interface {
	Transaction
	User
	Attendance
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

func configureMapper(engine *xorm.Engine) {
	engine.SetMapper(names.GonicMapper{})
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

func InitDatabase() SQLStore {
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
	configureMapper(engine)
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
	configureMapper(engine)
	return engine, nil
}

func InitTestDatabase() SQLStore {
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
	configureMapper(engine)
	return engine, nil
}
