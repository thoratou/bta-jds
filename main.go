package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/astaxie/beego"
	"github.com/boltdb/bolt"
	"github.com/thoratou/cgi-jds/controllers"
)

func main() {
	db, err := bolt.Open("bolt.db", 0600, &bolt.Options{Timeout: 10 * time.Second})
	if err != nil {
		log.Fatal(err)
		return
	}

	db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists([]byte("User"))
		return nil
	})

	controllers.SetDB(db)

	defer db.Close()

	rand.Seed(time.Now().UnixNano())

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/signin", &controllers.HomeController{}, "post:SignInQuery")
	beego.Router("/signup", &controllers.HomeController{}, "post:SignUpQuery")
	beego.Run()
}
