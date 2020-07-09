package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/edtap14/proyecto-productos-back/models"
	"github.com/labstack/echo/v4"
)

/*ProductSearch ruta que captura la respuesta de la API*/
func ProductSearch(c echo.Context) error {
	// var ps models.ProductSearchResponse
	// err := json.NewDecoder(r.Body).Decode(&ps)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Obtención de tareas
	res, err := http.Get(fmt.Sprintf("https://api.rainforestapi.com/request?api_key=2AFBB75C890447B4B31D97B70028C9DA&type=search&amazon_domain=amazon.com.mx&search_term=bulova"))
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	var psr models.ProductSearchResponse
	if err := json.NewDecoder(res.Body).Decode(&psr); err != nil {
		log.Fatal(err)
	}

	// js, err := json.MarshalIndent(psr, "", "\t")
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)
	return c.JSON(http.StatusOK, psr)
}
