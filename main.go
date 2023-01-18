package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/steven230500/teamcorego/models"
)

var client *http.Client


func GetDataQuestion(w http.ResponseWriter, r *http.Request) {
	url := "https://us-central1-teamcore-retail.cloudfunctions.net/test_mobile/api/questions"
	token := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJPbmxpbmUgSldUIEJ1aWxkZXIiLCJpYXQiOjE2NzM0NzU4MTEsImV4cCI6MTcwNTAxMTgxMSwiYXVkIjoid3d3LmV4YW1wbGUuY29tIiwic3ViIjoianJvY2tldEBleGFtcGxlLmNvbSIsIkdpdmVuTmFtZSI6IkpvaG5ueSIsIlN1cm5hbWUiOiJSb2NrZXQiLCJFbWFpbCI6Impyb2NrZXRAZXhhbXBsZS5jb20iLCJSb2xlIjpbIk1hbmFnZXIiLCJQcm9qZWN0IEFkbWluaXN0cmF0b3IiXX0.9wqriO_2Q8Xfwc9VcgMpr-2c4WVdLRJ5G6NcNaXdpuY"

	var questions models.Question
	var infoBody []models.InfoBody
	var info models.InfoBody

	err := GetDataJson(url, token,&questions)

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		for _, element := range questions.Data {
			info = models.InfoBody {
				PreguntaID: element.QuestionID,
				Pregunta: element.Question,
			}
			infoBody = append(infoBody, info)
		}
		question := models.ResponseData {
			Titulo: "Test Preguntas teamcore",
			Dia: questions.Date,
			Info: infoBody,
			APIVersion: 1,
		}
    	json.NewEncoder(w).Encode(question)
	}
}


func GetDataJson(url string, token string,target interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)
	
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	client = &http.Client{Timeout: 10 * time.Second}

	router := mux.NewRouter()

    router.HandleFunc("/", GetDataQuestion)

    http.ListenAndServe(":8080", router)
}