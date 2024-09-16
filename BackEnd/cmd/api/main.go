package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
	"github.com/trigo-William-HOANG/ProjetGOVUEJS/internal/handlers"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)
	fmt.Println("Starting GO API Service")
	fmt.Println(`
 ______     ______        ______     ______   __     
/\  ___\   /\  __ \      /\  __ \   /\  == \ /\ \   
\ \ \__ \  \ \ \/\ \     \ \  __ \  \ \  _-/ \ \ \  
 \ \_____\  \ \_____\     \ \_\ \_\  \ \_\    \ \_\ 
  \/_____/   \/_____/      \/_/\/_/   \/_/     \/_/ 
  ______   ______     __     ______     ______       
 /\__  _\ /\  == \   /\ \   /\  ___\   /\  __ \      
 \/_/\ \/ \ \  __<   \ \ \  \ \ \__ \  \ \ \/\ \     
    \ \_\  \ \_\ \_\  \ \_\  \ \_____\  \ \_____\    
     \/_/   \/_/ /_/   \/_/   \/_____/   \/_____/    
                                                     `)
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
