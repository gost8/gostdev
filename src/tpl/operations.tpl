package main

import "net/http"
{{range $key, $function := .}}
{{ range $oneCommentLine := split $function.Description }}
// {{ $oneCommentLine }}	{{ end }}
func {{ $function.Name }}(w http.ResponseWriter, r *http.Request) {
	// TODO func body
}
{{ end }}
