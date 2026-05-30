package store

import "testing"

func TestSaveAndGet(t *testing.T) {
	s := New()
	code := s.Save("hola mundo secreto")
	if code == "" {
		t.Fatal("el código no debería estar vacío")
	}
	got, ok := s.Get(code)
	if !ok {
		t.Fatal("debería existir el código recién guardado")
	}
	if got != "hola mundo secreto" {
		t.Errorf("texto = %q, esperaba 'hola mundo secreto'", got)
	}
}

func TestGetUnknown(t *testing.T) {
	s := New()
	if _, ok := s.Get("noexiste"); ok {
		t.Error("un código inexistente no debería encontrarse")
	}
}

func TestCodesAreDifferent(t *testing.T) {
	s := New()
	a := s.Save("uno")
	b := s.Save("dos")
	if a == b {
		t.Errorf("dos textos distintos no deberían tener el mismo código: %q", a)
	}
}
