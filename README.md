# C – Lightweight Go Web Utilities

`C` is a simple Go package that provides helper functions and abstractions for building lightweight web applications.  
It includes utilities for handling HTTP requests/responses, parsing configuration files, and serving static assets.

---

## ✨ Features
- **HTTP Helpers**
  - Send JSON, plain text, or error responses easily.
  - Render HTML templates with data binding.
  - Get and set headers.
  - Handle query parameters.
  - Work with secure cookies.
  - Bind JSON body directly to structs.

- **HTTP Router**
  - Minimal wrapper over `http.ServeMux`.
  - Supports `GET`, `POST`, `PUT`, and `DELETE` routes.
  - Static file serving with `/static/`.

- **Configuration Loader**
  - Load `.set` files with `key=value` pairs.
  - Retrieve values by key.

---

## 📦 Installation
```bash
go get github.com/your-username/C
