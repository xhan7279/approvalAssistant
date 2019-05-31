package controllers

import (
	"net/http"
)

var (
	projectnamemaxlen int8 = 100
	projectnameminlen int8 = 1
)

// CreateProject allows a user to create a project
func CreateProject(w http.ResponseWriter, r *http.Request) {

	//project := &models.ProjectModel{}
	//err := json.NewDecoder(r.Body).Decode(project)
	return

}
