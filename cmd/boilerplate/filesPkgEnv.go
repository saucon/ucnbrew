package boilerplate

import (
	"log"
	"os"
)

func createPkgEnv() (err error) {
	f, err := os.Create("configs/env/env.go")
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
package env

import (
	"fmt"
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var Config ServerConfig

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func NewEnv(filenames ...string) {
	err := loadConfig(filenames...)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error_cause":   PrintErrorStack(err),
			"error_message": err.Error(),
		}).Fatal("config/env: load config")
	}
}

func loadConfig(filenames ...string) (err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		logrus.Fatal(err, " config/env: load gotdotenv")
	}

	err = env.Parse(&Config)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.DBConfig)
	if err != nil {
		return err
	}

	err = env.Parse(&Config.ElasticConfig)
	if err != nil {
		return err
	}

	return err
}

func PrintErrorStack(err error) string {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	return stFormat
}
`)

	if err != nil {
		return err
	}

	return nil
}

func createPkgEnvConfig() (err error) {
	f, err := os.Create("configs/env/config.go")
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
package env

import "time"

type ServerConfig struct {
	Name          string ` + "`env:\"NAME_SERVER\"`" + `
	Port          string ` + "`env:\"PORT_SERVER,required\"`" + `
	Host          string ` + "`env:\"HOST_SERVER,required\"`" + `
	JSONPathFile  string ` + "`env:\"JSON_PATHFILE,required\"`" + `
	DBConfig      DBConfig
	ElasticConfig ElasticConfig
}

type DBConfig struct {
	Name          string ` + "`env:\"NAME_POSTGRES,required\"`" + `
	Host          string ` + "`env:\"HOST_POSTGRES,required\"`" + `
	Port          string ` + "`env:\"PORT_POSTGRES,required\"`" + `
	User          string ` + "`env:\"USER_POSTGRES\"`" + `
	Password      string ` + "`env:\"PASS_POSTGRES\"`" + `
	NameLogDb     string ` + "`env:\"NAME_POSTGRES_LOG,required\"`" + `
	HostLogDb     string ` + "`env:\"HOST_POSTGRES_LOG,required\"`" + `
	PortLogDb     string ` + "`env:\"PORT_POSTGRES_LOG,required\"`" + `
	UserLogDb     string ` + "`env:\"USER_POSTGRES_LOG\"`" + `
	PasswordLogDb string ` + "`env:\"PASS_POSTGRES_LOG\"`" + `
}

type ElasticConfig struct {
	Host     string ` + "`env:\"HOST_ELASTICSEARCH,required\"`" + `
	Port     string ` + "`env:\"PORT_ELASTICSEARCH,required\"`" + `
	User     string ` + "`env:\"USER_ELASTICSEARCH\"`" + `
	Password string ` + "`env:\"PASS_ELASTICSEARCH\"`" + `
	Index    string ` + "`env:\"INDEX_ELASTICSEARCH,required\"`" + `
}

type Logs struct {
	ID           uint      ` + "`json:\"id\" gorm:\"column:id\"`" + `
	Level        string    ` + "`json:\"level\" gorm:\"column:level\"`" + `
	Message      string    ` + "`json:\"message\" gorm:\"column:message\"`" + `
	CreatedAt    time.Time ` + "`json:\"created_at\" gorm:\"column:created_at\"`" + `
	Request      string    ` + "`json:\"request\" gorm:\"type:JSONB NULL DEFAULT '{}'::JSONB\"`" + `
	Response     string    ` + "`json:\"response\" gorm:\"type:JSONB NULL DEFAULT '{}'::JSONB\"`" + `
	RequestBE    string    ` + "`json:\"request_be\" gorm:\"type:JSONB NULL DEFAULT '{}'::JSONB\"`" + `
	ResponseBE   string    ` + "`json:\"response_be\" gorm:\"type:JSONB NULL DEFAULT '{}'::JSONB\"`" + `
	PathError    string    ` + "`json:\"path_error\"`" + `
	ResponseTime string    ` + "`json:\"response_time\"`" + `
	TraceHeader  string    ` + "`json:\"trace_header\" gorm:\"type:JSONB NULL DEFAULT '{}'::JSONB\"`" + `
}
`)

	if err != nil {
		return err
	}

	return nil
}
