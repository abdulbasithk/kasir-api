package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Kategori struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

var kategori = []Kategori{
	{ID: 1, Nama: "Makanan"},
	{ID: 2, Nama: "Minuman"},
	{ID: 3, Nama: "Bumbu"},
}

func getKategoriByID(w http.ResponseWriter, r *http.Request) {
	// Parse ID dari URL path
	// URL: /api/categories/123 -> ID = 123
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	// Cari kategori dengan ID tersebut
	for _, k := range kategori {
		if k.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(k)
			return
		}
	}

	// Kalau tidak found
	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func updateKategori(w http.ResponseWriter, r *http.Request) {
	// get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	// get data dari request
	var updateKategori Kategori
	err = json.NewDecoder(r.Body).Decode(&updateKategori)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// loop kategori, cari id, ganti sesuai data dari request
	for i := range kategori {
		if kategori[i].ID == id {
			updateKategori.ID = id
			kategori[i] = updateKategori

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateKategori)
			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}

func deleteKategori(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")

	// ganti id int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Kategori ID", http.StatusBadRequest)
		return
	}

	// loop kategori cari ID, dapet index yang mau dihapus
	for i, k := range kategori {
		if k.ID == id {
			// bikin slice baru dengan data sebelum dan sesudah index
			kategori = append(kategori[:i], kategori[i+1:]...)

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	http.Error(w, "Kategori belum ada", http.StatusNotFound)
}
