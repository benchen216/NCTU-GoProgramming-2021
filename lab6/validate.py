#!/usr/bin/env python3
import time
import os

from selenium.webdriver import Chrome
from selenium.webdriver.chrome.options import Options
from selenium.common.exceptions import TimeoutException, WebDriverException
from webdriver_manager.chrome import ChromeDriverManager
time.sleep(5)
options = Options()
options.headless = True
options.add_argument('--disable-gpu')
options.add_argument('--no-sandbox') # https://stackoverflow.com/a/45846909
options.add_argument('--disable-dev-shm-usage') # https://stackoverflow.com/a/50642913
chrome = Chrome(executable_path=ChromeDriverManager().install(),options=options)
#"/usr/lib/chromium-browser/chromedriver",
#ChromeDriverManager(version="83.0.4103.39").install()
chrome.get(f"http://0.0.0.0/index.html")