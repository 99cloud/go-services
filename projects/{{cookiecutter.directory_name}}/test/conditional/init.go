package conditional

import (
	"{{cookiecutter.project_slug}}/cmd/{{cookiecutter.app_slug}}/app"
	testUtil "{{cookiecutter.project_slug}}/test/util"
	"encoding/json"
	flag "github.com/spf13/pflag"
	"log"
	"strings"
	"testing"
	"time"
)

var dbname, sshHost string

func init() {
	flag.StringVar(&dbname, "db-name", "/tmp/{{cookiecutter.project_slug}}.sqlite", "dbname")
	flag.StringVar(&sshHost, "ssh-host", "172.16.30.34", "ssh-host")
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
			cmd.SetArgs([]string{"--db-name", dbname, "--ssh-host", sshHost})
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
