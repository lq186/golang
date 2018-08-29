package common

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/lq186/golang/lq186.com/apiserver/response"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func RandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func MD5(text string) string{
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

func JsonUnmarshal(request *http.Request, value interface{}) error {
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return err
	}
	defer request.Body.Close()
	return json.Unmarshal(body, value)
}

func CheckEmptyParam(writer http.ResponseWriter, paramName string, paramVal string) bool {
	if "" == strings.Trim(paramVal, " ") {
		response.WriteJsonData(writer, response.Data{Code: response.ParamError, Message: fmt.Sprintf("Parameter %s should not be empty.", paramName)})
		return false
	}
	return true
}