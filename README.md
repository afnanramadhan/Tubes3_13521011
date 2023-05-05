# Tubes 3 IF2211 STRATEGI ALGORITMA
> Aplikasi Chat GPT Sederhana dengan Menerapkan Algoritma String Matching
> Live demo [_here_](https://www.example.com). <br>
> App [_here_](https://tubes3-13521056-fe.vercel.app/)

## Table of Contents
* [General Info](#general-information)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Screenshots](#screenshots)
* [Usage](#usage)
* [Project Status](#project-status)
* [Acknowledgements](#acknowledgements)
* [Author](#contact)
<!-- * [License](#license) -->


## General Information
> Aplikasi ChatGPT sederhana dengan menerapkan pendekatan QA yang paling sederhana. Untuk menemukan pertanyaan yang paling mirip dengan input pengguna, memanfaatkan algoritma pencocokan string Knuth-Morris-Pratt (KMP) dan Boyer-Moore (BM) serta Regex untuk menentukan format dari pertanyaan.  

<!-- You don't have to answer all the questions - just the ones relevant to your project. -->


## Technologies Used/ Requirements
- Next.js   (https://nextjs.org/docs/getting-started/installation)
- Node.js   (https://nodejs.org/en/download)
- Mysql     (https://www.mysql.com/downloads/)
- Vs code   (https://code.visualstudio.com/download)


## Features
List the ready features here:
- Fitur pertanyaan teks (didapat dari database)
- Fitur kalkulator
- Fitur tanggal
- Tambah pertanyaan dan jawaban ke database
- Hapus pertanyaan dari database


## Usage
### Alternativ 1: 
- Buka [link vercel app ini](https://tubes3-13521056-fe.vercel.app/)
#### nb: apabila terdapat kegagalan/kejanggalan pada hosting baik pada back-end, front-end atau database, harap gunakan alternativ 2.
### Alternativ 2:
- Clone repository / download zip dari repository dari link yang ada pada lampiran
- Pastikan anda sudah menginstall Golang
- Masukkan file dump sql ke dalam database lokal yang ada pada device Anda dan atur file .env yang ada pada folder src/backend/models sehingga sesuai dengan database lokal Anda
- Jalankan Backend command : 
cd src/backend
go run main.go
- Jalankan Frontend command berikut :
cd src/backend
npm install
npm run dev


## Project Status
Project is:  _complete_ 


## Acknowledgements
Give credit here.
- This project was inspired by Chat GPT
- Terima Kasih banyak Kepada Asistean dan Dosen Mata Kuliah IF2211 Strategi Algoritma


## Author

|Nama |NIM|
|---|---|
|Afnan Edsa|13521011|
|Laila Bilbina|13521016|
|Syarifa Dwi P.|13521125|


