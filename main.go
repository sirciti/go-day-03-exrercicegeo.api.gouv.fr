package main

import "log"

var (
	articles   []string
	courseList = make(map[string]int)
)

func main() {
	displayCourse()
	addArticle("pomme")
	addArticle("pomme")
	delArticle("pomme")
	addArticle("poire")
	displayCourse()
}

func delArticle(article string) {
	for i, art := range articles {
		if art == article {
			//articles[:index] : prend tous les éléments avant l'index que l'on veut supprimer
			//articles[index+1:] : prend tous les éléments après l'index que l'on veut supprimer.
			articles = append(articles[:i], articles[(i+1):]...)
		}
	}
	delete(courseList, article)
}

func addArticle(article string) {
	found := false
	for _, i := range articles {
		if i == article {
			found = true
			break
		}
	}
	if !found {
		articles = append(articles, article)
	}
	courseList[article] = courseList[article] + 1
}

func displayCourse() {
	log.Printf("%#v", articles)
	log.Printf("%#v", courseList)
}
