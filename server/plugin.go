package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/public/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

func (p *Plugin) changeUserNotificationPreferences() {
	users, _ := p.API.GetUsers(&model.UserGetOptions{})
	for _, user := range users {
		prefs := []model.Preference{
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     "desktop_threads",
				Value:    "all",
			},
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     "email_threads",
				Value:    "all",
			},
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     "push",
				Value:    "all",
			},
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     "push_status",
				Value:    "online",
			},
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     "push_threads",
				Value:    "all",
			},
			{
				UserId:   user.Id,
				Category: model.PreferenceCategoryNotifications,
				Name:     model.PreferenceNameEmailInterval, // Adjusted to match provided path
				Value:    "30",                              // Send email notifications immediately
			},
		}
		p.API.UpdatePreferencesForUser(user.Id, prefs)
	}
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
