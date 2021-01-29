// github
package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	path = `c:\GoApp\chromedriver.exe`
	port = 8080
)

func main() {

	opts := []selenium.ServiceOption{}

	server, err := selenium.NewChromeDriverService(path, port, opts...)
	if err != nil {
		fmt.Println(err)
	}
	defer server.Stop()

	c := selenium.Capabilities{"browserName": "chrome"}

	wd, err := selenium.NewRemote(c, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		fmt.Println(err)
	}
	defer wd.Quit()

	err = wd.Get("https://github.com/ImDaeseong")
	if err != nil {
		panic(err)
	}

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()

	fmt.Println(url)
	fmt.Println(title)

	time.Sleep(10 * time.Second)
}
