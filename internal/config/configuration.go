package config

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	defCfg      map[string]string
	initialized = false
)

func init() {
	defCfg = make(map[string]string)
}

// LoadConfig loads configuration file
func LoadConfig() {

	log.Info("loading config...")

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	defCfg["app.id"] = "wallet-go-img"
	defCfg["app.version"] = "1.0.0"
	defCfg["app.env"] = "development" // set to production

	defCfg["server.host"] = "localhost"
	defCfg["server.port"] = "7000"
	defCfg["server.log.level"] = "debug" // valid values are trace, debug, info, warn, error, fatal
	defCfg["server.timeout.write"] = "15 seconds"
	defCfg["server.timeout.read"] = "15 seconds"
	defCfg["server.timeout.idle"] = "60 seconds"
	defCfg["server.timeout.graceshut"] = "15 seconds"

	defCfg["server.context.timeout"] = "30" // seconds

	defCfg["db.host"] = "localhost"
	defCfg["db.port"] = "3306"
	defCfg["db.user"] = "wallet_user"
	defCfg["db.password"] = "wallet"
	defCfg["db.name"] = "wallet"

	defCfg["health.local"] = "https://httpbin.org/status/200"
	defCfg["health.delay"] = "1"     // seconds
	defCfg["health.interval"] = "30" // seconds

	defCfg["hmac.secret"] = "th1s?MusT#b3!4*veRY%d33p#53creT"
	defCfg["hmac.age.minute"] = "10"

	for k := range defCfg {
		err := viper.BindEnv(k)
		if err != nil {
			log.Errorf("Failed to bind env \"%s\" into configuration. Got %s", k, err)
		}
	}

	initialized = true
}

// SetConfig put configuration key value
func SetConfig(key, value string) {
	viper.Set(key, value)
}

// Get fetch configuration as string value
func Get(key string) string {
	if !initialized {
		LoadConfig()
	}
	ret := viper.GetString(key)
	if len(ret) == 0 {
		if ret, ok := defCfg[key]; ok {
			return ret
		}
		log.Debugf("%s config key not found", key)
	}
	return ret
}

// GetBoolean fetch configuration as boolean value
func GetBoolean(key string) bool {
	if len(Get(key)) == 0 {
		return false
	}
	b, err := strconv.ParseBool(Get(key))
	if err != nil {
		panic(err)
	}
	return b
}

// GetInt fetch configuration as integer value
func GetInt(key string) int {
	if len(Get(key)) == 0 {
		return 0
	}
	i, err := strconv.ParseInt(Get(key), 10, 64)
	if err != nil {
		panic(err)
	}
	return int(i)
}

// GetFloat fetch configuration as float value
func GetFloat(key string) float64 {
	if len(Get(key)) == 0 {
		return 0
	}
	f, err := strconv.ParseFloat(Get(key), 64)
	if err != nil {
		panic(err)
	}
	return f
}

// Set configuration key value
func Set(key, value string) {
	defCfg[key] = value
}
