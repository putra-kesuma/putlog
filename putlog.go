package putlog

import (
	"log"
	"os"
)

// CreateLog is open or create file for save log history
// create a new file if does't exist
// open file if exist
func CreateLog(nameFile string) error {

	f, err := os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(f)
	return nil
}
