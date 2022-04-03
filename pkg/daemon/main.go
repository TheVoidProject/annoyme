package daemon

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	d "github.com/takama/daemon"
)

const (

    // name of the service
    name        = "annoymed"
    description = "annoyme daemon that handles scheduling and sending notifications"

    // port which daemon should be listen
    port = ":9977"
)

//    dependencies that are NOT required by the service, but might be used
var dependencies = []string{}

var stdlog, errlog *log.Logger

// Service has embedded daemon
type Service struct {
    d.Daemon
}

// Manage by daemon commands or run the daemon
func (service *Service) Manage(command string) (string, error) {

    // usage := "Usage: myservice install | uninstall | start | stop | status"
		usage := "annoyme [--daemon|-d] install | uninstall | start | stop | status"

    // if received any kind of command, do it
    if command != "run" {
    //     command := os.Args[2]
			switch command {
			case "install":
					return service.Install("--daemon run")
			case "uninstall":
					return service.Remove()
			case "start":
					return service.Start()
			case "stop":
					return service.Stop()
			case "status":
					return service.Status()
			default:
					return usage, nil
			}
    }

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

    // loop work cycle with accept connections or interrupt
    // by system signal
    for {
        select {
        case conn := <-listen:
            go handleClient(conn)
        case killSignal := <-interrupt:
            stdlog.Println("Got signal:", killSignal)
            stdlog.Println("Stoping listening on ", listener.Addr())
            listener.Close()
            if killSignal == os.Interrupt {
                return "Daemon was interruped by system signal", nil
            }
            return "Daemon was killed", nil
        }
    }

    // never happen, but need to complete code
    return usage, nil
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

func handleClient(client net.Conn) {
    for {
        buf := make([]byte, 4096)
				buf = append(buf, "echo: "...)
        numbytes, err := client.Read(buf)
        if numbytes == 0 || err != nil {
            return
        }
        client.Write(buf[:numbytes])
				client.Write([]byte("hello world"))
    }
}

func init() {
    stdlog = log.New(os.Stdout, "", log.Ldate|log.Ltime)
    errlog = log.New(os.Stderr, "", log.Ldate|log.Ltime)
}

func Control(cmd string) {
    srv, err := d.New(name, description, d.SystemDaemon, dependencies...)
    if err != nil {
        errlog.Println("Error: ", err)
        os.Exit(1)
    }

    service := &Service{srv}
    status, err := service.Manage(cmd)
    if err != nil {
        errlog.Println(status, "\nError: ", err)
        os.Exit(1)
    }
    fmt.Println(status)
}