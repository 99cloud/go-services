package controller

import (
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"

	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/cmd/APP_46ea591951824d8e9376b0f98fe4d48a/app"
	testUtil "PROJECT_46ea591951824d8e9376b0f98fe4d48a/test/util"

	flag "github.com/spf13/pflag"
)

var dbname string
var dbhost string
var dbusername string
var dbpassword string
var dbtype string
var dbport string

func init() {
	flag.StringVar(&dbname, "db-name", "/tmp/dbAPP_46ea591951824d8e9376b0f98fe4d48a", "dbname")
	flag.StringVar(&dbhost, "db-host", "", "dbhost")
	flag.StringVar(&dbusername, "db-username", "", "dbusername")
	flag.StringVar(&dbpassword, "db-password", "", "dbpassword")
	flag.StringVar(&dbtype, "db-type", "sqlite3", "dbtype")
	flag.StringVar(&dbport, "db-port", "", "dbport")
	flag.Parse()
	time.Sleep(time.Second * 3)
	go func() {
		cmd := app.NewAPIServerCommand()
		if len(flag.Args()) > 0 {
			var cmdFlags []string
			for _, iFlag := range flag.Args() {
				keyVal := strings.Split(iFlag, "=")
				cmdFlags = append(cmdFlags, "--"+keyVal[0], keyVal[1])
			}
			cmd.SetArgs(cmdFlags)
		} else {
			cmd.SetArgs([]string{"--db-name", dbname})
		}
		if err := cmd.Execute(); err != nil {
			log.Fatalln(err)
		}
	}()
}

// copy a common request to prevent test request and test target function request always timeout at save time

func requestTest(t *testing.T, requestURL, httpMethod string, postBody json.RawMessage, header map[string]string, skipTLSCheck, disableKeepAlive bool) ([]byte, int) {
	body, statusCode, err := testUtil.CommonRequest(requestURL,
		httpMethod, postBody, header,
		skipTLSCheck, disableKeepAlive, 100*time.Second)
	if err != nil {
		t.Errorf("Failed To Request url %s with data %s", requestURL, postBody)
		return nil, -1
	}
	return body, statusCode
}
