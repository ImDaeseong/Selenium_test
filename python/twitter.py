from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://twitter.com/login/")

    # 아이디 입력
    browser.implicitly_wait(10)
    id = browser.find_elements_by_name('session[username_or_email]')[0]
    id.send_keys("메일주소")

    # 비밀번호 입력
    browser.implicitly_wait(10)
    pwd = browser.find_elements_by_name('session[password]')[0]
    pwd.send_keys("비밀번호")

    # 트위터 로그인 클릭
    serachbutton = browser.find_element_by_xpath('//*[@id="react-root"]/div/div/div[2]/main/div/div/div[2]/form/div/div[3]/div/div')
    serachbutton.click()



    # 페이지 읽기
    # print(browser.page_source)

    # 드라이브 종료
    browser.quit()
