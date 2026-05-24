import unittest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

BASE_URL = "http://localhost:5173/payments"


class PaymentsXssTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        options = webdriver.ChromeOptions()
        options.add_argument("--headless=new")
        cls.driver = webdriver.Chrome(options=options)
        cls.wait = WebDriverWait(cls.driver, 10)

    @classmethod
    def tearDownClass(cls) -> None:
        cls.driver.quit()

    def setUp(self) -> None:
        self.driver.get(BASE_URL)
        self.wait.until(EC.presence_of_element_located((By.CSS_SELECTOR, "form")))

    def test_xss_payload_is_not_executed(self) -> None:
        payload = '<img src=x onerror="window.xssFlag=1">'

        self.driver.execute_script("window.xssFlag = 0")

        form = self.driver.find_element(By.CSS_SELECTOR, "form")
        email_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Email"]')
        amount_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Kwota"]')
        submit_button = self.driver.find_element(By.CSS_SELECTOR, 'button[type="submit"]')

        self.driver.execute_script("arguments[0].noValidate = true", form)
        email_input.clear()
        email_input.send_keys(payload)
        amount_input.clear()
        amount_input.send_keys("1")
        submit_button.click()

        def payload_present(driver: webdriver.Chrome) -> bool:
            items = driver.find_elements(By.CSS_SELECTOR, "ul li")
            return any(payload in item.text for item in items)

        self.wait.until(payload_present)
        xss_flag = self.driver.execute_script("return window.xssFlag")
        self.assertEqual(xss_flag, 0)


if __name__ == "__main__":
    unittest.main()
