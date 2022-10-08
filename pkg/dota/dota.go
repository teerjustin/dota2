package dota

import (
	"fmt"
	//"html"
    "log"
    "net/http"
	"strconv"
    "sync"
	"io"
	"encoding/json"
	"bytes"
)
const OPENDOTA_KEY = "b8eb4eb9-a8ef-4bd3-bf9e-f25d2b800e04"
/* Your Key
b8eb4eb9-a8ef-4bd3-bf9e-f25d2b800e04
To use your key, add api_key=XXXX as a query parameter to your API request:

https://api.opendota.com/api/matches/271145478?api_key=b8eb4eb9-a8ef-4bd3-bf9e-f25d2b800e04
The current billing cycle ends on 12/31/2021. We'll automatically bill the American Express ending in 2008.

Need support? Email api@opendota.com.



Your Usage
Month	# API calls	Estimated Fees
11	12	$0.01
Get started for free. Keep going at a ridiculously low price.
Free Tier	Premium Tier
Price	Free	$0.01 per 100 calls
Key Required?	No	Yes -- requires payment method
Call Limit	50000 per month	Unlimited
Rate Limit	60 calls per minute	1200 calls per minute
Support	Community support via Discord group	Priority support from core developers
Details
You're charged $0.0001 per call, rounded up to the nearest cent.
Getting an API key requires a linked payment method. We'll automatically charge the card at the beginning of the month for any fees owed.
500 errors don't count as usage, since that means we messed
*/

const MANDRILL_KEY = "j9H9bH6FxMUooFXnvStdTA"

func sendEmail(w http.ResponseWriter, r *http.Request) {
	to := To{Email: "teerjustin@gmail.com", Name: "Justin"}
    message := Message{
		Text: "Some Text Here",
		FromEmail: "kevinanthony.teer@gmail.com",
		FromName: "Kevin",
		Subject: "testEmail",
		To: []To{to},
	}
	m := MandrillPost{
		Key: MANDRILL_KEY,
		Message: message,
	}
	fmt.Println("This is our message object: ", m);

    json_data, err := json.Marshal(m)
    if err != nil {
        log.Fatal(err)
    }
    resp, err := http.Post("https://mandrillapp.com/api/1.0/messages/send", "application/json", bytes.NewBuffer(json_data))
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("THIS IS THE RESPONSE FROM MANDRILL: ", resp)
    var res map[string]interface{}


    json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println("THIS IS RES: ", res)
    fmt.Println(res["json"])
}

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "hello")
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    counter++
    fmt.Fprintf(w, strconv.Itoa(counter))
    mutex.Unlock()
}

func getAllHeroes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp, err := http.Get("https://api.opendota.com/api/heroes?api_key=" + OPENDOTA_KEY)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var heroes = &[]Hero{}
	err = json.Unmarshal(body, heroes)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(heroes)
}

func getAllHeroIcons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("HI IM JUSTIN")
	resp, err := http.Get("https://api.opendota.com/api/herostats?api_key=" + OPENDOTA_KEY)
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(resp.Body)
	var icons = &[]Icon{}
	err = json.Unmarshal(body, icons)
	if err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(icons)
}


func Run() {
	fmt.Println("It works")
	fmt.Println("Hello")

    http.HandleFunc("/", echoString)
	http.HandleFunc("/increment", incrementCounter)

    http.HandleFunc("/heroes", getAllHeroes)
	http.HandleFunc("/sendEmail", sendEmail)
	http.HandleFunc("/icons", getAllHeroIcons)
    log.Fatal(http.ListenAndServe(":8081", nil))
}