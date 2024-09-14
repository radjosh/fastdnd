package dataProcessors

/* todo:
 * server. should we keep the map loader in a separate file?
 */

/*
 * it appears that go is super picky about where files live
 * I couldn't get this to work until I moved it into its own package
 * AND SUBDIRECTORY
 *
 * also I needed to run "go build" in that subdir
 * leaving these sloppy notes here until I can get some intuition for it
 */

/*
 * This one uses the 3rd-party module "etree" for more fast and loose (and better?) xml parsing
 * etree is very forgiving of mismatches between struct and xml source
 */

import (
	// "fmt"
	"strings"

	"github.com/felixge/etree"
	// "utils/etree"
)


type Monster struct {
	Name string
	Size string
	Alignment string
	Ac string
	Hp string
	Speed string 
	Str string
	Dex string
	Con string
	Int string
	Wis string
	Cha string
	Skill string
	Passive string
	Languages string
	Cr string
	Spells string
	Description string
	Traits []map[string]string
	Actions []map[string]string
	Legendary []map[string]string
	Immune string
	Senses string
	Vulnerable string
	Resist string
}

var Monsters = make(map[string]Monster)
func init() {
	doc := etree.NewDocument()
	DATADIR := "../../data/"
	FILENAME := "mm.xml"
	if err := doc.ReadFromFile(DATADIR + FILENAME); err != nil {
		panic(err)
	}

  root := doc.SelectElement("compendium")

  for _, m := range root.SelectElements("monster") {
		monsterPtr := new(Monster)
    for _, e := range m.ChildElements() {
      if e.Tag != "trait" && e.Tag != "action" && e.Tag != "legendary" {
				switch e.Tag {
					case "name":
						monsterPtr.Name = strings.ToLower(strings.Trim(e.Text(), " "))
					case "size":
						monsterPtr.Size = strings.Trim(e.Text(), " ")
					case "alignment":
						monsterPtr.Alignment = strings.Trim(e.Text(), " ")
					case "ac":
						monsterPtr.Ac = strings.Trim(e.Text(), " ")
					case "hp":
						monsterPtr.Hp = strings.Trim(e.Text(), " ")
					case "str":
						monsterPtr.Str = strings.Trim(e.Text(), " ")
					case "dex":
						monsterPtr.Dex = strings.Trim(e.Text(), " ")
					case "con":
						monsterPtr.Con = strings.Trim(e.Text(), " ")
					case "int":
						monsterPtr.Int = strings.Trim(e.Text(), " ")
					case "wis":
						monsterPtr.Wis = strings.Trim(e.Text(), " ")
					case "cha":
						monsterPtr.Cha = strings.Trim(e.Text(), " ")
					case "skill":
						monsterPtr.Skill = strings.Trim(e.Text(), " ")
					case "passive":
						monsterPtr.Passive = strings.Trim(e.Text(), " ")
					case "languages":
						monsterPtr.Languages = strings.Trim(e.Text(), " ")
					case "cr":
						monsterPtr.Cr = strings.Trim(e.Text(), " ")
					case "spells":
						monsterPtr.Spells = strings.Trim(e.Text(), " ")
					case "description":
						monsterPtr.Description = strings.Trim(e.Text(), " ")
					case "immune":
						monsterPtr.Immune = strings.Trim(e.Text(), " ")
					case "senses":
						monsterPtr.Senses = strings.Trim(e.Text(), " ")
					case "vulnerable":
						monsterPtr.Vulnerable = strings.Trim(e.Text(), " ")
					case "resist":
						monsterPtr.Resist = strings.Trim(e.Text(), " ")
				}
      } else if e.Tag == "trait" {
				trait := make(map[string]string)	
        for _, f := range e.ChildElements() {
          switch f.Tag {
            case "name":
							trait["name"] = f.Text()
            case "text":
							trait["text"] = f.Text()
            case "attack":
							trait["attack"] = "Attack: " + f.Text()
          }
        }
				monsterPtr.Traits = append(monsterPtr.Traits, trait)
      } else if e.Tag == "action" {
				action := make(map[string]string)	
        for _, f := range e.ChildElements() {
          switch f.Tag {
            case "name":
							action["name"] = f.Text()
            case "text":
							action["text"] = f.Text()
            case "attack":
							action["attack"] = "Attack: " + f.Text()
          }
        }
				monsterPtr.Actions = append(monsterPtr.Actions, action)
      } else if e.Tag == "legendary" {
				legendary := make(map[string]string)	
        for _, f := range e.ChildElements() {
          switch f.Tag {
            case "name":
							legendary["name"] = f.Text()
            case "text":
							legendary["text"] = f.Text()
            case "attack":
							legendary["attack"] = "Attack: " + f.Text()
          }
        }
				monsterPtr.Legendary = append(monsterPtr.Legendary, legendary)
      }
		Monsters[monsterPtr.Name] = *monsterPtr
  }
}
}