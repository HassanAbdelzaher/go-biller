package migration

import (
	"encoding/xml"
	"errors"
	//"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/HassanAbdelzaher/lama"
	"github.com/MaisrForAdvancedSystems/mas-db-models/dbmodels"

	_ "embed"
)

//go:embed Migration.xml
var migrationFileData []byte

type Migs struct {
	///XMLName xml.Name `xml:"migs"`
	Migs []Mig `xml:"mig"`
}

type Mig struct {
	Version string `xml:"version,attr"`
	Stms    Stms   `xml:"stms"`
}

func (m *Mig) GetVersion() (float64, error) {
	if m.Version == "" {
		return 0, errors.New("invalied mig version")
	}
	f, err := strconv.ParseFloat(m.Version, 64)
	if err != nil {
		return 0, err
	}
	f = math.Round(1000*f) / 1000
	return f, nil
}

type Stms struct {
	Stms []Stm `xml:"stm"`
}

type Stm struct {
	Sql  string `xml:"sql,attr"`
	Skip string `xml:"skip,attr"`
}

func (s *Stm) IsSkip() bool {
	if strings.TrimSpace(s.Skip) == "1" || strings.ToLower(strings.TrimSpace(s.Skip)) == "true" {
		return true
	}
	return false
}

func Migrate(db *lama.Lama) error {
	log.Println("validate database version")
	var sys dbmodels.SYSTEM

	err := db.Model(dbmodels.SYSTEM{}).First(&sys)
	if err != nil {
		return err
	}
	dbVersion := sys.DATABASE_VERSION
	/*data, err := ioutil.ReadFile("Migration.xml")
	if err != nil {
		return err
	}*/

	var migs Migs
	err = xml.Unmarshal(migrationFileData, &migs)
	if err != nil {
		return err
	}
	log.Println("migration loaded")
	if migs.Migs == nil {
		return errors.New("invalied migration file")
	}
	for _, mig := range migs.Migs {
		ver, err := mig.GetVersion()
		if err != nil {
			return err
		}
		if dbVersion >= ver {
			//log.Println("skip ")
			continue
		}
		if len(mig.Stms.Stms) == 0 {
			continue
		}
		for _, stm := range mig.Stms.Stms {
			if stm.Sql == "" {
				continue
			}
			log.Println(stm.Sql)
			err := exec(db, stm.Sql)
			if err != nil {
				if !stm.IsSkip() {
					return err
				}else{
					log.Println("skip")
					log.Println(err)
				}
			}
			arcS := arcStm(stm.Sql)
			if arcS != nil && *arcS != "" {
				log.Println(*arcS)
				err := exec(db, *arcS)
				if err != nil {
					if !stm.IsSkip() {
						return err
					}else{
						log.Println("skip")
						log.Println(err)
					}
				}
			}
		}
		sys.DATABASE_VERSION = ver
		err = db.Save(sys)
		if err != nil {
			return err
		}
	}
	return nil
}

func exec(db *lama.Lama, stm string) error {
	_, err := db.DB.Exec(stm)
	return err
}

func arcStm(stm string) *string {
	if stm == "" {
		return nil
	}
	stm = strings.TrimSpace(strings.ToLower(stm))
	if strings.Contains(stm, "alter ") {
		if strings.Contains(stm, " hand_mh_st ") || strings.Contains(stm, " dbo.hand_mh_st ") {
			arcStm := strings.ReplaceAll(stm, "hand_mh_st", "arc_hand_mh_st")
			return &arcStm
		}
	}

	return nil
}
