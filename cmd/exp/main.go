package main

import (
	"fmt"

	"github.com/SHREYAS-KS-CS104/shotDetect/models"
)

const (
	host     = "sandbox.smtp.mailtrap.io"
	port     = 587
	username = "02228f02835c1f"
	password = "760d9e9c5203a6"
)

func main() {
	es := models.NewEmailService(models.SMTPConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	})
	err := es.ForgotPassword("shreyasks.cs18@bmsce.ac.in", "https://shreyas.com/rest-pw?")
	if err != nil {
		panic(err)
	}
	fmt.Println("Email sent")
}

/*
	msg := mail.NewMessage()
	msg.SetHeader("To", to)
	msg.SetHeader("From", from)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/plain", plaintext)
	msg.AddAlternative("text/html", html)
	msg.WriteTo(os.Stdout)

	dialer := mail.NewDialer(host, port, username, password)
	err := dialer.DialAndSend(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Message sent")
}
*/
/*

	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	us := models.UserService{
		DB: db,
	}

	user, err := us.Create("shreyOK@shrey.com", "shrey OK")
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}


	// Create a table...
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id 	SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
	);

	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
	);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")



	// Insert some data..
	name := "Shrey DAMN"
	email := "shreyDamn@shrey.com"

	row := db.QueryRow(`
			INSERT INTO users (name, email)
			VALUES ($1, $2) RETURNING id;`, name, email)
	row.Err()
	var id int
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("User created! id =", id)

	id := 1
	row := db.QueryRow(`
		SELECT name, email
		FROM users
		WHERE id=$1;`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User information: name=%s, email=%s\n", name, email)

	userID := 1
	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err := db.Exec(`
			INSERT INTO orders(user_id, amount, description)
			VALUES($1,$2,$3)`, userID, amount, desc)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("Created fake orders")



	type Order struct {
		ID          int
		UserID      int
		Amount      int
		Description string
	}
	var orders []Order
	userID := 1
	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1`, userID)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.Description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Orders:", orders)
}
*/
