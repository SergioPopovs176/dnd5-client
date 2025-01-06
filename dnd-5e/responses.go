package dnd5e

import "fmt"

type monstersResponse struct {
	Count    int64     `json:"count"`
	Monsters []Monster `json:"results"`
}

type Monster struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	Url   string `json:"url"`
}

func (m Monster) Info() string {
	return fmt.Sprintf("[index] %s | [name] %s | [url] %s",
		m.Index, m.Name, m.Url)
}
