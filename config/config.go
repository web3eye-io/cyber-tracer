package config

import (
	"bytes"
	_ "embed"
	"fmt"
	"os"
	"reflect"

	"github.com/BurntSushi/toml"
)

//go:embed config.toml
var ConfigFS string

type Config struct {
	Version  string   `toml:"version" env:"version"`
	NFTMeta  NFTMeta  `toml:"nft-meta" env:"nft_meta"`
	BlockETL BlockETL `toml:"block-etl" env:"block_etl"`
	ETH      ETH      `toml:"eth" env:"eth"`
	IPFS     IPFS     `toml:"ipfs" env:"ipfs"`
	MySQL    MySQL    `toml:"mysql" env:"mysql"`
}

type NFTMeta struct {
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogDir   string `toml:"log-dir" env:"log_dir"`
}

type BlockETL struct {
	HTTPPort int    `toml:"http-port" env:"http_port"`
	GrpcPort int    `toml:"grpc-port" env:"grpc_port"`
	LogDir   string `toml:"log-dir" env:"log_dir"`
}

type MySQL struct {
	IP       string `toml:"ip" env:"ip"`
	Port     int    `toml:"port" env:"port"`
	User     string `toml:"user" env:"user"`
	Password string `toml:"password" env:"password"`
	Database string `toml:"database" env:"database"`
}

type ETH struct {
	Wallets string `toml:"wallets" env:"wallets"`
}

type IPFS struct {
	HTTPGateway string `toml:"http-gateway" env:"http_gateway"`
}

// set default config
var config = &Config{
	NFTMeta: NFTMeta{
		HTTPPort: 30100,
		GrpcPort: 30101,
		LogDir:   "/var/log",
	},
}

type envMatcher struct {
	envMap map[string]string
}

func DetectEnv(co *Config) (err error) {
	e := &envMatcher{}
	e.envMap = make(map[string]string)
	ct := reflect.TypeOf(co)
	e.detectEnv(ct, "", "")
	_, err = toml.Decode(e.toToml(), co)
	return err
}

func (e *envMatcher) detectEnv(t reflect.Type, preffix, _preffix string) {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	for i := 0; i < t.NumField(); i++ {
		it := t.Field(i)
		envKey := fmt.Sprintf("%v%v", preffix, it.Tag.Get("env"))
		_envKey := fmt.Sprintf("%v%v", _preffix, it.Tag.Get("toml"))
		if it.Type.Kind() != reflect.Struct {
			if envValue, ok := os.LookupEnv(envKey); ok {
				if it.Type.Kind() == reflect.String {
					e.envMap[_envKey] = fmt.Sprintf("\"%v\"", envValue)
				} else {
					e.envMap[_envKey] = envValue
				}
			}
			continue
		}
		envKey = fmt.Sprintf("%v%v_", preffix, it.Tag.Get("env"))
		_envKey = fmt.Sprintf("%v%v.", _preffix, it.Tag.Get("toml"))
		e.detectEnv(it.Type, envKey, _envKey)
	}
}

func (e *envMatcher) toToml() string {
	var b bytes.Buffer

	for v := range e.envMap {
		b.WriteString(fmt.Sprintf("%v=%v\n", v, e.envMap[v]))
	}

	return b.String()
}

func init() {
	md, err := toml.Decode(ConfigFS, config)
	if err != nil {
		panic(fmt.Sprintf("failed to parse config file, %v", err))
	}
	if len(md.Undecoded()) > 0 {
		fmt.Printf("cannot parse [%v] to config\n", md.Undecoded())
	}
	err = DetectEnv(config)
	if err != nil {
		fmt.Printf("environment variable parse failed， %v", err)
	}
}

func GetConfig() *Config {
	return config
}
