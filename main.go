package main

import (
	"fmt"
	"strings"
)

// this variable become packaged scope
var score = 99.5

func main() {

	// when declared here the PackageScope.go cannot access here because this score belong to main()
	// var score = 99.5

	//sayHello("mario")
	//for _, value := range points {
	//	fmt.Println(value)
	//}
	//
	//showScore()
	//
	//Maps()
	//
	//passByValue()
	//
	//PointersLearn()
	//
	//// Structs (Entity)
	//mybill := newBill("marios's bill")
	//
	//mybill.updateTip(10)
	//mybill.addItem("coffee", 4.99)
	//mybill.addItem("pudding", 10)
	//
	//fmt.Println(mybill)
	//
	//billObject := bill{
	//	name:  "name",
	//	items: map[string]float64{},
	//	tip:   0,
	//}
	//
	//fmt.Println(billObject)
	//
	//fmt.Println(mybill.format())
	//
	//// ########################  Bill part - start ########################
	//billing := bills.CreateBill()
	//
	//bills.PromtOptions(billing)
	//
	//fmt.Println(billing)
	//
	//// ######################## Bill part end - End ########################
	//
	//// ######################## Interface - Start ########################
	//shapes := []shaper{
	//	square{length: 15.2},
	//	circle{radius: 7.5},
	//	circle{radius: 12.3},
	//	square{length: 4.9},
	//}
	//
	//for _, value := range shapes {
	//	printShapeInfo(value)
	//	fmt.Println("---")
	//}
	//// ######################## Interface - End ########################
	//
	////######################## Filepath - Start ########################
	//err := filepath.Walk(".",
	//	func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(path, info.Size(), info.IsDir(), info.Sys())
	//		return nil
	//	})
	//if err != nil {
	//	log.Println(err)
	//}

	//type FileInfo interface {
	//	Name() string       // base name of the file
	//	Size() int64        // length in bytes for regular files; system-dependent for others
	//	Mode() FileMode     // file mode bits
	//	ModTime() time.Time // modification time
	//	IsDir() bool        // abbreviation for Mode().IsDir()
	//	Sys() any           // underlying data source (can return nil)
	//}

	//######################## Filepath - End ########################

	sendSbomURL := strings.SplitN("PostSBOM=http://localhost:3000/sbom", "PostSBOM=", -1)
	fmt.Println(sendSbomURL[1])
}
