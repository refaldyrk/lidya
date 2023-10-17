package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/refaldyrk/lidya/constant"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func GenerateCode(filename, lang string) {
	if filename == "" {
		panic("Arg Can't Be Empty")
	}

	readFile, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	data := struct {
		Code       string `json:"code"`
		Lang       string `json:"lang"`
		Format     string `json:"format"`
		Theme      string `json:"theme"`
		Upscale    int    `json:"upscale"`
		LineNumber bool   `json:"lineNumber"`
		Border     struct {
			Thickness int    `json:"thickness"`
			Colour    string `json:"colour"`
			Radius    int    `json:"radius"`
		} `json:"border"`
	}{
		Code:       string(readFile),
		Lang:       lang,
		Theme:      "github-dark-dimmed",
		Format:     "jpeg",
		Upscale:    3,
		LineNumber: true,
		Border: struct {
			Thickness int    `json:"thickness"`
			Colour    string `json:"colour"`
			Radius    int    `json:"radius"`
		}{
			Thickness: 20,
			Colour:    "#A0ADB6",
			Radius:    4,
		},
	}

	// Marshal data ke JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	client := http.Client{}
	client.Timeout = 1 * time.Minute
	req, err := http.NewRequest("POST", constant.URL, bytes.NewBuffer(jsonData))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	cli, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(cli.Body)

	byteBody, err := io.ReadAll(cli.Body)
	if err != nil {
		panic(err)
	}

	// Simpan gambar ke file
	err = ioutil.WriteFile(fmt.Sprintf("%d.jpeg", time.Now().Unix()), byteBody, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("[SUCCESS]")
}
