package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func execPath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func main() {
	err := godotenv.Load(execPath() + "/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	CLOUDFLARE_UID, c_ok := os.LookupEnv("CLOUDFLARE_UID")
	CLOUDFLARE_API, api_ok := os.LookupEnv("CLOUDFLARE_API")
	KV_NAMESPACE, kn_ok := os.LookupEnv("KV_NAMESPACE")
	KV_KEY, kk_ok := os.LookupEnv("KV_KEY")

	if !c_ok || !api_ok || !kn_ok || !kk_ok {
		log.Fatal("Environment variables are not set")
	}

	client := &http.Client{}
	url := fmt.Sprintf("https://api.cloudflare.com/client/v4/accounts/%s/storage/kv/namespaces/%s/values/%s", CLOUDFLARE_UID, KV_NAMESPACE, KV_KEY)
	current := time.Now().Unix()
	current_str := strconv.FormatInt(current, 10)
	req, err := http.NewRequest("PUT", url, strings.NewReader(current_str))

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", CLOUDFLARE_API)
	req.Header.Set("Content-Type", "*/*")
	client.Do(req)
}
