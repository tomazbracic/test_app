package main

import (
	"fmt"
	"net/http"
	"log"
	"html/template"
	"time"

	"github.com/gocql/gocql"
)

type Item struct {
	ID string `json:"uuid"`
	Cas string `json:"cas"`
}

type ItemsPageData struct {
	PageTitle string
	Items []Item
}


var Session *gocql.Session
var tmpl *template.Template

func init() {
	var err error
	
	cluster := gocql.NewCluster("cassandra:9042")
	cluster.Port = 9042
	cluster.Keyspace = "projekt1"
	
	Session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	
	fmt.Println("cassandra init done")
}

func setTime(w http.ResponseWriter, r *http.Request) {
	var errs []string
	gocqlUuid := gocql.TimeUUID()
	cajt := time.Now()

	if err := Session.Query(`
		INSERT INTO projekt1.test (id, cas) VALUES (?, ?)`,
		gocqlUuid, cajt).Exec(); err != nil {
			errs = append(errs, err.Error())
	} 

	tmpl, _ = template.ParseFiles("set.html")
	tmpl.Execute(w, cajt)
}

func getTime(w http.ResponseWriter, r *http.Request) {
	var Items []Item
	
	query := "SELECT id, cas FROM projekt1.test"
	var tmpid string
	var tmpcas time.Time 

	iter := Session.Query(query).Iter()
	for iter.Scan(&tmpid, &tmpcas) {
		newitem := Item{ID: tmpid, Cas: tmpcas.String()}
		Items = append(Items, newitem)
	}

	err := iter.Close()
	if err != nil {
		log.Fatal(err)
	}

	data := ItemsPageData{
		PageTitle: "Moj casovni seznam",
		Items: Items,
	}

	fmt.Println(data)

	tmpl, err = template.ParseFiles("show.html")
	tmpl.Execute(w, data)
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	var err error
	
	tmpl, err = template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, nil)
}


func main() {
	var err error

	defer Session.Close()
	
	
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", mainPage)
	http.HandleFunc("/set", setTime)
	http.HandleFunc("/get", getTime)

	log.Fatal(http.ListenAndServe(":5555", nil))

}



// USE projekt1

// CREATE KEYSPACE IF NOT EXISTS projekt1 WITH REPLICATION = { 'class': 'SimpleStrategy', 'replication_factor': 1 };

// CREATE TABLE IF NOT EXISTS test ( id uuid PRIMARY KEY, cas timestamp );

// INSERT INTO test (id, cas) VALUES (uuid(), dateof(now()));

// SELECT * FROM test;

// id                                   | cas
// --------------------------------------+---------------------------------
//  01e77154-95e6-436a-a447-8ee0b7b7e238 | 2018-04-12 06:39:45.208000+0000
