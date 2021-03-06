package main

import ("log" 
				"fmt"
				"net/http"
				"github.com/gorilla/mux"
				"github.com/AKarakehayova/go-workshop/internal/diagnostics"
)

func main(){
	log.Print("Hello, World")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func(){
	err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}
	}()
	diagnostics := diagnostics.NewDiagnostics()
err := http.ListenAndServe(":8585", diagnostics)
	if err != nil {
		log.Fatal(err)
	}

}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}