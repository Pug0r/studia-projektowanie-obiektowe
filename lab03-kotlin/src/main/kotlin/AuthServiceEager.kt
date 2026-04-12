package main

class AuthServiceEager private constructor() {

    fun authenticate(username: String, password: String): Boolean {
        return username == "admin" && password == "admin"
    }

    companion object {
        val instance: AuthServiceEager = AuthServiceEager()
    }
}