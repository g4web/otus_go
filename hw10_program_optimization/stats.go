package hw10programoptimization

import (
	"bufio"
	"io"
	"log"
	"strings"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	scanner := bufio.NewScanner(r)

	result := make(DomainStat)

	for scanner.Scan() {
		separatedByFirstLevelDomain := strings.SplitN(scanner.Text(), "."+domain, 2)
		hasRequiredDomain := len(separatedByFirstLevelDomain) == 1
		if hasRequiredDomain {
			continue
		}

		separatedByEt := strings.SplitN(separatedByFirstLevelDomain[0], "@", 2)
		secondLevelDomain := strings.ToLower(separatedByEt[1])
		fullDomain := secondLevelDomain + "." + domain

		num := result[fullDomain]
		num++
		result[fullDomain] = num
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result, nil
}
