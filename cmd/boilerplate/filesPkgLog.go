package boilerplate

import (
	"log"
	"os"
)

func createPkgLogConfig(pkgname string) (err error) {
	f, err := os.Create("configs/log/log.go")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString(`// This file is generated using ucnbrew tool. 
// Check out for more info "https://github.com/saucon/ucnbrew"
package log

import (
	"` + pkgname + `/configs/env"
	

	"sync"
	"fmt"
	"github.com/olivere/elastic/v7" //nolint:typecheck
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7" //nolint:typecheck
)

type LogCustom struct {
	Logrus *logrus.Logger
	LogDb  *LogDbCustom
	WhoAmI iAm
}

type iAm struct {
	Name string
	Host string
	Port string
}

var instance *LogCustom
var once sync.Once

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewLogCustom(configServer env.ServerConfig) *LogCustom {
	var log *logrus.Logger

	configElstc := configServer.ElasticConfig

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL( //nolint:typecheck
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.User, configElstc.Password)) //nolint:typecheck
	if err != nil {
		selfLogError(err, "config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook( //nolint:typecheck
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "config/log: elastic client", log)
		}
		log.Hooks.Add(hook)
	}

	once.Do(func() {
		instance = &LogCustom{
			Logrus: log,
			WhoAmI: iAm{
				Name: configServer.Name,
				Host: configServer.Host,
				Port: configServer.Port,
			},
		}
	})
	return instance
}

func (l *LogCustom) Success(req, resp, reqBe, respBe interface{}, description, respTime string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"request":      req,
		"response":     resp,
		"request_be":   reqBe,
		"response_be":  respBe,
	}).Info(description)

	l.LogDb.SuccessLogDb(req, resp, reqBe, respBe, description, respTime, traceHeader)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Info(description string, traceHeader map[string]string, data ...interface{}) {
	l.Logrus.WithFields(logrus.Fields{
		"whoami":       l.WhoAmI,
		"trace_header": traceHeader,
		"message":      data,
	}).Info(description)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Error(err error, description string, respTime string, traceHeader map[string]string, req interface{},
	resp interface{}, reqBE interface{}, respBE interface{}) {

	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
	}).Error(description)

	l.LogDb.ErrorLogDb(err, description, respTime, stFormat, traceHeader, req, resp, reqBE, respBE)
}

// for description please use format for example
// "usecase: sync data"
func (l *LogCustom) Fatal(err error, description string, traceHeader map[string]string) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}

// for description please use format for example
// "usecase: sync data"
func selfLogFatal(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Fatal(description)
}
`)

	if err != nil {
		return err
	}

	return nil
}

func createPkgLogDbConfig() (err error) {
	f, err := os.Create("configs/log/logDb.go")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString(`// This file is generated using ucnbrew tool. 
// Check out for more info "https://github.com/saucon/ucnbrew"
package log

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	pglogrus "gopkg.in/gemnasium/logrus-postgresql-hook.v1"
	"gorm.io/gorm"
	"sync"
)

type LogDbCustom struct {
	Logrus *logrus.Logger
	WhoAmI iAm
	Db     *gorm.DB
}

var instanceDb *LogDbCustom
var onceDb sync.Once

func NewLogDbCustom(db *gorm.DB) *LogDbCustom {

	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	sqlDb, err := db.DB()

	if err != nil {
		//logger.Error(err, "db/NewLogDbCustom", nil, nil, nil)
		fmt.Println("adsgf")
	}

	hook := pglogrus.NewAsyncHook(sqlDb, map[string]interface{}{})
	hook.InsertFunc = func(sqlDb *sql.Tx, entry *logrus.Entry) error {
		level := entry.Level.String()
		if level == "info" {
			level = "success"
		}

		err = db.Debug().Exec("INSERT INTO logs(level, message, path_error, trace_header, request, response, request_be, response_be, created_at, response_time) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
			level, entry.Message, entry.Data["error_cause"], entry.Data["trace_header"], entry.Data["request"], entry.Data["response"], entry.Data["request_be"], entry.Data["response_be"], entry.Time, entry.Data["response_time"]).Error
		if err != nil {
			err := sqlDb.Rollback()
			if err != nil {
				return err
			}
		}
		return err
	}

	log.AddHook(hook)

	onceDb.Do(func() {
		instanceDb = &LogDbCustom{
			Logrus: log,
		}
		instance.LogDb = instanceDb
	})

	return instanceDb
}

func (l *LogDbCustom) ErrorLogDb(err error, description, respTime, strFormat string, traceHeader map[string]string, req, resp, reqBE, respBE interface{}) {
	err = errors.WithStack(err)

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"error_cause":   strFormat,
		"error_message": err.Error(),
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Error(description)
}

func (l *LogDbCustom) SuccessLogDb(req, resp, reqBE, respBE interface{}, description, respTime string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"whoami":        l.WhoAmI,
		"trace_header":  traceHeader,
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Info(description)
}
`)

	if err != nil {
		return err
	}

	return nil
}
