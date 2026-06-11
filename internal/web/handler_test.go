package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func postForm(h *Handler, path string, form url.Values) *httptest.ResponseRecorder {
	req := httptest.NewRequest(http.MethodPost, path, strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)
	return rec
}

// TestCrearYVer crea un paste y luego lo abre por su enlace /p/{code},
// comprobando que el texto se conserva.
func TestCrearYVer(t *testing.T) {
	h, err := New()
	if err != nil {
		t.Fatalf("New: %v", err)
	}

	rec := postForm(h, "/create", url.Values{"text": {"hola codigo"}})
	if rec.Code != http.StatusSeeOther {
		t.Fatalf("create: código %d, esperaba 303", rec.Code)
	}
	loc := rec.Header().Get("Location") // /p/{code}
	code := strings.TrimPrefix(loc, "/p/")
	if code == "" || code == loc {
		t.Fatalf("create no redirigió a /p/{code}, fue a %q", loc)
	}

	req := httptest.NewRequest(http.MethodGet, "/p/"+code, nil)
	view := httptest.NewRecorder()
	h.Routes().ServeHTTP(view, req)
	if view.Code != http.StatusOK {
		t.Fatalf("view: código %d, esperaba 200", view.Code)
	}
	if !strings.Contains(view.Body.String(), "hola codigo") {
		t.Errorf("la página no contiene el texto guardado")
	}
}

// TestCrearVacio: un paste sin texto no se crea y redirige a la página de inicio.
func TestCrearVacio(t *testing.T) {
	h, _ := New()
	rec := postForm(h, "/create", url.Values{"text": {"   "}})
	if loc := rec.Header().Get("Location"); loc != "/" {
		t.Errorf("esperaba redirección a '/', fue a %q", loc)
	}
}

// TestVerDesconocido devuelve 404 para un código que no existe.
func TestVerDesconocido(t *testing.T) {
	h, _ := New()
	req := httptest.NewRequest(http.MethodGet, "/p/noexiste", nil)
	rec := httptest.NewRecorder()
	h.Routes().ServeHTTP(rec, req)
	if rec.Code != http.StatusNotFound {
		t.Errorf("código %d, esperaba 404", rec.Code)
	}
}
