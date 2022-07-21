package main

import (
	"ageCat/ageApi"
	"ageCat/email"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	account       string
	password      string
	path          string
	emailPassword string
	smtpPath      string
	smtpPort      string
)

const emailTemplate = `your anime have updated!<br/>`

func main() {
	initInfos()
	testInfos()
	// test eamil function
	err := email.SendEmail([]string{path}, path, emailPassword, smtpPath, smtpPort, []byte("HelloWorld, there is ageCat, this email is function test"))
	if err != nil {
		log.Fatalln(err)
	}
	me := ageApi.AgeHandleConstructor(account, password)
	_, err = me.UpdateData()
	if err != nil {
		log.Fatal(err)
	}
	tick := time.Tick(time.Second * 5)
	for range tick {
		data, err := me.UpdateData()
		if err != nil {
			log.Println(err)
			continue
		}
		if data != nil {
			var info = emailTemplate
			for i := 0; i < len(data); i++ {
				info += "<p>" + strconv.Itoa(i+1) + ". " + data[i].Title + " " + data[i].NewTitle + "</p><br/>"
			}
			err := email.SendEmail([]string{path}, path, emailPassword, smtpPath, smtpPort, []byte(info))
			for err != nil {
				log.Println(err)
				time.Sleep(time.Second * 5)
				err = email.SendEmail([]string{path}, path, emailPassword, smtpPath, smtpPort, []byte(info))
			}
		}
	}

}
func initInfos() {
	initAccountInfo()
	initEmailInfo()
}

func testInfos() {
	testAccountInfo()
	testEmailInfo()
}

func initAccountInfo() {
	account = os.Getenv("account")
	password = os.Getenv("password")
}

func initEmailInfo() {
	path = os.Getenv("email")
	emailPassword = os.Getenv("emailPassword")
	smtpPath = os.Getenv("smtp")
	smtpPort = os.Getenv("smtpPort")
}

func testAccountInfo() {
	fmt.Println("account:", account)
	fmt.Println("password:", password)
	if account == "" || password == "" {
		log.Fatal("account or password is empty")
	}
}
func testEmailInfo() {
	fmt.Println("email:", path)
	fmt.Println("emailPassword:", emailPassword)
	fmt.Println("smtp:", smtpPath)
	fmt.Println("smtpPort:", smtpPort)
	if path == "" || emailPassword == "" || smtpPath == "" || smtpPort == "" {
		log.Fatal("email info is empty")
	}
}
