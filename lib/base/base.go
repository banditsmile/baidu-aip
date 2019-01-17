package base

import "crypto/md5"

type AuthObj struct {
	time int64
	is_cloud_user bool

}
func Auth(refresh bool) {

	if !refresh {

	}
}

func getAuthFileName (apiKey string) [md5.Size]byte{
	apiKeySlice := []byte(apiKey)
	return md5.Sum(apiKeySlice)
}


func writeAuthObj()