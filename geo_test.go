package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Données de test pour la réponse simulée de l'API
var mockLocation = []Location{
	{
		Name:    "Union Island",
		Lat:     12.5972,
		Lon:     -61.4335,
		Country: "VC",
	},
}

// Fonction de test pour getCoordinates
func TestGetCoordinates(t *testing.T) {
	// Création d'un serveur de test HTTP
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Vérifier que la requête contient bien les bons paramètres
		q := r.URL.Query()
		if q.Get("q") != "Union Island" || q.Get("limit") != "1" {
			t.Fatalf("Paramètres de requête incorrects : %v", q)
		}

		// Encodage de la réponse JSON simulée
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockLocation)
	}))
	defer server.Close()

	// Appel de la fonction avec les données de test, en utilisant le serveur local
	city := "Union Island"
	apiKey := "test_api_key" // Clé API fictive pour le test
	locations, err := getCoordinates(city, apiKey, server.URL) // Passer l'URL du serveur de test
	if err != nil {
		t.Fatalf("Erreur lors de l'appel de getCoordinates : %v", err)
	}

	// Vérification des résultats
	if len(locations) == 0 {
		t.Fatal("Aucune localisation renvoyée")
	}
	location := locations[0]
	if location.Name != mockLocation[0].Name ||
		location.Lat != mockLocation[0].Lat ||
		location.Lon != mockLocation[0].Lon ||
		location.Country != mockLocation[0].Country {
		t.Errorf("Résultat inattendu : obtenu %+v, attendu %+v", location, mockLocation[0])
	}
}
