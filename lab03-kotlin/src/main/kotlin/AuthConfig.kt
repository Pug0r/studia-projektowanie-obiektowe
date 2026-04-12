package main

import org.springframework.context.annotation.Bean
import org.springframework.context.annotation.Configuration

@Configuration
class AuthConfig {

    @Bean
    fun authServiceEager(): AuthServiceEager = AuthServiceEager.instance
}