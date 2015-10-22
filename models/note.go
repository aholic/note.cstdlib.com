package models

import "time"

type NoteEntry struct {
	content string
	author  string
	date    string
	url     string
	salt    string
}

func NewNoteEntry(author string, content string) *NoteEntry {
	ne := &NoteEntry{}
	ne.content = content
	ne.author = author
	ne.date = time.Now().Format("2006-01-02 15:04:05")
	return ne
}

func (n *NoteEntry) Save() bool {
	return true
}

func (n *NoteEntry) GetUrl() string {
	return n.url
}
