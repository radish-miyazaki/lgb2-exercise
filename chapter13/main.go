package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"
)

func IPAddrMiddleware(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			remoteAddr := r.RemoteAddr
			ip, _, err := net.SplitHostPort(remoteAddr)
			if err != nil {
				ip = remoteAddr
			}

			logger.Info("IP address", "ip", ip)

			next.ServeHTTP(w, r)
		})
	}
}

type NowResponse struct {
	DayOfWeek  string `json:"day_of_week"`
	DayOfMonth int    `json:"day_of_month"`
	Month      string `json:"month"`
	Year       int    `json:"year"`
	Hour       int    `json:"hour"`
	Minute     int    `json:"minute"`
	Second     int    `json:"second"`
}

func main() {
	// ルーティング
	mux := http.NewServeMux()
	mux.HandleFunc("GET /now", func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		fmt.Println(r.Header.Get("Accept"))
		if ok := slices.Contains(strings.Split(r.Header.Get("Accept"), ","), "application/json"); ok == false {
			// if Accept does not contain application/json, return text/plain
			// FIXME: not consider quality (e.g. application/json;q=0.9) and white space (e.g. application/json; charset=utf-8)
			w.Write([]byte(now.Format(time.RFC3339)))
			return
		}

		resp := NowResponse{
			DayOfWeek:  now.Weekday().String(),
			DayOfMonth: int(now.Month()),
			Month:      now.Month().String(),
			Year:       now.Year(),
			Hour:       now.Hour(),
			Minute:     now.Minute(),
			Second:     now.Second(),
		}
		b, err := json.Marshal(resp)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			slog.Error("parse error", "error", err)
			w.Write([]byte("Internal Server Error"))
		}
		w.Write(b)
	})

	// ミドルウェア
	options := &slog.HandlerOptions{Level: slog.LevelDebug}
	handler := slog.NewJSONHandler(os.Stderr, options)
	sl := slog.New(handler)
	ipAddressMiddleware := IPAddrMiddleware(sl)
	wrappedMux := ipAddressMiddleware(mux)

	// サーバの起動
	server := http.Server{
		Addr:         ":3000",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      wrappedMux,
	}
	sl.Info("Start server", "port", ":3000")
	err := server.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}
