package apitool

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

func getSecretKey() string {
	return beego.AppConfig.String("SecretKey")

}

func getAccount() string {
	return beego.AppConfig.String("Account")
}

func SignatureHeader() (map[string]interface{}, error) {
	timestamp := time.Now().Unix()
	payload := `{"account": "%s", "secret_key": "%s", "timestamp": %d}`
	hashedRequestPayload := sha256hex(fmt.Sprintf(payload, getAccount(), getSecretKey(), timestamp))

	signature := hex.EncodeToString([]byte(hmacSha256(getAccount(), hashedRequestPayload)))
	params := make(map[string]interface{})
	params["signature"] = signature
	params["account"] = getAccount()
	params["timestamp"] = timestamp
	return params, nil
}

func CustomSignatureHeader(account string, pwd string) (map[string]interface{}, error) {
	timestamp := time.Now().Unix()
	payload := `{"account": "%s", "secret_key": "%s", "timestamp": %d}`
	hashedRequestPayload := sha256hex(fmt.Sprintf(payload, account, pwd, timestamp))
	signature := hex.EncodeToString([]byte(hmacSha256(account, hashedRequestPayload)))
	params := make(map[string]interface{})
	params["signature"] = signature
	params["account"] = account
	params["timestamp"] = timestamp
	return params, nil
}

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func hmacSha256(s string, key string) string {
	hashed := hmac.New(sha256.New, []byte(s))
	hashed.Write([]byte(key))
	return string(hashed.Sum(nil))
}
