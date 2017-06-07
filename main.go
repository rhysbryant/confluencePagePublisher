package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/howeyc/gopass"
	"github.com/rhysbryant/go-confluence"
)

type Config struct {
	TempplateName string                `json:"tempplateName"`
	User          string                `json:"user"`
	Pass          string                `json:"password"`
	Url           string                `json:"url"`
	SpaceKey      string                `json:"spaceKey"`
	PageTitle     string                `json:"pageTitle"`
	ParentPages   []confluence.Ancestor `json:"parentPages"`
	Labels        []confluence.Label    `json:"labels"`
}

const (
	FieldSeperator     = ";"
	DefaultLabelPrefix = "Global"
	LabelSeperator     = " "
)

func getLineFromStdIn(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	return text
}

func getPasswordFromStdIn() string {
	fmt.Printf("Password: ")
	pass, err := gopass.GetPasswdMasked()
	if err != nil {
		return ""
	}

	return string(pass)
}

func loadConfig(fileName string) (*Config, error) {
	config := Config{}
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	e := json.Unmarshal(file, &config)
	if e != nil {
		return nil, e
	}
	return &config, nil
}

func loadFile(file string) (string, error) {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func main() {
	var fieldList, configPath, labels string

	flag.StringVar(&configPath, "config", "", "the config path")
	flag.StringVar(&fieldList, "fieldList", "", "macro name"+FieldSeperator+"replacement")
	flag.StringVar(&labels, "labels", "", "optional extra labels space seperated - added with global prefix")
	flag.Parse()

	replacements := strings.Split(fieldList, FieldSeperator)

	if (len(replacements) % 2) != 0 {
		log.Fatal("field name value missmatch")
	}

	if configPath == "" {
		log.Fatalln("config not found")
	}

	cfg, err := loadConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}

	if cfg.User == "" {
		cfg.User = getLineFromStdIn("User Name:")
	}

	if cfg.Pass == "" {
		cfg.Pass = getPasswordFromStdIn()
	}

	tempplate, err := loadFile(cfg.TempplateName)
	if err != nil {
		log.Fatalln(err)
	}

	userMacroReplacer := strings.NewReplacer(replacements...)

	html := userMacroReplacer.Replace(tempplate)
	fmt.Println(html)

	wiki, err := confluence.NewWiki(cfg.Url, confluence.BasicAuth(cfg.User, cfg.Pass))
	if err != nil {
		log.Fatal(err)
	}

	c := confluence.Content{Title: userMacroReplacer.Replace(cfg.PageTitle)}
	c.Body.Storage.Value = html
	c.Type = "page"
	c.Space.Key = cfg.SpaceKey
	c.Body.Storage.Representation = "storage"
	c.Meta.Labels = cfg.Labels
	c.Ancestors = cfg.ParentPages

	if labels != "" {

		prefix := DefaultLabelPrefix
		l := strings.Split(labels, LabelSeperator)

		for _, item := range l {
			labelName := userMacroReplacer.Replace(item)

			c.Meta.Labels = append(c.Meta.Labels, confluence.Label{Prefix: &prefix, Name: &labelName})
		}
	}

	newContent, err := wiki.AddContent(&c)
	if err != nil {
		log.Fatalf("error:%s", err)
	}

	fmt.Printf("content:%v+", newContent)

}
