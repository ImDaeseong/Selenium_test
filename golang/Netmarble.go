// Netmarble
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
	err = wd.Get("https://finance.naver.com/item/main.nhn?code=251270")
	if err != nil {
		panic(err)
	}

	//1분에 한번씩 refresh
	Ticker := time.NewTicker(60 * time.Second)
	defer Ticker.Stop()

	bDoneChan := make(chan bool)

	//종료
	go func() {
		time.Sleep(3600 * time.Second)
		bDoneChan <- true
	}()

	for {
		select {
		case <-Ticker.C:
			//웹페이지 refresh
			wd.Refresh()

		case <-bDoneChan:
			//리턴하면 종료됨
			return
		}
	}

}
