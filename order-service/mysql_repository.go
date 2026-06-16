// package main

// type MySQLRepository struct{}

// func (r MySQLRepository) Save(order Order) error {

// 	query := `
// 	INSERT INTO orders
// 	(user_id, nama_barang, berat, dimensi,
// 	jenis, alamat_pengirim, alamat_penerima, status)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	_, err := DB.Exec(
// 		query,
// 		order.UserID,
// 		order.NamaBarang,
// 		order.Berat,
// 		order.Dimensi,
// 		order.Jenis,
// 		order.AlamatPengirim,
// 		order.AlamatPenerima,
// 		order.Status,
// 	)

// 	return err
// }

// package main

// type MySQLRepository struct{}

// func (r MySQLRepository) Save(order Order) error {

// 	query := `
// 	INSERT INTO orders
// 	(
// 		user_id,
// 		resi,
// 		nama_barang,
// 		berat,
// 		dimensi,
// 		jenis,
// 		alamat_pengirim,
// 		alamat_penerima,
// 		status,
// 		eta
// 	)
// 	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
// 	`

// 	_, err := DB.Exec(
// 		query,
// 		order.UserID,
// 		order.Resi,
// 		order.NamaBarang,
// 		order.Berat,
// 		order.Dimensi,
// 		order.Jenis,
// 		order.AlamatPengirim,
// 		order.AlamatPenerima,
// 		order.Status,
// 		order.ETA,
// 	)

// 	return err
// }

package main

type MySQLRepository struct{}

func (r MySQLRepository) GetByID(id int) (*Order, error) {

	query := `
	SELECT
		order_id,
		user_id,
		resi,
		nama_barang,
		berat,
		dimensi,
		jenis,
		alamat_pengirim,
		alamat_penerima,
		status,
		eta
	FROM orders
	WHERE order_id = ?
	`

	var o Order

	err := DB.QueryRow(query, id).Scan(
		&o.OrderID,
		&o.UserID,
		&o.Resi,
		&o.NamaBarang,
		&o.Berat,
		&o.Dimensi,
		&o.Jenis,
		&o.AlamatPengirim,
		&o.AlamatPenerima,
		&o.Status,
		&o.ETA,
	)

	if err != nil {
		return nil, err
	}

	return &o, nil
}

func (r MySQLRepository) UpdateStatus(
	id int,
	status string,
) error {

	query := `
	UPDATE orders
	SET status = ?
	WHERE order_id = ?
	`

	_, err := DB.Exec(
		query,
		status,
		id,
	)

	return err
}

func (r MySQLRepository) Save(order Order) error {

	query := `
	INSERT INTO orders
	(
		user_id,
		resi,
		nama_barang,
		berat,
		dimensi,
		jenis,
		alamat_pengirim,
		alamat_penerima,
		status,
		eta
	)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := DB.Exec(
		query,
		order.UserID,
		order.Resi,
		order.NamaBarang,
		order.Berat,
		order.Dimensi,
		order.Jenis,
		order.AlamatPengirim,
		order.AlamatPenerima,
		order.Status,
		order.ETA,
	)

	return err
}
