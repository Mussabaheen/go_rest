package articles

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var Articles []Article

func HandleRequests(httpRouter *mux.Router) {
	Articles = []Article{
		{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	httpRouter.HandleFunc("/", getHomePage).Methods("GET")
	httpRouter.HandleFunc("/articles", getAllArticles).Methods("GET")
	httpRouter.HandleFunc("/article/{id}", getSingleArticle).Methods("GET")
	httpRouter.HandleFunc("/article/{id}", deleteSingleArticle).Methods("DELETE")
	httpRouter.HandleFunc("/article", createNewArticle).Methods("POST")
}

// getHomePage will display string on homepage
func getHomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

// getAllArticles will retrive all the articles
func getAllArticles(w http.ResponseWriter, r *http.Request) {
	if len(Articles) == 0 {
		http.Error(w, "No Articles", 404)
	}
	json.NewEncoder(w).Encode(Articles)
}

// getSingleArticle will retrieve single article using the path param {id}
func getSingleArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]

	for _, article := range Articles {
		if article.Id == articleID {
			json.NewEncoder(w).Encode(article)
			return
		}
	}
	http.Error(w, "no such article exist", 404)
}

// createNewArticle will append a new article using the article in body
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "unable to parse data into article", 400)
	}
	var article Article
	err = json.Unmarshal(reqBody, &article)
	if err != nil {
		http.Error(w, "unable to parse data into article", 400)
	}
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(Articles)
}

//deleteSingleArticle will delete a single article using the path param {id}
func deleteSingleArticle(w http.ResponseWriter, r *http.Request) {
	articleID := mux.Vars(r)["id"]

	for index, article := range Articles {
		if article.Id == articleID {
			Articles = append(Articles[:index], Articles[index+1:]...)
			json.NewEncoder(w).Encode(Articles)
			return
		}
	}
	http.Error(w, "no such article exist", 404)
}
