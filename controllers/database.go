package controllers

import "github.com/boltdb/bolt"

var globalDB *bolt.DB

//SetDB register global DB instance
func SetDB(db *bolt.DB) {
	globalDB = db
}

//GetDB get registered DB instance
func GetDB() *bolt.DB {
	return globalDB
}
