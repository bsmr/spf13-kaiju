// Copyright © 2014 Steve Francia <spf@spf13.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License./

package commands

import (
    "fmt"
    "github.com/spf13/cobra"
    "os"
    "net/http"
    "strconv"
    "strings"
    "github.com/codegangsta/martini"
)

var Root = &cobra.Command{
    Use:   "pandora",
    Short: "Pandora is an open source comment server",
    Long:  `Pandora is an open source comment server`,
    Run: RootRun, 
}

func Execute() {
    err := Root.Execute()
    if err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }

}

func RootRun(cmd *cobra.Command, args []string) {
    m := martini.New()
	r := martini.NewRouter()

    r.Get("/", index)
    r.Get("/comments/:forum/:post", comments)

	m.Action(r.Handle)

    http.ListenAndServe(":" + strconv.Itoa(Port), m)
}

var Verbose bool
var Port int

func init() {
    Root.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
    Root.Flags().IntVarP(&Port, "port", "p", 2714, "port number to run on")
}

func index() string {
    return "What up?"
}

func comments(parms martini.Params) (int, string) {
	forum := parms["forum"]
	post := parms["post"]
    return http.StatusOK, strings.Join([]string {"ah, yeah: ", forum, post} , " ")
}
