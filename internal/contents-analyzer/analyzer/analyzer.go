package analyzer

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/lib/pq"

	"backend/internal/database/connection"

	"github.com/eadium/contents-analyzer/brackets"
	"github.com/jackc/pgx"
)

var database *pgx.ConnPool
var reg *regexp.Regexp
var bracketsReg *regexp.Regexp
var separatorReg *regexp.Regexp
var recursiveReg *regexp.Regexp

type argError struct {
	arg  string
	prob string
}

type Ingredient struct {
	Name        string        `json:"name"`
	ID          int           `json:"id"`
	Danger      int           `json:"danger"`
	Groups      []int64       `json:"groups"`
	Ingredients *[]Ingredient `json:"ingredients"`
	// Description string        `json:"description"`
	// WikiLink    string        `json:"wiki_link"`
}

func init() {
	var err, err1, err2, err3 error
	database = connection.Connect()
	reg, err = regexp.Compile("[^a-zA-Z0-9А-Яа-я[:space:]]+")
	if err != nil {
		log.Fatal(err)
	}
	bracketsReg, err1 = regexp.Compile("[)]+")
	if err1 != nil {
		log.Fatal(err1)
	}
	separatorReg, err2 = regexp.Compile("[,.:;]+")
	if err2 != nil {
		log.Fatal(err2)
	}
	recursiveReg, err3 = regexp.Compile("[(]+")
	if err3 != nil {
		log.Fatal(err3)
	}
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
			danger, id, groups, e := getDangerLevel(curWord)
			if e != nil {
				return e
			}
			*ings = append(*ings, Ingredient{curWord, id, danger, groups, nil})
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
			danger, id, groups, e := getDangerLevel(curWord)
			if e != nil {
				return err
			}
			*ings = append(*ings, Ingredient{curWord, id, danger, groups, &subIngs})
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

func getDangerLevel(ing string) (int, int, []int64, error) {
	var danger, id int
	var groups []int64
	err := database.QueryRow(
		`SELECT i.id, i.danger, coalesce(ing_groups.groups, '{}') FROM ingredients AS i 
		FULL JOIN ing_groups ON i.id = ing_groups.id
		WHERE i.name = $1
		ORDER BY frequency LIMIT 1;
	`, ing).Scan(&id, &danger, pq.Array(&groups))
	// println(groups[0)
	if err != nil {
		if err == pgx.ErrNoRows {
			return -1, 0, groups, nil
		}
		// println(ing)
		log.Println("ERROR analyzer.go:132: getDangerLevel()", err.Error())
		return -1, -1, groups, err
	}
	// if len(groups) == 0 {
	// 	groups = "NULL"
	// }
	return danger, id, groups, nil
}
