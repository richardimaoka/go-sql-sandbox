package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql",
		"root:my-secret-pw@tcp(127.0.0.1:13306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	rows, err := db.Query("CREATE TABLE `entry` (`id` INTEGER NOT NULL AUTO_INCREMENT,`title` VARCHAR(100) NOT NULL,`public` TINYINT(1) NOT NULL DEFAULT 0,`content` TEXT NOT NULL,`created_at` DATETIME NOT NULL,`updated_at` DATETIME NOT NULL,FULLTEXT `full_text_idx` (`content`) WITH PARSER `ngram`,INDEX `created_at_idx` (`created_at`),INDEX `title_idx` (`title`),UNIQUE `created_at_uniq_idx` (`created_at`),PRIMARY KEY (`id`, `created_at`)) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4;")
	log.Println(rows)
	defer db.Close()
}
