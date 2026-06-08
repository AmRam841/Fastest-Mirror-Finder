package cli

import (
	"fmt"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/v6/progress"

	"github.com/jedib0t/go-pretty/v6/table"
	//	"github.com/jedib0t/go-pretty/v6/table"
	// "github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/list"
)

// TODO Make a Cli
// TODO Read the incoming arguments.
// TODO Decide what the first real word means.
// TODO Route to the matching command handler.
// TODO Let that handler inspect the remaining words.
// TODO Parse flags and values.
// TODO Run the logic.
// TODO Print results or errors.



func Arg_reader(){
	OsArgs := os.Args[1:]
	println(&OsArgs)
}


func Help_printer(){
	fmt.Printf("")
}



func Tables() {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)
    t.AppendHeader(table.Row{" ", "Name", "Job", "Salary"})
    t.AppendRows([]table.Row{
        {1, "Alice", "Developer", 60000},
        {2, "Bob", "Designer", 55000},
    })
    t.AppendFooter(table.Row{" ", " ", "Total", 115000})
    t.SetStyle(table.StyleLight)
    t.Render()
}


func Progresses() {
    pw := progress.NewWriter()
    pw.SetOutputWriter(os.Stdout)

    tracker1 := pw.AppendTracker(progress.Tracker{Name: "File1.zip", Total: 100})
    tracker2 := pw.AppendTracker(progress.Tracker{Name: "File2.zip", Total: 100})

    go pw.Render()

    for i := 0; i < 100; i += 10 {
        time.Sleep(500 * time.Millisecond)
        tracker1.Increment(10)
        tracker2.Increment(10)
    }

    time.Sleep(time.Second) // Allow for rendering completion
}

func Lists() {
    l := list.NewWriter()
    l.SetOutputMirror(os.Stdout)

    l.AppendItem("Project Kickoff")
    l.Indent()
    l.AppendItem("Plan")
    l.AppendItem("Code")
    l.Indent()
    l.AppendItem("Frontend")
    l.AppendItem("Backend")
    l.UnIndent()
    l.AppendItem("Launch")
    l.UnIndent()

    l.SetStyle(list.StyleConnectedLight)
    l.Render()
}