package models

import (
	"encoding/json"
	"errors"
	"strconv"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/syndtr/goleveldb/leveldb"
)

var urlCountGuard sync.Mutex
var urlCount int64

func init() {
	appDataDb, err := leveldb.OpenFile(Configs["conf/db.conf"].String("app::path"), nil)
	if err != nil {
		beego.Critical("unable to open " + Configs["conf/db.conf"].String("app::path"))
	}
	defer appDataDb.Close()

	urlCountStr, err := appDataDb.Get([]byte("urlCount"), nil)
	if err != nil {
		urlCount = int64(1014)
	} else {
		urlCount, _ = strconv.ParseInt(string(urlCountStr), 10, 64)
	}
}

type NoteEntry struct {
	content string
	author  string
	date    string
	url     string
	salt    string
}

type JsonNoteEntry struct {
	Content string
	Author  string
	Date    string
	Url     string
	Salt    string
}

func (n *NoteEntry) toPublic() *JsonNoteEntry {
	return &JsonNoteEntry{
		Content: n.content,
		Author:  n.author,
		Date:    n.date,
		Url:     n.url,
		Salt:    n.salt,
	}
}

func (j *JsonNoteEntry) toPrivate() *NoteEntry {
	return &NoteEntry{
		content: j.Content,
		author:  j.Author,
		date:    j.Date,
		url:     j.Url,
		salt:    j.Salt,
	}
}

func GetNoteEntryByUrl(url string) (*NoteEntry, error) {
	urlDb, err := leveldb.OpenFile(Configs["conf/db.conf"].String("url::path"), nil)
	if err != nil {
		beego.Critical("unable to open " + Configs["conf/db.conf"].String("url::path"))
		return nil, errors.New("unable to open database")
	}
	defer urlDb.Close()

	if buf, err := urlDb.Get([]byte(url), nil); err != nil {
		return nil, err
	} else {
		ne := &JsonNoteEntry{}
		if err := json.Unmarshal(buf, &ne); err != nil {
			beego.Error("unable to unmarshal" + string(buf))
			return nil, errors.New("unable to unmarshal")
		} else {
			return ne.toPrivate(), nil
		}
	}

	return nil, nil
}

func NewNoteEntry(author string, content string) (*NoteEntry, error) {
	ne := &NoteEntry{}
	ne.content = content
	ne.author = author
	ne.date = time.Now().Format("2006-01-02 15:04:05")
	ne.salt = GenerateRandomStr62(beego.AppConfig.DefaultInt("saltlen", 4))

	urlCountGuard.Lock()
	defer urlCountGuard.Unlock()

	ne.url = Int64ToStr62(urlCount) + ne.salt
	urlCount++

	if ne.save() {
		return ne, nil
	} else {
		return nil, errors.New("unable to save")
	}
	return nil, nil
}

func (n *NoteEntry) save() bool {
	appDataDb, err := leveldb.OpenFile(Configs["conf/db.conf"].String("app::path"), nil)
	if err != nil {
		beego.Error(err)
		beego.Critical("unable to open " + Configs["conf/db.conf"].String("app::path"))
		return false
	}
	defer appDataDb.Close()

	tmp := int64(1014)
	if urlCountStr, err := appDataDb.Get([]byte("urlCount"), nil); err == nil {
		tmp, _ = strconv.ParseInt(string(urlCountStr), 10, 64)
	}

	if tmp < urlCount && appDataDb.Put([]byte("urlCount"), []byte(strconv.FormatInt(urlCount, 10)), nil) != nil {
		beego.Critical("unable to save to db")
		return false
	}

	urlDb, err := leveldb.OpenFile(Configs["conf/db.conf"].String("url::path"), nil)
	if err != nil {
		beego.Critical("unable to open " + Configs["conf/db.conf"].String("url::path"))
		return false
	}
	defer urlDb.Close()

	if buf, err := json.Marshal(n.toPublic()); err != nil {
		beego.Error("unable to Marshal object", n)
		return false
	} else {
		urlDb.Put([]byte(n.url), buf, nil)
	}

	return true
}

func (n *NoteEntry) GetUrl() string {
	return n.url
}
func (n *NoteEntry) GetContent() string {
	return n.content
}
func (n *NoteEntry) GetDate() string {
	return n.date
}
