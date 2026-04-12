package main

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration
import org.springframework.core.env.Environment

@Configuration
class AuthConfig(
    private val environment: Environment
) {

    @Bean
    fun authService(): AuthService {
        val mode = environment.getProperty("app.auth.singleton-mode", "eager")
        return if (mode.equals("lazy", ignoreCase = true)) {
            AuthServiceLazy.instance
        } else {
            AuthServiceEager.instance
        }
    }
}