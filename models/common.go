package models

import (
	"math/rand"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

var Configs map[string]config.ConfigContainer
var base62Indices map[byte]int64
var pow62s []int64

const base62Candidates string = "1RXuxsb8Aqadlv2nTjUyQ9fcrhz4GPg3LN6poeK5VImJCFYt7BwHZSWMDOk0Ei"
const maxInt64 int64 = 839299365868340223 // (62 ^ 10) - 1
const maxStr62 string = "iiiiiiiiii"      // correspond to (62 ^ 10) - 1

func init() {
	for i := int64(62); i <= maxInt64+1; i *= 62 {
		pow62s = append(pow62s, i)
	}

	base62Indices = make(map[byte]int64)
	for i := len(base62Candidates) - 1; i >= 0; i-- {
		base62Indices[byte(base62Candidates[i])] = int64(i)
	}

	Configs = make(map[string]config.ConfigContainer)
	if cf, err := config.NewConfig("ini", "conf/db.conf"); err == nil {
		Configs["conf/db.conf"] = cf
	} else {
		beego.Critical("unable to open conf/db.conf")
	}
}

func GenerateRandomStr62(length int) string {
	if length > len(maxStr62) || length <= 0 {
		return ""
	}

	i := rand.Int63n(pow62s[length-1])
	ret := Int64ToStr62(i)

	for len(ret) < length {
		ret = base62Candidates[0:1] + ret
	}

	return ret
}

func Int64ToStr62(n int64) string {
	if n < 62 {
		return base62Candidates[n : n+1]
	}

	ret := []byte{}
	for n > 0 {
		ret = append(ret, byte(base62Candidates[n%62]))
		n /= 62
	}

	for i, j := 0, len(ret)-1; i < j; i, j = i+1, j-1 {
		ret[i], ret[j] = ret[j], ret[i]
	}

	return string(ret)
}

func Str62ToInt64(s string) int64 {
	var ret int64 = 0

	for i := 0; i < len(s); i++ {
		ret *= 62
		ret += base62Indices[byte(s[i])]
	}

	return ret
}
