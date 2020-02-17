package main

import (
	"context"
	"io/ioutil"
	"runtime/trace"
	"sync"

	"github.com/pkg/profile"
)

func main() {

	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	ctx, task := trace.NewTask(context.Background(), "main start")

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		r := trace.StartRegion(ctx, "reading file")
		defer r.End()
		trace.Log(ctx, "category", "I/O file")
		ioutil.ReadFile(`n1.txt`)
	}()

	go func() {
		defer wg.Done()
		r := trace.StartRegion(ctx, "writing file")
		defer r.End()

		trace.Log(ctx, "goroutine", "2")
		ioutil.WriteFile(`n2.txt`, []byte(`42`), 0644)
	}()

	wg.Wait()

	defer task.End()
}
