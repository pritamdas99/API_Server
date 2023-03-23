package data

type Book struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	GenreIds []int  `json:"genreids"`
}

type Genre struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type bookDB map[int]Book
type genreDB map[int]Genre
type userDB map[string]User

var BookList bookDB
var GenreList genreDB
var UserList userDB

func init() {
	BookList = make(bookDB)
	GenreList = make(genreDB)
	UserList = make(userDB)
	UserInit()
	GenreInit()
	BookInit()
}

func UserInit() {
	UserList["prishan076@gmail.com"] = User{
		ID:       "prishan076@gmail.com",
		Password: "1111",
	}
	UserList["pritam@appscode.com"] = User{
		ID:       "pritam@appscode.com",
		Password: "1111",
	}
	UserList["pd17021999@gmail.com"] = User{
		ID:       "pd17021999@gmail.com",
		Password: "1111",
	}
}

func GenreInit() {
	GenreList[1] = Genre{
		ID:   "1",
		Name: "Thrill",
	}
	GenreList[2] = Genre{
		ID:   "2",
		Name: "Comedy",
	}
	GenreList[3] = Genre{
		ID:   "3",
		Name: "Biography",
	}
}

func BookInit() {
	BookList[1] = Book{
		ID:   "1",
		Name: "Harry Potter",
		GenreIds: []int{
			1, 2,
		},
	}
	BookList[2] = Book{
		ID:   "2",
		Name: "Mr. Bean",
		GenreIds: []int{
			2,
		},
	}
	BookList[3] = Book{
		ID:   "3",
		Name: "Wings of Fire",
		GenreIds: []int{
			3,
		},
	}
}
