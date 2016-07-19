package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Title           string
	Uid             int64
	Category        string
	Attachment      string
	Content         string    `orm:"size(5000)"`
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type Comment struct {
	Id      int64
	Tid     int64
	Name    string
	Content string    `orm:"size(5000)"`
	Created time.Time `orm:"index"`
}

func init() {
	orm.RegisterModel(new(Category), new(Topic), new(Comment))
}

func AddComment(tid, name, content string) error {
	o := orm.NewOrm()

	cid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	comment := &Comment{
		Tid:     cid,
		Name:    name,
		Content: content,
		Created: time.Now(),
	}
	_, err = o.Insert(comment)

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", cid).One(topic)
	if err != nil {
		return err
	}

	topic.ReplyCount++
	_, err = o.Update(topic)
	return err
}

func AddTopic(title, category, content, attachment string) (string, error) {
	o := orm.NewOrm()

	topic := &Topic{
		Title:      title,
		Category:   category,
		Content:    content,
		Attachment: attachment,
		Created:    time.Now(),
		Updated:    time.Now(),
		ReplyTime:  time.Now(),
	}
	_, err := o.Insert(topic)
	o.Read(topic)
	id := topic.Id
	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err != nil {
		beego.Error(err)
	}
	cate.TopicCount++
	_, err = o.Update(cate)
	if err != nil {
		beego.Error(err)
	}
	return strconv.FormatInt(id, 10), err
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}

	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTopics(isDesc bool) (topics []*Topic, err error) {
	o := orm.NewOrm()

	topics = make([]*Topic, 0)

	qs := o.QueryTable("topic")
	if isDesc {
		_, err = qs.OrderBy("-created").All(&topics)
	} else {
		_, err = qs.All(&topics)
	}
	return topics, err
}

func GetCateTopics(category string) ([]*Topic, error) {
	o := orm.NewOrm()
	topic := make([]*Topic, 0)

	qs := o.QueryTable("topic")
	_, err := qs.Filter("category", category).All(&topic)
	if err != nil {
		return nil, err
	}
	return topic, err
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)

	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}

func GetAllComment(tid string) ([]*Comment, error) {
	o := orm.NewOrm()
	cid, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		beego.Error(err)
	}

	comm := make([]*Comment, 0)
	qs := o.QueryTable("comment")
	_, err = qs.Filter("tid", cid).All(&comm)
	if err != nil {
		return nil, err
	}

	return comm, err
}

func DelTopic(id, category string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	cate := new(Category)
	qs := o.QueryTable("category")
	err = qs.Filter("title", category).One(cate)
	if err != nil {
		return err
	}
	cate.TopicCount--
	_, err = o.Update(cate)
	if err != nil {
		return err
	}

	topic := &Topic{Id: cid}

	_, err = o.Delete(topic)
	return err
}

func DelComment(id, tid string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := new(Topic)
	qs := o.QueryTable("topic")
	err = qs.Filter("id", tid).One(topic)
	topic.ReplyCount--
	if err != nil {
		return err
	}
	_, err = o.Update(topic)

	comment := &Comment{Id: cid}
	_, err = o.Delete(comment)
	return err
}
func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()

	cate := &Category{Id: cid}

	_, err = o.Delete(cate)

	return err
}

func ModifyTopic(id, title, category, content, attachment string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	topic := &Topic{Id: cid}

	var oldcate string
	oldcate = topic.Category

	if o.Read(topic) == nil {
		topic.Title = title
		topic.Category = category
		topic.Attachment = attachment
		topic.Content = content
		topic.Updated = time.Now()
	}

	o.Update(topic)

	if oldcate != topic.Category {
		cate := new(Category)
		qs := o.QueryTable("category")
		err = qs.Filter("title", oldcate).One(cate)
		if err != nil {
			return err
		}
		cate.TopicCount--
		o.Update(cate)

		cate = new(Category)
		qs = o.QueryTable("category")
		err = qs.Filter("title", topic.Category).One(cate)
		if err != nil {
			return err
		}
		cate.TopicCount++
		o.Update(cate)
	}
	return nil

}
func ViewTopic(id string) (*Topic, error) {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	topic := new(Topic)

	qs := o.QueryTable("topic")
	err = qs.Filter("id", cid).One(topic)
	if err != nil {
		return nil, err
	}
	topic.Views++
	_, err = o.Update(topic)
	return topic, nil
}
