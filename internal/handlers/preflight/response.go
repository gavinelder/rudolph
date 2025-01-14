package preflight

import (
	"github.com/airbnb/rudolph/pkg/model/machineconfiguration"
	"github.com/airbnb/rudolph/pkg/types"
)

// PreflightResponse represents sync response returned to a Santa client by the sync server.
//
// WARNING: The PreflightResponse copies its format directly from the database; make sure this struct's fields
//
//	are consistent with the fields of the store.MachineConfiguration type
//
// Use Santa defined constants
// https://github.com/google/santa/blob/main/Source/santactl/Commands/sync/SNTCommandSyncConstants.m#L32-L35
type PreflightResponse struct {
	ClientMode             types.ClientMode `json:"client_mode"`
	BlockedPathRegex       string           `json:"blocked_path_regex"`
	AllowedPathRegex       string           `json:"allowed_path_regex"`
	BatchSize              int              `json:"batch_size"`
	EnableBundles          bool             `json:"enable_bundles"`
	EnabledTransitiveRules bool             `json:"enable_transitive_rules"`
	CleanSync              bool             `json:"clean_sync,omitempty"`
	FullSyncInterval       int              `json:"full_sync_interval,omitempty"`
	UploadLogsURL          string           `json:"upload_logs_url,omitempty"`
}

// ConstructPreflightResponse converts a MachineConfiguration pulled from the database into the corresponding
// response to be return as an API response.
func ConstructPreflightResponse(machineConfiguration machineconfiguration.MachineConfiguration, cleanSync bool) *PreflightResponse {
	return &PreflightResponse{
		ClientMode:             machineConfiguration.ClientMode,
		BlockedPathRegex:       machineConfiguration.BlockedPathRegex,
		AllowedPathRegex:       machineConfiguration.AllowedPathRegex,
		BatchSize:              machineConfiguration.BatchSize,
		EnableBundles:          machineConfiguration.EnableBundles,
		EnabledTransitiveRules: machineConfiguration.EnabledTransitiveRules,
		UploadLogsURL:          machineConfiguration.UploadLogsURL,
		FullSyncInterval:       machineConfiguration.FullSyncInterval,
		CleanSync:              cleanSync,
		// Notably, we do not grab the clean sync from the config
		// FYI: Sending down a clean sync to the client instructs it to erase all rules, so be careful!
	}

	// return (*PreflightResponse)(&machineConfiguration)
}
