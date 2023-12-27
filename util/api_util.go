package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/yumu-uw/switchbot-api-client/model"
)

const base_url = "https://api.switch-bot.com/v1.1"

var (
	token      string
	nonce      uuid.UUID
	time_stamp int64
	sign       string
)

func init() {
	loadEnv()
	token = os.Getenv("TOKEN")
	secret := os.Getenv("SECRET")
	nonce, _ = uuid.NewUUID()
	time_stamp = time.Now().Unix() * 1000
	data := token + strconv.FormatInt(time_stamp, 10) + nonce.String()
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	sign = base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("can't read .env file")
	}
}

func invoke(method string, end_point_path ...string) []byte {
	client := &http.Client{}
	u, _ := url.JoinPath(base_url, end_point_path...)
	req, err := http.NewRequest(method, u, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("sign", sign)
	req.Header.Set("nonce", nonce.String())
	req.Header.Set("t", strconv.Itoa(int(time_stamp)))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func GetDeviceList() []model.Device {
	p := "devices"
	b := invoke(http.MethodGet, p)
	var result model.DeviceListResponseModel
	if err := json.Unmarshal(b, &result); err != err {
		log.Fatal(err)
	}
	return result.Body.DeviceList
}

func GetPlugMiniStatus(device_id string) {
	p := []string{"devices", device_id, "status"}
	b := invoke(http.MethodGet, p...)
	var plugmini_status model.PlugMiniResponseModel
	if err := json.Unmarshal(b, &plugmini_status); err != err {
		log.Fatal(err)
	}
	fmt.Println(plugmini_status)
}
