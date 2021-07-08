package helmclient

import (
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/cli"
	"helm.sh/helm/v3/pkg/getter"
	"helm.sh/helm/v3/pkg/repo"
)

func (c *HelmClient) GetSetings() *cli.EnvSettings {
	return c.Settings
}

func (c *HelmClient) GetProviders() getter.Providers {
	return c.Providers
}

func (c *HelmClient) GetStorage() *repo.File {
	return c.storage
}

func (c *HelmClient) GetActionConfig() *action.Configuration {
	return c.ActionConfig
}

func (c *HelmClient) GetLinting() bool {
	return c.linting
}