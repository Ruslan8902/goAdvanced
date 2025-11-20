package verify

import (
	"fmt"
	"os"
)

type Storage struct {
	path string
}

func NewStorage(path string) *Storage {
	return &Storage{
		path: path,
	}
}

func (s *Storage) WriteStorage(content []byte) {
	file, err := os.Create(s.path)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись прошла успешно!")
}

func (s *Storage) ReadStorage() ([]byte, error) {
	if IsJson(s.path) {
		file, err := os.ReadFile(s.path)
		if err != nil {
			fmt.Println(err)
			return []byte{}, err
		}
		return file, nil
	}
	return []byte{}, nil
}

func (s *Storage) SaveListJSON(binList *[]byte) {
	if IsJson(s.path) {
		s.WriteStorage(*binList)
	} else {
		fmt.Print("Not JSON storage")
	}
}

// func ToBytes(binList *bins.BinList) ([]byte, error) {
// 	file, err := json.Marshal(binList)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return file, nil
// }
