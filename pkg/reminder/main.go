package reminder

import (
	"time"
	"net"
	"encoding/gob"

	"github.com/gen2brain/beeep"
)


type Reminder struct {
	Title 			string
	Message 		string
	Datetime 		time.Time
	Recurring 	bool
	Repeat 			int
	Delay 			int
}

func Encode(r Reminder, conn net.Conn) {
	e := gob.NewEncoder(conn)
	e.Encode(r)
}

func Decode(conn net.Conn) *Reminder {
	r := &Reminder{}
	d := gob.NewDecoder(conn)
	d.Decode(r)
	return r
}

func (r Reminder) Notify() error {
	return beeep.Notify(r.Title, r.Message, "assets/reminder.png")
}