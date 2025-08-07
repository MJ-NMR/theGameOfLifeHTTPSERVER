package routers

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/MJ-NMR/theGameOfLife"
)

type handlefuc func(w http.ResponseWriter, r *http.Request)

var gridtmpl = template.Must(template.ParseFiles("./templates/grid.html"))

func StartHandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("/start Request")
	st := theGameOfLife.CreateState(30, 30)
	st[18][15] = true
	st[19][15] = true
	st[19][14] = true
	st[20][15] = true
	st[20][16] = true
	
	refreshchan = theGameOfLife.PlayRoundsChan(st)
	<-refreshchan
	<-refreshchan
	<-refreshchan
	err := gridtmpl.Execute(w, <-refreshchan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var refreshchan chan theGameOfLife.State

func RefreshHandeler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("/refreash Request")
	err := gridtmpl.Execute(w, <-refreshchan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
