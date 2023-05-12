package cmd

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/guidomantilla/go-feather-api-sample/cmd/serve"
)

func ExecuteAppCmd() {

	appCmd := &cobra.Command{}
	appCmd.AddCommand(createServeCmd(), createMigrateCmd(), createTestCmd())

	if err := appCmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}

func createServeCmd() *cobra.Command {

	serveCmd := &cobra.Command{
		Use: "serve",
		Run: serve.ExecuteCmdFn,
	}

	return serveCmd
}

func createMigrateCmd() *cobra.Command {

	migrateUpCmd := &cobra.Command{
		Use: "up",
		//Run: migrate.UpCmdFn,
	}

	migrateDownCmd := &cobra.Command{
		Use: "down",
		//Run: migrate.DownCmdFn,
	}

	migrateCmd := &cobra.Command{
		Use: "migrate",
	}

	migrateCmd.AddCommand(migrateUpCmd, migrateDownCmd)

	return migrateCmd
}

func createTestCmd() *cobra.Command {

	testCmd := &cobra.Command{
		Use: "test",
		//Run: test.ExecuteCmdFn,
	}

	return testCmd
}
