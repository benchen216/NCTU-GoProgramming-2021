#!/usr/bin/env python3
import time
import os

from selenium.webdriver import Chrome
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import TimeoutException, WebDriverException
from webdriver_manager.chrome import ChromeDriverManager

def check(n,ans):
    chrome.find_element_by_id('value').clear()
    chrome.find_element_by_id('value').send_keys(str(n))
    chrome.find_element_by_id('check').click()
    print('Your answer of '+ str(n) +' is "'+chrome.find_element_by_id('answer').get_attribute('innerHTML') +'"')
    if ans in chrome.find_element_by_id('answer').get_attribute('innerHTML'):
        print('get POINT 1')
        return 1
    else:
        print('wrong answer')
        return 0

#from webdriver_manager.chrome import ChromeDriverManager
time.sleep(5)
options = Options()
options.headless = True
options.add_argument('--disable-gpu')
options.add_argument('--no-sandbox') # https://stackoverflow.com/a/45846909
options.add_argument('--disable-dev-shm-usage') # https://stackoverflow.com/a/50642913
chrome = Chrome(executable_path=ChromeDriverManager().install(),options=options)
#"/usr/lib/chromium-browser/chromedriver",
#ChromeDriverManager(version="83.0.4103.39").install()
chrome.get(f"http://0.0.0.0:8080")

i = 0
i += check(923174692024939,'is prime')
i += check(987,'is not prime')
i += check(918257,'is prime')
i += check(8,'is not prime')

if i==4:
    print('success! You get all point')
else:
    print('There is something wrong!')