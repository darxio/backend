package analyzer

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"backend/internal/database/connection"

	"github.com/eadium/contents-analyzer/brackets"
	"github.com/jackc/pgx"
)

var database *pgx.ConnPool

func init() {
	database = connection.Connect()
}

type Ingredient struct {
	Name        string        `json:"name"`
	Danger      int           `json:"danger"`
	Ingredients *[]Ingredient `json:"ingredients"`
	// Description string        `json:"description"`
	// WikiLink    string        `json:"wiki_link"`
}

type argError struct {
	arg  string
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}

// Analyze analyses product ingredients
func Analyze(s string) (*[]Ingredient, error) {
	bracketsBalanced, _ := brackets.Bracket(s)
	if !bracketsBalanced {
		log.Println("ALERT! Brackets are not balanced!")
		return nil, &argError{"SYNTAX_ERROR", "Brackets are not balanced"}
	}
	letters := strings.Split(s, "")
	ings := make([]Ingredient, 0)
	err := parse(letters, &ings)
	if err != nil {
		return nil, err
	}
	// json, _ := json.Marshal(ings)
	// ioutil.WriteFile("./prod_contents.json", json, 0644)
	return &ings, nil
}

// Parse parses
func parse(letters []string, ings *[]Ingredient) error {
	reg, err := regexp.Compile("[^a-zA-Z0-9А-Яа-я[:space:]]+")
	if err != nil {
		log.Fatal(err)
	}
	bracketsReg, err1 := regexp.Compile("[)]+")
	if err1 != nil {
		log.Fatal(err1)
	}
	separatorReg, err2 := regexp.Compile("[,.:;]+")
	if err2 != nil {
		log.Fatal(err2)
	}
	recursiveReg, err3 := regexp.Compile("[(]+")
	if err3 != nil {
		log.Fatal(err3)
	}
	curWord := ""
	for i := 0; i < len(letters); i++ {
		if separatorReg.MatchString(letters[i]) || i == len(letters)-1 {
			if i == len(letters)-1 || len(curWord) <= 2 {
				curWord += letters[i]
				if reg.MatchString(curWord) == true {
					curWord = ""
					continue
				}
			}
			curWord = strings.TrimSpace(bracketsReg.ReplaceAllString(curWord, ""))
			danger := getDangerLevel(curWord)
			*ings = append(*ings, Ingredient{curWord, danger, nil})
			curWord = ""
			continue
		}

		if recursiveReg.MatchString(letters[i]) {
			closePos := findClosingParen(letters, i+1)
			substring := letters[i+1 : closePos-1]
			subIngs := make([]Ingredient, 0)
			err := parse(substring, &subIngs)
			if err != nil {
				return err
			}
			curWord = strings.TrimSpace(bracketsReg.ReplaceAllString(curWord, ""))
			danger := getDangerLevel(curWord)
			*ings = append(*ings, Ingredient{curWord, danger, &subIngs})
			if i < len(letters)-1 {
				i = closePos - 1
			} else {
				i = closePos + 1
			}
			curWord = ""
		}
		curWord += letters[i]
	}
	return nil
}

func findClosingParen(text []string, openPos int) int {
	closePos := openPos
	counter := 1
	for counter > 0 {
		c := text[closePos]
		closePos++
		if c == "(" {
			counter++
		} else if c == ")" {
			counter--
		}
	}
	return closePos
}

func getDangerLevel(ing string) int {
	var danger int
	err := database.QueryRow(
		`SELECT danger FROM ingredients WHERE name = $1 LIMIT 1
	`, ing).Scan(&danger)
	if err != nil {
		// log.Println("ERROR analyzer.go:132: getDangerLevel()", err.Error())
		return -1
	}
	return danger
}
