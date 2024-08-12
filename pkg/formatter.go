package pkg

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
	"time"

	model "github.com/hazaloolu/git_get/models"
)

func PrintData(data model.JSONData) {
	log.Printf("Repo found: %d", data.Count)
	const format = "%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprint(tw, format, "Repo ", "Stars ", "Created at ", "Description")
	fmt.Fprint(tw, format, "------", "--------", "---------", "-------")

	for _, i := range data.Items {
		desc := i.Description
		if len(desc) > 50 {
			desc = string(desc[:50]) + "...."
		}

		t, err := time.Parse(time.RFC3339, i.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(tw, format, i.FullName, i.StargazersCount, t.Year(), desc)
	}
	tw.Flush()
}
