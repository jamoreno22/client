package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	l3 "github.com/jamoreno22/client/pkg/proto"
	"google.golang.org/grpc"
)

type consistency struct {
	zfName string
	rv     l3.VectorClock
	ip     string
}

var cons []consistency

func main() {

	var brokerIP string
	brokerIP = "10.10.28.20:8000"
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(brokerIP, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	cc := l3.NewBrokerClient(conn)

	var command string
	var comm l3.Command

	defer conn.Close()
	for {
		fmt.Println("Ingrese comando")
		fmt.Scanln(&command)
		split := strings.Split(command, " ")
		split2 := strings.Split(split[0], ".")
		switch split[0] {
		case "Get":
			comm = l3.Command{Action: 4, Name: split2[0], Domain: split2[1], Option: "", Parameter: ""}
		default:
			log.Println("Ingrese un comando válido")
			continue
		}
		runDNSIsAvailable(cc, command)

	}

}

func runDNSIsAvailable(cc l3.BrokerClient, comm string) error {
	msg := l3.Message{Text: comm}
	_, err := cc.DNSIsAvailable(context.Background(), &msg)
	return err
}

func pingDataNode(ip string) bool {
	timeOut := time.Duration(10 * time.Second)
	_, err := net.DialTimeout("tcp", ip, timeOut)
	if err != nil {
		return false
	}
	return true
}

func mReads(zfName string, rv *l3.VectorClock, ip *string, com *l3.Command){
	if len(consList) != 0 {
		for _, s := range consList {
			if s.zfName == zfName {
				spl := strings.Split(*ip, ".")
				switch spl[3] {
				case "17":
					if s.rv.Rv1 <= /*no se de donde sacar el rvector que deberia tener*/ {
						s.ip = *ip
						s.rv = *actionRv
						s.com = *com
					} else {
						log.Println("Existe un error en la consistencia")
					}
				case "18":
					if s.rv.Rv2 <= /*no se de donde sacar el rvector que deberia tener*/ {
						s.ip = *ip
						s.rv = *actionRv
						s.com = *com
					} else {
						log.Println("Existe un error en la consistencia")
					}
				case "19":
					if s.rv.Rv3 <= /*no se de donde sacar el rvector que deberia tener*/ {
						s.ip = *ip
						s.rv = *actionRv
						s.com = *com
					} else {
						log.Println("Existe un error en la consistencia")
					}
				}
			}
		}
	} else {
		consList = append(consList, &Consistency{zfName: zfName, rv: *rv, ip: *ip, com: *com})
	}
}