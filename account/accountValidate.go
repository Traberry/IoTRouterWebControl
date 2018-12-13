package account

import (
	"bytes"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/pbkdf2"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	dsn = "postgres://loraserver_as:dbpassword@localhost/loraserver_as?sslmode=disable"
)

type User struct {
	ID           int64     `db:"id"`
	Username     string    `db:"username"`
	IsAdmin      bool      `db:"is_admin"`
	IsActive     bool      `db:"is_active"`
	SessionTTL   int32     `db:"session_ttl"`
	CreatedAt    time.Time `db:"created_at"`
	UpdatedAt    time.Time `db:"updated_at"`
	PasswordHash string    `db:"password_hash"`
	Email        string    `db:"email"`
	Note         string    `db:"note"`
}


func AccountValidate(username, password string) bool {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}

	user := User{}

	err = sqlx.Get(db, &user, "select " + "*" + " from \"user\" where username = $1", username)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return hashCompare(password, user.PasswordHash)

}

func IsGlobalAdmin(username string) bool {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		fmt.Println(err)
	}

	user := User{}

	err = sqlx.Get(db, &user, "select " + "*" + " from \"user\" where username = $1", username)
	if err != nil {
		log.Println("get user information from db error:", err)
	}

	return user.IsAdmin
}

func hashWithSalt(password string, salt []byte, iterations int) string {
	// Generate the hash.  This should be a little painful, adjust ITERATIONS
	// if it needs performance tweeking.  Greatly depends on the hardware.
	// NOTE: We store these details with the returned hash, so changes will not
	// affect our ability to do password compares.
	hash := pbkdf2.Key([]byte(password), salt, iterations, sha512.Size, sha512.New)

	// Build up the parameters and hash into a single string so we can compare
	// other string to the same hash.  Note that the hash algorithm is hard-
	// coded here, as it is above.  Introducing alternate encodings must support
	// old encodings as well, and build this string appropriately.
	var buffer bytes.Buffer

	buffer.WriteString("PBKDF2$")
	buffer.WriteString("sha512$")
	buffer.WriteString(strconv.Itoa(iterations))
	buffer.WriteString("$")
	buffer.WriteString(base64.StdEncoding.EncodeToString(salt))
	buffer.WriteString("$")
	buffer.WriteString(base64.StdEncoding.EncodeToString(hash))

	return buffer.String()
}

// HashCompare verifies that passed password hashes to the same value as the
// passed passwordHash.
func hashCompare(password string, passwordHash string) bool {
	// Split the hash string into its parts.
	hashSplit := strings.Split(passwordHash, "$")

	// Get the iterations and the salt and use them to encode the password
	// being compared.cre
	iterations, _ := strconv.Atoi(hashSplit[2])
	salt, _ := base64.StdEncoding.DecodeString(hashSplit[3])
	newHash := hashWithSalt(password, salt, iterations)
	return newHash == passwordHash
}

