// Package web sirve el pastebin (crear y ver textos).
package web

import (
	"embed"
	"html/template"
	"net/http"
	"strings"

	"github.com/samuelcatalanz123/pastebin-go/internal/store"
)

//go:embed templates/*.html static/*
var files embed.FS

// Handler sirve la web.
type Handler struct {
	tmpl  *template.Template
	store *store.Store
}

// New crea el Handler con un store vacío.
func New() (*Handler, error) {
	tmpl, err := template.ParseFS(files, "templates/*.html")
	if err != nil {
		return nil, err
	}
	return &Handler{tmpl: tmpl, store: store.New()}, nil
}

// Routes monta las rutas.
func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(files))
	mux.HandleFunc("GET /{$}", h.home)
	mux.HandleFunc("POST /create", h.create)
	mux.HandleFunc("GET /p/{code}", h.view)
	return mux
}

func (h *Handler) home(w http.ResponseWriter, r *http.Request) {
	if err := h.tmpl.ExecuteTemplate(w, "home.html", nil); err != nil {
		http.Error(w, "error del servidor", http.StatusInternalServerError)
	}
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("text")
	if strings.TrimSpace(text) == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	code := h.store.Save(text)
	http.Redirect(w, r, "/p/"+code, http.StatusSeeOther)
}

// viewData son los datos para la página que muestra un texto.
type viewData struct {
	Text string
	URL  string
}

func (h *Handler) view(w http.ResponseWriter, r *http.Request) {
	code := r.PathValue("code")
	text, ok := h.store.Get(code)
	if !ok {
		http.NotFound(w, r)
		return
	}
	data := viewData{Text: text, URL: "http://" + r.Host + "/p/" + code}
	if err := h.tmpl.ExecuteTemplate(w, "view.html", data); err != nil {
		http.Error(w, "error del servidor", http.StatusInternalServerError)
	}
}
