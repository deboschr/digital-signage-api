# Digital Signage API for Airport


# Cara pakai sehari-hari

1. **Deploy perubahan**

   * Commit & push ke `master` → Actions otomatis:

     * build → push image ke GHCR → SSH ke VPS → `git pull` → `docker compose pull && up -d`.

2. **Cek status**

   ```bash
   # di VPS
   cd /opt/digital-signage/digital-signage-api
   sudo docker compose ps
   sudo docker compose logs -f
   curl http://127.0.0.1:8080/health
   # atau dari luar
   curl http://31.97.223.168/api/health
   ```

3. **Restart layanan (kalau perlu)**

   ```bash
   sudo docker compose restart
   ```

4. **Hotfix cepat langsung di server (darurat)**

   > (Tidak ideal, tapi kadang perlu)

   ```bash
   git status         # pastikan bersih
   # edit file kecil → test
   sudo docker compose up -d --build  # hanya jika compose masih pakai build; kalau image:latest, cukup restart
   ```

   Setelah darurat selesai, **commit & push** dari lokal supaya ga “nyangkut” di server.

5. **Bersihkan disk (aman)**

   ```bash
   sudo docker image prune -f         # hapus image tak terpakai
   sudo docker builder prune -f       # hapus cache build
   sudo docker system df              # lihat pemakaian
   ```

6. **Start otomatis saat reboot**

   ```bash
   sudo systemctl enable docker
   ```

# Hal yang **sudah beres**

* API jalan di container, bind ke `127.0.0.1:8080`.
* Nginx reverse proxy memaparkan `http://<IP>/api/...`.
* CI/CD: push ke `master` ⇒ otomatis build & deploy.

# Hal yang **boleh nanti**

* HTTPS + domain (Let’s Encrypt via nginx/certbot).
* Rollback versi image (pakai tag `sha-<short>`).
* Endpoint `/version` yang menampilkan SHA build.
* CORS, rate-limit, dan `.env` untuk konfigurasi.

# Saran kecil biar rapi

* Tambahkan ke **README** repo:

  * Endpoint publik (`http://<IP>/api/health`)
  * Cara deploy (push ke `master`)
  * Cara cek log & restart (perintah di atas)
* Simpan **SSH deploy key** cadangan (mis. `id_ed25519_vps.bak`) di password manager.

Kalau mau, aku bisa tulis **README.md** ops-guide < 30 baris buat timmu, atau bantu pasang **HTTPS + domain** kapan pun kamu siap.
