package apis

import (
	"encoding/json"
	"fmt"
	"github.com/PritamDas17021999/API-server/auth"
	"github.com/PritamDas17021999/API-server/data"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"strconv"
	"time"
)

var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

var router = mux.NewRouter()

func Homepage(w http.ResponseWriter, r *http.Request) {
	//	pingCounter.Inc()
	fmt.Fprintf(w, "Hello World")
}

func Ping(w http.ResponseWriter, r *http.Request) {
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	avail := 1
	for idx, _ := range data.BookList {
		if idx > avail {
			break
		}
		avail++
	}
	var Newbook data.Book
	err := json.NewDecoder(r.Body).Decode(&Newbook)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.BookList[avail] = Newbook
	err = json.NewEncoder(w).Encode(data.BookList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	avail := 1
	for idx, _ := range data.GenreList {
		if idx > avail {
			break
		}
		avail++
	}
	var NewGenre data.Genre
	err := json.NewDecoder(r.Body).Decode(&NewGenre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.GenreList[avail] = NewGenre
	err = json.NewEncoder(w).Encode(data.GenreList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Bookids := params["id"]
	Bookid, _ := strconv.Atoi(Bookids)
	Book := data.BookList[Bookid]
	err := json.NewDecoder(r.Body).Decode(&Book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.BookList[Bookid] = Book
	err = json.NewEncoder(w).Encode(Book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func UpdateGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Genreids := params["id"]
	Genreid, _ := strconv.Atoi(Genreids)
	Genre := data.GenreList[Genreid]
	err := json.NewDecoder(r.Body).Decode(&Genre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.GenreList[Genreid] = Genre
	err = json.NewEncoder(w).Encode(Genre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Bookids := params["id"]
	Bookid, _ := strconv.Atoi(Bookids)
	delete(data.BookList, Bookid)
	err := json.NewEncoder(w).Encode(data.BookList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Genreids := params["id"]
	Genreid, _ := strconv.Atoi(Genreids)
	var lst []int
	for idx, val := range data.BookList {
		genres := val.GenreIds
		for _, j := range genres {
			if j == Genreid {
				lst = append(lst, idx)
			}

		}
	}
	delete(data.GenreList, Genreid)
	fmt.Println(lst)
	fmt.Println(data.BookList)
	for _, i := range lst {
		_, ok := data.BookList[i]
		if !ok {
			continue
		}
		delete(data.BookList, i)
	}
	err := json.NewEncoder(w).Encode(data.GenreList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Bookids := params["id"]
	Bookid, _ := strconv.Atoi(Bookids)
	Book := data.BookList[Bookid]
	err := json.NewEncoder(w).Encode(Book)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	Genreids := params["id"]
	Genreid, _ := strconv.Atoi(Genreids)
	Genre := data.GenreList[Genreid]
	err := json.NewEncoder(w).Encode(Genre)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	User := data.UserList[id]
	err := json.NewEncoder(w).Encode(User)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func GetAllBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.BookList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllGenre(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.GenreList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewUser data.User
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	fmt.Println(NewUser)
	data.UserList[NewUser.ID] = NewUser
	token, err := auth.GenerateJWT()
	if err != nil {
		log.Fatal(err)
	}
	cookie := http.Cookie{
		Name:  "Token",
		Value: token,
		Path:  "/",
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewUser)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewUser data.User
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	val, ok := data.UserList[NewUser.ID]
	fmt.Println(val)
	if !ok {
		http.Error(w, "user does not exist", 400)
		return
	}
	if val.Password != NewUser.Password {
		http.Error(w, "wrong username or password", 400)
		return
	}
	fmt.Println(NewUser)
	data.UserList[NewUser.ID] = NewUser
	token, err := auth.GenerateJWT()
	if err != nil {
		log.Fatal(err)
	}
	cookie := http.Cookie{
		Name:    "Token",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Add(15 * time.Minute),
	}
	http.SetCookie(w, &cookie)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(NewUser)

}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "Token",
		Expires: time.Now(),
	})
	w.WriteHeader(http.StatusOK)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var NewUser data.User
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	NewUser = data.UserList[id]
	err := json.NewDecoder(r.Body).Decode(&NewUser)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	data.UserList[NewUser.ID] = NewUser
	json.NewEncoder(w).Encode(NewUser)
	w.WriteHeader(http.StatusOK)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	fmt.Println(id)
	delete(data.UserList, id)
	fmt.Println(data.UserList)
	err := json.NewEncoder(w).Encode(data.UserList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data.UserList)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func callFunc() {
	prometheus.MustRegister(pingCounter)
	router.HandleFunc("/", Homepage)
	router.HandleFunc("/ping", Ping)
	router.Handle("/metrics", promhttp.Handler())
	router.HandleFunc("/addbook", auth.IsAuthenticated(AddBook)).Methods("POST")
	router.HandleFunc("/addgenre", auth.IsAuthenticated(AddGenre)).Methods("POST")
	router.HandleFunc("/adduser", AddUser).Methods("POST")
	router.HandleFunc("/loginuser", LoginUser).Methods("POST")
	router.HandleFunc("/logout", Logout).Methods("GET")
	router.HandleFunc("/updatebook/{id}", auth.IsAuthenticated(UpdateBook)).Methods("PUT")
	router.HandleFunc("/updategenre/{id}", auth.IsAuthenticated(UpdateGenre)).Methods("PUT")
	router.HandleFunc("/updateuser/{id}", auth.IsAuthenticated(UpdateUser)).Methods("PUT")
	router.HandleFunc("/deletebook/{id}", auth.IsAuthenticated(DeleteBook)).Methods("DELETE")
	router.HandleFunc("/deletegenre/{id}", auth.IsAuthenticated(DeleteGenre)).Methods("DELETE")
	router.HandleFunc("/deleteuser/{id}", auth.IsAuthenticated(DeleteUser)).Methods("DELETE")
	router.HandleFunc("/getbook/{id}", auth.IsAuthenticated(GetBook)).Methods("GET")
	router.HandleFunc("/getgenre/{id}", auth.IsAuthenticated(GetGenre)).Methods("GET")
	router.HandleFunc("/getallbook", auth.IsAuthenticated(GetAllBook)).Methods("GET")
	router.HandleFunc("/getallgenre", auth.IsAuthenticated(GetAllGenre)).Methods("GET")
	router.HandleFunc("/getalluser", GetAllUser).Methods("GET")
}

func StartServer(port int) {
	log.Printf("-------------starting sever at %d -------\n", port)
	callFunc()
	Server := &http.Server{
		Addr:    ":" + strconv.Itoa(port),
		Handler: router,
	}

	fmt.Println(Server.ListenAndServe())

}

// curl -X POST -d '{"id":"www.gmail.com","password":"1111"}' http://localhost:8080/loginuser
