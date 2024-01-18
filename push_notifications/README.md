# Push Notification gRPC Service

## About
Service ini menyediakan 2 program, yaitu:
1. Program untuk melihat penjelasan dan nutrisi dari buah
2. Program untuk melihat puisi sesuai dengan total baris puisi yang diinginkan

## Requirements
- OS Windows
- Golang 1.21.4 ver
- Python 3.11.0 ver
- Github Account
- Git 2.43.0 ver
- VSCode

## Setup
1. Clone repository ini
```bash
git clone https://github.com/Ziayaya/grpc-posttest.git
```
2. Buka folder push_notification di dalam VSCode 
3. Nyalakan virtual environtment
```bash
venv\Scripts\activate
```
4. Buka dual cmd di dalam vscode
5. Untuk cmd pertama, run server golang di dalam file server
```bash
go run main.go
```
6. Untuk cmd kedua dengan menggunakan cmd yang telah menyalakan venv, run client python di dalam file client
```bash
python main.py
```