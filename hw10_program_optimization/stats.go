package hw10programoptimization

import (
	"bufio"
	"errors"
	"io"
	"strings"

	"github.com/valyala/fastjson"
)

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	reader := bufio.NewReader(r)

	result := make(DomainStat)

	for {
		lineBytes, _, err := reader.ReadLine()

		if errors.Is(err, io.EOF) {
			return result, nil
		}

		if err != nil {
			return result, err
		}

		userEmail := fastjson.GetString(lineBytes, "Email")

		hasRequiredDomain := strings.HasSuffix(userEmail, "."+domain)
		if !hasRequiredDomain {
			continue
		}

		separatedByEt := strings.SplitN(userEmail, "@", 2)
		secondLevelDomain := strings.ToLower(separatedByEt[1])

		num := result[secondLevelDomain]
		num++
		result[secondLevelDomain] = num
	}
}
