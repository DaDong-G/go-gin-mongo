package mongo_connect

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	Id        bson.ObjectId `bson:"_id"`
	Username  string        `bson:"name"`
	Pass      string        `bson:"pass"`
	Regtime   int64         `bson:"regtime"`
	Interests []string      `bson:"interests"`
}

const URL string = "127.0.0.1:27017"

var (
	c       *mgo.Collection
	session *mgo.Session
)

//var session *mgo.Session

func init() {
	fmt.Println("aaaa")
	session, _ = mgo.Dial(URL)
	//切换到数据库
	db := session.DB("blog")
	//切换到collection
	c = db.C("mgotest")
}

func (user User) ToString() string {
	return fmt.Sprintf("%+v", user)
}

func add() {
	stu1 := new(User)
	stu1.Id = bson.NewObjectId()
	fmt.Println(stu1)
	fmt.Println(stu1.Id)
	stu1.Username = "stu1_name"
	stu1.Pass = "stu1_pass"
	stu1.Regtime = time.Now().Unix()
	stu1.Interests = []string{"象棋", "游泳"}
	err := c.Insert(stu1)
	if err == nil {
		fmt.Println("插入成功")
	} else {
		fmt.Println(err.Error())
		defer panic(err)
	}
}

func Find() string {
	defer session.Close()
	var users []User
	c.Find(bson.M{"name": "stu1_name"}).All(&users)
	for _, value := range users {
		fmt.Println(value.ToString())
		//	fmt.Println(typeof(value))
	}
	idStr := "5d53aee20496b73d44bd1bcd"
	objectId := bson.ObjectIdHex(idStr)
	user := new(User)
	c.Find(bson.M{"_id": objectId}).One(user)
	fmt.Println(user)
	return user.Username

}

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func del() {
	err := c.Remove(bson.M{"_id": bson.ObjectIdHex("577fb2d1cde67307e819133d")})
	if err != nil {
		fmt.Println("删除失败" + err.Error())
	} else {
		fmt.Println("删除成功")
	}
}

//func main()  {
//	//add()
//	find()
//}
