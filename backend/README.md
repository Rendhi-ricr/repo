# Backend - Repository UN

Backend API untuk aplikasi Repository UN menggunakan Go.

## 📁 Struktur Folder

```
backend/
├── cmd/                        # Command-line applications
│   └── server/
│       └── main.go            # 🚀 Entry point utama server
│
├── internal/                   # Kode internal aplikasi
│   ├── config/
│   │   └── database.go        # Konfigurasi koneksi database
│   │
│   ├── handlers/              # HTTP Handlers (Controllers)
│   │   ├── auth.go           # Handler login, register, get me
│   │   ├── document.go       # Handler CRUD dokumen
│   │   └── user.go           # Handler manajemen user
│   │
│   ├── middleware/            # Middleware
│   │   ├── auth.go           # JWT authentication & authorization
│   │   └── cors.go           # CORS handling
│   │
│   ├── models/                # Data structures
│   │   ├── document.go       # Struktur Document
│   │   └── user.go           # Struktur User
│   │
│   └── utils/                 # Utility functions
│       └── pdf.go            # PDF validation & splitting
│
├── migrations/                 # SQL Migration files
│   └── 001_create_users_table.sql
│
├── uploads/                    # File yang diupload
│   └── split/                 # Hasil split PDF per halaman
│
├── go.mod                      # Go module definition
├── go.sum                      # Dependencies checksum
└── README.md                   # File ini
```

## 🚀 Cara Menjalankan

### Development
```bash
# Dari folder backend
go run cmd/server/main.go
```

### Production Build
```bash
# Build executable
go build -o server.exe cmd/server/main.go

# Jalankan
./server.exe
```

## 📚 API Endpoints

### Auth (Public)
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| POST | `/api/auth/login` | Login user |
| POST | `/api/auth/register` | Register user baru |
| GET | `/api/auth/me` | Get data user yang login |

### Users (Admin Only)
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/users` | List semua user |
| POST | `/api/users` | Buat user baru |
| GET | `/api/users/:id` | Get user by ID |
| PUT | `/api/users/:id` | Update user |
| DELETE | `/api/users/:id` | Hapus user |

### Documents
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/api/documents` | List semua dokumen |
| POST | `/api/documents` | Upload dokumen baru |
| GET | `/api/documents/:id` | Get dokumen by ID |
| PUT | `/api/documents/:id` | Update dokumen |
| DELETE | `/api/documents/:id` | Hapus dokumen |
| GET | `/api/documents/pages/:id` | Get halaman PDF |

### Files
| Method | Endpoint | Deskripsi |
|--------|----------|-----------|
| GET | `/download/:id` | Download dokumen |
| GET | `/preview/split/:id/:page` | Preview halaman PDF |

## 🔧 Konfigurasi

### Database
Set environment variable `DATABASE_URL`:
```bash
export DATABASE_URL="postgres://user:password@localhost:5432/repository_db"
```

Atau akan menggunakan default:
```
postgres://postgres:rendhi123@localhost:5432/repository_db
```

### JWT Secret
Edit file `internal/middleware/auth.go` untuk mengubah JWT secret:
```go
var JWTSecret = []byte("your-super-secret-key-change-in-production")
```

## 📝 Penjelasan Struktur

### Mengapa menggunakan folder `internal/`?
Folder `internal` adalah konvensi Go yang mencegah package di dalamnya diimport oleh project lain. Ini membantu menjaga enkapsulasi kode.

### Mengapa memisahkan `cmd/`?
Folder `cmd` berisi entry point aplikasi. Jika nanti ada CLI tool atau multiple services, bisa ditambahkan di sini:
```
cmd/
├── server/     # HTTP server
├── cli/        # Command line tool
└── worker/     # Background worker
```

### Mengapa `handlers` bukan `controllers`?
Di Go, istilah yang umum digunakan adalah "handler" karena kita menangani HTTP requests, bukan "controller" seperti di MVC frameworks.

## 🧑‍💻 Untuk Pemula

1. **Mulai dari `cmd/server/main.go`** - Ini adalah entry point. Lihat bagaimana routes didefinisikan.

2. **Pelajari `internal/handlers/`** - Lihat bagaimana request ditangani.

3. **Pahami `internal/middleware/`** - Pelajari cara authentication bekerja.

4. **Lihat `internal/models/`** - Memahami struktur data yang digunakan.

## 📦 Dependencies

- `github.com/jackc/pgx/v5` - PostgreSQL driver
- `github.com/golang-jwt/jwt/v5` - JWT authentication
- `github.com/google/uuid` - UUID generation
- `github.com/pdfcpu/pdfcpu` - PDF processing
- `golang.org/x/crypto` - Password hashing
