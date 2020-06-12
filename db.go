package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "os"
import "github.com/joho/godotenv"

func loadDatabase() *sql.DB {
  if godotenv.Load("credentials.env") != nil {
    fmt.Println("Database failed to open")
    return nil
  }

  uname := os.Getenv("DB_USER")
  pword := os.Getenv("DB_PASS")

  fmt.Printf("Username: %s, Password: %s\n", uname, pword)

  db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/classroom", uname, pword))

  if err != nil {
    fmt.Println(err)
    return nil
  }

  return db
}

func createTables(db *sql.DB) {
  db.Exec("CREATE TABLE cameras ( id INT NOT NULL, address VARCHAR(255) NOT NULL, room VARCHAR(20) NOT NULL, framerate INT, bitrate VARCHAR(10), hlsTime INT, hlsWrap INT, codec VARCHAR(20), PRIMARY KEY (id) );")
  db.Exec("CREATE TABLE people ( id INT NOT NULL, uname VARCHAR(20) NOT NULL, fname VARCHAR(20) NOT NULL, lname VARCHAR(20) NOT NULL, role CHAR NOT NULL, PRIMARY KEY(id) );")
  db.Exec("CREATE TABLE classes ( id INT NOT NULL, name VARCHAR(40) NOT NULL, room VARCHAR(20) NOT NULL, stime INT NOT NULL, etime INT NOT NULL, PRIMARY KEY(id) );")
  db.Exec("CREATE TABLE roster ( pid INT NOT NULL, cid INT NOT NULL, FOREIGN KEY (pid) REFERENCES people(id), FOREIGN KEY (cid) REFERENCES classes(id) );")
}

func populateSomeData(db *sql.DB) {
  db.Exec("INSERT INTO cameras ( id, address, room ) VALUES ( 84257, 'rtsp://170.93.143.139/rtplive/470011e600ef003a004ee33696235daa', '4103' );")
  db.Exec("INSERT INTO people VALUES ( 18427, 'jeegan21', 'Joseph', 'Egan', 'S' );")
  db.Exec("INSERT INTO people VALUES ( 659244, 'mtegan22', 'Max', 'Egan', 'S' );")
  db.Exec("INSERT INTO people VALUES ( 472662, 'regan', 'Rose', 'Egan', 'T' );")
  db.Exec("INSERT INTO classes VALUES ( 88231, 'Spanish III X', '3301', 955, 1050 );")
  db.Exec("INSERT INTO roster VALUES ( 18427, 88231 );")
  db.Exec("INSERT INTO roster VALUES ( 659244, 88231 );")
  db.Exec("INSERT INTO roster VALUES ( 472662, 88231 );")
}
