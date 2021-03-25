package apilog

import (
	"encoding/json"
	"fmt"
	"github.com/jasonlvhit/gocron"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func init() {
	infFPCleaned := filepath.Clean(logPath + infoFilePath)
	serFPCleaned := filepath.Clean(logPath + serviceFilePath)
	sumFPCleaned := filepath.Clean(logPath + summaryFilePath)

	gocron.Every(1).Second().Do(completeLog(infFPCleaned))
	gocron.Every(1).Second().Do(completeLog(serFPCleaned))
	gocron.Every(1).Second().Do(completeLog(sumFPCleaned))
}

// completeLog returns function that appends current log (if exist) file name
// with current time when current time minute is 0 or divisible by 15.
func completeLog(filePath string) func() {
	return func() {
		t := time.Now()
		if t.Minute() == 0 || t.Minute() % 15 == 0 {
			os.Rename(filePath, filePath + t.Format("20060102_0304"))
		}
	}
}

var timestampFmt = "2006-01-0215:04:05.999"
var defaultLogPath = "./log"
var logPath = defaultLogPath

// Log formats.
const (
	logFmtInfo = "TIMESTAMP|%s|LOG_TYPE|%s|IP|%s|URI|%s|REQUEST_ID|%s|SESSION_ID|%s|TRAN_ID|%s|METHOD|%s|REQUEST_PARAM|%s|RESPONSE_PARAM|%s|RESULT|%s|RESULT_CODE|%s|RESP_TIME|%d"
	logFmtService = "TIMESTAMP|%s|LOG_TYPE|%s|NODE|%s|REQUEST_ID|%s|TRAN_ID|%s|USER_ID|%s|ACTION|%s|COMMAND|%s|REQUEST_PARAM|%s|RESPONSE_PARAM|%s|RESULT|%s|RESULT_CODE|%s|RESULT_DESC|%s|RESP_TIME|%d"
	logFmtSummary = "TIMESTAMP|%s|RESP_TIME|%d|TID|%s|MSISDN|%s|FBBID|%s|NTYPE|%s|URI|%s|DESCRIPTION|%s|ACTION|%s"
)

// LOG_TYPE values.
const (
	logTypeInfo = "INFO"
	logTypeError = "ERROR"
)

// RESULT_CODE values.
const (
	resultSuccess = "SUCCESS"
	resultError = "ERROR"
)

// Relative log file paths.
const (
	infoFilePath = "/info/log.info"
	serviceFilePath = "/service/log.service"
	summaryFilePath = "/summary/log.sum"
)

// SetPath sets path for log files.
//
// If not set, use default path "./logs".
func SetPath(p string) {
	logPath = filepath.Clean(p)
}

// timestamp returns current time in "2006-01-0215:04:05.999" format.
func timestamp() string {
	return time.Now().Format(timestampFmt)
}

// info user for logging client (incoming) requests.
func info(logType, ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, result, resCode string, respTime time.Duration) {
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
		respTime.Milliseconds())

	writeln(log, infoFilePath)
}

// InfoSuccess used for logging success client (incoming) requests.
func InfoSuccess(ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, resCode string, respTime time.Duration) {
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
func InfoError(ip, uri, reqID, sessionID, tranID, method string, reqBody, respBody interface{}, resCode string, respTime time.Duration) {
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
func service(logType, node, reqID, tranID, usrID, action, cmd string, reqBody, respBody interface{}, result, resCode, resDesc string, respTime time.Duration) {
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
		resCode,
		resDesc,
		respTime.Milliseconds())

	writeln(log, serviceFilePath)
}

// ServiceSuccess used for logging success outgoing requests.
func ServiceSuccess(node, reqID, tranID, usrID, action, cmd string, reqBody, respBody interface{}, resCode, resDesc string, respTime time.Duration) {
	service(logTypeInfo,
	node,
	reqID,
	tranID,
	usrID,
	action,
	cmd,
	reqBody,
	respBody,
	resultSuccess,
	resCode,
	resDesc,
	respTime)
}

// ServiceError used for logging failed outgoing requests.
func ServiceError(node, reqID, tranID, usrID, action, cmd string, reqBody, respBody interface{}, resCode, resDesc string, respTime time.Duration) {
	service(logTypeError,
		node,
		reqID,
		tranID,
		usrID,
		action,
		cmd,
		reqBody,
		respBody,
		resultError,
		resCode,
		resDesc,
		respTime)
}

// Summary used for incoming request, outgoing request, and outgoing response.
func Summary(respTime time.Duration, tranID, msisdn, fbbID, netwkType, uri, desc, action string) {
	log := fmt.Sprintf(logFmtSummary,
		timestamp(),
		respTime.Milliseconds(),
		tranID,
		msisdn,
		fbbID,
		netwkType,
		uri,
		desc,
		action)

	writeln(log, summaryFilePath)
}

// writeln writes log to file in path.
func writeln(log, filePath string) {
	absPath, _ := filepath.Abs(logPath)
	fpJoined := filepath.Join(absPath, filePath)
	dirPath := filepath.Dir(fpJoined)
	// Create directory if not exists.
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	f, _ := os.OpenFile(fpJoined, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	defer f.Close()
	f.WriteString(log + "\n")
}
