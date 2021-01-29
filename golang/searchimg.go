// searchimg
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
	err = wd.Get("https://www.google.co.kr/imghp?hl=ko&tab=wi&ogbl")
	if err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//검색창에 검색어 입력후 엔터
	id, _ := wd.FindElement(selenium.ByName, `q`)
	id.Clear()
	id.SendKeys("이미지 캐릭터")
	id.SendKeys(selenium.EnterKey)

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//링크 클릭
	link, _ := wd.FindElement(selenium.ByLinkText, "애니메이션")
	link.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	pagesource, _ := wd.PageSource()

	fmt.Println(url)
	fmt.Println(title)
	fmt.Println(pagesource)

	time.Sleep(60 * time.Second)
}
