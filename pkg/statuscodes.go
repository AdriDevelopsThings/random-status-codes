package pkg

import (
	"math/rand"
	"time"
)

func getRandomStatusCode() int {
	statusCodes := [53]int{
		200,
		201,
		202,
		203,
		204,
		205,
		208,
		226,
		300,
		301,
		302,
		303,
		304,
		305,
		306,
		307,
		308,
		400,
		401,
		402,
		403,
		404,
		405,
		406,
		407,
		408,
		409,
		410,
		411,
		412,
		413,
		414,
		415,
		416,
		417,
		418,
		421,
		426,
		428,
		429,
		431,
		451,
		500,
		501,
		502,
		503,
		504,
		505,
		506,
		507,
		508,
		510,
		511,
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return statusCodes[r.Int()%53]
}
