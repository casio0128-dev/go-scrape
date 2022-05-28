package main

import (
	"encoding/json"
	"fmt"
	"github.com/sclevine/agouti"
	"go-scrape/browser"
	"go-scrape/profile"
	"go-scrape/web"
	"net/http"
	"os"
	"strings"
)

func main() {
	prof := parse("profile.json")
	targetProfile := &(prof[0])

	web.SetRouting(map[string]interface{}{
		web.Index: nil,
	})
	http.ListenAndServe(":8080", nil)

	d := agouti.ChromeDriver(targetProfile.BrowserOption.ChromeOption())

	if err := d.Start(); err != nil {
		panic(err)
	}
	defer func() {
		if err := d.Stop(); err != nil {
			panic(err)
		}
	}()

	page, err := d.NewPage(agouti.Browser("chrome"))
	if err != nil {
		panic(err)
	}

	if err := browser.Do(page, targetProfile); err != nil {
		panic(err)
	}
}

func init() {
	if err := setDriverPath(); err != nil {
		panic(err)
	}
}

func setDriverPath() error {
	current, err := os.Getwd()
	if err != nil {
		return err
	}

	pathEnv := []string{os.Getenv("PATH"), fmt.Sprintf("%s%sdrivers", current, string(os.PathSeparator))}
	fmt.Println("PATH=>", pathEnv, " os.PathSeparator=>", string(os.PathSeparator), " os.PathListSeparator=>", string(os.PathListSeparator))

	return os.Setenv("PATH", strings.Join(pathEnv, string(os.PathListSeparator)))
}

func loadJSON(path string) (interface{}, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var result interface{}
	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func parse(path string) profile.Profiles {
	bytes, _ := os.ReadFile(path)
	var profiles profile.Profiles
	if err := json.Unmarshal(bytes, &profiles); err != nil {
		panic(err)
	}
	return profiles
}

//func loadJSON(path string, obj interface{}) (interface{}, error) {
//	bytes, err := os.ReadFile(path)
//	if err != nil {
//		return nil, err
//	}
//	//var result interface{}
//	if err := json.Unmarshal(bytes, &obj); err != nil {
//		return nil, err
//	}
//	return obj, nil
//}
