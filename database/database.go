package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"orid19.com/ecommerce/api/types"
)

type Store interface {
	CreateUser(user types.User) error
	GetUser(username string) (types.User, error)
	DoesUserExist(username string) (bool, error)
}

// SQLDatabase implictly implements UserStore
type SQLDatabase struct {
	databaseStore *sql.DB
}

func setupDB(db *sql.DB) error {
	db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL
		);`)

	db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		price REAL NOT NULL
		);`)

	return nil
}

func NewDB() SQLDatabase {
	databaseStore, err := sql.Open("sqlite3", "./database.db")

	setupDB(databaseStore)

	if err != nil {
		return SQLDatabase{}
	}

	return SQLDatabase{databaseStore: databaseStore}
}

func (db SQLDatabase) DoesUserExist(username string) (bool, error) {
	var count int
	err := db.databaseStore.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", username).Scan(&count)
	if err != nil {
		return true, err
	}

	if count == 1 {
		return true, nil
	}

	return false, nil
}

func (db SQLDatabase) GetUser(username string) (types.User, error) {
	var user types.User

	// returns zero or one row
	// if the result is zero rows, Scan will throw an error
	err := db.databaseStore.QueryRow("SELECT * FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.PasswordHash)

	if err != nil {
		// return the empty user var, plus an error
		return user, err
	}
	return user, nil
}

func (db SQLDatabase) CreateUser(user types.User) error {
	// user has been previously processed to hash the password
	_, err := db.databaseStore.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (db SQLDatabase) InsertProduct(product types.Product) error {
	// insert a new product into the database
	_, err := db.databaseStore.Exec("INSERT INTO products (name, price) VALUES (?, ?)", product.Name, product.Price)

	if err != nil {
		return err
	}

	return nil
}

func (db SQLDatabase) GetProducts(offset, limit int) ([]types.Product, error) {
	var products []types.Product

	// get the limit and offset numbers from the request query parameters
	rows, err := db.databaseStore.Query("SELECT * FROM products LIMIT ? OFFSET ?", limit, offset)

	if err != nil {
		return products, err
	}

	for rows.Next() {
		var product types.Product
		err = rows.Scan(&product.ID, &product.Name, &product.Price)

		if err != nil {
			return products, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (db SQLDatabase) GetProduct(id int) (types.Product, error) {
	var product types.Product

	err := db.databaseStore.QueryRow("SELECT * FROM products WHERE id = ?", id).Scan(&product.ID, &product.Name, &product.Price)

	if err != nil {
		return product, err
	}

	return product, nil
}
