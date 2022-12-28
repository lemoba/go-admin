package controller

import (
	"net/http"
	"text/template"
	"time"
)

func registerTestRoutes() {
	http.HandleFunc("/test", handleTest)
}

func handleTest(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("assets/test.html")
	//t.Execute(w, "hello")
	//rand.Seed(time.Now().Unix())
	//t.Execute(w, rand.Intn(10) > 5)
	//daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sta", "Sun"}
	//t.Execute(w, daysOfWeek)

	//t, _ := template.ParseFiles("assets/t1.html", "assets/t2.html")
	//t.Execute(w, "hello world")

	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("t3.html").Funcs(funcMap)

	t.ParseFiles("assets/t3.html")
	t.Execute(w, time.Now())

}

func formatDate(t time.Time) string {
	layout := "2006-01-02 15:04:05"
	return t.Format(layout)
}
