# Pastebin (Go)

[![CI](https://github.com/samuelcatalanz123/pastebin-go/actions/workflows/ci.yml/badge.svg)](https://github.com/samuelcatalanz123/pastebin-go/actions/workflows/ci.yml)
![Go](https://img.shields.io/badge/Go-1.26-00ADD8?logo=go&logoColor=white)
![License](https://img.shields.io/badge/license-MIT-green)

App web para pegar texto o código y obtener un **enlace único** para compartirlo
(como Pastebin o GitHub Gist). Hecha en **Go**.

## Uso

```bash
go run .
```

Abre **http://localhost:8080**, pega tu texto, pulsa **Crear enlace** y comparte
la dirección `/p/{code}` que te da.

## Cómo funciona

- `internal/store`: guarda cada texto con un **código corto** aleatorio (base62,
  `crypto/rand`) y lo recupera por ese código. Protegido con un `sync.Mutex`.
- `internal/web`: el formulario (`GET /`), crear (`POST /create`) y ver
  (`GET /p/{code}`). El texto se muestra con `html/template`, que **escapa el
  HTML** por seguridad.

## Estructura

```
main.go                 arranque del servidor
internal/store/         guardar/recuperar textos por código + pruebas
internal/web/           formulario, crear y ver (handlers + plantillas)
```

## Pruebas

```bash
go test ./...
```

La prueba comprueba que un texto guardado se recupera igual por su código, y que
un código inexistente no se encuentra.

## Stack

Go (net/http, html/template, go:embed, crypto/rand, sync).
