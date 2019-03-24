package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	retry "github.com/avast/retry-go"
	log "github.com/sirupsen/logrus"

	"net/http"
)


type Client interface {
	GetUrl(string, string) ([]byte, int, error)
}

type RealClient struct{}

func (c RealClient) GetUrl(url, path string) ([]byte, int, error) {
	var body []byte

	url = fmt.Sprint(url, path)
	var retrySum uint
	var code int

	err := retry.Do(
		func() error {
			log.Print("Doing request to URL: ", url, "Retry Number: ", retrySum)
			resp, err := http.Get(url)
			if err != nil {
				log.Error("error doing request number: ", retrySum, "Error: ", err)
				retrySum += 1
				code = http.StatusInternalServerError
				return err
			}

			if resp.StatusCode > 299 {
				err = errors.New(fmt.Sprint("Server Response Error: ", resp.StatusCode, "Request Number: ", retrySum))
				log.Error(err)
				code = resp.StatusCode
				retrySum += 1
				return err
			}

			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Error("error reading body number: ", retrySum, "Error: ", err)
				retrySum += 1
				return err
			}

			code = resp.StatusCode
			return nil
		},
		retry.Attempts(5),
	)

	if err != nil {
		log.Error("Error after all retrys ", err)
		return nil, code, err
	}

	return body, code, nil
}

type TestClient struct{}

func (t TestClient) GetUrl(id, path string) ([]byte, int, error) {
	log.Print("I'm the fake client bruh")
	return []byte("1"), 200, nil
}

func InitClient() Client {
	if os.Getenv("ENV") == "TEST" {
		return new(TestClient)
	}
	return new(RealClient)
}
