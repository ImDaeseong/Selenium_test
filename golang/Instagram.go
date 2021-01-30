// Instagram
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

	/*
		//사이트 호출
		err = wd.Get("https://www.instagram.com/")

		if err != nil {
			panic(err)
		}

		wd.SetImplicitWaitTimeout(10 * time.Second)

		//아이디 입력
		id, _ := wd.FindElement(selenium.ByName, `username`)
		id.Clear()
		id.SendKeys("메일주소")

		wd.SetImplicitWaitTimeout(10 * time.Second)

		//비밀번호 입력
		pwd, _ := wd.FindElement(selenium.ByName, `password`)
		pwd.Clear()
		pwd.SendKeys("비밀번호")

		wd.SetImplicitWaitTimeout(10 * time.Second)

		//로그인 버튼 클릭
		btnlogin, _ := wd.FindElement(selenium.ByID, "loginForm")
		//btnlogin, _ := wd.FindElement(selenium.ByXPATH, "//div[@id='loginForm']")
		btnlogin.Click()

		wd.SetImplicitWaitTimeout(10 * time.Second)

		//페이스북으로 로그인하기
		search_button, _ := wd.FindElement(selenium.ByCSSSelector, `span.KPnG0`)
		search_button.Click()
	*/

	//사이트 호출
	err = wd.Get("https://www.facebook.com/login.php?skip_api_login=1&api_key=124024574287414&kid_directed_site=0&app_id=124024574287414&signed_next=1&next=https%3A%2F%2Fwww.facebook.com%2Fdialog%2Foauth%3Fclient_id%3D124024574287414%26redirect_uri%3Dhttps%253A%252F%252Fwww.instagram.com%252Faccounts%252Fsignup%252F%26state%3D%257B%2522fbLoginKey%2522%253A%2522m3vg001kmr4mzxxdnng1b0qderfivxoohtra1r1v8r4io1ulswe3%2522%252C%2522fbLoginReturnURL%2522%253A%2522%252F%2522%257D%26scope%3Demail%26response_type%3Dcode%252Cgranted_scopes%26locale%3Dko_KR%26ret%3Dlogin%26fbapp_pres%3D0%26logger_id%3D3dd4f19e-30be-41f9-afb0-b3aaf36c7c7b%26tp%3Dunspecified&cancel_url=https%3A%2F%2Fwww.instagram.com%2Faccounts%2Fsignup%2F%3Ferror%3Daccess_denied%26error_code%3D200%26error_description%3DPermissions%2Berror%26error_reason%3Duser_denied%26state%3D%257B%2522fbLoginKey%2522%253A%2522m3vg001kmr4mzxxdnng1b0qderfivxoohtra1r1v8r4io1ulswe3%2522%252C%2522fbLoginReturnURL%2522%253A%2522%252F%2522%257D%23_%3D_&display=page&locale=ko_KR&pl_dbl=0")

	if err != nil {
		panic(err)
	}

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//아이디 입력
	id, _ := wd.FindElement(selenium.ByName, `email`)
	id.Clear()
	id.SendKeys("메일주소")

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//비밀번호 입력
	pwd, _ := wd.FindElement(selenium.ByName, `pass`)
	pwd.Clear()
	pwd.SendKeys("비밀번호")

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//로그인 버튼 클릭
	btnlogin, _ := wd.FindElement(selenium.ByID, "loginbutton")
	btnlogin.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//인스타 계속하기 버튼 클릭
	serachbutton, _ := wd.FindElement(selenium.ByXPATH, "//button[@type='button']")
	serachbutton.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	//나중에 햐기 버튼 클릭
	nextbutton, _ := wd.FindElement(selenium.ByXPATH, "//button[@tabindex='0']")
	nextbutton.Click()

	wd.SetImplicitWaitTimeout(10 * time.Second)

	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	pagesource, _ := wd.PageSource()

	fmt.Println(url)
	fmt.Println(title)
	fmt.Println(pagesource)

	time.Sleep(60 * time.Second)
}
