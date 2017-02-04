package plugin

import (
	"errors"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	plugin "github.com/hashicorp/go-plugin"
)

type PluginInfo struct {
	ID     string
	Path   string
	Client *plugin.Client
}

func NewManager(ptype, glob, dir string, pluginImpl plugin.Plugin) *Manager {

	manager := &Manager{
		Type:       ptype,
		Glob:       glob,
		Path:       dir,
		Plugins:    map[string]*PluginInfo{},
		pluginImpl: pluginImpl,
	}

	return manager
}

// Manager handles lifecycle of plugin mgmt for different plugin types. In This
// example we have two plugin types: greeter and clubber, both of which have
// multiple implementations.
type Manager struct {
	Type        string                 // id for types of plugins this manager deals with
	Glob        string                 // glob match for plugin filenames
	Path        string                 // path for plugins
	Plugins     map[string]*PluginInfo // Info for foudn plugins
	initialized bool                   // has been initialized
	pluginImpl  plugin.Plugin          // Plugin implementation dummy interface
}

func (m *Manager) Init() error {

	// discover plugin abs paths
	plugins, err := plugin.Discover(m.Glob, m.Path)
	if err != nil {
		return err
	}

	// grab all PluginInfos
	for _, plugin := range plugins {

		// use glob trailing * to get us a friendly plugin id name
		_, file := filepath.Split(plugin)
		globAsterix := strings.LastIndex(m.Glob, "*")
		trim := m.Glob[0:globAsterix]
		id := strings.TrimPrefix(file, trim)

		// add to our slice of plugin info
		m.Plugins[id] = &PluginInfo{
			ID:   id,
			Path: plugin,
		}

	}

	m.initialized = true

	return nil
}

func (m *Manager) Launch() error {

	for id, info := range m.Plugins {

		log.Printf("Registering plugin client for type=%s, id=%s, impl=%s", m.Type, id, info.Path)
		// create new client
		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: HandshakeConfig,
			Plugins:         m.pluginMap(id),
			Cmd:             exec.Command(info.Path),
		})

		if _, ok := m.Plugins[id]; !ok {
			// if not found, ignore?
			continue
		}
		pinfo := m.Plugins[id]
		pinfo.Client = client

	}

	return nil
}

func (m *Manager) Dispose() {
	var wg sync.WaitGroup
	for _, pinfo := range m.Plugins {
		wg.Add(1)

		go func(client *plugin.Client) {
			client.Kill()
			wg.Done()
		}(pinfo.Client)
	}

	wg.Wait()

}

func (m *Manager) GetInterface(id string) (interface{}, error) {

	if _, ok := m.Plugins[id]; !ok {
		return nil, errors.New("Plugin ID not found in registered plugins!")
	}

	// Grab registerd plugin.Client
	client := m.Plugins[id].Client

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return nil, err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense(id)
	if err != nil {
		return nil, err
	}

	return raw, nil
}

// pluginMap should be used by clients for the map of plugins.
// This maps types of plugins to interface implementations. Since each
// plugin.Client we have registers one binary, we set a 1:1 mapping between
// plugin type (here treated as a name/id) to its implementation.
func (m *Manager) pluginMap(id string) map[string]plugin.Plugin {
	pmap := map[string]plugin.Plugin{}

	// for _, pinfo := range m.Plugins {
	// 	pmap[pinfo.ID] = m.pluginImpl
	// }

	pmap[id] = m.pluginImpl

	return pmap
}
