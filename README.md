# Golang Fiber Rest API

Bu proje, Go programlama dili kullanılarak geliştirilmiş bir REST API uygulamasıdır. Uygulama, kullanıcı yönetimi için temel CRUD (Create, Read, Update, Delete) işlemlerini gerçekleştiren bir API sunar. Fiber web framework'ü kullanılarak geliştirilmiştir.

Folksdev youtube kanalında yayınlanan [Golang Fiber Rest API](https://www.youtube.com/watch?v=zTq6ei6E9uY) vidosundan yararlanılarak geliştirilmiştir.

#### İçindekiler

- [Kurulum](#kurulum)
- [Çalıştırma](#çalıştırma)
- [API Kullanımı](#api-kullanımı)
- [Proje Yapısı](#proje-yapısı)

#### Kurulum

1. **Go Kurulumu** : İlk olarak, bilgisayarınıza Go programlama dilini kurmanız gerekmektedir. Go'nun resmi web sitesinden indirme işlemini gerçekleştirebilirsiniz. [Go İndirme Linki](https://golang.org/dl/)
2. **Proje İndirme** : Proje dosyalarını bilgisayarınıza indirin.
```bash
git clone https://github.com/BerkayMehmetSert/golang-fiber-rest-api.git
cd golang-fiber-rest-api
```
3. **Gerekli Paketlerin Yüklenmesi** : Proje dosyalarını indirdikten sonra, proje dizininde aşağıdaki komutu çalıştırarak gerekli paketleri yükleyin.
```bash
go mod tidy
```

#### Çalıştırma

Proje dizininde aşağıdaki komutu çalıştırarak uygulamayı başlatabilirsiniz.
```bash
go run main.go
```

#### API Kullanımı

API'yi test etmek için Postman veya benzeri bir araç kullanabilirsiniz. Aşağıda API'nin kullanımı hakkında bilgiler yer almaktadır.

Sağlık Kontrolü
- **Endpoint**: /healthcheck
- **Metot**: GET
- **Açıklama**: Sunucunun çalışıp çalışmadığını kontrol eder.
- **Örnek İstek**:`curl -X GET http://localhost:8080/healthcheck `

Kullanıcı Oluşturma
- **Endpoint**: /api/v1/user
- **Metot**: POST
- **Açıklama**: Yeni bir kullanıcı oluşturur.
- **Örnek İstek**:`curl -X POST http://localhost:8080/api/v1/user \ -H "Content-Type: application/json" \ -d '{ "firstName": "John", "lastName": "Doe", "email": "john.doe@example.com", "password": "password", "age": 30 }'`

Kullanıcı Bilgilerini Getirme
- **Endpoint**: /api/v1/user/:userId
- **Metot**: GET
- **Açıklama**: Belirli bir kimliğe sahip kullanıcıyı döndürür.
- **Örnek İstek**:`curl -X GET http://localhost:8080/api/v1/user/1`

Tüm Kullanıcıları Getirme
- **Endpoint**: /api/v1/user
- **Metot**: GET
- **Açıklama**: Tüm kullanıcıları döndürür.
- **Örnek İstek**:`curl -X GET http://localhost:8080/api/v1/user`

E-posta ile Kullanıcı Getirme
- **Endpoint**: /api/v1/user/email/:email
- **Metot**: GET
- **Açıklama**: Belirli bir e-posta adresine sahip kullanıcıyı döndürür.
- **Örnek İstek**:`url -X GET http://localhost:8080/api/v1/user/email/john.doe@gmail.com`

#### Proje Yapısı

```
golang-fiber-rest-api/
├── configuration/
│   └── appConfigs.go
├── internal/
│   ├── golang-fiber-rest-api/
│   │   ├── application/
│   │   │   ├── controller/
│   │   │   │   ├── request/
│   │   │   |   |   └── userCreateRequest.go
│   │   │   │   ├── response/
│   │   │   |   |   └── userResponse.go
│   │   │   │   └── userController.go
│   │   │   ├── handler/
│   │   │   │   └── user/
│   │   │   │   |   ├── command.go
│   │   │   │   |   ├── commandHandler.go
│   │   │   ├── query/
│   │   │   │   └── userQueryService.go
│   │   │   └── repository/
│   │   │   |   └── userRepository.go
│   │   |   └── web/
│   │   |   |   └── router.go
│   │   ├── domain/
│   │   │   ├── user.go
│   │   ├── pkg/
│   │   │   ├── server/
│   │   │   │   └── server.go
│   │   │   └── utils/
│   │   │   |   └── stub.go
│   │   |   └── messages.go
│   │   └── test/
│   │   |   └── controller_test.go
├── go.mod
├── go.sum
└── main.go
```
