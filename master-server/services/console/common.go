package console

import (
	"fmt"
	"log"
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

func writeBadRequestError(w *http.ResponseWriter, requestBody string) {
	(*w).WriteHeader(http.StatusBadRequest)
	errorMessage := fmt.Sprintf("{\"message\": \"bad request [%s]\"}", requestBody)
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
}

func writeWrongHttpMethodResponse(w *http.ResponseWriter, method string) {
	(*w).WriteHeader(http.StatusNotFound)
	errorMessage := fmt.Sprintf("{\"message\": \"unsupported http method [%s]\"}", method)
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
}

func writeServerErrorResponse(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	errorMessage := "{\"message\": \"internal server error\"}"
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
	log.Print(err)
}
