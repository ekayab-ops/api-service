package helper

import (
	"math/rand"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

// GetRandomUUID returns a random UUID
func GetRandomUUID() string {
	rand.Seed(time.Now().UnixNano())
	uuid := make([]byte, 16)
	rand.Read(uuid)
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

// GetJSONFromResponse extracts JSON from an HTTP response
func GetJSONFromResponse(res *http.Response) (*map[string]interface{}, error) {
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		return nil, err
	}

	return &jsonData, nil
}