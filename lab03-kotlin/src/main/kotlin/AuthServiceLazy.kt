package main

class AuthServiceLazy private constructor() : AuthService {

    override
    fun authenticate(username: String, password: String): Boolean {
        return username == "admin" && password == "admin"
    }

    companion object {
        val instance: AuthServiceLazy by lazy { AuthServiceLazy() }
    }
}