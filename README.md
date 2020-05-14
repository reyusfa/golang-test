## Getting Started

#### Clone Project
```
git clone https://github.com/reyusfa/golang-test.git
cd golang-test
```

#### Build & Run Docker
```
docker build . -t golang-test
docker run -p 8080:8080 golang-test
```

#### Web Service Address

```
http://localhost:8080
http://localhost:8080/palindrome-number?input=1 10
http://localhost:8080/organize-book?input=3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14
http://localhost:8080/missing-number?input=23242526272830
```

## Demo

[Index Page](http://3.87.58.95:8080)

[Soal 0 - Palindrome Number](http://3.87.58.95:8080/palindrome-number?input=1%2010)

[Soal 1 - Organize Book](http://3.87.58.95:8080/organize-book?input=3A13%205X19%209Y20%202C18%201N20%203N20%201M21%201F14%209A21%203N21%200E13%205G14%208A23%209E22%203N14)

[Soal 2 - Missing Number](http://3.87.58.95:8080/missing-number?input=23242526272830)
