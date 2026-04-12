Zadanie 3 Wzorce kreacyjne
Spring Boot (Kotlin)
Proszę stworzyć prosty serwis do autoryzacji, który zasymuluje
autoryzację użytkownika za pomocą przesłanej nazwy użytkownika oraz
hasła. Serwis powinien zostać wstrzyknięty do kontrolera (4.5).
Aplikacja ma oczywiście zawierać jeden kontroler i powinna zostać
napisana w języku Kotlin. Oparta powinna zostać na frameworku Spring
Boot. Serwis do autoryzacji powinien być singletonem.

- [x] 3.0 Należy stworzyć jeden kontroler wraz z danymi wyświetlanymi z
listy na endpoint’cie w formacie JSON - Kotlin + Spring Boot
- [x] 3.5 Należy stworzyć klasę do autoryzacji (mock) jako Singleton w
formie eager
- [x] 4.0 Należy obsłużyć dane autoryzacji przekazywane przez użytkownika
- [x] 4.5 Należy wstrzyknąć singleton do głównej klasy via @Autowired lub
kontruktor (constructor injection)
- [ ] 5.0 Obok wersji Eager do wyboru powinna być wersja Singletona w wersji
lazy