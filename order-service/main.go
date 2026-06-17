package main

import (
	"fmt"
	"net/http"
)

func main() {
	ConnectDB()

	// 1. Daftarkan tanpa dibungkus enableCORS
	http.HandleFunc("/order", createOrderHandler)
	http.HandleFunc("/order/get", getOrderHandler)
	http.HandleFunc("/order/status", updateOrderHandler)
	http.HandleFunc("/order/eta", etaHandler)

	// 2. Ambil default mux
	mux := http.DefaultServeMux

	// 3. Bungkus global
	globalCORS := enableCORS(mux.ServeHTTP)

	fmt.Println("Order Service running on :8083")
	// 4. Jalankan dengan globalCORS
	http.ListenAndServe(":8083", globalCORS)
}
