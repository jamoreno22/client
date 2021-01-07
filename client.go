package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	l3 "github.com/jamoreno22/client/pkg/proto"
	"google.golang.org/grpc"
)

//Consistency struct
type Consistency struct {
	zfName string
	rv     l3.VectorClock
	ip     string
	com    l3.Command
}

var consList []*Consistency

var getIPRv *l3.VectorClock

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
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			command = scanner.Text()
			split := strings.Split(command, " ")
			split2 := strings.Split(split[1], ".")
			switch split[0] {
			case "Get":
				comm = l3.Command{Action: 4, Name: split2[0], Domain: split2[1], Option: "", Parameter: "", Ip: ""}
			default:
				log.Println("Ingrese un comando válido")
				continue
			}

			localPageInfo, _ := cc.GetIP(context.Background(), &comm)

			log.Println(mReads(comm.Domain, localPageInfo.DnsIP, &comm))
		}
	}

}

func runDNSIsAvailable(cc l3.BrokerClient, comm string) string {
	msg := l3.Message{Text: comm}
	state, err := cc.DNSIsAvailable(context.Background(), &msg)
	if err != nil {
		fmt.Println("DNSIsAvailable error")
	}
	dnsIps := []string{"10.10.28.17:8000", "10.10.28.18:8000", "10.10.28.19:8000"}
	if state.Dns1 == true {
		return dnsIps[0]
	} else if state.Dns2 == true {
		return dnsIps[1]
	} else if state.Dns3 == true {
		return dnsIps[2]
	}
	log.Fatalln("Dns servers not available")
	return "Dns not available"
}

func pingDataNode(ip string) bool {
	timeOut := time.Duration(10 * time.Second)
	_, err := net.DialTimeout("tcp", ip, timeOut)
	if err != nil {
		return false
	}
	return true
}

func mReads(zfName string, ip string, comm *l3.Command) string {
	if len(consList) != 0 {
		for _, s := range consList {
			if s.zfName == zfName {
				spl := strings.Split(ip, ".")
				switch spl[3] {
				case "17":
					if s.rv.Rv1 <= getIPRv.Rv1 {
						s.ip = ip
						s.rv.Rv1 = getIPRv.Rv1
						s.rv.Rv2 = getIPRv.Rv2
						s.rv.Rv3 = getIPRv.Rv3
						s.com.Domain = comm.Domain
						s.com.Ip = comm.Ip
						s.com.Name = comm.Name
						s.com.Option = comm.Option
						s.com.Parameter = comm.Parameter
					} else {
						log.Println("Existe un error en la consistencia")
					}
				case "18":
					if s.rv.Rv2 <= getIPRv.Rv2 {
						s.ip = ip
						s.rv.Rv1 = getIPRv.Rv1
						s.rv.Rv2 = getIPRv.Rv2
						s.rv.Rv3 = getIPRv.Rv3
						s.com.Domain = comm.Domain
						s.com.Ip = comm.Ip
						s.com.Name = comm.Name
						s.com.Option = comm.Option
						s.com.Parameter = comm.Parameter
					} else {
						log.Println("Existe un error en la consistencia")
					}
				case "19":
					if s.rv.Rv3 <= getIPRv.Rv3 {
						s.ip = ip
						s.rv.Rv1 = getIPRv.Rv1
						s.rv.Rv2 = getIPRv.Rv2
						s.rv.Rv3 = getIPRv.Rv3
						s.com.Domain = comm.Domain
						s.com.Ip = comm.Ip
						s.com.Name = comm.Name
						s.com.Option = comm.Option
						s.com.Parameter = comm.Parameter
					} else {
						log.Println("Existe un error en la consistencia")
					}
				}
				return "ip: " + s.ip
			}
		}
	} else {
		consList = append(consList, &Consistency{zfName: comm.Domain,
			rv: l3.VectorClock{Name: comm.Domain, Rv1: 0, Rv2: 0, Rv3: 0},
			ip: ip, com: l3.Command{Action: comm.Action, Name: comm.Name, Domain: comm.Domain, Option: comm.Option, Parameter: comm.Parameter, Ip: comm.Ip}})
	}
	return ""
}
