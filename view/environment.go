package view

import (
	"fmt"
	"github.com/andreluzz/cas-xog/log"
	"github.com/andreluzz/cas-xog/model"
	"github.com/howeyc/gopass"
	"os"
	"strconv"
	"github.com/andreluzz/cas-xog/constant"
)

var targetEnvInput string

func Environments(action string, environments *model.Environments) bool {
	if action == "m" {
		return true
	}

	sourceInput := "-1"
	targetInput := "-1"

	username, password := "", ""

	log.Info("\n")
	log.Info("Available environments:\n")

	if environments.Available == nil || len(environments.Available) == 0 {
		log.Info("\n[CAS-XOG][red[ERROR]] - None available environments found!\n")
		log.Info("\n[CAS-XOG][red[FATAL]] - Check your xogEnv.xml file. Press enter key to exit...")
		scanExit := ""
		fmt.Scanln(&scanExit)
		os.Exit(0)
	}

	for i, e := range environments.Available {
		log.Info("%d - %s\n", i+1, e.SelectAttrValue("name", "Unknown environment name"))
	}

	if action == "r" {
		log.Info("Choose reading environment [1]: ")
		sourceInput = "1"
		fmt.Scanln(&sourceInput)

		var err error
		var noLogin bool
		envIndex, err := strconv.Atoi(sourceInput)

		if err != nil || envIndex <= 0 || envIndex > len(environments.Available) {
			log.Info("\n[CAS-XOG][red[ERROR]] - Invalid reading environment index!\n")
			return false
		}

		log.Info("[CAS-XOG]Processing environment login")
		err, noLogin = environments.InitSource(sourceInput, username, password)
		if err != nil {
			log.Info("\n[CAS-XOG][red[ERROR]] - %s", err.Error())
			log.Info("\n[CAS-XOG][red[FATAL]] - Check your xogEnv.xml file. Press enter key to exit...")
			scanExit := ""
			fmt.Scanln(&scanExit)
			os.Exit(0)
		}
		if noLogin {
			requestLogin(sourceInput, username, password, constant.INPUT_SOURCE)
		}
		log.Info("\r[CAS-XOG][green[Login successfully]] - Environment: %s \n", environments.Source.Name)
	}

	log.Info("Choose writing environment [1]: ")
	targetInput = "1"
	fmt.Scanln(&targetInput)

	envIndex, err := strconv.Atoi(targetInput)

	if action == "r" {
		targetEnvInput = targetInput
	} else if action == "w" && targetEnvInput != targetInput {
		log.Info("\n[CAS-XOG][yellow[Warning]]: Trying to write files read from a different target environment!")
		log.Info("\n[CAS-XOG]Do you want to continue anyway? (y = Yes, n = No) [n]: ")
		input := "n"
		fmt.Scanln(&input)
		if input != "y" {
			return false
		}
	}

	if err != nil || envIndex <= 0 || envIndex > len(environments.Available) {
		log.Info("\n[CAS-XOG][red[ERROR]] - Invalid writing environment index!\n")
		return false
	}

	if sourceInput == targetInput {
		environments.CopyTargetFromSource()
	} else {
		var noLogin bool
		username, password = "", ""
		log.Info("[CAS-XOG]Processing environment login")
		err, noLogin = environments.InitTarget(targetInput, username, password)
		if err != nil {
			log.Info("\n[CAS-XOG][red[ERROR]] - %s", err.Error())
			log.Info("\n[CAS-XOG][red[FATAL]] - Check your xogEnv.xml file. Press enter key to exit...")
			scanExit := ""
			fmt.Scanln(&scanExit)
			os.Exit(0)
		}
		if noLogin {
			requestLogin(targetInput, username, password, constant.INPUT_TARGET)
		}
		log.Info("\r[CAS-XOG]Environment: %s - [green[Login successfully]]\n", environments.Target.Name)
	}

	return true
}

func requestLogin(envIndex string, username string, password string, inputType string)  {
	log.Info("\r[CAS-XOG][yellow[Request Login]] - Environment: %s \n", environments.Source.Name)
	log.Info("Username: ")
	fmt.Scanln(&username)

	log.Info("Password: ")
	passwordTemp, _ := gopass.GetPasswd()
	password = string(passwordTemp[:])

	log.Info("[CAS-XOG]Processing environment login")

	var err error

	if inputType == constant.INPUT_SOURCE {
		err, _ = environments.InitSource(envIndex, username, password)
	} else {
		err, _ = environments.InitTarget(envIndex, username, password)
	}

	if err != nil {
		log.Info("\n[CAS-XOG][red[ERROR]] - %s", err.Error())
		log.Info("\n[CAS-XOG][red[FATAL]] - Check your xogEnv.xml file. Press enter key to exit...")
		scanExit := ""
		fmt.Scanln(&scanExit)
		os.Exit(0)
	}
}