// Package store guarda textos pegados y los recupera por un código corto.
package store

import (
	"crypto/rand"
	"sync"
)

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Store guarda los textos en memoria (código → texto).
type Store struct {
	mu     sync.Mutex
	pastes map[string]string
}

// New crea un Store vacío.
func New() *Store {
	return &Store{pastes: make(map[string]string)}
}

// Save guarda un texto y devuelve su código único.
func (s *Store) Save(text string) string {
	s.mu.Lock()
	defer s.mu.Unlock()

	code := newCode()
	for _, existe := s.pastes[code]; existe; _, existe = s.pastes[code] {
		code = newCode() // si por casualidad ya existe, generamos otro
	}
	s.pastes[code] = text
	return code
}

// Get devuelve el texto de un código (y si existe o no).
func (s *Store) Get(code string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	text, ok := s.pastes[code]
	return text, ok
}

// newCode genera un código aleatorio de 6 caracteres base62.
func newCode() string {
	b := make([]byte, 6)
	_, _ = rand.Read(b)
	for i := range b {
		b[i] = alphabet[int(b[i])%len(alphabet)]
	}
	return string(b)
}
