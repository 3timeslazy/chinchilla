package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

type Handler struct {
	cache Cache
}

type Cache interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type Redis struct {
	cli *redis.Client
}

func (r *Redis) Get(key string) (string, error) {
	return r.cli.Get(key).Result()
}

func (r *Redis) Set(key, value string) error {
	_, err := r.cli.Set(key, value, time.Minute*5).Result()
	return err
}

func NewRedis(addr string) *Redis {
	return &Redis{cli: redis.NewClient(&redis.Options{Addr: addr})}
}

// TODO: Temporary name
type Hi_PETRA_REQS struct {
	URL string `json:"url"`
}

// TODO: Temporary name
type Hi_PETRA_RESP struct {
	URL  string `json:"url"`
	Tail string `json:"tail"`
}

func main() {
	h := Handler{cache: NewRedis("localhost:6379")}

	http.HandleFunc("/", h.Add)
	// http.HandleFunc("/daa", handlers.Daa)
	http.ListenAndServe(":8080", nil)
}

func (h *Handler) Add(w http.ResponseWriter, r *http.Request) {
	body := new(Hi_PETRA_REQS)
	err := getBodyJSON(r, body)
	if err != nil {
		log.Printf("getBodyJSON: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error. pls try later"))
		return
	}

	tail := randString(2 + rand.Intn(6))

	log.Printf("url: %s, tail: %s", body.URL, tail)
	err = h.cache.Set(tail, body.URL)
	if err != nil {
		log.Printf("redis set: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error. pls try later"))
		return
	}

	w.Write([]byte("success"))
}

func getBodyJSON(r *http.Request, body interface{}) error {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("read from request body: %v", err)
	}
	// after applying ioutil.ReadAll, r.Body will become nil
	// so reinitialize it
	r.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	err = json.Unmarshal(bodyBytes, body)
	if err != nil {
		return fmt.Errorf("unmarshal request body: %v", err)
	}

	return nil
}

func getTail(path string) string {
	return strings.Split(path, "/")[1]
}

func randString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
