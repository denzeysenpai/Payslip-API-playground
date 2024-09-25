package routes

import (
	"html/template"
	"log"
	"net/http"
)

type Path struct {
	Data string
}

func CreateView(path string) *Path {
	return &Path{Data: path}
}

// Handler responsible for attaching views to the path handler
func (path *Path) Handler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Data)
	if err != nil { // check if template is valid
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	// Render the template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}
