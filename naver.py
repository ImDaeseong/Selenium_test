import time
from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    browser.implicitly_wait(10)
    browser.get("https://www.naver.com")

    # 로그인 페이지로 이동
    loginpage = browser.find_element_by_class_name("link_login")
    loginpage.click()

    # 아이디/비밀번호 입력
    sid = ''
    spwd = ''

    # 아이디 입력
    browser.implicitly_wait(10)
    id = browser.find_element_by_id('id')
    id.send_keys(sid)

    # 비밀번호 입력
    browser.implicitly_wait(10)
    pw = browser.find_element_by_id('pw')
    pw.send_keys(spwd)

    # 로그인 버튼 클릭
    btnlogin = browser.find_element_by_id("log.login")
    btnlogin.click()

    # 페이지 읽기
    # print(browser.page_source)

    # 3초후 종료
    time.sleep(3)

    # 드라이브 종료
    browser.quit()
