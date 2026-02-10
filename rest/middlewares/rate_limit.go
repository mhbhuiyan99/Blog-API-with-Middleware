package middleware

import (
	"blogAPI/util"
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func (m *Middlewares) RateLimit(next http.Handler) http.Handler {

	// Define a client struct to hold the rate limiter and last seen time for each client.
	type client struct {
		limiter *rate.Limiter
		lastSeen time.Time
	}

	var (
		mu sync.Mutex
		// Update the map so the values are pointers to a client struct.
		clients = make(map[string]*client)
	)

	// Launch a background goroutine which removes old entries from the clients map once every minute.
	go func() {
		for {
			time.Sleep(time.Minute)

			// Lock the mutex to prevent any rate limiter checks from happening while the cleanup is taking place.
			mu.Lock()

			// Loop through all clients. If they haven't been seen within the last three
            // minutes, delete the corresponding entry from the map.

			for ip, client := range clients {
				if time.Since(client.lastSeen) > 3*time.Minute {
					delete(clients, ip)
				}
			}

			// Importantly, unlock the mutex when the cleanup is complete.
			mu.Unlock()
		}
	}()

	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request) {
		
		if m.cnf.RateLimiter.Enabled{
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			mu.Lock()

			if _, found := clients[ip]; !found {

				// Create and add a new client struct to the map if it doesn't already exist.
				clients[ip] = &client{
					limiter: rate.NewLimiter(
						rate.Limit(m.cnf.RateLimiter.RPS), 
						m.cnf.RateLimiter.Burst), 
				}
			}

				// Update the last seen time for the client.
				clients[ip].lastSeen = time.Now()

				if !clients[ip].limiter.Allow() {
					mu.Unlock()
					util.SendError(w, "Too Many Requests", http.StatusTooManyRequests)
					return
				}

				mu.Unlock()
		}

			next.ServeHTTP(w, r)
		})
}