package permission

type Permission string

const (
	Camera     Permission = "camera"
	Location   Permission = "location"
	Storage    Permission = "storage"
	Microphone Permission = "microphone"
)

type PermissionStatus string

const (
	Granted PermissionStatus = "granted"
	Denied  PermissionStatus = "denied"
	Pending PermissionStatus = "pending"
)

type requestPayload struct {
	Permission Permission `json:"permission"`
}

type responsePayload struct {
	Status PermissionStatus `json:"status"`
}

// RequestPermission pede ao runtime que solicite a permissão ao sistema.
//func RequestPermission(p Permission, callback func(PermissionStatus)) {
//	core.InvokeNative("request_permission", requestPayload{Permission: p}, func(response string) {
//		var out responsePayload
//		_ = json.Unmarshal([]byte(response), &out)
//		callback(out.Status)
//	})
//}
//
//// CheckPermission consulta o status atual da permissão.
//func CheckPermission(p Permission, callback func(PermissionStatus)) {
//	core.InvokeNative("check_permission", requestPayload{Permission: p}, func(response string) {
//		var out responsePayload
//		_ = json.Unmarshal([]byte(response), &out)
//		callback(out.Status)
//	})
//}
