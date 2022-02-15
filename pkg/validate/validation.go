package validate

import (
	"fmt"
	"github.com/Tayduro/registration-web-server/pkg/config"
	"github.com/jmoiron/sqlx"
	"regexp"
)

type ValidationErr struct {
	FieldValue string
	ErrMassage string
}

func (v *ValidationErr) Error() string {
	return fmt.Sprintf("filed %s, %s", v.FieldValue, v.ErrMassage)
}

func Length(min, max int, str string) string {
	l := len(str)
	if l < min || l > max {
		return fmt.Sprintf("filed should be not less that %v and not greater that %v", min, max)
	}
	return ""
}

func Email(email string) string {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !emailRegex.MatchString(email) {
		return fmt.Sprintf("incorrect email format")
	}
	return ""
}

func UniqueEmail (email string) string {
	connstring := config.ConfigServer()

	db, err := sqlx.Connect("postgres", connstring)

	if err != nil {
		panic(err)
	}

	defer db.Close()


var dbEmail string
	 err = db.QueryRow("SELECT email FROM users WHERE email= $1", email).Scan(&dbEmail)
	 if dbEmail == email {
		 return "this email is already in use"
	 }

	return ""
}
