import time
from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('E:\Selenium_test\chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://www.google.com")

    # 페이지 읽기
    # print(browser.page_source)

    # 3초후 종료
    time.sleep(3)

    # 드라이브 종료
    browser.quit()
