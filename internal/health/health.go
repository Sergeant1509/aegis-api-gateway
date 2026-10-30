// Package health provides a process-liveness endpoint.
package health
import ("encoding/json"; "net/http"; "time")
type response struct { Status string `json:"status"`; Time string `json:"time"` }
func Handler() http.Handler { return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request) { w.Header().Set("Content-Type","application/json"); _=json.NewEncoder(w).Encode(response{Status:"ok",Time:time.Now().UTC().Format(time.RFC3339)}) }) }
