package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Relation struct {
	Rid  int32
	Type string
	With int32
	//From string
}

type Person struct {
	Pid       int32
	Name      string
	Relations map[int32]int32 //maps relationid to personid
}

//-------------------------------------------------------------
func (p *Person) AddRelation(q *Person, name string) *Person {
	rel := &Relation{101, name, q.Pid} //store in stack too
	//Realtion with her identity! not whole her
	p.Relations = make(map[int32]int32)
	p.Relations[rel.Rid] = q.Pid
	//Above relationship defined with me :)
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
