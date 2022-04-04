package daemon

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	inner_logger "github.com/TheVoidProject/annoyme/pkg/logger"
	"github.com/TheVoidProject/annoyme/pkg/reminder"
	"github.com/go-co-op/gocron"
	"github.com/sirupsen/logrus"
)

var (
	stdout logrus.Logger
	log logrus.Logger
	scheduler gocron.Scheduler
)
const port string = ":9977"


func init() {
	stdout, log = inner_logger.New("daemon")
	scheduler = *gocron.NewScheduler(time.Now().Location())
	scheduler.StartAsync()
}

func (p *program) run() (string, error) {
    // Do something, call your goroutines, etc

    // Set up channel on which to send signal notifications.
    // We must use a buffered channel or risk missing the signal
    // if we're not ready to receive when the signal is sent.
    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

    // Set up listener for defined host and port
    listener, err := net.Listen("tcp", port)
    if err != nil {
        return "Possibly was a problem with the port binding", err
    }

    // set up channel on which to send accepted connections
    listen := make(chan net.Conn, 100)
    go acceptConnection(listener, listen)


		schedule := make(chan reminder.Reminder, 100)
    // loop work cycle with accept connections or interrupt
    // by system signal
    for {
        select {
				case r := <-schedule:
						go handleJob(r)
        case conn := <-listen:
            go handleClient(conn, schedule)
        case killSignal := <-interrupt:
            stdout.Println("Got signal:", killSignal)
            stdout.Println("Stoping listening on ", listener.Addr())
            listener.Close()
            if killSignal == os.Interrupt {
                return "Daemon was interruped by system signal", nil
            }
            return "Daemon was killed", nil
        }
    }
}

// Accept a client connection and collect it in a channel
func acceptConnection(listener net.Listener, listen chan<- net.Conn) {
    for {
        conn, err := listener.Accept()
        if err != nil {
            continue
        }
        listen <- conn
    }
}

func handleClient(client net.Conn, schedule chan<- reminder.Reminder) {
		schedule <- reminder.Decode(client)
		client.Close()
}

func handleJob(r reminder.Reminder) {
	// scheduler.Every(1).Day().At("18:41").LimitRunsTo(r.Repeat).Do(r.Notify)
	// t := time.Now().Add(10 * time.Second).Format("15:04:05") 
	// stdout.Warning(t)
	scheduler.Every(1).Day().At(r.Time).Do(r.Notify)
	stdout.Warning(scheduler.Jobs())
}