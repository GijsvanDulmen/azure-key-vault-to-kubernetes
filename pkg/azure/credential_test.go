/*
Copyright Sparebanken Vest

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

package azure

import (
	"os"
	"testing"

	"github.com/Azure/go-autorest/autorest/azure"
)

// func TestAuthDefault(t *testing.T) {
// 	creds, err := NewAzureKeyVaultCredentialsFromEnvironment()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	_, err = creds.
// 	if err == nil {
// 		t.Fail()
// 	}
// }

func ensureIntegrationEnvironment(t *testing.T) {
	if os.Getenv("AKV2K8S_CLIENT_ID") == "" {
		t.Skip("Skipping integration test - no credentials")
	}

	os.Setenv("AZURE_CLIENT_ID", os.Getenv("AKV2K8S_CLIENT_ID"))
	os.Setenv("AZURE_CLIENT_SECRET", os.Getenv("AKV2K8S_CLIENT_SECRET"))
	os.Setenv("AZURE_TENANT_ID", os.Getenv("AKV2K8S_CLIENT_TENANT_ID"))
}

// func TestChinaCloud(t *testing.T) {
// 	ensureIntegrationEnvironment(t)

// 	os.Setenv("AZURE_ENVIRONMENT", "AzureChinaCloud")

// 	creds, err := NewAzureKeyVaultCredentialsFromEnvironment()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	token := creds.(*azureKeyVaultCredentials).Token
// 	err = token.Refresh()
// 	if err != nil {
// 		t.Error(err)
// 	}

// 	if token.Token().Resource != azure.ChinaCloud.ResourceIdentifiers.KeyVault {
// 		t.Errorf("Endpoint incorrect. Expected '%s', but got '%s'", azure.ChinaCloud.ResourceIdentifiers.KeyVault, token.Token().Resource)
// 	}
// }

func TestAudience(t *testing.T) {
	ensureIntegrationEnvironment(t)

	creds, err := NewFromEnvironment()
	if err != nil {
		t.Error(err)
	}

	token := creds.(*credentials).Token
	err = token.Refresh()
	if err != nil {
		t.Error(err)
	}

	if token.Token().Resource != azure.PublicCloud.ResourceIdentifiers.KeyVault {
		t.Error()
	}
}
