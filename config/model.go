package main

import "strings"
import "fmt"
import "strconv"
import "os"
import "log"

const binfname = "./data.bin"

type Model struct {
	prog string
	args []string
	opts Serialisable
}

func write_binary_file(buffer *[]byte, fname string) {
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// hmm, look more closely at this 'defer'
	_, err = f.Write([]byte(*buffer))
	if err != nil {
		log.Fatal(err)
	}
}

//func read_binary_file(buffer *[]byte, fname string) {
//	//
//}

//func zero_buffer(buffer *[]byte) {
//	(*buffer)[0] = 0
//}

func output_buffer(buffer *[]byte) {
	for _, value := range *buffer {
		fmt.Printf("%d ", value)
	}
	fmt.Println()
}

func (re Model) begin() {
	// if config file doesn't exist, create it
	var size uint16 = 0
	{
		s := Sizer{&size}
		re.opts.serialise(s)
	}
	fmt.Println("Size: " + strconv.Itoa(int(size)))
	{
		buffer := make([]byte, size)
		re.opts.serialise(Saver{&buffer, 0})
		output_buffer(&buffer)
		write_binary_file(&buffer, binfname)
		//zero_buffer(&buffer)
		//read_binary_file(&buffer, binfname)
		re.opts.serialise(Loader{&buffer, 0})
	}
}

func (re Model) end() {
}

func (re Model) getProgramName() string {
	p := strings.Split(re.prog, "/")
	return LastItemOfStringArray(p)
}
