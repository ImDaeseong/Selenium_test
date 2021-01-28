from selenium import webdriver

if __name__ == '__main__':

    # Chrome WebDriver 로 chrome.exe 실행
    browser = webdriver.Chrome('./chromedriver.exe')

    """
    # 사이트 호출 - 인스타 계정으로 로그인
    browser.implicitly_wait(10)
    browser.get("https://www.instagram.com/")

    # 아이디 입력
    browser.implicitly_wait(10)
    id = browser.find_element_by_name('username')
    id.send_keys("메일주소")

    # 비밀번호 입력
    browser.implicitly_wait(10)
    pwd = browser.find_element_by_name('password')
    pwd.send_keys("비밀번호")

    # 로그인 버튼 클릭
    btnlogin = browser.find_element_by_xpath('//*[@id="loginForm"]/div/div[3]')
    btnlogin.click()

    # 페이스북으로 로그인하기
    search_button = browser.find_element_by_css_selector("span.KPnG0")
    search_button.click()
    """

    # 사이트 호출 - 페이스북 계정으로 로그인
    browser.implicitly_wait(10)
    browser.get("https://www.facebook.com/login.php?skip_api_login=1&api_key=124024574287414&kid_directed_site=0&app_id=124024574287414&signed_next=1&next=https%3A%2F%2Fwww.facebook.com%2Fdialog%2Foauth%3Fclient_id%3D124024574287414%26redirect_uri%3Dhttps%253A%252F%252Fwww.instagram.com%252Faccounts%252Fsignup%252F%26state%3D%257B%2522fbLoginKey%2522%253A%2522m3vg001kmr4mzxxdnng1b0qderfivxoohtra1r1v8r4io1ulswe3%2522%252C%2522fbLoginReturnURL%2522%253A%2522%252F%2522%257D%26scope%3Demail%26response_type%3Dcode%252Cgranted_scopes%26locale%3Dko_KR%26ret%3Dlogin%26fbapp_pres%3D0%26logger_id%3D3dd4f19e-30be-41f9-afb0-b3aaf36c7c7b%26tp%3Dunspecified&cancel_url=https%3A%2F%2Fwww.instagram.com%2Faccounts%2Fsignup%2F%3Ferror%3Daccess_denied%26error_code%3D200%26error_description%3DPermissions%2Berror%26error_reason%3Duser_denied%26state%3D%257B%2522fbLoginKey%2522%253A%2522m3vg001kmr4mzxxdnng1b0qderfivxoohtra1r1v8r4io1ulswe3%2522%252C%2522fbLoginReturnURL%2522%253A%2522%252F%2522%257D%23_%3D_&display=page&locale=ko_KR&pl_dbl=0")

    # 아이디 입력
    browser.implicitly_wait(10)
    id = browser.find_element_by_name('email')
    id.send_keys("메일주소")

    # 비밀번호 입력
    browser.implicitly_wait(10)
    pwd = browser.find_element_by_name('pass')
    pwd.send_keys("비밀번호")

    # 로그인 버튼 클릭
    browser.implicitly_wait(10)
    btnlogin = browser.find_element_by_id("loginbutton")
    btnlogin.click()

    # 인스타 계속하기 버튼 클릭
    browser.implicitly_wait(10)
    serachbutton = browser.find_element_by_xpath('//*[@id="react-root"]/section/main/article/div[2]/div/div/div[2]/button/div/div')
    serachbutton.click()

    # 나중에 햐기 버튼 클릭
    browser.implicitly_wait(10)
    serachbutton = browser.find_element_by_xpath('/html/body/div[4]/div/div/div/div[3]/button[2]')
    serachbutton.click()


    # 페이지 읽기
    # print(browser.page_source)

    # 드라이브 종료
    browser.quit()
