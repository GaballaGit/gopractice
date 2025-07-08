package intermediate 

import(
	"log"
	"os"
)

func main() {

	log.Println("This message has been logged")

	log.SetPrefix("LOG: ")
	log.Println("Help, I am unda tha wata")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("Please note this activity.")

	infoLog.Println("Counting, neurons activated: 2")
	errorLog.Println("No brain activity")

	file, err := os.OpenFile("app.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("There was an error opening log: %v\n", err)
	}
	defer file.Close()

	debugLog := log.New(file, "Debug: ", log.Ldate | log.Ltime | log.Llongfile)

	debugLog.Println("New activity in the cells")

}

var (
	infoLog = log.New(os.Stdout, "Info: ", log.Ldate | log.Ltime | log.Lshortfile)
	errorLog = log.New(os.Stdout, "ERROR: ", log.Ldate | log.Ltime | log.Llongfile)
	warnLog = log.New(os.Stdout, "Warning: ", log.Ldate | log.Ltime | log.Llongfile)
)
