package format

import (
	"cmp"
	"encoding/json"
	"fmt"
	"github.com/VitalySR/task13_mod/internal/entity"
	"log"
	"os"
	"slices"
)

func Do(filePathRead string, filePathWrite string) error {
	var ps = make([]entity.Patient, 0, 3)

	err := readFile(filePathRead, &ps)

	if err != nil {
		return err
	}

	slices.SortFunc(ps, func(a, b entity.Patient) int {
		return cmp.Compare(a.Age, b.Age)
	})

	err = writeFile(filePathWrite, &ps)

	if err != nil {
		return err
	}

	return nil
}

func readFile(filePath string, ps *[]entity.Patient) error {
	fileRead, err := os.Open(filePath)

	if err != nil {
		return fmt.Errorf("open file %s error: %w", filePath, err)
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("close file %s error: %v\n", filePath, err)
		}
	}(fileRead)

	dec := json.NewDecoder(fileRead)

	for dec.More() {
		var p entity.Patient
		err := dec.Decode(&p)
		if err != nil {
			return fmt.Errorf("decode file %s error: %w", filePath, err)
		}
		*ps = append(*ps, p)
	}

	log.Printf("Read from file: %+v\n", ps)

	return nil
}

func writeFile(filePath string, ps *[]entity.Patient) error {
	fileWrite, err := os.Create(filePath)
	defer func(fileWrite *os.File) {
		err := fileWrite.Close()
		if err != nil {
			log.Fatalf("close file %s error: %v\n", filePath, err)
		}
	}(fileWrite)

	if err != nil {
		return fmt.Errorf("create file %s error: %w", filePath, err)
	}

	err = json.NewEncoder(fileWrite).Encode(ps)
	if err != nil {
		return fmt.Errorf("encode file %s error: %w", filePath, err)
	}

	log.Printf("End of write to file %s\n", filePath)

	return nil
}
