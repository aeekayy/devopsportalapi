package vault

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Vault struct {
	Url		string `json:"url"`
	Name	string `json:"name"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/{vaultID}", GetAVault)
	router.Delete("/{vaultID}", DeleteVault)
	router.Post("/", CreateVault)
	return router
}

func GetAVault(w http.ResponseWriter, r *http.Request) {
	vaultID := chi.URLParam(r, "vaultID")
	vaults := Vault{
		Slug:  vaultID,
		Url: "http://vault.aeekay.co",
		Title: "Hello world",
	}
	render.JSON(w, r, vaults) // A chi router helper for serializing and returning json
}

func DeleteVault(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Deleted TODO successfully"
	render.JSON(w, r, response) // Return some demo response
}

func CreateVault(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string)
	response["message"] = "Created Vault successfully"
	render.JSON(w, r, response) // Return some demo response
}
