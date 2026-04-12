package main

interface AuthService {
    fun authenticate(username: String, password: String): Boolean
}