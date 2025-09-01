package C

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
)

func RenderHtml(w http.ResponseWriter, filename string, mapData interface{}) {
	parseTemplate, err := template.ParseFiles(filename)
	if err != nil {
		http.Error(w, "Html could not be render", http.StatusBadGateway)
		return
	}
	err = parseTemplate.Execute(w, mapData)
	if err != nil {
		http.Error(w, "Error while executing html bytes", 500)
	}
}

func Json(w http.ResponseWriter, jsonData interface{}) {
	json.NewEncoder(w).Encode(jsonData)
}

func SendString(w http.ResponseWriter, value string) {
	w.Header().Set("Content-type", "text/plain")
	w.Write([]byte(value))
}

func GetHeader(r *http.Request, head string) string {
	if head != "" {
		return r.Header.Get(head)
	}
	return "no header"
}

func SetHeader(w http.ResponseWriter, head map[string]string) bool {
	if head != nil {
		for k, v := range head {
			w.Header().Set(k, v)
		}
		return true
	}
	return false
}

func Query(r *http.Request, value string) string {
	return r.URL.Query().Get(value)
}

func SendError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func SetSecureCookie(w http.ResponseWriter, name string, value string, path string) bool {
	if name != "" && value != "" && path != "" {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Value:    value,
			Path:     path,
			HttpOnly: true,
			Secure:   true,
		})
		return true
	} else {
		return false
	}
}

func GetSecureCookie(w http.ResponseWriter, r *http.Request, name string) string {
	a, err := r.Cookie(name)
	if err != nil {
		http.Error(w, err.Error(), 404)
	}
	return a.Value
}

func BindsJson(w http.ResponseWriter, r *http.Request, toData any) error {
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, toData)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return err
	}
	return nil
}
