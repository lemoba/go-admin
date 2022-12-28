package controller

import (
	"net/http"
	"text/template"
)

func registerTestRoutes() {
	http.HandleFunc("/test", handleTest)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("assets/week.html")
	//rand.Seed(time.Now().Unix())
	//t.Execute(w, rand.Intn(10) > 5)
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	t.Execute(w, daysOfWeek)
}
