package global

import (
	"fmt"
	"io/ioutil"
	"os"
)

var accessTokenCacheDirectory = fmt.Sprintf("%s/kdrc.json", os.TempDir())

func SaveTokenToDisk(byteArray []byte) {
	os.Remove(accessTokenCacheDirectory)
	ioutil.WriteFile(accessTokenCacheDirectory, byteArray, 0644)
}
