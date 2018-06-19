package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const JSON = "application/json; charset=UTF-8"

const testBody = "These are the voyages of the Starship Enterprise. Its continuing mission, to explore strange new worlds, to seek out new life and new civilizations, to boldly go where no one has gone before. We need to neutralize the homing signal. Each unit has total environmental control, gravity, temperature, atmosphere, light, in a protective field. Sensors show energy readings in your area. We had a forced chamber explosion in the resonator coil. Field strength has increased by 3,000 percent. \nNow what are the possibilities of warp drive? Cmdr Riker's nervous system has been invaded by an unknown microorganism. The organisms fuse to the nerve, intertwining at the molecular level. That's why the transporter's biofilters couldn't extract it. The vertex waves show a K-complex corresponding to an REM state. The engineering section's critical. Destruction is imminent. Their robes contain ultritium, highly explosive, virtually undetectable by your transporter. \nResistance is futile."

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "")
	fmt.Fprintf(w, "<!DOCTYPE html><html><body>Coming soon!<br><br>-v<br><br><br><a href='https://github.com/vvmk'>GitHub</a><br><a href='https://linkedin.com/in/vincentmasiello'>LinkedIn</a><br /><a href='mailto:v@complexaesthetic.com'>Email</a></body></html>")
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	postId, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	post := Post{
		postId,
		"Test Post",
		testBody,
		[]string{"test", "tags", "go"},
		time.Now(),
	}

	w.Header().Set("Content-Type", JSON)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		panic(err)
	}
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: GetAllPosts")
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: CreatePost")
}

func EditPost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: EditPost")
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO: DeletePost")
}
