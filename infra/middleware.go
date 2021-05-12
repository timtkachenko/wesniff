package infra

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/go-openapi/runtime/middleware"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func Middleware(handler http.Handler) http.Handler {
	return Views(HmacVerify(handler))
}

func HmacVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		const HMAC_HEADER = "x-hmac-signature"
		const PRIVATE_KEY = "cb5473c8-a928-4500-a5f6-a1e0e2a41d6e"
		signature := r.Header.Get(HMAC_HEADER)

		buf := GetBuffer()
		defer PutBuffer(buf)

		_, err := buf.ReadFrom(r.Body)
		r.Body.Close()
		if err != nil {
			log.Println(err)
		} else {
			val := buf.Bytes()
			log.Println("Received request [", r.URL.Path, "] with body: ", string(val))
			//
			key := []byte(PRIVATE_KEY)
			h := hmac.New(sha256.New, key)
			h.Write(val)
			recreated := hex.EncodeToString(h.Sum(nil))
			if recreated != signature {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				w.Write([]byte(`{"code":-1,"message":"Forbidden","data":null}`))
				return
			}
			r.Body = ioutil.NopCloser(bytes.NewReader(val))
		}
		next.ServeHTTP(w, r)
	})
}

func ApiDoc(handler http.Handler, path string, spec []byte) http.Handler {
	specUrl := path + "/swagger.json"
	opts := middleware.RedocOpts{BasePath: path, SpecURL: specUrl}
	opts.EnsureDefaults()
	return middleware.Redoc(opts,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.URL.Path == specUrl {
				w.WriteHeader(http.StatusOK)
				w.Write(spec)
				return
			}
			handler.ServeHTTP(w, r)
		}))
}
func Views(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
		} else {
			//http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			verify, err := getPackrFile("verify")
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(verify)
			return
		}
	})
}
