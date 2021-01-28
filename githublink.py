import time
from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    options = webdriver.ChromeOptions()

    # 브라우저 안 띄우기
    # options.add_argument("headless")

    # 그래픽 카드 사용 안하기
    options.add_argument("disable-gpu")

    # 브라우저 크기 지정
    options.add_argument("window-size=1920,1080")

    browser = webdriver.Chrome('./chromedriver.exe', options=options)

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://github.com/ImDaeseong?tab=repositories")

    # DaeseongSwift link 페이지 호출
    browser.implicitly_wait(10)
    browser.find_element_by_link_text('DaeseongSwift').click()

    # 페이지 읽기
    # print(browser.page_source)

    # 3초후 종료
    time.sleep(3)

    # 드라이브 종료
    browser.quit()
