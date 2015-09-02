package simplersa

import "log"
import "crypto/rsa"
import "crypto/rand"
import "../key"

func Simplersabase64(msg_to_enc string) {
	//msg := "ip:127.0.0.1 username:root password:jialin,0204"
    msg := msg_to_enc
	private_key := key.Get_private_key()
	publick_key := private_key.PublicKey

	encrypted_msg, err := rsa.EncryptPKCS1v15(rand.Reader, &publick_key, []byte(msg))
	if err != nil {
		log.Fatal(err.Error())
	}
	new_msg, err := rsa.DecryptPKCS1v15(rand.Reader, private_key, encrypted_msg)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		log.Printf("%s\n", new_msg)
	}
}
