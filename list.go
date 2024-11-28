package alist

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func (c *Client) List(req ListReq) (*ListResp, error) {
	jsonBytes, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	request, _ := http.NewRequest("POST", c.config.ServerUrl+"/api/fs/list", bytes.NewBuffer(jsonBytes))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", c.config.Token)
	result, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}
	resp := &ListResp{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ListReq struct {
	Path     string `json:"path"`
	Password string `json:"password"`
	Page     int    `json:"page"`
	PerPage  int    `json:"per_page"`
	Refresh  bool   `json:"refresh"`
}

type ListResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Content []struct {
			Name     string      `json:"name"`
			Size     int         `json:"size"`
			IsDir    bool        `json:"is_dir"`
			Modified time.Time   `json:"modified"`
			Created  time.Time   `json:"created"`
			Sign     string      `json:"sign"`
			Thumb    string      `json:"thumb"`
			Type     int         `json:"type"`
			Hashinfo string      `json:"hashinfo"`
			HashInfo interface{} `json:"hash_info"`
		} `json:"content"`
		Total    int    `json:"total"`
		Readme   string `json:"readme"`
		Header   string `json:"header"`
		Write    bool   `json:"write"`
		Provider string `json:"provider"`
	} `json:"data"`
}
