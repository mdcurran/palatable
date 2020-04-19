package cmd

import (
	"net/http"

	"github.com/spf13/cobra"

	"github.com/mdcurran/palatable/internal/pkg/server"
)

var rootCmd = &cobra.Command{
	Use: "palatable",
	Short: "palatable enables reminiscence about food.",
	Run: func(cmd *cobra.Command, args []string) {
		s := server.New()

		s.Logger.Info("starting HTTP server")
		err := http.ListenAndServe(":8080", s.Router)
		if err != nil {
			s.Logger.Panicf("unable to initialise HTTP server: %v", err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
