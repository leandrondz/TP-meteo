package main
 
import (
    "fmt"
    "log"
    "net/http"
)
 
func accueil(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Mon application météo fonctionne !")
}
 
func main() {
    http.HandleFunc("/", accueil)
 
    log.Println("Serveur lancé sur http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}