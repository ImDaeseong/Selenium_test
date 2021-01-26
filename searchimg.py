from selenium import webdriver
from selenium.webdriver.common.keys import Keys

if __name__ == '__main__':
    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://www.google.co.kr/imghp?hl=ko&tab=wi&ogbl")

    # 검색창에 검색어 입력후 엔터
    serachstring = browser.find_element_by_name("q")
    serachstring.send_keys('이미지 캐릭터')
    serachstring.send_keys(Keys.RETURN)

    # 링크 클릭
    browser.find_element_by_link_text('애니메이션').click()

    # 페이지 읽기
    print(browser.page_source)

    # 드라이브 종료
    browser.quit()
