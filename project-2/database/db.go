package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
	err error
)

func handleDatabaseConnection() {

	appConfig := AppConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		appConfig.DbHost,
		appConfig.DbPort,
		appConfig.DbUser,
		appConfig.DbPassword,
		appConfig.DbName,
	)

	db, err = sql.Open(appConfig.DbDialect, dsn)

	if err != nil {
		log.Panicln("error occured while trying to validate database arguments: ", err.Error())
		return
	}

	if err := db.Ping(); err != nil {
		log.Panicln("error occured while trying to connect to database :", err.Error())
		return
	}

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(10 * time.Minute)
}

func handleRequiredTables() {
	const (
		createTableUsersQuery = `
			CREATE TABLE IF NOT EXISTS
				users
					(
						id SERIAL PRIMARY KEY,
						username VARCHAR(25) NOT NULL,
						email VARCHAR(50) NOT NULL,
						password TEXT NOT NULL,
						age INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							unique_email
								UNIQUE(email),
						CONSTRAINT
							unique_username
								UNIQUE(username)
					)
		`
		createTablePhotosQuery = `
			CREATE TABLE IF NOT EXISTS
				photos
					(
						id SERIAL PRIMARY KEY,
						title TEXT NOT NULL,
						caption TEXT NOT NULL,
						photo_url TEXT NOT NULL,
						user_id INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_photos_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE
					)
		`
		createTableSocialMediaQuery = `
			CREATE TABLE IF NOT EXISTS
				social_media
					(
						id SERIAL PRIMARY KEY,
						name VARCHAR(50) NOT NULL,
						social_media_url TEXT NOT NULL,
						user_id INT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_social_media_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE
					)
		`
		createTableCommentsQuery = `
			CREATE TABLE IF NOT EXISTS
				comments
					(
						id SERIAL PRIMARY KEY,
						user_id INT NOT NULL,
						photo_id INT NOT NULL,
						message TEXT NOT NULL,
						created_at TIMESTAMPTZ DEFAULT now(),
						updated_at TIMESTAMPTZ DEFAULT now(),
						CONSTRAINT
							fk_comments_user_id
								FOREIGN KEY(user_id)
									REFERENCES
										users(id)
											ON DELETE CASCADE,
						CONSTRAINT 
							fk_comments_photo_id
								FOREIGN KEY(photo_id)
									REFERENCES
										photos(id)
											ON DELETE CASCADE
					)
		`
	)

	_, err = db.Exec(createTableUsersQuery)

	if err != nil {
		log.Panic("error while create table users : ", err.Error())
		return
	}

	_, err = db.Exec(createTablePhotosQuery)

	if err != nil {
		log.Panic("error while create table photos : ", err.Error())
		return
	}

	_, err = db.Exec(createTableSocialMediaQuery)

	if err != nil {
		log.Panic("error while create table social_media : ", err.Error())
		return
	}

	_, err = db.Exec(createTableCommentsQuery)

	if err != nil {
		log.Panic("error while create table comments : ", err.Error())
		return
	}

}

func InitializeDatabase() {
	handleDatabaseConnection()
	handleRequiredTables()
}

func GetInstanceDatabaseConnection() *sql.DB {
	return db
}
