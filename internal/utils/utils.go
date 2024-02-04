package utils

import (
	"auth/orm"
	"encoding/json"
	"net/http"
)

func OrmInit(dbName string) *orm.ORM {
	gorm := orm.NewORM()
	gorm.InitDB(dbName)
	return gorm
}

func RespondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Erreur lors de la conversion en JSON", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	if _, err := w.Write(jsonData); err != nil {
		return
	}
}
