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

func writeJsonInResponse(w *http.ResponseWriter, data string) {
	(*w).WriteHeader(http.StatusOK)
	mustWrite := []byte(data)
	writeBytesInResponse(w, &mustWrite)
}

func writeForbiddenError(w *http.ResponseWriter) {
	writeAnyErrorResponse(w, nil, http.StatusForbidden, "{\"message\": \"request forbidden\"}")
}

func writeBadRequestError(w *http.ResponseWriter, requestBody string) {
	writeAnyErrorResponse(w, nil, http.StatusBadRequest, fmt.Sprintf("{\"message\": \"bad request [%s]\"}", requestBody))
}

func writeWrongHttpMethodResponse(w *http.ResponseWriter, method string) {
	writeAnyErrorResponse(w, nil, http.StatusNotFound, fmt.Sprintf("{\"message\": \"unsupported http method [%s]\"}", method))
}

func writeServerErrorResponse(w *http.ResponseWriter, err error) {
	writeAnyErrorResponse(w, err, http.StatusInternalServerError, "{\"message\": \"internal server error\"}")
}

func writeBadRequestErrorResponse(w *http.ResponseWriter, err error) {
	writeAnyErrorResponse(w, err, http.StatusBadRequest, "{\"message\": \"bad request\"}")
}

func writeAnyErrorResponse(w *http.ResponseWriter, err error, errorCode int, errorMessage string) {
	(*w).WriteHeader(errorCode)
	mustWrite := []byte(errorMessage)
	writeBytesInResponse(w, &mustWrite)
	if err != nil {
		log.Print(err)
	}
}
