package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/wwqdrh/launcher/internal/app"
	"github.com/wwqdrh/launcher/internal/keeper"
	"github.com/wwqdrh/launcher/internal/updater"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "launcher",
		Short: "Application launcher with auto-update and keep-alive features",
	}

	// 初始化配置目录
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	configDir := filepath.Join(homeDir, ".launcher")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		log.Fatal(err)
	}

	appManager := app.NewManager(configDir)
	updaterService := updater.NewUpdater(appManager)
	keeperService := keeper.NewKeeper(appManager)

	// 添加启动命令
	var startCmd = &cobra.Command{
		Use:   "start [app-name]",
		Short: "Start an application",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := appManager.StartApp(args[0]); err != nil {
				log.Fatal(err)
			}
		},
	}

	// 添加配置命令
	var configCmd = &cobra.Command{
		Use:   "config [app-name] [command]",
		Short: "Configure an application",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := appManager.ConfigureApp(args[0], args[1]); err != nil {
				log.Fatal(err)
			}
		},
	}

	// 添加设置版本命令
	var setVersionCmd = &cobra.Command{
		Use:   "setversion [app-name] [version]",
		Short: "Set application version",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			if err := appManager.SetAppVersion(args[0], args[1]); err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(startCmd, configCmd, setVersionCmd)

	// 启动更新检查和保活服务
	go updaterService.Start()
	go keeperService.Start()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
