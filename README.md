# Mygram-go

Tugas Final FGA Golang 2022 digitalent kominfo bersama Hacktiv8

MyGram adalah aplikasi yang dapat menyimpan foto maupun membuat comment untuk foto orang lain. Aplikasi ini dibuat dengan menggunakan bahasa pemrograman Go dan framework Gin dan orm Gorm. Aplikasi ini juga menggunakan database PostgreSQL.

## Rancangan Database

![Rancangan Database](/assets/images/drawSQL-export-2022-10-16_13_06.png)

## Cara menjalankan aplikasi

Apabila ingin menjalankan aplikasi ini, silahkan ikuti langkah-langkah berikut:

1. Clone repository ini
2. Masuk ke folder repository yang sudah di clone
3. Jalankan perintah `go mod tidy` untuk menginstall semua dependency yang dibutuhkan
4. Copy dan paste .env.example menjadi .env
5. Jalankan perintah `go run main.go` untuk menjalankan aplikasi
6. Aplikasi dapat diakses melalui `localhost:8080`

## Cara menjalankan test di postman

1. Buka postman
2. Import file postman yang ada di folder `postman` yang ada dalam folder assets
3. Jalankan test yang ada di postman
