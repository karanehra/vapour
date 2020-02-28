package middlewares

import "net/http"

//CORSMiddleware adds the access control headers to requests
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		res.Header().Add("Access-Control-Allow-Origin", "http://localhost:8000")
		res.Header().Add("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS, PUT")
		if req.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(res, req)
	})
}
