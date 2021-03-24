# API API Log

There are 3 types of logs
1. Info
2. Service
3. Summary

---

## Info Log
Logs client (incoming) requests to an app.

Format: `TIMESTAMP|{}|LOG_TYPE|{}|IP|{}|URI|{}|REQUEST_ID|{}|{}|TRAN_ID|{}|METHOD|{}|REQUEST_PARAM|{}|RESPONSE_PARAM|{}|RESULT|{}|RESULT_CODE|{}|RESP_TIME|{}`

`{}` is a placeholder.

---

## Service Log

Logs outgoing requests of an app.

Format: `TIMESTAMP|{}|LOG_TYPE|{}|NODE|{}|REQUEST_ID|{}|TRAN_ID|{}|USER_ID|{}|ACTION|{}|COMMAND|{}|REQUEST_PARAM|{}|RESPONSE_PARAM|{}|RESULT|{}|RESULT_CODE|{}|RESULT_DESC|{}|RESP_TIME|{}`

`{}` is a placeholder.

---

## Summary Log

Logs incoming request.

Logs outgoing request and response.

Format: `TIMESTAMP|{}|RESP_TIME|{}|TID|{}|MSISDN|{}|FBBID|{}|NTYPE|{}|URI|{}|DESCRIPTION|{}|ACTION`

`{}` is a placeholder.
