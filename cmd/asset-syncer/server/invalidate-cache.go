/*
Copyright 2021 VMware. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"fmt"

	"github.com/kubeapps/kubeapps/pkg/dbutils"
	log "k8s.io/klog/v2"
)

func InvalidateCache(serveOpts Config, args []string) error {
	if len(args) != 0 {
		return fmt.Errorf("This command does not take any arguments (got %v)", len(args))
	}

	dbConfig := dbutils.Config{URL: serveOpts.DatabaseURL, Database: serveOpts.DatabaseName, Username: serveOpts.DatabaseUser, Password: serveOpts.DatabasePassword}
	kubeappsNamespace := serveOpts.KubeappsNamespace
	manager, err := newManager(dbConfig, kubeappsNamespace)
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	err = manager.Init()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	defer manager.Close()

	err = manager.InvalidateCache()
	if err != nil {
		return fmt.Errorf("Error: %v", err)
	}
	log.Infof("Successfully invalidated cache")
	return nil
}
