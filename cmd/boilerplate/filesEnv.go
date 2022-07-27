package boilerplate

import (
	"log"
	"os"
)

func createEnv() (err error) {
	f, err := os.Create(".env")
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	_, err = f.WriteString(`
PORT_SERVER=""
HOST_SERVER=""
NAME_SERVER=""
JSON_PATHFILE=""

HOST_POSTGRES=""
PORT_POSTGRES=""
USER_POSTGRES=""
PASS_POSTGRES=""
NAME_POSTGRES=""

HOST_POSTGRES_LOG=""
PORT_POSTGRES_LOG=""
USER_POSTGRES_LOG=""
PASS_POSTGRES_LOG=""
NAME_POSTGRES_LOG=""

HOST_ELASTICSEARCH=""
PORT_ELASTICSEARCH=""
INDEX_ELASTICSEARCH=""
USER_ELASTICSEARCH=""
PASS_ELASTICSEARCH=""
`)

	if err != nil {
		return err
	}

	return nil
}
