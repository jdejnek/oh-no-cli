package http_client

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"

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

func CallApi(method string, path string, queryParams QueryParams) (string, error) {
	client := &http.Client{}
	u, err := url.Parse("http://localhost:8443")
	if err != nil {
		return "", err
	}

	u.Path = path

	q := u.Query()
	for key, value := range queryParams.ToMap() {
		q.Set(key, value)
	}
	u.RawQuery = q.Encode()

	fmt.Print(u.String())

	req, err := http.NewRequest(method, u.String(), nil)
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
	fmt.Printf("%s", body)
	return u.String(), nil
}
