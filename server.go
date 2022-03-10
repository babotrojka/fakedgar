package main
import(
	"net/http"
	"html/template"
	"fmt"
	"io/ioutil"
	// "os"
	"os/exec"
	"bytes"
)

// var templates *template.Template

func provide(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return;
	}
	switch r.Method {
	case "GET" :
		t, _ := template.ParseFiles("templates/main.html")
		err := t.ExecuteTemplate(w, "main.html", nil)
		if err != nil {
			panic(err)
		}
	default:
		fmt.Fprintf(w, "Only GET and POST are supported")
	
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	var response []byte
	writeToFile(r.FormValue("Input"))
	byt, err := compile("gcc", "file.c")
	if len(byt) > 0  {
		// fmt.Printf("%s", byt)
		response = byt
	}
	
	if err != nil {
		fmt.Printf("%s\n", err)
		w.Write([]byte(response))
		return
	}
	byt, err = run("file")
	// fmt.Printf("%s", byt)
	
	
	if err != nil {
		response = append(response, []byte(err.Error())...)
		w.Write([]byte(response))
		return
	}
	response = append(response, byt...)
	// fmt.Println("Dobio sam request")
	w.Write([]byte(response))
	
}

func writeToFile(input string) {
	ioutil.WriteFile("file.c", []byte(input), 0644)
}
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
		return slurp, err
	}
	return slurp, nil
}

func run(filename string) ([]byte, error) {
	var out bytes.Buffer
	cmd := exec.Command("./" + filename)
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}


func main() {
	// templates = template.Must(template.ParseFiles("templates/main.html"))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))	
	http.HandleFunc("/", provide)
	http.HandleFunc("/calculate/", calculate)
	http.ListenAndServe(":8888", nil)
	

}

