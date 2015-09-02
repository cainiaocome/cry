package main

//import "log"
import "crypto/rsa"
import "crypto/rand"
import "encoding/base64"

// rsa and then base64, operate on []byte
func RsaBase64(msg_to_enc []byte) (msg_encrypted string, err error) {
	private_key := get_private_key()
	publick_key := private_key.PublicKey

	msg_rsaed, err := rsa.EncryptPKCS1v15(rand.Reader, &publick_key, msg_to_enc)
    if err!=nil {
        return
    }
    msg_rsaed_base64ed := base64.StdEncoding.EncodeToString(msg_rsaed)
    msg_encrypted = msg_rsaed_base64ed
    return
}
// dont put it in client
// debase4 and then dersa, return []byte, well, mostly should be json
func DeRsaBase64(msg_encrypted string) (msg_decrypted []byte, err error) {
	private_key := get_private_key()
	//publick_key := private_key.PublicKey

	msg_debase64ed, err := base64.StdEncoding.DecodeString(msg_encrypted)
    if err!=nil {
        return
    }
	msg_debase64ed_dersaed, err := rsa.DecryptPKCS1v15(rand.Reader, private_key, msg_debase64ed)
    msg_decrypted = msg_debase64ed_dersaed
    return
}
