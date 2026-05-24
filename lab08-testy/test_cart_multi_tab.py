import unittest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

BASE_URL = "http://localhost:5173"


class CartMultiTabTests(unittest.TestCase):
    @classmethod
    def setUpClass(cls) -> None:
        options = webdriver.ChromeOptions()
        options.add_argument("--headless=new")
        cls.driver = webdriver.Chrome(options=options)
        cls.wait = WebDriverWait(cls.driver, 10)

    @classmethod
    def tearDownClass(cls) -> None:
        cls.driver.quit()

    def test_cart_state_is_not_shared_between_tabs(self) -> None:
        self.driver.get(f"{BASE_URL}/products")
        self.wait.until(
            EC.presence_of_element_located((By.CSS_SELECTOR, "section ul li"))
        )
        add_button = self.driver.find_element(
            By.XPATH, "//button[normalize-space() = 'Dodaj do koszyka']"
        )
        add_button.click()

        self.driver.execute_script("window.open(arguments[0], '_blank')", f"{BASE_URL}/cart")
        self.driver.switch_to.window(self.driver.window_handles[1])

        self.wait.until(
            EC.presence_of_element_located((By.XPATH, "//h2[normalize-space()='Koszyk']"))
        )
        empty_cart_message = self.driver.find_element(
            By.XPATH, "//p[normalize-space()='Koszyk jest pusty']"
        )
        self.assertTrue(empty_cart_message.is_displayed())


if __name__ == "__main__":
    unittest.main()
