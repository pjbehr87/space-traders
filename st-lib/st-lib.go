package stlib

import (
	"fmt"
	"io"
	"net/http"
)

type StLib struct {
	root  string
	token string
}

func NewStLib(token string) StLib {
	return StLib{
		root:  "https://api.spacetraders.io/v2/",
		token: token,
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
