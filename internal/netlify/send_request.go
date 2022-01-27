package netlify

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
)

const apiUrlPattern string = "https://api.netlify.com/api/v1/sites/%s"

func GetSite(siteId, deployToken string) error {
	uri, _ := url.Parse(fmt.Sprintf(apiUrlPattern, siteId))
	request := buildRequest(http.MethodGet, uri.String(), nil, deployToken)
	return sendRequest(request)
}

func PostDeploy(siteId, deployToken, zipFilePath string) error {
	zipFile, err := os.Open(zipFilePath)
	if err != nil {
		handleError(err, "Error opening site zip")
	}
	defer zipFile.Close()

	// url := path.Join(fmt.Sprintf(apiUrlPattern, siteId), "deploys")
	uri, _ := url.Parse(fmt.Sprintf(apiUrlPattern, siteId))
	uri.Path = path.Join(uri.Path, "deploys")
	request := buildRequest(http.MethodPost, uri.String(), zipFile, deployToken)
	request.Header.Add("Content-Type", "application/zip")

	return sendRequest(request)
}

func buildRequest(httpMethod string, uri string, body io.Reader, deployToken string) *http.Request {
	request, err := http.NewRequest(httpMethod, uri, body)
	if err != nil {
		handleError(err, "Error creating deploy request")
	}
	request.Header.Add("Authorization", "Bearer "+deployToken)
	return request
}

func sendRequest(request *http.Request) error {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		return errors.New("error sending request; response code " + strconv.Itoa(response.StatusCode))
	}
	return nil
}

func handleError(err error, context string) {

}
