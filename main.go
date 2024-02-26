package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/google/uuid"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

func main() {
	typeGen := flag.String("t", "ulid", "type of value to generate")
	convertType := flag.String("c", "", "type to convert to")

	flag.Parse()

	var res string

	switch *typeGen {
	case "ulid":
		switch *convertType {
		case "uuid":
			ulidVal := generateUlid()
			res = convertUlidToUuidString(ulidVal)
		default:
			res = generateUlid().String()
		}
	case "uuid":
		res = generateUuid()
	case "symfony":
		res = generateRandomString()
	}

	fmt.Println(res)

	copyToClipboard(res)
}

func generateUlid() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.New(rand.NewSource(time.Now().UnixNano())))
}

func generateUuid() string {
	return uuid.NewString()
}

func convertUlidToUuidString(ulidVal ulid.ULID) string {
	ulidBytes := ulidVal[:]

	if len(ulidBytes) != 16 {
		log.Fatal("data must be exactly 16 bytes long")
	}

	hexStr := hex.EncodeToString(ulidBytes)
	return fmt.Sprintf("%s-%s-%s-%s-%s", hexStr[:8], hexStr[8:12], hexStr[12:16], hexStr[16:20], hexStr[20:])
}

func generateRandomString() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"

	bytes := make([]byte, 64)

	_, err := rand.Read(bytes)

	if err != nil {
		log.Fatal(err)
	}

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}

	return string(bytes)
}

func copyToClipboard(s string) {
	if err := clipboard.WriteAll(s); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Copied to clipboard!")
}
