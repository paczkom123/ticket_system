package main 

/*
	TODO because of the time limit everything is in one file - so next version should have classes
	TODO make the functionalities microservices and work with RESTful http request
*/

import (
	"fmt"
	"log"
    "os"
	"math/rand"
	"time"
	"github.com/gofrs/uuid"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Ticket struct {
	Uuid string
}

const (
	maxRandom = 5 //max number of generated tickets
	totalTicket = 1000 // max number of inserted tickets
)

func main(){
	ticketQueue := make([]string,0)
	var numberOfRejected = 0
	var totalNumber = 0
	
	//creating log file
	f, err := os.OpenFile("queue_size.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
    	log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	db, err := gorm.Open("sqlite3", "ticket.db")
  	if err != nil {
    	panic("failed to connect database")
  	}
	defer db.Close()
	db.AutoMigrate(&Ticket{})

	for numberOfRejected < totalTicket {
		log.Printf("Queu size %d" , len(ticketQueue) ) //logging the size of the queue
		go generateUUID(&ticketQueue,&numberOfRejected,&totalNumber,db)
		go allocateTicket(&ticketQueue,db)
		time.Sleep(1 * time.Second)
	}

	//print final results
	fmt.Println(numberOfRejected)
	fmt.Println(totalNumber)
	fmt.Println(totalNumber-numberOfRejected)
}

/*
	add uuid to the queue. Adds if the queue has less than 10 elements and 
	if the uuid is not in the queue and in the database
	@param ticketQueue reference of the queue slice
	@numberOfRejected reference of numberOfRejected variable
	@totalNumber reference of totalNumber variable 
	@param db reference to gorm object 
*/
func generateUUID(ticketQueue *[]string,numberOfRejected *int,totalNumber *int,db *gorm.DB){
	size := randomInt()
	for i := 0; i < size; i++ {
		if (len(*ticketQueue) >= 10){
			(*numberOfRejected)++
			(*totalNumber)++
		}else{
			u := uuid.Must(uuid.NewV4())
			_,exist := getTicket(db,u.String())
			if (!existInQueue(ticketQueue,u.String()) && !exist){
				(*totalNumber)++
				*ticketQueue = append(*ticketQueue,u.String())
			}
		}
	}

}

/*
	checks if there is any ticket to allocate and insert them to the database
	inserts only 1 when the size of the queue is 1 and 2 when size > 2
	@param ticketQueue reference of the queue slice
	@param db reference to gorm object
	returns 1 if there was any allocation and 0 if there was not (not used in this version) 
*/
func allocateTicket(ticketQueue *[]string,db *gorm.DB)int{
	//making sure that it will not go out of bounds
	if (len(*ticketQueue)==0){
		return 0
	}
	//making sure that it will not go out of bounds
	if (len(*ticketQueue)==1){
		first := (*ticketQueue)[0]
		db.Create(&Ticket{Uuid: first})
		copy((*ticketQueue), (*ticketQueue)[1:])
		*ticketQueue = (*ticketQueue)[:len((*ticketQueue)) - 1]
		return 1
	}
	first,second := (*ticketQueue)[0],(*ticketQueue)[1]
	db.Create(&Ticket{Uuid: first})
	db.Create(&Ticket{Uuid: second})
	copy((*ticketQueue), (*ticketQueue)[2:])
	*ticketQueue = (*ticketQueue)[:len((*ticketQueue)) - 2]


	return 1
}

/*	
	checks if uuid exist in the queue
	@param ticketQueue reference of the queue slice
	@param u string id to check
	@return bool true exist false doesnt exist
*/
func existInQueue(ticketQueue *[]string,u string)bool{
	for _, a := range (*ticketQueue) {
        if a == u {
			return true
			break
        }
    }
    return false
}

/*
	checks if ticket with a specific uuid exist in the database 
	and returns the ticket/or empty Ticket struct and if it exist
	@param db reference to gorm object
	@param u string id to check
	@return Ticket bool
*/
func getTicket(db *gorm.DB,u string)(Ticket,bool){
	var ticket Ticket
	db.First(&ticket, "uuid = ?", u)
	if (ticket != Ticket{}) {
		return ticket,true
	}
	return ticket,false

}

// Returns an int >= min, < max
func randomInt() int {
	rand.Seed(time.Now().UnixNano())
    return 1 + rand.Intn(maxRandom-1)
}