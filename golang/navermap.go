// navermap
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

	//사이트 호출
	err = wd.Get("https://v4.map.naver.com")
	if err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//안내메시지 끄기
	noticebtn, _ := wd.FindElement(selenium.ByCSSSelector, `span.img`)
	noticebtn.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//검색창에 검색어 입력후 엔터
	id, _ := wd.FindElement(selenium.ByID, `search-input`)
	id.Clear()
	id.SendKeys("치킨")

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//검색 버튼 클릭
	searchbtn, _ := wd.FindElement(selenium.ByCSSSelector, `button.spm`)
	searchbtn.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	pagesource, _ := wd.PageSource()

	fmt.Println(url)
	fmt.Println(title)
	fmt.Println(pagesource)

	time.Sleep(60 * time.Second)
}
