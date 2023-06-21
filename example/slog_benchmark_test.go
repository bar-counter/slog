package example

import (
	"fmt"
	"github.com/bar-counter/slog"
	"github.com/bar-counter/slog/lager"
	"testing"
)

func BenchmarkSlogStdout(b *testing.B) {
	// mock SlogStdout
	lagerDefinition := slog.DefaultLagerDefinition()
	lagerDefinition.LogHideLineno = true
	err := slog.InitWithConfig(lagerDefinition)
	if err != nil {
		b.Fatal(err)
	}
	// reset counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// do SlogStdout
		slog.Infof("Hi %s, system is starting up ...", "paas-bot")
		slog.Info("check-info", lager.Data{
			"info": "something",
		})

		slog.Debug("check-info", lager.Data{
			"info": "something",
		})

		slog.Warn("failed-to-do-something", lager.Data{
			"info": "something",
		})

		err = fmt.Errorf("this is an error")
		slog.Error("failed-to-do-something", err)
		slog.Info("shutting-down")
	}
	b.StopTimer()
}

func BenchmarkParallelSlogStdout(b *testing.B) {
	// mock Parallel SlogStdout
	lagerDefinition := slog.DefaultLagerDefinition()
	lagerDefinition.LogHideLineno = true
	err := slog.InitWithConfig(lagerDefinition)
	if err != nil {
		b.Fatal(err)
	}

	// reset counter
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// do Parallel SlogStdout

			slog.Infof("Hi %s, system is starting up ...", "paas-bot")
			slog.Info("check-info", lager.Data{
				"info": "something",
			})

			slog.Debug("check-info", lager.Data{
				"info": "something",
			})

			slog.Warn("failed-to-do-something", lager.Data{
				"info": "something",
			})

			err = fmt.Errorf("this is an error")
			slog.Error("failed-to-do-something", err)
			slog.Info("shutting-down")
		}
	})
	b.StopTimer()
}
