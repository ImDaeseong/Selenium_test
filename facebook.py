from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://www.facebook.com/login/?privacy_mutation_token=eyJ0eXBlIjowLCJjcmVhdGlvbl90aW1lIjoxNjExODI1NjIwLCJjYWxsc2l0ZV9pZCI6NzMxOTQxNDIwNzMyOTEwfQ%3D%3D")

    # 아이디 입력
    browser.implicitly_wait(10)
    id = browser.find_element_by_id('email')
    id.send_keys("메일주소")

    # 비밀번호 입력
    browser.implicitly_wait(10)
    pwd = browser.find_element_by_id('pass')
    pwd.send_keys("비밀번호")

    # 로그인 버튼 클릭
    browser.implicitly_wait(10)
    btnlogin = browser.find_element_by_id("loginbutton")
    btnlogin.click()


    # 페이지 읽기
    # print(browser.page_source)

    # 드라이브 종료
    browser.quit()
