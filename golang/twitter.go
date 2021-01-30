// twitter
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

	//사이트 호출
	err = wd.Get("https://twitter.com/login/")
	if err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//아이디 입력
	id, _ := wd.FindElement(selenium.ByName, `session[username_or_email]`)
	id.Clear()
	id.SendKeys("메일주소")

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//비밀번호 입력
	pwd, _ := wd.FindElement(selenium.ByName, `session[password]`)
	pwd.Clear()
	pwd.SendKeys("비밀번호")

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//트위터 로그인 클릭
	serachbutton, _ := wd.FindElement(selenium.ByXPATH, "//div[@data-testid='LoginForm_Login_Button']")
	serachbutton.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	pagesource, _ := wd.PageSource()

	fmt.Println(url)
	fmt.Println(title)
	fmt.Println(pagesource)

	time.Sleep(60 * time.Second)
}
