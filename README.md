# Kasir API - Belajar Golang Bareng Mas Umam 🚀

Selamat datang di project **Kasir API**! Project ini adalah API sederhana untuk manajemen kategori barang dalam sistem kasir, dibangun menggunakan bahasa pemrograman Go (Golang) sebagai bagian dari seri pembelajaran "Belajar Golang Bareng Mas Umam".

## 📋 Fitur
- **Project Information**: Menampilkan status penyelesaian project.
- **Health Check**: Memastikan API berjalan dengan baik.
- **CRUD Categories**: 
  - List semua kategori.
  - Detail kategori berdasarkan ID.
  - Tambah kategori baru.
  - Update data kategori.
  - Hapus kategori.
- **In-Memory Storage**: Data disimpan dalam memory (reset saat server dimatikan).

## 🛠️ Tech Stack
- **Language**: [Go (Golang)](https://golang.org/)
- **Router**: Standard Library `net/http`
- **Format**: JSON

## 🚀 Memulai

### Prasyarat
Pastikan Anda sudah menginstal Go di mesin Anda. Cek dengan perintah:
```bash
go version
```

### Instalasi & Menjalankan
1. Clone repository ini (atau download code-nya).
2. Jalankan server menggunakan perintah:
   ```bash
   go run main.go
   ```
3. Server akan berjalan di: `http://localhost:8080`

## 📡 API Documentation

### 1. Base URL / Info
Menampilkan informasi tentang project.
- **URL**: `/`
- **Method**: `GET`
- **Response**:
  ```json
  {
      "completed": "Yes",
      "completed_at": "Kamis, 3 Sya'ban 1447 H / 22 Januari 2026 M At 09:47:13",
      "learn": "Code With Umam - Belajar Membuat API Dengan Golang",
      "task": "Kasir API - Categories"
  }
  ```

### 2. Health Check
Mengecek status API.
- **URL**: `/health`
- **Method**: `GET`
- **Response**: `200 OK`

### 3. Get All Categories
Mengambil semua daftar kategori.
- **URL**: `/categories`
- **Method**: `GET`

### 4. Create Category
Menambahkan kategori baru.
- **URL**: `/categories`
- **Method**: `POST`
- **Body**:
  ```json
  {
      "name": "Nama Kategori",
      "description": "Deskripsi Kategori"
  }
  ```

### 5. Get Category by ID
Mengambil detail kategori berdasarkan ID.
- **URL**: `/categories/{id}`
- **Method**: `GET`

### 6. Update Category
Memperbarui data kategori.
- **URL**: `/categories/{id}`
- **Method**: `PUT`
- **Body**:
  ```json
  {
      "name": "Nama Baru",
      "description": "Deskripsi Baru"
  }
  ```

### 7. Delete Category
Menghapus kategori.
- **URL**: `/categories/{id}`
- **Method**: `DELETE`

## 📝 Testing via REST Client
Jika Anda menggunakan VS Code, Anda bisa menggunakan file `kasir-api.rest` untuk mencoba semua endpoint secara langsung.

---
Dibuat dengan ❤️ untuk belajar Golang.
