package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"github.com/oklog/ulid"
	"log"
	"math/rand"
	"time"
)

func main() {
	typeGen := flag.String("type", "ulid", "type of value to generate")
	convertType := flag.String("convert", "", "type to convert to")

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
	}

	fmt.Println(res)
}

func generateUlid() ulid.ULID {
	return ulid.MustNew(ulid.Timestamp(time.Now()), rand.New(rand.NewSource(time.Now().UnixNano())))
}

func convertUlidToUuidString(ulidVal ulid.ULID) string {
	ulidBytes := ulidVal[:]

	if len(ulidBytes) != 16 {
		log.Fatal("data must be exactly 16 bytes long")
	}

	hexStr := hex.EncodeToString(ulidBytes)
	return fmt.Sprintf("%s-%s-%s-%s-%s", hexStr[:8], hexStr[8:12], hexStr[12:16], hexStr[16:20], hexStr[20:])
}
