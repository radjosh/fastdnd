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

	"utils/etree"
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
      if e.Tag != "trait" && e.Tag != "action" {
				switch e.Tag {
					case "name":
						monsterPtr.Name = strings.Trim(e.Text(), " ")
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
				}
      //} else {
        //for _, f := range e.ChildElements() {
          //switch f.Tag {
            //case "name":
            //fmt.Println(e.Tag, f.Text()) // e is not a typo
            //case "text":
              //fmt.Println(" ", f.Text())
            //case "attack":
              //fmt.Println(" ", f.Tag, f.Text())
          //}
        //}
			//}
			}
    }
		Monsters[monsterPtr.Name] = *monsterPtr
  }
	//for _, monster := range(Monsters) {
		//fmt.Println(monster)
	//}

	// fmt.Println(Monsters["Werebear"])
}
