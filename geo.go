package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

// Structure pour stocker les coordonnées d'une ville
type Location struct {
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	Country   string  `json:"country"`
	State     string  `json:"state"`
}

// Fonction pour obtenir les coordonnées d'une ville
func getCoordinates(city string, apiKey string) ([]Location, error) {
	baseURL := "http://api.openweathermap.org/geo/1.0/direct"
	params := url.Values{}
	params.Add("q", city)
	params.Add("limit", "1")      // Limite à un résultat
	params.Add("appid", apiKey)   // Ajoute la clé d'API

	// Construction de l'URL complète
	requestURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	// Envoi de la requête HTTP GET
	resp, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Vérification du code de statut de la réponse
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("échec de la requête, statut HTTP: %d", resp.StatusCode)
	}

	// Décodage de la réponse JSON
	var locations []Location
	if err := json.NewDecoder(resp.Body).Decode(&locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func main() {
	// Nom de la ville à rechercher
	city := "UNION ISLAND"
	// Clé d'API
	apiKey := "c732a4f732342956ec521490b59a7dce"

	// Appel de la fonction pour obtenir les coordonnées
	locations, err := getCoordinates(city, apiKey)
	if err != nil {
		fmt.Println("Erreur :", err)
		os.Exit(1)
	}

	// Affichage des coordonnées de la ville
	if len(locations) > 0 {
		location := locations[0]
		fmt.Printf("Ville : %s, Pays : %s, Latitude : %.4f, Longitude : %.4f\n",
			location.Name, location.Country, location.Lat, location.Lon)
	} else {
		fmt.Println("Aucune correspondance trouvée pour cette ville.")
	}
}
