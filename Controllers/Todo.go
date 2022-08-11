package Controllers

import (
	"TREgitim/Config"
	"TREgitim/Models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Model struct {
	Mentee   Models.Mentee
	Userprof Models.UserProfile
	User     Models.User
	About    Models.About
	Skills   []Models.SkillCatalog
}

type TodoModel struct {
	Mentor         Models.Mentor
	Mentee         Models.Mentee
	UserMentee     Models.User
	UserMentor     Models.User
	UserProfMentee Models.UserProfile
	UserProfMentor Models.UserProfile
	Todo0          []Models.Todo
	Todo1          []Models.Todo
	Todo2          []Models.Todo
	Todo3          []Models.Todo
	Who            uint
	Percent        int
}

type AddTodoModel struct {
	Title    string
	Menteeid uint
	Mentorid uint
}

type Todo struct {
	Title       string
	Description string
	Enddate     string
}

type DragModel struct {
	Id         uint
	Status     int
	Actiondate string
}

type CommentModel struct {
	Description string
	Todoid      uint
}

func MyMentees(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var user Models.User
	var mentor Models.Mentor
	var mentees []Models.Mentee
	Config.DB.Where("mail = ?", controll).First(&user)
	Config.DB.Where("user_id", user.ID).First(&mentor)
	Config.DB.Where("mentor_id", mentor.ID).Find(&mentees)

	var jsonmodel []Model

	for _, element := range mentees {
		var userr Models.User
		var solomodel Model
		Config.DB.Where("id = ?", element.UserID).First(&userr)
		solomodel.User = userr
		var prof Models.UserProfile
		Config.DB.Where("user_id = ?", element.UserID).First(&prof)
		solomodel.Userprof = prof
		var about Models.About
		Config.DB.Where("user_id = ?", element.UserID).First(&about)
		solomodel.Mentee = element

		var skillsfirst []Models.Skill
		Config.DB.Where("user_id= ?", element.UserID).Find(&skillsfirst)
		var skillsfinal []Models.SkillCatalog
		skillsfinal = nil
		for index, _ := range skillsfirst {
			var skillssecond Models.SkillCatalog
			Config.DB.Where("id= ?", skillsfirst[index].SkillCatalogID).First(&skillssecond)
			skillsfinal = append(skillsfinal, skillssecond)
		}
		solomodel.Skills = skillsfinal

		jsonmodel = append(jsonmodel, solomodel)
	}

	c.JSON(200, jsonmodel)

}

func GetTodo(c *gin.Context) {
	// menteeid - mentorid
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	number, _ := strconv.ParseUint(c.Param("menteeid"), 10, 32)
	numberr, _ := strconv.ParseUint(c.Param("mentorid"), 10, 32)
	controll := fmt.Sprintf("%s", control)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var mentee Models.Mentee
	var mentor Models.Mentor
	Config.DB.Where("id = ?", uint(number)).First(&mentee)
	Config.DB.Where("id = ?", uint(numberr)).First(&mentor)
	var usermentee Models.User
	var usermentor Models.User
	Config.DB.Where("id = ?", mentee.UserID).First(&usermentee)
	Config.DB.Where("id = ?", mentor.UserID).First(&usermentor)
	var profmentee Models.UserProfile
	var profmentor Models.UserProfile
	Config.DB.Where("user_id = ?", mentee.UserID).First(&profmentee)
	Config.DB.Where("user_id = ?", mentor.UserID).First(&profmentor)

	var jsonmodel TodoModel
	jsonmodel.Mentee = mentee
	jsonmodel.Mentor = mentor
	jsonmodel.UserMentee = usermentee
	jsonmodel.UserMentor = usermentor
	jsonmodel.UserProfMentee = profmentee
	jsonmodel.UserProfMentor = profmentor

	var todos []Models.Todo
	var todo0 []Models.Todo
	var todo1 []Models.Todo
	var todo2 []Models.Todo
	var todo3 []Models.Todo
	Config.DB.Where("mentee_id = ? AND is_deleted = ?", mentee.ID, false).Find(&todos)

	//var comment Models.Comment
	//var comments []Models.Comment
	var totaltodo int
	var successtodo int
	for _, element := range todos {
		//var comments []Models.Comment
		totaltodo += 1
		if element.Issuccessfull == true {
			successtodo += 1
		}
		if element.Status == 0 {
			todo0 = append(todo0, element)
		}
		if element.Status == 1 {
			todo1 = append(todo1, element)
		}
		if element.Status == 2 {
			todo2 = append(todo2, element)
		}
		if element.Status == 3 {
			todo3 = append(todo3, element)
		}

	}
	var percent int
	var x = 100 / totaltodo
	percent = successtodo * x
	if successtodo == totaltodo {
		percent = 100
	}
	jsonmodel.Percent = percent

	var tododelete1 Models.Todo
	Config.DB.Where("id = ?", 11).First(&tododelete1)
	todo0 = append(todo0, tododelete1)
	var tododelete2 Models.Todo
	Config.DB.Where("id = ?", 12).First(&tododelete2)
	todo1 = append(todo1, tododelete2)
	var tododelete3 Models.Todo
	Config.DB.Where("id = ?", 13).First(&tododelete3)
	todo2 = append(todo2, tododelete3)
	var tododelete4 Models.Todo
	Config.DB.Where("id = ?", 14).First(&tododelete4)
	todo3 = append(todo3, tododelete4)
	jsonmodel.Todo0 = todo0
	jsonmodel.Todo1 = todo1
	jsonmodel.Todo2 = todo2
	jsonmodel.Todo3 = todo3
	var userrr Models.User
	var mentorrr Models.Mentor
	Config.DB.Where("mail = ?", controll).First(&userrr)
	Config.DB.Where("user_id = ?", userrr.ID).First(&mentorrr)
	if mentorrr.ID != 0 {
		jsonmodel.Who = mentorrr.ID
	}
	if mentorrr.ID == 0 {
		var menteeee Models.Mentee
		Config.DB.Where("user_id = ?", userrr.ID).First(&menteeee)
		jsonmodel.Who = menteeee.ID
	}

	c.JSON(200, jsonmodel)
}

func AddTodo(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	if controll == "" {
		c.JSON(200, "Session boş")
	}

	var model AddTodoModel
	var todo Models.Todo
	c.BindJSON(&model)
	todo.Title = model.Title
	todo.Status = 0
	todo.MenteeID = model.Menteeid
	todo.MentorID = model.Mentorid
	//application.Date = time.Now().Format("02-01-2006")
	Config.DB.Create(&todo)
	c.JSON(200, &todo)
}

func DeleteTodo(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var todo Models.Todo
	Config.DB.Where("id = ?", uint(number)).First(&todo)
	todo.IsDeleted = true
	Config.DB.Save(&todo)
	c.JSON(200, "Silme işlemi başarılı")
}

func UpdateTodo(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if control == nil {
		c.Redirect(301, "/Login")
	}

	var todo Todo
	var todowillchange Models.Todo
	c.BindJSON(&todo)
	Config.DB.Where("id = ?", number).First(&todowillchange)
	todowillchange.Title = todo.Title
	todowillchange.Description = todo.Description
	todowillchange.EndDate = todo.Enddate
	Config.DB.Save(todowillchange)
	c.JSON(200, todowillchange)
}

func Dragged(c *gin.Context) {
	//todoid
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]

	if control == nil {
		c.Redirect(301, "/Login")
	}
	var todomodel DragModel
	c.BindJSON(&todomodel)

	var todo Models.Todo
	Config.DB.Where("id = ?", todomodel.Id).First(&todo)
	todo.Status = todomodel.Status

	//action, _ := time.Parse("02.01.2006", todomodel.Actiondate)
	//end, _ := time.Parse("02.01.2006", todo.EndDate)
	//action, _ := time.Parse("2006.01.02", todomodel.Actiondate)
	//end, _ := time.Parse("2006.01.02", todo.EndDate)
	var str1 = todomodel.Actiondate
	var str2 = todo.EndDate
	var strgunend = str2[8:10]
	var strgunact = str1[0:2]
	var strayend = str2[5:7]
	var strayact = str1[3:5]
	strayendd, _ := strconv.ParseUint(strayend, 10, 32)
	strgunendd, _ := strconv.ParseUint(strgunend, 10, 32)
	strgunactt, _ := strconv.ParseUint(strgunact, 10, 32)
	strayactt, _ := strconv.ParseUint(strayact, 10, 32)
	fmt.Println(strayendd)
	if todomodel.Status == 3 && strayactt < strayendd {
		todo.Issuccessfull = true
		Config.DB.Save(todo)
	}
	if todomodel.Status == 3 && strayactt == strayendd && strgunactt < strgunendd {
		todo.Issuccessfull = true
		Config.DB.Save(todo)
	}
	if todomodel.Status == 3 && strayactt == strayendd && strgunactt == strgunendd {
		todo.Issuccessfull = true
		Config.DB.Save(todo)
	}
	todo.ActionDate = todomodel.Actiondate
	Config.DB.Save(todo)

	c.JSON(200, "Statü değişti.")
}

func AddComment(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	controll := fmt.Sprintf("%s", control)

	if control == nil {
		c.Redirect(301, "/Login")
	}
	var jsonmodel CommentModel
	c.BindJSON(&jsonmodel)

	var comment Models.Comment
	var user Models.User
	var mentee Models.Mentee
	var mentor Models.Mentor
	Config.DB.Where("mail = ?", controll).First(&user)
	Config.DB.Where("user_id = ?", user.ID).First(&mentee)
	comment.Description = jsonmodel.Description
	comment.TodoID = jsonmodel.Todoid
	if mentee.ID == 0 {
		Config.DB.Where("user_id = ?", user.ID).First(&mentor)
		comment.MentorID = mentor.ID
		comment.MenteeID = 36
	}
	if mentee.ID != 0 {
		comment.MenteeID = mentee.ID
		comment.MentorID = 56
	}
	Config.DB.Create(&comment)

	c.JSON(202, comment)
}

type CommentModell struct {
	Comment Models.Comment
	Profile Models.UserProfile
}

func TodoComments(c *gin.Context) {
	session, _ := store.Get(c.Request, "sessioncontrol")
	control := session.Values["sessionmail"]
	number, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	if control == nil {
		c.Redirect(301, "/Login")
	}
	var comments []Models.Comment
	Config.DB.Where("todo_id = ?", uint(number)).Find(&comments)

	var jsonnn []CommentModell
	for _, element := range comments {
		var jsonmodell CommentModell
		jsonmodell.Comment = element
		if element.MentorID != 56 {
			var mentor Models.Mentor
			var userr Models.User
			var prof Models.UserProfile
			Config.DB.Where("id = ?", element.MentorID).First(&mentor)
			Config.DB.Where("id = ?", mentor.UserID).First(&userr)
			Config.DB.Where("user_id = ?", userr.ID).First(&prof)
			jsonmodell.Profile = prof
		}
		if element.MentorID == 56 {
			var mentee Models.Mentee
			var userr Models.User
			var prof Models.UserProfile
			Config.DB.Where("id = ?", element.MenteeID).First(&mentee)
			Config.DB.Where("id = ?", mentee.UserID).First(&userr)
			Config.DB.Where("user_id = ?", userr.ID).First(&prof)
			jsonmodell.Profile = prof
		}
		jsonnn = append(jsonnn, jsonmodell)
	}

	c.JSON(200, jsonnn)
}
