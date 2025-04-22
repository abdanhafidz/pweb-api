package utils

import (
	"log"
	"os"

	"pweb-api.abdanhafidz.com/config"
)

func LogError(errorLogged error) {
	log.Println("Error Log :", errorLogged)

	_, err := os.Stat(config.LOG_PATH + "/error_log.txt")
	if os.IsNotExist(err) {
		_, err = os.Create(config.LOG_PATH + "/error_log.txt")
		if err != nil {
			log.Fatalf("Gagal buka file log: %v", err)
		}
	}

	file, err := os.OpenFile(config.LOG_PATH+"/error_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatalf("Gagal buka file log: %v", err)
	}

	log.SetOutput(file)
}
