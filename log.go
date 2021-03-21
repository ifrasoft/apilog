package apilog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"
)

var timestampFmt = "2006-01-0215:04:05.999"
var defaultLogPath = "./log"
var logPath = defaultLogPath

// Log formats.
const (
	logFmtInfo = "TIMESTAMP|%s|LOG_TYPE|%s|IP|%s|URI|%s|REQUEST_ID|%s|SESSION_ID|%s|TRAN_ID|%s|METHOD|%s|REQUEST_PARAM|%s|RESPONSE_PARAM|%s|RESULT|%s|RESULT_CODE|%s|RESP_TIME|%d"
	logFmtService = "TIMESTAMP|%s|LOG_TYPE|%s|NODE|%s|REQUEST_ID|%s|TRAN_ID|%s|USER_ID|%s|ACTION|%s|COMMAND|%s|REQUEST_PARAM|%s|RESPONSE_PARAM|%s|RESULT|%s|RESULT_CODE|%s|RESULT_DESC|%s|RESP_TIME|%s"
	logFmtSummary = "TIMESTAMP|%s|RESP_TIME|%s|TID|%s|MSISDN|%s|FBBID|%s|NTYPE|%s|URI|%s|DESCRIPTION|%s|ACTION"
)

// LOG_TYPE values.
const (
	logTypeInfo = "INFO"
	logTypeError = "ERROR"
)

// RESULT_CODE values.
const (
	resultSuccess = "SUCCESS"
	resultError = "FAILED"
)

// SetPath sets path for log files.
//
// If not set, use default path "./logs".
func SetPath(p string) {
	logPath = path.Clean(p)
}

// timestamp returns current time in "2006-01-0215:04:05.999" format.
func timestamp() string {
	return time.Now().Format(timestampFmt)
}

// toMilli returns time in milliseconds.
func toMilli(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// info user for logging client (incoming) requests.
func info(logType, ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, result, resCode string, respTime time.Time) {
	reqBodyJsonBytes, _ := json.Marshal(reqBody)
	respBodyJsonBytes, _ := json.Marshal(respBody)

	log := fmt.Sprintf(logFmtInfo,
		timestamp(),
		logType,
		ip,
		uri,
		reqID,
		sessionID,
		tranID,
		strings.ToUpper(method),
		string(reqBodyJsonBytes),
		string(respBodyJsonBytes),
		result,
		resCode,
		toMilli(respTime))

	writeln(log, "/info")
}

// InfoSuccess used for logging success client (incoming) requests.
func InfoSuccess(ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, resCode string, respTime time.Time) {
	info(logTypeInfo,
		ip,
		uri,
		reqID,
		sessionID,
		tranID,
		strings.ToUpper(method),
		reqBody,
		respBody,
		resultSuccess,
		resCode,
		respTime)
}

// InfoError used for logging failed client (incoming) requests.
func InfoError(ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, resCode string, respTime time.Time) {
	info(logTypeError,
		ip,
		uri,
		reqID,
		sessionID,
		tranID,
		strings.ToUpper(method),
		reqBody,
		respBody,
		resultError,
		resCode,
		respTime)
}

// service used for logging outgoing requests.
func service(logType, node, reqID, tranID, usrID, action, cmd string, reqBody, respBody interface{}, result, resCode string) {
	// TODO: Implement.
	reqBodyJsonBytes, _ := json.Marshal(reqBody)
	respBodyJsonBytes, _ := json.Marshal(respBody)

	log := fmt.Sprintf(logFmtService,
		timestamp(),
		logType,
		node,
		reqID,
		tranID,
		usrID,
		action,
		cmd,
		string(reqBodyJsonBytes),
		string(respBodyJsonBytes),
		result,
		resCode)

	writeln(log, "/service/service.log")
}

// ServiceSuccess used for logging success outgoing requests.
func ServiceSuccess() {
	// TODO: Implement.
	log := fmt.Sprintf(logFmtService,
		timestamp(),
		logTypeInfo)
	writeln(log, "/service")
}

// ServiceError used for logging failed outgoing requests.
func ServiceError() {
	// TODO: Implement.
	log := fmt.Sprintf(logFmtService,
		timestamp(),
		logTypeError)
	writeln(log, "/service")
}

// Summary used for incoming request, outgoing request, and outgoing response.
func Summary() {
	// TODO: Implement.
	log := fmt.Sprintf(logFmtSummary,
		timestamp())
	writeln(log, "/summary/sum.log")
}

// writeln writes log to file in path.
//
// fdName may omit leading "/".
func writeln(log, filePath string) {
	// TODO: Implement
	pathJoined := path.Join(logPath, filePath)
	ioutil.WriteFile(pathJoined, []byte(log + "\n"), os.ModePerm)
}
