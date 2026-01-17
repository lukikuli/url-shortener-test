package mysql

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"time"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
)

func InitMysqlDB(cfg *MysqlConfig) *sql.DB {
	var (
		errMysql error
		dbConn   *sql.DB
	)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		val.Encode(),
	)
	log.Printf("DSN: %s\n", dsn)

	if os.Getenv("APP_ENV") == "production" {
		dbConn, errMysql = sql.Open(`nrmysql`, dsn)
	} else {
		logrus.Infof("Connecting to MySQL DSN: %s", dsn)
		dbConn, errMysql = sql.Open(`mysql`, dsn)
	}
	if errMysql != nil {
		logrus.Panicf("Failed, Connection to Mysql Database: %v", errMysql)
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbConn.SetMaxOpenConns(100)
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbConn.SetMaxIdleConns(20)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Should be less than MySQL's wait_timeout (default 28800s/8h)
	dbConn.SetConnMaxLifetime(15 * time.Minute)
	// SetConnMaxIdleTime sets the maximum amount of time a connection may be idle.
	// Closes idle connections before MySQL does (prevents EOF)
	dbConn.SetConnMaxIdleTime(5 * time.Minute)

	// Ping to test
	if err := dbConn.Ping(); err != nil {
		logrus.Panicf("mysql ping failed: %s", err.Error())
	}

	return dbConn
}
