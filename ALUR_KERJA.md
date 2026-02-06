# 🔄 Alur Kerja Sistem CRUD Dokumen

## 1️⃣ CREATE - Menambah Dokumen Baru

### Frontend Flow:
```
User → Klik "Tambah Dokumen" 
     → Isi form (judul, penulis, jenis file, status)
     → Upload file
     → Centang konfirmasi
     → Klik "Simpan Dokumen"
     → Loading...
     → Success → Redirect ke list dokumen
```

### Backend Process:
```
POST /api/documents
  ↓
Parse multipart form data
  ↓
Validate: judul, penulis, jenis_file, file
  ↓
Generate UUID untuk ID
  ↓
Generate UUID untuk nama file
  ↓
Save file ke uploads/[uuid].[ext]
  ↓
Insert ke database:
  - id, judul, penulis, jenis_file, file_path, status
  ↓
Return JSON response dengan data dokumen
```

---

## 2️⃣ READ - Melihat Dokumen

### A. List Semua Dokumen

#### Frontend Flow:
```
User → Buka halaman /documents
     → Loading...
     → Fetch data dari API
     → Tampilkan dalam grid cards
     → User bisa search/filter
```

#### Backend Process:
```
GET /api/documents
  ↓
Query database: SELECT * FROM documents ORDER BY created_at DESC
  ↓
Loop semua rows
  ↓
Build array of Document objects
  ↓
Return JSON array
```

### B. Lihat Detail Satu Dokumen

#### Frontend Flow:
```
User → Klik "Edit" pada dokumen
     → Navigate ke /documents/edit/:id
     → Loading...
     → Fetch data dokumen by ID
     → Populate form dengan data
```

#### Backend Process:
```
GET /api/documents/:id
  ↓
Query database: SELECT * FROM documents WHERE id = $1
  ↓
Scan row ke Document struct
  ↓
Return JSON object
```

---

## 3️⃣ UPDATE - Mengubah Dokumen

### Frontend Flow:
```
User → Klik "Edit" pada dokumen
     → Form terisi dengan data existing
     → User ubah data yang diinginkan
     → (Opsional) Upload file baru
     → Centang konfirmasi
     → Klik "Simpan Perubahan"
     → Loading...
     → Success → Redirect ke list dokumen
```

### Backend Process:
```
PUT /api/documents/:id
  ↓
Parse multipart form data
  ↓
Validate: judul, penulis, jenis_file
  ↓
Cek apakah ada file baru?
  ├─ YA:
  │   ↓
  │   Query file_path lama dari DB
  │   ↓
  │   Hapus file lama dari disk
  │   ↓
  │   Generate UUID untuk file baru
  │   ↓
  │   Save file baru ke uploads/
  │   ↓
  │   UPDATE database dengan file_path baru
  │
  └─ TIDAK:
      ↓
      UPDATE database (metadata saja)
  ↓
Return JSON response
```

---

## 4️⃣ DELETE - Menghapus Dokumen

### Frontend Flow:
```
User → Klik tombol "🗑️" pada dokumen
     → Confirm dialog: "Apakah Anda yakin?"
     → User klik OK
     → Loading...
     → Success → Refresh list dokumen
```

### Backend Process:
```
DELETE /api/documents/:id
  ↓
Query file_path dari database WHERE id = $1
  ↓
Hapus file fisik dari disk: os.Remove(file_path)
  ↓
Delete dari database: DELETE FROM documents WHERE id = $1
  ↓
Return success message
```

---

## 5️⃣ DOWNLOAD - Mengunduh File

### Frontend Flow:
```
User → Klik tombol "⬇️ Download"
     → Open new tab dengan URL download
     → Browser auto download file
```

### Backend Process:
```
GET /download/:id
  ↓
Query file_path dari database WHERE id = $1
  ↓
Set header: Content-Disposition: attachment
  ↓
Serve file menggunakan http.ServeFile()
  ↓
Browser download file
```

---

## 🔍 SEARCH & FILTER

### Frontend Only (Client-side):
```
User → Ketik di search box
     ↓
     Filter array documents:
       - Match judul (case-insensitive)
       - Match penulis (case-insensitive)
     ↓
     Re-render grid dengan hasil filter

User → Pilih filter status
     ↓
     Filter array documents:
       - all: tampilkan semua
       - draft: status === 'draft'
       - publish: status === 'publish'
     ↓
     Re-render grid dengan hasil filter
```

---

## 📊 Data Flow Diagram

```
┌──────────┐         ┌──────────┐         ┌──────────┐
│          │  HTTP   │          │  SQL    │          │
│ Frontend │ ◄─────► │ Backend  │ ◄─────► │ Database │
│ (Svelte) │ Request │   (Go)   │  Query  │ (Postgres)│
│          │Response │          │         │          │
└──────────┘         └────┬─────┘         └──────────┘
                          │
                          │ File I/O
                          ▼
                     ┌──────────┐
                     │  uploads/│
                     │  folder  │
                     └──────────┘
```

---

## 🎯 State Management

### Frontend States:

#### DocumentList.svelte
```javascript
documents = []        // Array semua dokumen
loading = true       // Loading state
error = ""          // Error message
searchQuery = ""    // Search input
filterStatus = "all" // Filter dropdown
```

#### DocumentAdd.svelte
```javascript
title = ""          // Form input
author = ""         // Form input
fileType = ""       // Form select
status = "draft"    // Form select
file = null         // File object
confirm = false     // Checkbox
loading = false     // Submit loading
error = ""          // Error message
```

#### DocumentEdit.svelte
```javascript
title = ""          // Form input (populated)
author = ""         // Form input (populated)
fileType = ""       // Form select (populated)
status = "draft"    // Form select (populated)
file = null         // New file (optional)
currentFileName = "" // Info text
confirm = false     // Checkbox
loading = false     // Submit loading
loadingData = true  // Initial load
error = ""          // Error message
```

---

## 🔐 Error Handling

### Frontend:
```javascript
try {
  const response = await fetch(API);
  if (!response.ok) throw new Error("Gagal...");
  return response.json();
} catch (e) {
  error = e.message;
  // Display error di UI
}
```

### Backend:
```go
if err != nil {
    http.Error(w, "Pesan error", http.StatusCode)
    return
}
```

### Error Types:
- **400 Bad Request**: Data tidak lengkap, ID tidak valid
- **404 Not Found**: Dokumen tidak ditemukan
- **405 Method Not Allowed**: HTTP method salah
- **500 Internal Server Error**: Database error, file I/O error

---

## ✅ Validation

### Frontend Validation:
```javascript
if (!title || !author || !fileType || !file || !confirm) {
  error = "Semua field wajib diisi dan dikonfirmasi.";
  return;
}
```

### Backend Validation:
```go
if judul == "" || penulis == "" || jenisFile == "" {
    http.Error(w, "Metadata tidak lengkap", http.StatusBadRequest)
    return
}
```

---

## 🎨 UI States

### Loading State:
```html
<div class="animate-spin rounded-full h-16 w-16 
     border-4 border-blue-500 border-t-transparent">
</div>
```

### Empty State:
```html
<div class="text-center py-20">
  <div class="text-6xl mb-4">📭</div>
  <h3>Tidak ada dokumen</h3>
  <p>Mulai dengan menambahkan dokumen baru</p>
  <button>+ Tambah Dokumen Pertama</button>
</div>
```

### Error State:
```html
<div class="p-4 bg-red-50 border-l-4 border-red-500 
     text-red-700 rounded-lg">
  <p class="font-semibold">Error:</p>
  <p>{error}</p>
</div>
```

---

## 🚀 Performance

### Frontend:
- ✅ Client-side filtering (no API calls)
- ✅ Lazy loading dengan onMount
- ✅ Debouncing untuk search (bisa ditambahkan)
- ✅ Optimistic UI updates

### Backend:
- ✅ Single query untuk list
- ✅ Indexed database queries
- ✅ Efficient file I/O
- ✅ Connection pooling (PostgreSQL)

---

## 📱 Responsive Design

### Breakpoints:
- **Mobile**: < 768px → 1 column grid
- **Tablet**: 768px - 1024px → 2 columns grid
- **Desktop**: > 1024px → 3 columns grid

### Adaptive UI:
```css
.grid {
  grid-template-columns: 1fr;
}

@media (min-width: 768px) {
  .grid {
    grid-template-columns: repeat(2, 1fr);
  }
}

@media (min-width: 1024px) {
  .grid {
    grid-template-columns: repeat(3, 1fr);
  }
}
```

---

## 🎯 User Journey

### Happy Path - Tambah Dokumen:
```
1. User buka aplikasi
2. Klik "Tambah Dokumen"
3. Isi semua field
4. Upload file
5. Centang konfirmasi
6. Klik "Simpan"
7. Lihat success message
8. Redirect ke list
9. Dokumen baru muncul di grid
```

### Happy Path - Edit Dokumen:
```
1. User lihat list dokumen
2. Klik "Edit" pada dokumen
3. Form terisi otomatis
4. Ubah data yang perlu
5. (Opsional) Upload file baru
6. Centang konfirmasi
7. Klik "Simpan Perubahan"
8. Lihat success message
9. Redirect ke list
10. Perubahan terlihat di grid
```

### Happy Path - Hapus Dokumen:
```
1. User lihat list dokumen
2. Klik tombol "🗑️"
3. Confirm dialog muncul
4. User klik OK
5. Dokumen hilang dari grid
6. File terhapus dari server
```

---

## 🔄 Complete CRUD Cycle

```
┌─────────────────────────────────────────────────┐
│                                                 │
│  CREATE → READ → UPDATE → DELETE → READ        │
│    ↓       ↓       ↓        ↓        ↓         │
│  Tambah  Lihat   Edit    Hapus    Refresh      │
│    ↓       ↓       ↓        ↓        ↓         │
│   DB     Query   Update  Delete   Query        │
│    ↓       ↓       ↓        ↓        ↓         │
│  File    Serve   Replace Remove   Serve        │
│                                                 │
└─────────────────────────────────────────────────┘
```

Sistem CRUD lengkap dan siap digunakan! 🎉
