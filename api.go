package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strings"
)

const (
	base = "https://api.massbots.xyz"
)

type API struct {
	botID  string
	token  string
	client *http.Client
}

func New(botID, token string) API {
	return API{
		botID:  botID,
		token:  token,
		client: &http.Client{},
	}
}

type Error struct {
	Code    int    `json:"-"`
	Message string `json:"error"`
}

func (e *Error) Error() string {
	return strings.TrimSpace(
		fmt.Sprintf("massbots/sdk: (%d) %s", e.Code, e.Message),
	)
}

func get[T any](a API, e string) (v T, _ error) {
	req, err := http.NewRequest(http.MethodGet, base+"/"+e, nil)
	if err != nil {
		return v, err
	}

	req.Header.Set("User-Agent", "massbots/sdk/go")
	req.Header.Set("X-Token", a.token)
	req.Header.Set("X-Bot-Id", a.botID)

	build, ok := debug.ReadBuildInfo()
	if ok {
		req.Header.Set("X-Version", build.Main.Version)
	}

	resp, err := a.client.Do(req)
	if err != nil {
		return v, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return v, err
	}

	if resp.StatusCode != http.StatusOK {
		e := Error{Code: resp.StatusCode}
		_ = json.Unmarshal(data, &e)
		return v, &e
	}

	return v, json.Unmarshal(data, &v)
}
