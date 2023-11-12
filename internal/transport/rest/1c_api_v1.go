package rest

import (
	"log"
	"net/http"
)

func InitialConnectionInfoBase(baseURL, login, pass string) (bool, *http.Client) {

	Get, _ := http.NewRequest(http.MethodGet, baseURL, nil)
	Get.SetBasicAuth(login, pass)

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       0,
	}
	response, err := client.Do(Get)

	result := false

	if response.StatusCode == http.StatusOK {
		result = true
	}

	if err != nil {
		log.Fatal(err)
	}

	return result, &client

}

func GetStatusInfoBase(baseURL, token string) {

}
