package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Relationship struct {
	Rid  int
	Type string
	With int
	//From string
}

type Person struct {
	Pid      int
	Name     string
	Relation map[string][]int
}

//-------------------------------------------------------------
func (p *Person) AddRelation(q *Person, name string) *Person {
	r := &Relationship{101, name, q.Pid} //store in stack too
	//relationship stored, pointing to her

	if p.Relation == nil {
		p.Relation = make(map[string][]int)
	}
	//Initializing map

	p.Relation[name] = append(p.Relation[name], r.Rid) //to relatinship id only
	return p
}

//-------------------------------------------------------------

func main() {
	//api request to be match with url SOMETHING LIKE CHAT INTERFACE
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Content-type", "application/json")

		p1 := &Person{Pid: 1}
		p1.Name = "Peeyush" //this id and other details may come from pages: our and other
		p2 := &Person{2, "Shreya", nil}

		p1.AddRelation(p2, "WIFE")
		p1.AddRelation(p2, "FRIEND")
		p2.AddRelation(p1, "HUSBAND")

		b, err := json.Marshal(p1)
		if err != nil {
			log.Fatal(err)
		}
		rw.Write(b)
	})
	log.Println("Running on http://localhost:9890")
	log.Fatal(http.ListenAndServe(":9890", nil))
}
