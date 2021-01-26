from selenium import webdriver


if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://youtube.com")

    # 유튜브 검색창에 검색어 입력
    serachstring = browser.find_element_by_xpath('//*[@id="search"]')
    serachstring.send_keys('아이유')

    # 유튜브 검색버튼 클릭
    serachbutton = browser.find_element_by_xpath('// *[ @ id = "search-icon-legacy"]')
    serachbutton.click()

    # 페이지 읽기
    print(browser.page_source)

    # 드라이브 종료
    browser.quit()
