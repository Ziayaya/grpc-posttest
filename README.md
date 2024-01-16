# Simple Calculator gRPC Service

## Requirement
OS Windows
Golang 1.21.4 ver
Github Account
Git 2.43.0 ver
VSCode

## Setup
1. Buka VSCode dan arahkan ke dalam folder baru
2. Run command di bawah untuk membuat project baru
```bash
go mod init github.com/Ziayaya/grpc-posttest
```
3. Buat file main.go, lalu masukkan code di bawah
```package main

import "google.golang.org/grpc"

func main() {}
```
4. Run command di bawah untuk membuat file go.sum secara otomatis
```bash
go mod tidy
```
5. Buat folder baru untuk meletakkan file calculator.proto, dalam project ini saya membuat folder bernama calculatorpb
6. Buat file calculator.proto dan masukkan code sesuai dengan yang ada di repo ini
7. Install beberapa package di bawah ini
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
```
```bash
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
8. Download zip [protoc] (https://github.com/protocolbuffers/protobuf/releases) melalui browser, pilih sesuai dengan OS yang digunakan
9. Extract folder tersebut di dalam 1 file lalu masukkan ke dalam Local Disk C, lalu copy path file /bin nya
```C:\protoc-25.0-rc-2-win64\bin```
10. Buka Edit the System Environment Variables lalu di dalam opsi Advanced, masuk ke dalam Environment Variables
11. Di dalam Environment Variables, klik 2x Variable Path di dalam System Variable
12. Klik opsi New lalu masukkan copy path yang telah dilakukan sebelumnya
13. Klik Ok untuk melakukan close System Environment Variables
14. Kembali ke VSCode, run command di bawah untuk melakukan konfigurasi proto pada gRPC
```bash
protoc --go_out=. --go-grpc_out=. calculatorpb/calculator.proto
```
15. Setelah berhasil, buat folder calculator_server dan calculator_client lalu buat file main.go dan masukkan code sesuai dengan repository ini.