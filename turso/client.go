package turso

import (
	"fmt"
	"goh/go-htmx/utils"
	"io"
	"net/http"
)

type Api struct {
	token   string
	orgName string
}

func CreateApi(settings utils.Settings) Api {
	return Api{token: settings.ApiToken, orgName: settings.OrgName}
}

func (api *Api) GetLocations() ([]byte, error) {
	url := fmt.Sprintf("https://api.turso.tech/v1/organizations/%s/locations", api.orgName)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return nil, err
	}

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return nil, err
	}
	fmt.Printf("client: response body: %s\n", resBody)
	return resBody, nil
}
