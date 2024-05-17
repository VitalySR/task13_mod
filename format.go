package format

import (
	"encoding/json"
	"fmt"
	"github.com/VitalySR/task13_mod/internal/entity"
	"log"
	"os"
)

func Do(filePathRead string, filePathWrite string) error {
	fileRead, err := os.Open(filePathRead)

	if err != nil {
		return fmt.Errorf("open file %s error: %w", filePathRead, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file %s error: %v\n", filePathRead, err)
		}
	}(fileRead)

	var ps = make([]entity.Patient, 0, 4)

	dec := json.NewDecoder(fileRead)

	for dec.More() {
		var p entity.Patient
		err := dec.Decode(&p)
		if err != nil {
			return fmt.Errorf("decode file %s error: %w", filePathRead, err)
		}
		ps = append(ps, p)
	}

	log.Printf("%+v\n", ps)

	fileWrite, err := os.Create(filePathWrite)
	defer func(fileWrite *os.File) {
		err := fileWrite.Close()
		if err != nil {
			log.Fatalf("close file %s error: %v\n", filePathWrite, err)
		}
	}(fileWrite)

	if err != nil {
		return fmt.Errorf("create file %s error: %w", filePathWrite, err)
	}

	err = json.NewEncoder(fileWrite).Encode(ps)
	if err != nil {
		return fmt.Errorf("encode file %s error: %w", filePathWrite, err)
	}

	return nil
}
