package routes

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

//Products ruta que envía la peticion de busqueda
func Products(c echo.Context) error {

}

// Search ruta que captura la respuesta de la API
func search(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	q := c.QueryParam("q")
	amazonURL := os.Getenv("AMZ_BASE_URL")
	return c.String(http.StatusOK, fmt.Sprintf(amazonURL, q))
	// var producto string
	// fmt.Println("¿Que producto buscas el día de hoy?")
	// fmt.Scanln(&producto)
	// var ps models.ProductSearchResponse
	// err := json.NewDecoder(r.Body).Decode(&ps)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Obtención de tareas
	// res, err := http.Get(fmt.Sprintf("https://api.rainforestapi.com/request?api_key=2AFBB75C890447B4B31D97B70028C9DA&type=search&amazon_domain=amazon.com.mx&search_term=%s", producto))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer res.Body.Close()
	// var psr models.ProductSearchResponse
	// if err := json.NewDecoder(res.Body).Decode(&psr); err != nil {
	// 	log.Fatal(err)
	// }

	// js, err := json.MarshalIndent(psr, "", "\t")
	// if err != nil {
	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.Write(js)

}
