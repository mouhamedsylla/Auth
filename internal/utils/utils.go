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

func RespondWithJSON(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
