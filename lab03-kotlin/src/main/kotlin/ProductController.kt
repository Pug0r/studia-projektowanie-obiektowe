package main

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/api")
class ProductController(
    private val authServiceEager: AuthServiceEager
) {

    private val products = listOf(
        Product(1, "Laptop", 4999.99),
        Product(2, "Mouse", 129.99),
        Product(3, "Keyboard", 259.99)
    )

    @GetMapping("/products")
    fun getProducts(): List<Product> = products

    @PostMapping("/auth/login")
    fun login(@RequestBody request: LoginRequest): LoginResponse {
        val authenticated = authServiceEager.authenticate(
            request.username,
            request.password
        )
        return LoginResponse(authenticated)
    }
}

data class LoginRequest(
    val username: String,
    val password: String
)

data class LoginResponse(
    val authenticated: Boolean
)