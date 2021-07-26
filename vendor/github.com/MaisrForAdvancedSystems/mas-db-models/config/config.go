package config

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"github.com/spf13/viper"
)
const MasDbConnName="mas_db_connection"
const MaxDbConnections="mas_max_db_connections"
const MaxDbIdleConnections="mas_idle_db_connections"

const key = ("a very very very very secret key") // 32 bytes
const filename = "./config.dat"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var HOST = flag.String("host", "", "host")
var PORT = flag.Int("port", 25500, "port")
var DEBUG = flag.String("debug", "false", "debug")

func init() {
	flag.Parse()
	viper.BindEnv(MasDbConnName)
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json") // REQUIRED if the config file does not have the extension in the name
	viper.SetDefault(MaxDbConnections,100)
	viper.SetDefault(MaxDbIdleConnections,10)
	cname:=viper.GetString(MasDbConnName)
	if DEBUG != nil && (*DEBUG == "true" || *DEBUG == "1") {
		viper.Set("debug",true)
		log.Printf("runing in debug mode %v", *DEBUG)
	}
	if cname!=""{
		return;
	}
	isFileExsists := fileExists(filename)
	var cnsStr string
	if !isFileExsists {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter Database Server Host Or Ip: ")
		dbServer, _ := reader.ReadString('\n')
		fmt.Print("Enter Database Name: ")
		dbName, _ := reader.ReadString('\n')
		fmt.Print("Enter sa user password: ")
		saPw, _ := reader.ReadString('\n')
		if dbServer!=""{
			cnsStr = fmt.Sprintf("server=%s;database=%s;user id=sa;password=%s;", dbServer, dbName, saPw)
			viper.SetDefault(MasDbConnName,cnsStr)
			encStr, err := encrypt([]byte(key), []byte(cnsStr))
			check(err)
			f, err := os.Create(filename)
			check(err)
			defer f.Close()
			f.Write(encStr)
		}
	} else {
		cbytes, err := ioutil.ReadFile(filename)
		check(err)
		if len(cbytes) == 0 {
			os.Remove(filename)
		}
		cnsStrytes, err := decrypt([]byte(key), cbytes)
		check(err)
		cnsStr = string(cnsStrytes)
		viper.SetDefault(MasDbConnName,cnsStr)
	}
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func encrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func decrypt(key, text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}
