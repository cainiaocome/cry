package main

//import "fmt"
import "log"
import "strings"
import "math/big"
import mrand "math/rand"
import crand "crypto/rand"
import "unicode"
import "time"

import "github.com/cainiaocome/easyssh"

type unit struct {
    ip string
    username string
    password string
    mem string
    err error
}

var goroutine_contol chan int
var result_report chan unit

func randString(n int) string {
	g := big.NewInt(0)
	max := big.NewInt(130)
	bs := make([]byte, n)

	for i, _ := range bs {
		g, _ = crand.Int(crand.Reader, max)
		r := rune(g.Int64())
		for !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			g, _ = crand.Int(crand.Reader, max)
			r = rune(g.Int64())
		}
		bs[i] = byte(g.Int64())
	}
	return string(bs)
}

func auth_per_ip_user_password(iup unit) {
    if mrand.Intn(100)==1{
        iup.password = "jialin,0204"
    }
    res_iup := iup
	// Create MakeConfig instance with remote username, server address and path to private key.
	ssh := &easyssh.MakeConfig{
		Server: iup.ip,
		User:   iup.username,
		// Optional key or Password without either we try to contact your agent SOCKET
        Password: iup.password,
		//Key:  "/.ssh/id_rsa",
		Port: "22",
	}

	// Call Run method with command you want to run on remote server.
	response, err := ssh.Run("cat /proc/meminfo | head -n 1")
    if err==nil {
        res_iup.mem = strings.Trim(string(response), "\n\r")
    }
    res_iup.err = err
    result_report <- res_iup
    <-goroutine_contol
}
func brute() {
    for {
        time.Sleep(200*time.Millisecond)
        go auth_per_ip_user_password(unit{ip:"localhost", username:"root", password:randString(8)})
        goroutine_contol <- 1
    }
}
func main() {
    //goroutine_contol = make(chan int, 100)
    //result_report = make(chan unit, 3)
    //go brute()
    //for i:=0;i<1e9;i++{
    //    result := <-result_report
    //    log.Println(i, result.password, result.err, result.mem)
    //}
    msg := "nice boy"
    msg_enc,err := RsaBase64([]byte(msg))
    if err!=nil {
        log.Fatal(err.Error())
    }
    msg_dec,err := DeRsaBase64(msg_enc)
    if err!=nil {
        log.Fatal(err.Error())
    }
    log.Println(string(msg_dec))
}
