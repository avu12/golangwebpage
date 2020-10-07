package database

import (
	"database/sql"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	temp      int
	date      string
	citycount int
	count     int
	email     string
	quote     string
)

func StartDatabaseUse(dbname string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbname)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func SelectStarExampleQuery() {

	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query(`select * from "WEATHERTABLE"`)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&temp, &date)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(temp, date)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func InsertTempDateCityNameQuery(tablename string, temp int, date time.Time, city string, name string) error {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()
	res, err := db.Exec(`INSERT INTO "WEATHERTABLE" ("Temperature", "Time","City","Name") VALUES ($1, $2, $3::text,$4::text)`, temp, date, city, name)
	if err != nil {
		log.Println(res, err)
		return err
	}

	return nil
}

func CityRateQuery(city string) (int, int, error) {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	cityrows, err := db.Query(`SELECT COUNT("City") FROM "WEATHERTABLE"  WHERE "City" = $1`, city)
	if err != nil {
		log.Println(cityrows, err)
		return -1, -1, err
	}
	for cityrows.Next() {
		err := cityrows.Scan(&citycount)
		if err != nil {
			log.Println(err)
			return -1, -1, err
		}

	}
	defer cityrows.Close()

	rows, err := db.Query(`SELECT COUNT("City") FROM "WEATHERTABLE"`)
	if err != nil {
		log.Println(rows, err)
		return -1, -1, err
	}
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Println(err)
			return -1, -1, err
		}

	}
	defer rows.Close()

	return citycount, count, nil
}

func InsertToMailTableWithoutConfirm(email string, hash string) error {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	res, err := db.Exec(`INSERT INTO "EMAILLIST" ("email","confirmed","emailhash") VALUES ($1::text,$2,$3::text)`, email, false, hash)
	if err != nil {
		log.Println(res, err)
		return err
	}
	return nil
}

func SelectAllConfirmedMail() []string {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	to := []string{}
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT email FROM "EMAILLIST" where confirmed = True`)
	if err != nil {
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			log.Println(err)
		}
		to = append(to, email)
	}
	return to
}
func SelectMail(email string) []string {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	to := []string{}
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT email FROM "EMAILLIST" where ( emailhash = $1 AND confirmed = False )`, email)
	if err != nil {
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&email)
		if err != nil {
			log.Println(err)
		}
		to = append(to, email)
	}
	return to
}

func UpdateConfirmData(emailhash string) error {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("SelectQuote StartDatabaseUse  error: ", err)
	}
	defer db.Close()
	res, err := db.Exec(`UPDATE "EMAILLIST" SET confirmed = True WHERE emailhash = $1 `, emailhash)
	if err != nil {
		log.Println(res, err)
		return err
	}
	return nil
}

//SelectQuote selecting a quote for the daily mail from the database randomly
func SelectQuote() (string, error) {
	db, err := StartDatabaseUse(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Println("SelectQuote StartDatabaseUse  error: ", err)
	}
	defer db.Close()
	numberofquote := 1 + rand.Intn(550)

	rows, err := db.Query(`SELECT quote FROM "QUOTES" where ser = $1`, numberofquote)
	if err != nil {
		log.Println("SelectQuote db.Query error: ", err)
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(&quote)
		if err != nil {
			log.Println("SelectQuote row.Scan error: ", err)
			return "", err
		}
	}
	return quote, nil
}
