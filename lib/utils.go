package lib

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mahfuz110244/project2/entity"
)

func GetTextInfo(message bytes.Buffer) (*entity.Response, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	url := "http://localhost:8080/api/v1/word/occurrence"
	responseBody := bytes.NewBuffer(message.Bytes())
	req, err := http.NewRequest("GET", url, responseBody)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("reasone: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		var res entity.Response
		err = json.Unmarshal(data, &res)
		if err != nil {
			return nil, err
		}
		return &res, nil
	}
	return nil, entity.ErrSomethingWentWrong
}
