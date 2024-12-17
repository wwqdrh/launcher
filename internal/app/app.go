package app

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type App struct {
	Name        string `json:"name"`
	Command     string `json:"command"`
	Version     string `json:"version"`
	UpdateURL   string `json:"update_url"`
	DownloadURL string `json:"download_url"`
}

type Manager struct {
	configDir string
	apps      map[string]App
}

func NewManager(configDir string) *Manager {
	m := &Manager{
		configDir: configDir,
		apps:      make(map[string]App),
	}
	m.loadConfig()
	return m
}

func (m *Manager) loadConfig() error {
	configFile := filepath.Join(m.configDir, "apps.json")
	data, err := os.ReadFile(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &m.apps)
}

func (m *Manager) saveConfig() error {
	configFile := filepath.Join(m.configDir, "apps.json")
	data, err := json.MarshalIndent(m.apps, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, 0644)
}

func (m *Manager) StartApp(name string) error {
	app, ok := m.apps[name]
	if !ok {
		return fmt.Errorf("app %s not found", name)
	}

	cmd := exec.Command("sh", "-c", app.Command)
	return cmd.Start()
}

func (m *Manager) ConfigureApp(name, command string) error {
	m.apps[name] = App{
		Name:    name,
		Command: command,
	}
	return m.saveConfig()
}

func (m *Manager) GetApp(name string) (App, bool) {
	app, ok := m.apps[name]
	return app, ok
}

func (m *Manager) SetAppVersion(name, version string) error {
	app, ok := m.apps[name]
	if !ok {
		return fmt.Errorf("app %s not found", name)
	}

	app.Version = version
	m.apps[name] = app
	return m.saveConfig()
}
