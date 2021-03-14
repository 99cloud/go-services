package conditional

import (
	"PROJECT_46ea591951824d8e9376b0f98fe4d48a/cmd/APP_46ea591951824d8e9376b0f98fe4d48a/app"
	testUtil "PROJECT_46ea591951824d8e9376b0f98fe4d48a/test/util"
	"encoding/json"
	"log"
	"strings"
	"testing"
	"time"

	flag "github.com/spf13/pflag"
)

var dbname string

func init() {
	flag.StringVar(&dbname, "db-name", "/tmp/PROJECT_46ea591951824d8e9376b0f98fe4d48a.sqlite", "dbname")
	flag.Parse()
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
	time.Sleep(time.Second * 3)
}

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
