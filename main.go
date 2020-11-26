package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type GithubRelease struct {
	Assets []GithubAsset `json:"assets"`
}

type GithubAsset struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func main() {
	if len(os.Args) != 5 {
		fmt.Printf("Usage is '%s repository_name tag asset_file output_file'\n", os.Args[0])
		os.Exit(1)
	}

	repo := os.Args[1]
	tag := os.Args[2]
	file := os.Args[3]
	outputFile := os.Args[4]

	releaseUrl := "https://api.github.com/repos/" + repo + "/releases/tags/" + tag
	token := os.Getenv("GITHUB_TOKEN")

	req, err := http.NewRequest("GET", releaseUrl, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req.Header.Add("User-Agent", "Gros")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	release := &GithubRelease{}
	err = json.Unmarshal(body, release)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	assetURL := ""
	for _, asset := range release.Assets {
		if file == asset.Name {
			assetURL = asset.URL
			break
		}
	}

	if assetURL == "" {
		fmt.Println("Cannot find asset " + file)
		os.Exit(1)
	}

	req, err = http.NewRequest("GET", assetURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	req.Header.Add("User-Agent", "Gros")
	req.Header.Add("Accept", "application/octet-stream")
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	err = ioutil.WriteFile(outputFile, body, 0755)
}
