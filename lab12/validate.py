#!/usr/bin/env python3
import time
import os

from selenium.webdriver import Chrome
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import TimeoutException, WebDriverException
from webdriver_manager.chrome import ChromeDriverManager 
from selenium.webdriver.common.keys import Keys
from selenium.webdriver.common.by import By
import re

def check(inpu,ans,the_chrome):
    the_chrome.find_element(By.ID,'chat_input').send_keys(inpu)
    the_chrome.find_element(By.ID,'chat_input').send_keys(Keys.ENTER)
    time.sleep(3)
    print('Your answer of '+ inpu +' is "'+the_chrome.find_element(By.XPATH,"(//b)[last()]").get_attribute('innerText') +'"')
    pattern = re.compile(f".*: {ans}")

    point = 0 

    if pattern.match(chrome.find_element(By.XPATH,"(//b)[last()]").get_attribute('innerText')):
        print('get POINT 1')
        point+=1
    else:
        print('wrong answer for browser1')

    if pattern.match(chrome2.find_element(By.XPATH,"(//b)[last()]").get_attribute('innerText')):
        print('get POINT 1')
        point+=1
    else:
        print('wrong answer for browser2')

    return point

#from webdriver_manager.chrome import ChromeDriverManager
time.sleep(3)
options = Options()
options.headless = True
options.add_argument('--disable-gpu')
options.add_argument('--no-sandbox') # https://stackoverflow.com/a/45846909
options.add_argument('--disable-dev-shm-usage') # https://stackoverflow.com/a/50642913
chrome = Chrome(executable_path=ChromeDriverManager().install(),options=options)
chrome2 = Chrome(executable_path=ChromeDriverManager().install(),options=options)

#"/usr/lib/chromium-browser/chromedriver",
#ChromeDriverManager(version="83.0.4103.39").install()
chrome.get(f"http://0.0.0.0:8899")
chrome2.get(f"http://0.0.0.0:8899")

point = 0
point += check("蔡英文ooo",'蔡\\*文ooo',chrome)
point += check(" nnn馬英九",' nnn馬\\*九',chrome2)
point += check("靠夭 ",' nnn馬\\*九',chrome)
point += check("幹你娘",' nnn馬\\*九',chrome2)
point += check('韓國瑜蔡英文','韓\\*瑜蔡\\*文',chrome)
point += check('三小黃宏成台灣阿成世界偉人財神總統','韓\\*瑜蔡\\*文',chrome)
point += check('三n小黃宏成台灣阿成世界偉人財神總統','三n小黃\\*成台灣阿成世界偉人財神總統',chrome)
#i += check(8,'is not prime')

if point==14:
    print('success! You get all point')
else:
    print(f'There is something wrong! get {point} of 14 points')