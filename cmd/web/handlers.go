package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"eventurely/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	events, err := app.events.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, event := range events {
		fmt.Fprintf(w, "%+v\n", event)
	}

	//templateFiles := []string{
	//	"./ui/html/base.gohtml",
	//	"./ui/html/pages/home.gohtml",
	//	"./ui/html/partials/nav.gohtml",
	//}
	//
	//templates, err := template.ParseFiles(templateFiles...)
	//if err != nil {
	//	app.serverError(w, r, err)
	//	return
	//}
	//
	//err = templates.ExecuteTemplate(w, "base", nil)
	//if err != nil {
	//	app.serverError(w, r, err)
	//}
}

func (app *application) eventCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", "POST")
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	ownerID := 1
	title := "My Birthday!"
	description := "It's my birthday and I'll cry if I want to."
	location := "The Party Place"
	uniqueLink := "my-birthday-" + strconv.Itoa(int(time.Now().Unix()))
	startsAt := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)
	endsAt := time.Date(2023, time.January, 1, 23, 59, 59, 0, time.UTC)
	privacySetting := models.PrivacySettingPublic

	id, err := app.events.Insert(ownerID, title, description, location, uniqueLink, startsAt, endsAt, privacySetting)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/event/view?id=%d", id), http.StatusSeeOther)
}

func (app *application) eventView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	event, err := app.events.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	fmt.Fprintf(w, "%+v", event)
}
