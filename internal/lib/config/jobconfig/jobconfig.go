package jobconfig

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/EagleLizard/ezd-daemon/internal/lib/constants"
	"github.com/EagleLizard/ezd-daemon/internal/lib/model/jobdef"
)

var JobDefs map[string]jobdef.JobDef = make(map[string]jobdef.JobDef)

func init() {
	fmt.Println("jobconfig")
	loadDefs()
}

func GetDef(repoName string) *jobdef.JobDef {
	if jd, ok := JobDefs[repoName]; ok {
		return &jd
	}
	return nil
}

func loadDefs() {
	bd := constants.BaseDir()
	cfgDir := filepath.Join(bd, "hook-configs")
	dirEnts, err := os.ReadDir(cfgDir)
	if err != nil {
		log.Fatalf("%v", err)
	}
	for _, d := range dirEnts {
		fmt.Printf("%s\n", d.Name())
		if strings.HasSuffix(d.Name(), ".toml") {
			fp := filepath.Join(cfgDir, d.Name())
			tomlData, err := os.ReadFile(fp)
			if err != nil {
				log.Fatalf("%v", err)
			}
			jd := jobdef.JobDef{}
			toml.Decode(string(tomlData), &jd)
			fmt.Printf("%+v\n", jd)
			if _, ok := JobDefs[jd.Repo.Name]; ok {
				panic(fmt.Errorf("JD_0.1: Job definition for repo %s already defined [duplicate]", jd.Repo.Name))
			}
			JobDefs[jd.Repo.Name] = jd
		}
	}
}
