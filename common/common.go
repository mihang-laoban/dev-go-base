package common

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Digit    = "0123456789"
	LowChar  = "abcdefghijklmnopqrstuvwxyz"
	HighChar = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func ErrorHandling(err error) {
	if err != nil {
		log.Printf("stack trace: %+v", err)
	}
}

func GetRandomString(l int, str ...string) string {
	builder := strings.Builder{}
	for _, v := range str {
		builder.WriteString(v)
	}
	bytes := builder.String()
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func NumGetter(text string) (num int, err error) {
	reg := regexp.MustCompile(`[\d]+`)
	r := fmt.Sprintf("%s\n", reg.FindAllString(text, -1)) //[Golang]
	if num, err = strconv.Atoi(r[1 : len(r)-2]); err != nil {
		return num, errors.New(fmt.Sprintf("%s", num))
	}
	return num, nil
}
