package diagnostics
import ( 
				"fmt"
				"net/http"
				"github.com/gorilla/mux")


func NewDiagnostics() *mux.Router{
	router := mux.NewRouter()
	router.HandleFunc("/healtz", healtz)
	router.HandleFunc("/info", ready)
	return router
}

func healtz(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}

func ready(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}