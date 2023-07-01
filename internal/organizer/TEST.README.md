#



```
go test -v ./... 2>&1
go test -json ./... | test-summary
```

https://github.com/jstemmer/go-junit-report
```
# go mod init
go mod init github.com/moudrick/kbot
go mod tidy
```

```
# go get -u                         github.com/jstemmer/go-junit-report/v2

go test -v    ./... 2>&1  |  go run github.com/jstemmer/go-junit-report/v2                -set-exit-code | tee go-junit-report.xml
go test -json ./... 2>&1  |  go run github.com/jstemmer/go-junit-report/v2 -parser gojson -set-exit-code | tee go-junit-report.xml

go test -v    ./... 2>&1  |  go run github.com/jstemmer/go-junit-report/v2                -set-exit-code  -out go-junit-report-plain.xml
go test -json ./... 2>&1  |  go run github.com/jstemmer/go-junit-report/v2 -parser gojson -set-exit-code  -out go-junit-report-json.xml
```

```
 ## xmlstarlet ed --inplace -u "//*[local-name()='failure']" -v "  !! no stack trace detected !! " junit-report.xml


## xmlstarlet ed --inplace -u "//*[local-name()='failure']" -v "  __ no stack trace detected __ " go-junit-report.xml
xmlstarlet ed --inplace -u "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" -v "  __ no stack trace detected __ " go-junit-report.xml

## xmlstarlet ed --inplace -u "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" -v "$(printf "<![CDATA[%s]]>" "  !! no stack trace detected !! ")" go-junit-report.xml

## xmlstarlet ed --inplace -u "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" -v "<![CDATA[  !! no stack trace detected !! ]]>" go-junit-report.xml
```

```
run: >-
      xmlstarlet ed --inplace 
        -u "//*[local-name()='failure'][not(text()) or normalize-space(text())='']"
        -v "  __ no stack trace detected __ "
        go-junit-report.xml


xmlstarlet ed --inplace --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" --value "  __ no stack trace detected __ " go-junit-report.xml

          args: >-
              ed --inplace
              --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']"
              --value "  __ no stack trace detected __ "
              go-junit-report.xml

hostname
  - name: delete all @hostname attributes
    uses: Mudlet/xmlstarlet-action@master
    with:
      args: ed -P -L --delete //@hostname go-junit-report.xml


                                        xmlstarlet ed --inplace --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" --value "  __ no stack trace detected __ " go-junit-report.xml
docker run 2a6c8c:5558e3b0193a4a998e754889eb6ce854 ed --inplace --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" --value "  __ no stack trace detected __ " go-junit-report.xml


      - name: delete all @hostname attribut
        if: success() || failure()       # run this step even if previous step failedes
        uses: Mudlet/xmlstarlet-action@master
        with:
          args: >
              ed -P -L --update "//@hostname" --value "__no_stack_trace_detected__"
              go-junit-report.xml
```



```
go install                        github.com/jstemmer/go-junit-report/v2@latest

go get -u                         github.com/jstemmer/go-junit-report
go get -u                         github.com/jstemmer/go-junit-report/v2

go run github.com/jstemmer/go-junit-report@v2 -version
go run github.com/jstemmer/go-junit-report    -version
                         ./go-junit-report    -version

go get github.com/jstemmer/go-junit-report@none
go get github.com/jstemmer/go-junit-report/v2@none
```

```
curl -L https://example.com/mygzip.tar.gz | tar zxv
curl -L https://github.com/jstemmer/go-junit-report/releases/download/v2.0.0/go-junit-report-v2.0.0-linux-amd64.tar.gz | tar zxv
chmod +x ./go-junit-report

go test -v    ./... 2>&1  |  ./go-junit-report -set-exit-code                     | tee go-junit-report.xml
go test -json ./... 2>&1  |  ./go-junit-report -parser gojson -set-exit-code 2>&1 | tee go-junit-report-json.xml
go test -json ./... 2>&1  |  ./go-junit-report -parser gojson -set-exit-code       -out go-junit-report-json.xml

```


```
moudrick@DESKTOP-J9A3R6I:/mnt/c/_/tt-devops/t-kbot/kbot$ go test -v 2>&1 ./...   |  go run github.com/jstemmer/go-junit-report -set-exit-code 
<?xml version="1.0" encoding="UTF-8"?>
<testsuites>
        <testsuite tests="4" failures="3" time="0.017" name="github.com/moudrick/kbot/internal/organizer/events">
                <properties>
                        <property name="go.version" value="go1.18.1"></property>
                </properties>
                <testcase classname="events" name="TestParseFirstDate" time="0.000">
                        <failure message="Failed" type=""></failure>
                </testcase>
                <testcase classname="events" name="TestParseFirstDate/Event_on_06/30/2023_at_10:00_AM" time="0.000"></testcase>
                <testcase classname="events" name="TestParseFirstDate/Meeting_at_2023-06-16_14:30" time="0.000">
                        <failure message="Failed" type="">    eventparser_test.go:50: </failure>
                </testcase>
                <testcase classname="events" name="TestParseFirstDate/Invalid_date_format" time="0.000">
                        <failure message="Failed" type="">    eventparser_test.go:48: </failure>
                </testcase>
        </testsuite>
</testsuites>
exit status 1
moudrick@DESKTOP-J9A3R6I:/mnt/c/_/tt-devops/t-kbot/kbot$ 
```


```
/usr/bin/docker run --name a6c8c5558e3b0193a4a998e754889eb6ce854_e60839 --label 2a6c8c --workdir /github/workspace --rm -e "INPUT_ARGS" -e "HOME" -e "GITHUB_JOB" -e "GITHUB_REF" -e "GITHUB_SHA" -e "GITHUB_REPOSITORY" -e "GITHUB_REPOSITORY_OWNER" -e "GITHUB_REPOSITORY_OWNER_ID" -e "GITHUB_RUN_ID" -e "GITHUB_RUN_NUMBER" -e "GITHUB_RETENTION_DAYS" -e "GITHUB_RUN_ATTEMPT" -e "GITHUB_REPOSITORY_ID" -e "GITHUB_ACTOR_ID" -e "GITHUB_ACTOR" -e "GITHUB_TRIGGERING_ACTOR" -e "GITHUB_WORKFLOW" -e "GITHUB_HEAD_REF" -e "GITHUB_BASE_REF" -e "GITHUB_EVENT_NAME" -e "GITHUB_SERVER_URL" -e "GITHUB_API_URL" -e "GITHUB_GRAPHQL_URL" -e "GITHUB_REF_NAME" -e "GITHUB_REF_PROTECTED" -e "GITHUB_REF_TYPE" -e "GITHUB_WORKFLOW_REF" -e "GITHUB_WORKFLOW_SHA" -e "GITHUB_WORKSPACE" -e "GITHUB_ACTION" -e "GITHUB_EVENT_PATH" -e "GITHUB_ACTION_REPOSITORY" -e "GITHUB_ACTION_REF" -e "GITHUB_PATH" -e "GITHUB_ENV" -e "GITHUB_STEP_SUMMARY" -e "GITHUB_STATE" -e "GITHUB_OUTPUT" -e "RUNNER_OS" -e "RUNNER_ARCH" -e "RUNNER_NAME" -e "RUNNER_ENVIRONMENT" -e "RUNNER_TOOL_CACHE" -e "RUNNER_TEMP" -e "RUNNER_WORKSPACE" -e "ACTIONS_RUNTIME_URL" -e "ACTIONS_RUNTIME_TOKEN" -e "ACTIONS_CACHE_URL" -e GITHUB_ACTIONS=true -e CI=true -v "/var/run/docker.sock":"/var/run/docker.sock" -v "/home/runner/work/_temp/_github_home":"/github/home" -v "/home/runner/work/_temp/_github_workflow":"/github/workflow" -v "/home/runner/work/_temp/_runner_file_commands":"/github/file_commands" -v "/home/runner/work/kbot/kbot":"/github/workspace" 2a6c8c:5558e3b0193a4a998e754889eb6ce854 ed --inplace --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" --value "  __ no stack trace detected __ " go-junit-report.xml
```

```
docker run 2a6c8c:5558e3b0193a4a998e754889eb6ce854 ed --inplace --update "//*[local-name()='failure'][not(text()) or normalize-space(text())='']" --value "  __ no stack trace detected __ " go-junit-report.xml
```