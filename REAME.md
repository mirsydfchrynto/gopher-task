# 📚 Repository Pengumpulan Tugas

Repository ini merupakan tempat pengumpulan tugas yang menggunakan sistem **Fork** dan **Pull Request**. Ikuti panduan di bawah ini untuk mengerjakan dan mengumpulkan tugas dengan benar.

---

## 🎯 Alur Pengerjaan Tugas

```
Repository Utama → Fork → Clone → Buat Branch → Commit → Push → Pull Request
```

---

## 📋 Cara Pengerjaan dan Pengumpulan Tugas

### 1️⃣ Fork Repository

1. Buka repository utama di GitHub
2. Klik tombol **Fork** di pojok kanan atas
3. Repository akan ter-copy ke akun GitHub kamu

### 2️⃣ Clone Repository Fork

Clone repository dari akun kamu (bukan repository utama):

```bash
git clone https://github.com/username-kamu/nama-repository.git
cd nama-repository
```

### 3️⃣ Buat Branch Baru

Buat branch dengan format nama yang jelas:

```bash
git checkout -b tugas-1-nama-kamu
```

**Format penamaan branch:**
- `tugas-1-nama-kamu`
- `tugas-2-nama-kamu`
- `assignment-1-nama-kamu`

### 4️⃣ Kerjakan Tugas

1. Buat folder dengan nama kamu (jika belum ada):
   ```
   nama-kamu/
   ├── tugas-1/
   │   ├── file1.ext
   │   └── file2.ext
   └── tugas-2/
       └── ...
   ```

2. Simpan semua file tugas di folder tersebut

### 5️⃣ Commit Perubahan

```bash
git add .
git commit -m "Menambahkan Tugas 1 - Nama Kamu"
```

**Format commit message:**
- `Menambahkan Tugas 1 - Nama Kamu`
- `Update Tugas 2 - Nama Kamu`
- `Fix: Perbaikan Tugas 1 - Nama Kamu`

### 6️⃣ Push ke Repository Fork

```bash
git push origin tugas-1-nama-kamu
```

### 7️⃣ Buat Pull Request (PR)

1. Buka repository fork kamu di GitHub
2. Klik tombol **Compare & pull request**
3. Pastikan:
   - **Base repository**: Repository utama
   - **Base branch**: `main` (atau `master`)
   - **Head repository**: Repository fork kamu
   - **Compare branch**: Branch tugas kamu
4. Isi judul PR: `[Tugas 1] Nama Kamu`
5. Isi deskripsi PR dengan detail:
   ```
   ## Informasi Tugas
   - Nama: [Nama Lengkap]
   - Tugas: [Nomor/Judul Tugas]
   
   ## Checklist
   - [ ] Kode sudah ditest
   - [ ] File sudah sesuai struktur folder
   - [ ] Tidak ada conflict
   ```
6. Klik **Create pull request**

---

## 📁 Struktur Folder

```
repository-utama/
├── README.md
├── nama-mahasiswa-1/
│   ├── tugas-1/
│   ├── tugas-2/
│   └── tugas-3/
├── nama-mahasiswa-2/
│   ├── tugas-1/
│   └── tugas-2/
└── nama-mahasiswa-3/
    └── tugas-1/
```

**Aturan Struktur:**
- Setiap mahasiswa punya folder sendiri
- Nama folder menggunakan format: `nama-lengkap` (lowercase, gunakan dash `-`)
- Setiap tugas dalam subfolder terpisah

---

## ⚠️ Hal yang Perlu Diperhatikan

### ✅ DO's (Lakukan)
- Selalu fork repository sebelum mengerjakan
- Buat branch baru untuk setiap tugas
- Commit dengan message yang jelas
- Test kode sebelum PR
- Buat PR sesuai deadline

### ❌ DON'Ts (Jangan)
- Jangan langsung commit ke branch `main`
- Jangan edit file milik orang lain
- Jangan lupa sync fork dengan repository utama
- Jangan force push tanpa alasan jelas

---

## 🔄 Update Fork dari Repository Utama

Sebelum mengerjakan tugas baru, pastikan fork kamu ter-update:

```bash
# Tambahkan remote upstream (sekali saja)
git remote add upstream https://github.com/username-pengajar/nama-repository.git

# Fetch perubahan dari upstream
git fetch upstream

# Pindah ke branch main
git checkout main

# Merge perubahan dari upstream
git merge upstream/main

# Push ke fork kamu
git push origin main
```

---

## 🆘 Troubleshooting

### Problem: Conflict saat PR
**Solusi:**
```bash
git checkout main
git pull upstream main
git checkout tugas-1-nama-kamu
git merge main
# Resolve conflict secara manual
git add .
git commit -m "Resolve conflict"
git push origin tugas-1-nama-kamu
```

### Problem: Salah commit ke branch main
**Solusi:**
```bash
git checkout main
git reset --hard HEAD~1
git push -f origin main
```

---

## 📞 Kontak

Jika ada pertanyaan atau kendala:
- Buka **Issue** di repository utama
- Hubungi [Nama Pengajar/Asisten]

---

## 📝 Deadline

| Tugas | Deadline | Status |
|-------|----------|--------|
| Tugas 1 | DD/MM/YYYY | 🟢 Open |
| Tugas 2 | DD/MM/YYYY | 🟡 Coming |
| Tugas 3 | DD/MM/YYYY | ⚪ TBA |

---

**Happy Coding! 🚀**
