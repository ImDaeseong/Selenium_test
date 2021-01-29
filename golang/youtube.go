package main

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	path = `E:\GoApp\src\chromedriver.exe`
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

	err = wd.Get("https://youtube.com")
	if err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)

	id, _ := wd.FindElement(selenium.ByCSSSelector, `#search`)
	id.Clear()
	id.SendKeys("아이유")

	searchbtn, _ := wd.FindElement(selenium.ByCSSSelector, `#search-icon-legacy > yt-icon`)
	searchbtn.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	pagesource, _ := wd.PageSource()

	fmt.Println(url)
	fmt.Println(title)
	fmt.Println(pagesource)

	time.Sleep(10 * time.Second)
}
