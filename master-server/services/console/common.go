package console

import (
	"net/http"
	"strconv"
	"strings"
)

func getBooleanFromQuery(param []string) bool {
	return param != nil && len(param) > 0 && strings.ToLower(param[0]) == "false"
}

func getIntFromQuery(param []string) (int, bool) {
	if param != nil {
		intParam, err := strconv.Atoi(param[0])
		if err != nil {
			return 0, false
		}
		return intParam, true
	}
	return 0, false
}

func writeBytesInResponse(w *http.ResponseWriter, mustWrite *[]byte) {
	mustWriteLen := len(*mustWrite)
	writtenBytes, err := (*w).Write(*mustWrite)
	if err != nil {
		panic(err)
	}
	for writtenBytes < mustWriteLen {
		curWrittenBytes, err := (*w).Write((*mustWrite)[writtenBytes:])
		writtenBytes += curWrittenBytes
		if err != nil {
			panic(err)
		}
	}
}
