package main

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api/products")
class ProductController {

    private val products = listOf(
        Product(1, "Laptop", 4999.99),
        Product(2, "Mouse", 129.99),
        Product(3, "Keyboard", 259.99)
    )

    @GetMapping
    fun getProducts(): List<Product> = products
}