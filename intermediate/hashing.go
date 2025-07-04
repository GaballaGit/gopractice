package intermediate 

import(
	"fmt"
	"crypto/rand"
	"io"
	"crypto/sha512"
	"encoding/base64"
)

func main() {

	// For security -> hash and then encode the hash to store it
	// Hash is one way, for security!
	// Encoding specalises in storing binary data!

	//user1 := "Ado"
	salt, err := generateSalt()
	if err != nil {
		panic(err)
	}

	// Encode and store into a database for saftey
	adoSalt := base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("Ado salt: %x\n", adoSalt)

	//user2 := "Teto"
	salt, err = generateSalt()
	if err != nil {
		panic(err)
	}
	tetoSalt := base64.StdEncoding.EncodeToString(salt)
	fmt.Printf("Teto salt: %x\n", tetoSalt)

	user1Pass := "Mikusnum1fan"
	user2Pass := "Mikusnum1fan"

	fmt.Printf("Ado: %x\n", regHash([]byte(user1Pass)))
	fmt.Printf("Teto: %x\n\n", regHash([]byte(user2Pass)))

	hashPass1 := saltedHash([]byte(adoSalt), []byte(user1Pass))
	hashPass2 := saltedHash([]byte(tetoSalt), []byte(user2Pass))

	fmt.Printf("Ado: %x\n", hashPass1)
	fmt.Printf("Teto: %x\n", hashPass2)
}

func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}
	return salt, err
}

func regHash(pass []byte) [64]byte{
	hash := sha512.Sum512(pass)
	return hash
}

func saltedHash(user []byte, pass []byte) string {
	saltedPassword := append(pass, user...)
	hash := sha512.Sum512(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
