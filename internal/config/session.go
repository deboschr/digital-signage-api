package config

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func SetupSession(r *gin.Engine) {
	store := cookie.NewStore([]byte("brOA51TcvnZNeQ9Y1idrJNZYX1g6OgwXTNPhpHJ30HYVkDO4Ki"))
	store.Options(sessions.Options{
    	// Path menentukan jalur URL di mana cookie berlaku.
    	// "/"      → cookie dikirim di semua endpoint (umum dipakai).
    	// "/api"   → cookie hanya terkirim kalau akses URL diawali "/api".
    	Path: "/",                       

    	// Domain menentukan domain tempat cookie berlaku.
    	// ""             → cookie hanya berlaku untuk domain persis (misalnya cms.pivods.com saja).
		// "pivods.com"   → cookie berlaku di domain utama, TIDAK otomatis berlaku ke subdomain.
		// ".pivods.com"  → cookie berlaku di semua subdomain (cms.pivods.com, admin.pivods.com, dll).
		Domain: ".pivods.com",

		// MaxAge menentukan umur cookie (dalam detik).
		//  >0 → umur cookie sesuai nilai (contoh: 86400 detik = 1 hari).
		//   0 → cookie hanya berlaku sampai browser ditutup (session cookie).
		//  <0 → cookie langsung dihapus.
		MaxAge: 86400,

		// Secure menentukan apakah cookie hanya boleh dikirim lewat HTTPS.
		// true  → cookie hanya dikirim di koneksi HTTPS (aman, wajib di production).
		// false → cookie juga bisa dikirim di HTTP (boleh untuk localhost/development).
		Secure: true,

		// HttpOnly melindungi cookie dari akses JavaScript di browser.
		// true  → cookie tidak bisa diakses via document.cookie (aman dari XSS).
		// false → cookie bisa diakses dari JavaScript (jarang dibutuhkan, rawan security).
		HttpOnly: true,

		// SameSite menentukan apakah cookie boleh dikirim di request lintas situs (cross-site).
		// http.SameSiteDefaultMode (0) → ikuti kebijakan default browser.
		// http.SameSiteNoneMode	 (1) → cookie selalu dikirim di cross-site request (⚠️ butuh Secure: true).
		// http.SameSiteLaxMode		 (2) → default aman, cookie terkirim di navigasi top-level GET (contoh: login redirect).
		// http.SameSiteStrictMode  (3) → paling ketat, cookie hanya dikirim kalau domain sama persis (bisa ganggu SSO atau embed iframe).
		SameSite: http.SameSiteNoneMode,
})



	r.Use(sessions.Sessions("cms_session", store))
}
