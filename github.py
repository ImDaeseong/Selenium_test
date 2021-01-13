from selenium import webdriver


if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://github.com/ImDaeseong")

    # 페이지 읽기
    print(browser.page_source)

    # 드라이브 종료
    browser.quit()
