package models

import (
	"log"

	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/pop"

	"github.com/pkg/errors"
	"fmt"
)

// DB is a connection to your database to be used
// throughout your application.
var DB *pop.Connection

func init() {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}
func Init () {
	var err error
	env := envy.Get("GO_ENV", "development")
	DB, err = pop.Connect(env)
	if err != nil {
		log.Fatal(err)
	}
	pop.Debug = env == "development"
}

func PopulateDB() error {

	return DB.Transaction(func(tx *pop.Connection) error {
		group1 := &Group{GroupName: "group1"}
		if err := tx.Create(group1); err != nil {
			return errors.WithStack(err)
		}

		group2 := &Group{GroupName: "group2"}
		if err := tx.Create(group2); err != nil {
			return errors.WithStack(err)
		}

		mark := &User{Username: "mark"}
		if err := tx.Create(mark); err != nil {
			return errors.WithStack(err)
		}
		peter := &User{Username: "peter"}
		if err := tx.Create(peter); err != nil {
			return errors.WithStack(err)
		}
		ben := &User{Username: "ben"}
		if err := tx.Create(ben); err != nil {
			return errors.WithStack(err)
		}

		ug1:=&UserGroup{
			GroupID:group1.ID,
			UserID:mark.ID,
		}
		if err := tx.Create(ug1); err != nil {
			return errors.WithStack(err)
		}

		ug2:=&UserGroup{
			GroupID:group1.ID,
			UserID:peter.ID,
		}
		if err := tx.Create(ug2); err != nil {
			return errors.WithStack(err)
		}

		ug3:=&UserGroup{
			GroupID:group2.ID,
			UserID:ben.ID,
		}
		if err := tx.Create(ug3); err != nil {
			return errors.WithStack(err)
		}
		return nil
	})

}


func GetUsers()error {
	return DB.Transaction(func(tx *pop.Connection) error {
		u:=Users{}
		err:=tx.Eager().Where("username = ?","mark").All(&u)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(u)
		return nil
	})

}

func GetGroups()error {
	return DB.Transaction(func(tx *pop.Connection) error {
		g:=Groups{}
		err:=tx.Eager("").Where("group_name = ?","group1").All(&g)
		if err!=nil {
			fmt.Println(err)
		}
		fmt.Println(g)
		return nil
	})

}