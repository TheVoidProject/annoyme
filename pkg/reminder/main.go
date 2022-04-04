package reminder

import (
	"encoding/gob"
	"os"
	"net"
	"time"
	"path/filepath"
	"log"

	// "github.com/TheVoidProject/annoyme/pkg/logger"
	"github.com/gen2brain/beeep"
	// "github.com/sirupsen/logrus"
)


type Reminder struct {
	Title 			string
	Message 		string
	Date				string
	Time 				string
	Recurring 	bool
	Repeat 			int
	Delay 			time.Duration
	Sound 			bool
}

// var (
// 	stdout logrus.Logger
// 	log logrus.Logger
// )

var ANNOYME_LOCAL_DIR string

func init() {
	// stdout, log = logger.New("reminder")

	HOME_DIR, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	// AppData/Roaming/annoyme/
	ANNOYME_LOCAL_DIR = filepath.Join(HOME_DIR, ".local", "share", "annoyme")
}

func New(t string, m string) Reminder {
	return Reminder{
		Title: t,
		Message: m,
		Date: "",
		Time: "",
		Recurring: false,
		Repeat: 1,
		Delay: 5 * time.Minute,
		Sound: false,
	}
}

func Encode(r Reminder, conn net.Conn) {
	e := gob.NewEncoder(conn)
	e.Encode(r)
}

func Decode(conn net.Conn) Reminder {
	r := &Reminder{}
	d := gob.NewDecoder(conn)
	d.Decode(r)
	return *r
}

func (r Reminder) Notify() {
	icon := filepath.Join(ANNOYME_LOCAL_DIR, "icon.png")
	var err error
	if (r.Sound) {
		err = beeep.Alert(r.Title, r.Message, icon)
	} else {
		err = beeep.Notify(r.Title, r.Message, icon)
	}
	if err != nil {
		log.Fatal(err)
	}
}