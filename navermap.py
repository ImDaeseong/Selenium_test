from selenium import webdriver
import time

from selenium.webdriver.common.keys import Keys

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://v4.map.naver.com")

    # 안내메시지 끄기
    serachbutton = browser.find_element_by_xpath('//*[@id="dday_popup"]/div[2]/button')
    serachbutton.click()

    # 검색창에 검색어 입력후 엔터
    serachstring = browser.find_element_by_id("search-input")
    serachstring.send_keys('치킨')
    # serachstring.send_keys(Keys.RETURN)

    # 검색 버튼 클릭
    search_button = browser.find_element_by_css_selector("button.spm")
    search_button.click()

    # 페이지 읽기
    print(browser.page_source)

    # 드라이브 종료
    browser.quit()