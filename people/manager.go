package people

import (
	"net/http"
	"encoding/json"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([...]string{"Penn", "Teller"})
}