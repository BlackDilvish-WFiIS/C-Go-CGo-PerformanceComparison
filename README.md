# C-Go-CGo Performance Comparison

Temat projektu: "Porównanie wydajności analogicznego kodu zaimplementowanego w go, c oraz cgo"

Projekt został zrealizowany w ramach przedmiotu akademickiego "Język GO"

## Opis udziału wykonawców w realizacji projektu
- Gabriel Naleźnik:
    - Konteneryzacja rozwiązań za pomocą dockera
    - Utworzenie skryptów do automatycznego uruchomienia testów wydajnościowych
    - Utworzenie porównania odczytywania plików i IO
    - Utworzenie porównania dodawania łańcuchów znakowych
    - Utworzenie porównania mnożenia macierzy
    - Utworzenie porównania permutacji łańcuchów znakowych
- Piotr Jasiński:
    - Utworzenie porównania dla "Hello world"
    - Utworzenie porównania dla ręcznej implementacji sortowania przez wstawianie
    - Utworzenie porównania dla rekurencyjnego ciągu Fibonacciego
    - Utworzenie porównania dla poszukiwania najdłuższego wspólnego ciągu znaków
    - Utworzenie porównania dla wyszukiwania binarnego
    - Dodanie uruchomienia implementacji w C w dwóch wersjach:
        - Brak optymalizacji dla prędkości wykonania (flaga `gcc -O0`)
        - Maksymalna optymalizacja dla prędkości wykonania (flaga `gcc -O3`)

## Uruchomienie projektu

### Wymagania
- Docker lub środowisko z zainstalowanym Go i GCC (system LINUX lub cygwin) 

### Instrukcja
- Uruchomić skrypt docker compose za pomocą polecenia: `docker-compose up --build`
- Alternatywnie, uruchomić główny skrypt z `run.sh` na urządzeniu ze skonfigurowanym środowiskiem