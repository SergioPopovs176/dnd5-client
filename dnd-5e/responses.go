package dnd5e

import "fmt"

type monstersResponse struct {
	Count    int64         `json:"count"`
	Monsters []MonsterShot `json:"results"`
}

type MonsterShot struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

func (m MonsterShot) Info() string {
	return fmt.Sprintf("[index] %s | [name] %s | [url] %s",
		m.Index, m.Name, m.Url)
}

type MonsterFullResponse struct {
	Index     string `json:"index"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Size      string `json:"size"`
	Type      string `json:"type"`
	Subtype   string `json:"subtype"`
	Alignment string `json:"alignment"`
}

func (m MonsterFullResponse) Info() string {
	return fmt.Sprintf("[index] %s | [name] %s | [url] %s | [size] %s | [type] %s | [subtype] %s | [alignment] %s",
		m.Index, m.Name, m.Url, m.Size, m.Type, m.Subtype, m.Alignment)
}
