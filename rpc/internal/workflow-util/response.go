package util2

import (
	"fmt"
	"net/http"
)

var fmtStr string = "{\"code\":\"%d\",\n\"msg\":\"%s\",\n\"data\":%s}"

// ResponseOk 返回信息给前台
func ResponseOk(writer http.ResponseWriter, data string) {
	fmt.Fprintf(writer, fmtStr, 200, "", data)
}

// ResponseErr 返回错误给前台
func ResponseErr(w http.ResponseWriter, msg string) {
	fmt.Fprintf(w, fmtStr, 400, msg, "")
}
