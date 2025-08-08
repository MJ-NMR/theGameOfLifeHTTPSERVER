package routers

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/MJ-NMR/GOL"
)

type handlefuc func(w http.ResponseWriter, r *http.Request)

var gridtmpl = template.Must(template.ParseFiles("./templates/grid.html"))


type temlateData struct {
	Get string
	Trigger string
	Grid GOL.State
}

var stoped bool = false
var started bool
var grid GOL.State

func StartHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("/start Request")
	st := GOL.StateExample()
	refreshchan = GOL.PlayRoundsChan(st)
	grid = <-refreshchan
	tempdata := temlateData{
		Get: "/refresh",
		Trigger: "every 1s",
		Grid: grid,
	}
	stoped = false
	started = true
	err := gridtmpl.Execute(w, tempdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var refreshchan chan GOL.State

var isStarted bool


func RefreshHandeler(w http.ResponseWriter, r *http.Request) {
	if stoped {
		w.WriteHeader(204)
		return
	}
	fmt.Println("/refreash Request")
	grid = <-refreshchan
	tempdata := temlateData{
		Get: "/refresh",
		Trigger: "every 1s",
		Grid: grid,
	}
	err := gridtmpl.Execute(w, tempdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StepHandler(w http.ResponseWriter, r *http.Request) {
	stoped = true
	fmt.Println("/stop Request")
	if !isStarted {
		st := GOL.StateExample()
		refreshchan = GOL.PlayRoundsChan(st)
		isStarted = true
	}
	tempdata := temlateData{
		Get: "/step",
		Trigger: "onclick",
		Grid: <-refreshchan,
	}
	err := gridtmpl.Execute(w , tempdata)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
