package workers

import "os"

func WriteToFile(data []byte) error {
	f, err := os.Create("goconcurrency/xkcd.json")
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	return nil
}
