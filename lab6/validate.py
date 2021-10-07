from selenium import webdriver
from selenium.webdriver.chrome.options import Options
import time
 
 
options = Options()
options.add_argument("--headless")
 
chrome = webdriver.Chrome('./chromedriver', chrome_options=options)
chrome.get("http://0.0.0.0:8080")