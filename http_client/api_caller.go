package http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/TylerBrock/colorjson"
	_ "github.com/joho/godotenv/autoload"
)

type QueryParams struct {
	Params map[string]interface{}
}

// ToMap converts QueryParams struct into a map[string]string
func (qp *QueryParams) ToMap() map[string]string {
	result := make(map[string]string)
	for key, value := range qp.Params {
		switch v := value.(type) {
		case string:
			result[key] = v
		case int:
			result[key] = strconv.Itoa(v)
			// Add more cases as needed for other types
		}
	}
	return result
}

func buildURL(path string, queryParams QueryParams) string {
	host := os.Getenv("HOST")
	u, err := url.Parse(host)
	if err != nil {
		return err.Error()
	}

	u.Path = path
	q := u.Query()
	for key, value := range queryParams.ToMap() {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	return u.String()
}

func parseJSON(body []byte) string {
	f := colorjson.NewFormatter()
	f.Indent = 2

	var responseData map[string]interface{}
	json.Unmarshal([]byte(body), &responseData)

	d, _ := f.Marshal(responseData)

	return string(d)
}

func CallApiWithParams(method string, path string, queryParams QueryParams) (string, error) {
	client := &http.Client{}

	urlString := buildURL(path, queryParams)
	req, err := http.NewRequest(method, urlString, nil)
	if err != nil {
		return "", err
	}

	auth := os.Getenv("API_KEY")
	req.Header.Add("authorization", auth)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	output := parseJSON(body)
	fmt.Println(string(output))

	return string(output), nil
}

func CallApiWithPath(method string, path string, subpath string) string {
	client := &http.Client{}

	host := os.Getenv("HOST")
	u, err := url.Parse(host)
	if err != nil {
		return err.Error()
	}

	extendedPath := fmt.Sprintf("%s%s", path, subpath)
	u.Path = extendedPath

	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return err.Error()
	}

	fmt.Println(u.String())
	auth := os.Getenv("API_KEY")
	req.Header.Add("authorization", auth)

	resp, err := client.Do(req)
	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	if path == "/connectors" {
		fmt.Println(string(body))
		return string(body)
	}

	output := parseJSON(body)
	fmt.Println(output)

	return output
}
