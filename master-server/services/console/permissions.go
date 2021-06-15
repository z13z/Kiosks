package console

import (
	"github.com/lib/pq"
	"github.com/z13z/Kiosks/master-server/crypto"
	"net/http"
)

const permissionUsersKey = "users"
const permissionImagesKey = "images"
const permissionKiosksKey = "kiosks"
const authenticationHeaderKey = "Authentication"

func CheckPermissionToCall(r *http.Request, w *http.ResponseWriter, permissionName string) bool {
	authenticationToken := r.Header.Get(authenticationHeaderKey)
	claims, ok := crypto.CheckJwt(authenticationToken)
	if !ok || claims == nil || claims.Permissions == "" {
		writeForbiddenError(w)
		return false
	}
	arr := pq.StringArray{}
	err := arr.Scan(claims.Permissions)
	if err != nil || arrayNotContainsString(arr, permissionName) {
		writeForbiddenError(w)
		return false
	}
	return true
}

func arrayNotContainsString(array pq.StringArray, permissionName string) bool {
	for _, elem := range array {
		if permissionName == elem {
			return false
		}
	}
	return true
}
