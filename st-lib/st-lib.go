package stlib

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type StLib struct {
	root  string
	token string
}

type MetaPagination struct {
	Total int
	Page  int
	Limit int
}

func NewStLib(token string) StLib {
	return StLib{
		root:  "https://api.spacetraders.io/v2/",
		token: token,
	}
}

func getWps(wp string) *WaypointSymbol {
	if wp == "" {
		return nil
	}
	wps := strings.Split(wp, "-")

	return &WaypointSymbol{
		Sector:   wps[0],
		System:   wps[0] + "-" + wps[1],
		Waypoint: wp,
	}
}

func (st *StLib) GetUrl(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", st.root+url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", st.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
func (st *StLib) PostUrl(url string, data *[]byte) ([]byte, error) {
	byteData := bytes.NewBuffer([]byte{})
	if data != nil {
		byteData = bytes.NewBuffer(*data)
	}
	req, err := http.NewRequest("POST", st.root+url, byteData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", st.token))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(string(body))
	}

	return body, nil
}
