import unittest
from selenium import webdriver
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

BASE_URL = "http://localhost:5173/payments"


def get_validation_message(driver: webdriver.Chrome, element) -> str:
    return driver.execute_script("return arguments[0].validationMessage", element)


def is_valid(driver: webdriver.Chrome, element) -> bool:
    return driver.execute_script("return arguments[0].checkValidity()", element)


class PaymentsFormValidationTests(unittest.TestCase):
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

    def test_required_fields_are_enforced(self) -> None:
        email_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Email"]')
        amount_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Kwota"]')
        submit_button = self.driver.find_element(By.CSS_SELECTOR, 'button[type="submit"]')

        submit_button.click()

        self.assertFalse(is_valid(self.driver, email_input))
        self.assertFalse(is_valid(self.driver, amount_input))
        self.assertNotEqual(get_validation_message(self.driver, email_input), "")
        self.assertNotEqual(get_validation_message(self.driver, amount_input), "")

    def test_invalid_email_format_is_rejected(self) -> None:
        email_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Email"]')
        amount_input = self.driver.find_element(By.CSS_SELECTOR, 'input[placeholder="Kwota"]')
        submit_button = self.driver.find_element(By.CSS_SELECTOR, 'button[type="submit"]')

        email_input.send_keys("not-an-email")
        amount_input.send_keys("10")
        submit_button.click()

        self.assertFalse(is_valid(self.driver, email_input))
        self.assertTrue(is_valid(self.driver, amount_input))
        self.assertNotEqual(get_validation_message(self.driver, email_input), "")


if __name__ == "__main__":
    unittest.main()
