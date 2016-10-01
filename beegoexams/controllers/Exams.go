
package controllers

import "github.com/astaxie/beego"
//import "encoding/json"


type Exams struct {
	beego.Controller
}

type ExamsResult struct {
	Questions []QuestionsResult
}

type QuestionsResult struct {
	Question string
	Answers []string
	UserAnswer string
}

// Use Get rather than Post so that we can simulate easier in the browser
func (this *Exams) Get() {
    var result ExamsResult
    
    question1 := QuestionsResult{}
    question1.Question = "Why do most people think that they are incapable of ________?"
    answers1 := []string{"draw", "drawing", "drawn"}
	question1.Answers = answers1
	question1.UserAnswer = ""
    
    question2 := QuestionsResult{}
    question2.Question = "Certificates and academic honors will________ at the graduation ceremony."
    answers2 := []string{"award","be award","be awarded"}
	question2.Answers = answers2
	question2.UserAnswer = ""
    
    question3 := QuestionsResult{}
    question3.Question = "I wish he would stop ___________ to be an expert in genetic engineering."
    answers3 := []string{"pretend","pretending","to pretend"}
	question3.Answers = answers3
	question3.UserAnswer = ""
    
    question4 := QuestionsResult{}
    question4.Question = "You’re not coming to the party tonight? You __________ let me know earlier."
    answers4 := []string{"would have","should have","will have"}
	question4.Answers = answers4
	question4.UserAnswer = ""
    
    question5 := QuestionsResult{}
    question5.Question = "Alex will need new furniture when ________ into his new apartment."
    answers5 := []string{"he moves","does he move","he will move"}
	question5.Answers = answers5
	question5.UserAnswer = ""

    question6 := QuestionsResult{}
    question6.Question = "Alex will need new furniture when ________ into his new apartment."
    answers6 := []string{"he moves","does he move","he will move"}
	question6.Answers = answers6
	question6.UserAnswer = ""
    
    question7 := QuestionsResult{}
    question7.Question = "The manager’s position had been vacant for more than three weeks when it ____________."
    answers7 := []string{"had filled","was filled","been filled"}
	question7.Answers = answers7
	question7.UserAnswer = ""

    question8 := QuestionsResult{}
    question8.Question = "His theory is very difficult; ________ people understand it."
    answers8 := []string{"few","a few","a great many"}
	question8.Answers = answers8
	question8.UserAnswer = ""
	
	question9 := QuestionsResult{}
    question9.Question = "When a law is ____________, each branch of the government votes on it."
    answers9 := []string{"been passed","been passing","being passed"}
	question9.Answers = answers9
	question9.UserAnswer = ""

    question10 := QuestionsResult{}
    question10.Question = "Randy was eager to know whether ____________ been admitted to Harvard."
    answers10 := []string{"he has","he had","had he"}
	question10.Answers = answers10
	question10.UserAnswer = ""

    question11 := QuestionsResult{}
    question11.Question = "I was supposed to write a ten-page paper, but I said what I _____ in only five pages."
    answers11 := []string{"must’ve said","had to say","better say"}
	question11.Answers = answers11
	question11.UserAnswer = ""

    question12 := QuestionsResult{}
    question12.Question = "______ go about the task should be everyone's concern."
    answers12 := []string{"What we will","How will","How we will"}
	question12.Answers = answers12
	question12.UserAnswer = ""

    question13 := QuestionsResult{}
    question13.Question = "If you __________ what she was going through, you wouldn’t be so critical."
    answers13 := []string{"knew","would known","have known"}
	question13.Answers = answers13
	question13.UserAnswer = ""

    question14 := QuestionsResult{}
    question14.Question = "Maybe this knowledge will help us understand many old rituals __________ have lasted over the centuries."
    answers14 := []string{"what will","which","that they"}
	question14.Answers = answers14
	

    question15 := QuestionsResult{}
    question15.Question = "The editor disregarded my request that he __________ any story concerning violence."
    answers15 := []string{"not print","didn´t print","don´t print"}
	question15.Answers = answers15
	question15.UserAnswer = ""

    question16 := QuestionsResult{}
    question16.Question = "Mac asked the students if learning about other cultures __________ contribute to international understanding in the long run."
    answers16 := []string{"would","will","had"}
	question16.Answers = answers16
	question16.UserAnswer = ""

    question17 := QuestionsResult{}
    question17.Question = "Once the math teacher had proposed the date for the final test, she couldn’t __________."
    answers17 := []string{"call off it","called off","call it off"}
	question17.Answers = answers17
	question17.UserAnswer = ""


    question18 := QuestionsResult{}
    question18.Question = "______________ knowing what to do, she resorted to therapy."
    answers18 := []string{"Do not","Not","To not"}
	question18.Answers = answers18
	question18.UserAnswer = ""

    question19 := QuestionsResult{}
    question19.Question = "Never ____________ in such a colossal traffic jam."
    answers19 := []string{"I have been","have been I","have I been"}
	question19.Answers = answers19
	question19.UserAnswer = ""

    question20 := QuestionsResult{}
    question20.Question = "Not only _________________ us with her presence, but she also gave us a memorable keepsake."
    answers20 := []string{"Jo honored","did Jo honor","Jo had honored"}
	question20.Answers = answers20
	question20.UserAnswer = ""

    question21 := QuestionsResult{}
    question21.Question = "How does she do it? She does not really follow a strict diet, nor _________________ work out!"
    answers21 := []string{"she","she doesn´t","does she"}
	question21.Answers = answers21
	question21.UserAnswer = ""


    
    aSlice := []QuestionsResult{question1,question2,question3,question4,question5,question6,question7,question8,question9,question10,question11,question12,question13,question14,question15,question16,question17,question18,question19,question20,question21}
    
    result.Questions = aSlice
	
	
	this.Data["json"] = result
	this.ServeJSON()
}
