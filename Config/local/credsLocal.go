package local
import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var (
	once  sync.Once
	creds *Config
)


type Config struct {
	Server struct {
		Port int 	`yaml:"port"`
		Host string `yaml:"host"`
	}
	Database struct {
		Port int 	`yaml:"port"`
		Host string `yaml:"host"`
	}
}



func Instance() *Config {
	once.Do(func() {
		creds = InitCreds()
	})
	return creds
}

func InitCreds() *Config {

	c := Config{}
	pwd, _ := os.Getwd()
	yamlFile, err := ioutil.ReadFile(pwd+"/Config/local/conf.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return &c
}
