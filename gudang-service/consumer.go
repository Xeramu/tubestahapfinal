package main

import (
	"encoding/json"
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Fungsi untuk mendengarkan pesanan masuk dari order-service
func StartOrderConsumer(repo *PackageRepository) {
	// Mengambil URL RabbitMQ dari environment variable seperti di main.go
	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		rabbitURL = "amqp://guest:guest@rabbitmq:5672/"
	}

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Printf("[Consumer Error] Gagal koneksi ke RabbitMQ: %v", err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[Consumer Error] Gagal membuka channel: %v", err)
		return
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"order_queue", // Harus sama persis dengan order-service
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("[Consumer Error] Gagal declare queue: %v", err)
		return
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true, // auto-ack
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("[Consumer Error] Gagal register consumer: %v", err)
		return
	}

	log.Println(" [*] Gudang Service stand-by menunggu data dari Order Service...")

	// Loop untuk membaca pesan terus-menerus
	for d := range msgs {
		var pkg Package
		err := json.Unmarshal(d.Body, &pkg)
		if err != nil {
			log.Printf("Gagal unmarshal data paket: %v", err)
			continue
		}

		// Set default status dan zona sebelum disimpan ke database
		pkg.Status = "pending"
		pkg.WarehouseZone = "ZONE-A"

		// Menggunakan repository milik temanmu untuk menyimpan data
		// Catatan: Pastikan di repository.go temanmu punya fungsi Create atau Save untuk Package
		err = repo.Create(&pkg)
		if err != nil {
			log.Printf("Gagal menyimpan paket ke DB via Repository: %v", err)
		} else {
			log.Printf("[Gudang] Sukses! Paket dengan Resi %s berhasil disimpan via Repository.", pkg.Resi)
		}
	}
}
