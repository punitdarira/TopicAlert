package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"
	"topicalert/Mail"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Topic struct {
	Email_id string `json:"email_id"`
	Topic    string `json:"topic"`
}

type EmailRowDb struct {
	Email_id string
	User_id  int
}

func main() {

	runTracker()
	Mail.Mail()
	/*
		router := gin.Default()
		router.GET("/get-topics", getTopics)
		router.POST("/new-topic", postTopic)
		router.Run("localhost:8080")
	*/
}

func getTopics(c *gin.Context) {
	dbConnection()
	//c.IndentedJSON(http.StatusOK, topics)
}

func dbConnection() {

}

func postTopic(c *gin.Context) {
	var newTopic Topic

	// Call BindJSON to bind the received JSON to
	// newTopic.
	if err := c.BindJSON(&newTopic); err != nil {
		return
	}

	// Add the new album to the slice.
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/topictracker")
	if err != nil {
		panic(err.Error())
	}
	var userId int
	var emailId string

	db.QueryRow("select email_id, user_id from topictracker.emails where email_id = ?",
		newTopic.Email_id).Scan(&emailId, &userId)
	fmt.Println("EmailId from DB=", emailId)
	defer db.Close()
	if emailId == "" {
		var insertedEmail EmailRowDb
		if _, err := db.Query("insert into topictracker.emails (email_id) values (?)",
			newTopic.Email_id); err != nil {
			c.IndentedJSON(http.StatusCreated, "error while insert")
		} else {
			emailInsertedRow := db.QueryRow("select email_id, user_id from topictracker.emails where email_id = ?",
				newTopic.Email_id)
			emailInsertedRow.Scan(&insertedEmail.Email_id, &insertedEmail.User_id)
			fmt.Println(insertedEmail.User_id)
			fmt.Println(insertedEmail.Email_id)

			//inserting into topic db
			insertTopic(db, insertedEmail.User_id, newTopic.Topic, c)
		}
	} else {
		insertTopic(db, userId, newTopic.Topic, c)
	}

}

func insertTopic(db *sql.DB, user_id int, topic string, c *gin.Context) {
	db.Query("insert into topictracker.topics (user_id, topic) values (?, ?)",
		user_id, topic)
	c.IndentedJSON(http.StatusCreated, topic)
}

func runTracker() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/topictracker")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("select * from topictracker.topics group by user_id, topic")
	var userId int
	var topic string
	userTopicsMap := make(map[int][]string)
	for rows.Next() {
		rows.Scan(&userId, &topic)
		userTopicsMap[userId] = append(userTopicsMap[userId], topic)
	}

	var userWaitGroup sync.WaitGroup
	userWaitGroup.Add(len(userTopicsMap))

	for userId, userTopics := range userTopicsMap {
		go runTopicsForUser(userId, userTopics, &userWaitGroup)
	}
	userWaitGroup.Wait()
}

func runTopicsForUser(userId int, userTopics []string, userWaitGroup *sync.WaitGroup) {
	defer userWaitGroup.Done()
	fmt.Println("Running for user ", userId)
	var topicForEachUser sync.WaitGroup
	topicForEachUser.Add(len(userTopics))
	for _, topic := range userTopics {
		go runTopic(topic, &topicForEachUser)
	}
	topicForEachUser.Wait()
}

func runTopic(topic string, topicForEachUser *sync.WaitGroup) {
	defer topicForEachUser.Done()
	fmt.Println("Running topic ", topic)
	scrape(topic)
}
