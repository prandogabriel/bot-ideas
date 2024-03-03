package models

import "encoding/json"

type User struct {
	ChatID   int64      `json:"chat_id"`
	Name     string     `json:"name"`
	IdeaList []IdeaList `json:"idea_list"`
}

type Idea struct {
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type IdeaList struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	HasCompleted bool   `json:"has_completed"`
	Ideas        []Idea `json:"ideas"`
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) FromJSON(data []byte) error {
	return json.Unmarshal(data, u)
}

func (il *IdeaList) ToJSON() ([]byte, error) {
	return json.Marshal(il)
}

func (il *IdeaList) FromJSON(data []byte) error {
	return json.Unmarshal(data, il)
}

func (i *Idea) ToJSON() ([]byte, error) {
	return json.Marshal(i)
}

func (i *Idea) FromJSON(data []byte) error {
	return json.Unmarshal(data, i)
}
