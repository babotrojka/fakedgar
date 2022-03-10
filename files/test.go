package main
import(
	"fmt"
	"os/exec"
	"io/ioutil"
	"bytes"
	"os"
)

func check(err error){
	if err != nil {
		panic(err)
	}
}
func compile(compiler string, filename string) ([]byte, error) {
	cmd := exec.Command(compiler, filename, "-o", filename[:(len(filename) - 2)])
	stderr, err := cmd.StderrPipe()
	check(err)
	err = cmd.Start()
	check(err)	
	slurp, _ := ioutil.ReadAll(stderr)
	err = cmd.Wait()
	if err != nil {
		fmt.Printf("%s\n", err)
		return slurp, err
	}
	return slurp, nil
}

func run(filename string) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("./" + filename[:(len(filename) - 2)])
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return out.Bytes(), nil
}

func main() {

	byt, err := compile("gcc", "file.c")
	if len(byt) > 0 {
		fmt.Printf("%s\n", byt)
	}
	
	if err != nil {
		os.Exit(0)
	}
	byt, err = run("file.c")
	if err != nil {
		os.Exit(0)
	}
	fmt.Printf("%s", byt)
		
}