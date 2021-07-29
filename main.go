package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func listRepos(username string) (resp *http.Response, err error) {
	enpoint := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	return http.Get(enpoint)
}

func main() {
	res, err := listRepos("throberto")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}
