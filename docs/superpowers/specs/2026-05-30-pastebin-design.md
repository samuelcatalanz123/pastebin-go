# Diseño: Pastebin (Go)

**Fecha:** 2026-05-30 · **Estado:** Aprobado · **Autor:** Samuel (14º proyecto)

## Objetivo

Una app web para pegar texto o código y obtener un **enlace único** para
compartirlo. Objetivo de aprendizaje: guardar contenido con un código corto y
servirlo por una ruta con parámetro.

## Pantalla y rutas

- **Inicio** — `GET /`: un formulario con un área de texto y un botón "Crear".
- **Crear** — `POST /create`: guarda el texto, genera un código y redirige a `/p/{code}`.
- **Ver** — `GET /p/{code}`: muestra el texto guardado + el enlace para compartir.
  Código inexistente → 404.

## Arquitectura

```
pastebin-go/
  main.go                 arranca el servidor (:8080 o PORT)
  internal/store/
    store.go              Store en memoria: Save(text)→code, Get(code)→(text,ok); código base62
    store_test.go         prueba: guardar y recuperar el mismo texto; código inexistente
  internal/web/
    handler.go            GET / (form), POST /create, GET /p/{code} (ver)
    templates/home.html   formulario
    templates/view.html   mostrar el texto + enlace
    static/style.css
  README.md
```

- **store.go:** `Store{ mu sync.Mutex; pastes map[string]string }`, `New()`,
  `Save(text) string` (genera un código de 6 caracteres base62 con `crypto/rand`,
  reintenta si choca, guarda y devuelve el código), `Get(code) (string, bool)`.
- **handler.go:** `home` muestra el formulario; `create` lee `text`, si no está
  vacío lo guarda y redirige a `/p/{code}`; `view` busca el código (404 si no
  existe) y muestra el texto + el enlace `http://{host}/p/{code}`.

## Pruebas

- **store_test.go:** `Save` + `Get` devuelve el mismo texto; `Get` de un código
  inexistente → `ok == false`. Concurrencia segura (Mutex).
- `go build/vet/test` limpios.

## Seguridad

El texto se muestra con `html/template`, que **escapa el HTML** automáticamente
(evita inyección de scripts).

## Fuera de alcance (YAGNI)

Base de datos, borrar/editar, expiración, resaltado de sintaxis, listado público.

## Criterios de éxito

1. `go run .` sirve en http://localhost:8080.
2. Pegar texto y "Crear" da un enlace `/p/{code}` que muestra ese texto.
3. Un código inexistente da 404.
4. La prueba del store pasa.
