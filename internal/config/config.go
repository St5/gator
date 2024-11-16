package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
} 

func(c *Config) SetUser(username string) error{

	c.CurrentUserName = username
	return write(*c)
}

func write(c Config) error{

	path, err := getFulPath()
	if err != nil {
		return nil
	}

	file,err := os.Create(path)

	if err != nil{
		return err		
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(c)
	if err != nil {
		return err
	}

	return nil
}

func getFulPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(home, CONFFILE)
	return path, nil
}

const CONFFILE = ".gatorconfig.json"

func Read() (Config, error) {
	cnfg := Config{}

	path, err := getFulPath()
	if err != nil {
		return cnfg, nil
	}

	file,err := os.Open(path)
	if(err != nil){
		return cnfg, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cnfg)

	if err != nil {
		return cnfg, err
	}

	return cnfg, nil
	
}

