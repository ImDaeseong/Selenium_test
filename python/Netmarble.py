from selenium import webdriver
import time

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    # 사이트 호출
    browser.implicitly_wait(10)
    browser.get("https://finance.naver.com/item/main.nhn?code=251270")

    try:
        i = 0
        while True:
            i = i + 1

            if i > 3600:
                break

            time.sleep(60)
            browser.refresh()
    except:
        print('Error...')

    finally:

        # 페이지 읽기
        print(browser.page_source)

        # 드라이브 종료
        browser.quit()
