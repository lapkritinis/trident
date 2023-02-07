// Copyright 2022 NetApp, Inc. All Rights Reserved.

package ontap

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"testing"

	"github.com/RoaringBitmap/roaring"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	mockapi "github.com/netapp/trident/mocks/mock_storage_drivers/mock_ontap"
	"github.com/netapp/trident/storage"
	sa "github.com/netapp/trident/storage_attribute"
	"github.com/netapp/trident/utils"

	tridentconfig "github.com/netapp/trident/config"
	. "github.com/netapp/trident/logging"
	drivers "github.com/netapp/trident/storage_drivers"
	"github.com/netapp/trident/storage_drivers/ontap/api"
)

// ToStringPointer takes a string and returns a string pointer
func ToStringPointer(s string) *string {
	return &s
}

func NewTestLUNHelper(storagePrefix string, driverContext tridentconfig.DriverContext) *LUNHelper {
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "iscsi_vs",
	    "aggregate":         "aggr1",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "%v",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}
	`, storagePrefix)
	// parse commonConfigJSON into a CommonStorageDriverConfig object
	commonConfig, err := drivers.ValidateCommonSettings(context.Background(), commonConfigJSON)
	if err != nil {
		Log().Errorf("could not decode JSON configuration: %v", err)
		return nil
	}
	config := &drivers.OntapStorageDriverConfig{}
	config.CommonStorageDriverConfig = commonConfig
	helper := NewLUNHelper(*config, driverContext)
	return helper
}

func TestSnapshotNames_DockerContext(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)

	snapName1 := helper.getInternalSnapshotName("snapshot-123")
	assert.Equal(t, "_snapshot_snapshot_123", snapName1, "Strings not equal")

	snapName2 := helper.getInternalSnapshotName("snapshot")
	assert.Equal(t, "_snapshot_snapshot", snapName2, "Strings not equal")

	snapName3 := helper.getInternalSnapshotName("_snapshot")
	assert.Equal(t, "_snapshot__snapshot", snapName3, "Strings not equal")

	snapName4 := helper.getInternalSnapshotName("_____snapshot")
	assert.Equal(t, "_snapshot______snapshot", snapName4, "Strings not equal")

	k8sSnapName1 := helper.getInternalSnapshotName("snapshot-0bf1ec69_da4b_11e9_bd10_000c29e763d8")
	assert.Equal(t, "_snapshot_snapshot_0bf1ec69_da4b_11e9_bd10_000c29e763d8", k8sSnapName1, "Strings not equal")
}

func TestSnapshotNames_KubernetesContext(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)

	k8sSnapName1 := helper.getInternalSnapshotName("snapshot-0bf1ec69_da4b_11e9_bd10_000c29e763d8")
	assert.Equal(t, "_snapshot_snapshot_0bf1ec69_da4b_11e9_bd10_000c29e763d8", k8sSnapName1, "Strings not equal")

	k8sSnapName2 := helper.getInternalSnapshotName("mySnap")
	assert.Equal(t, "_snapshot_mySnap", k8sSnapName2, "Strings not equal")
}

func TestHelperGetters(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)

	snapPathPattern := helper.GetSnapPathPattern("my-Bucket")
	assert.Equal(t, "/vol/my_Bucket/storagePrefix_*_snapshot_*", snapPathPattern, "Strings not equal")

	snapPathPatternForVolume := helper.GetSnapPathPatternForVolume("my-Vol")
	assert.Equal(t, "/vol/*/storagePrefix_my_Vol_snapshot_*", snapPathPatternForVolume, "Strings not equal")

	snapPath := helper.GetSnapPath("my-Bucket", "storagePrefix_my-Lun", "snap-1")
	assert.Equal(t, "/vol/my_Bucket/storagePrefix_my_Lun_snapshot_snap_1", snapPath, "Strings not equal")

	snapName1 := helper.GetSnapshotName("storagePrefix_my-Lun", "my-Snapshot")
	assert.Equal(t, "storagePrefix_my_Lun_snapshot_my_Snapshot", snapName1, "Strings not equal")

	snapName2 := helper.GetSnapshotName("my-Lun", "snapshot-123")
	assert.Equal(t, "my_Lun_snapshot_snapshot_123", snapName2, "Strings not equal")

	internalVolName := helper.GetInternalVolumeName("my-Lun")
	assert.Equal(t, "storagePrefix_my_Lun", internalVolName, "Strings not equal")

	lunPath := helper.GetLUNPath("my-Bucket", "my-Lun")
	assert.Equal(t, "/vol/my_Bucket/storagePrefix_my_Lun", lunPath, "Strings not equal")

	lunName := helper.GetInternalVolumeNameFromPath(lunPath)
	assert.Equal(t, "storagePrefix_my_Lun", lunName, "Strings not equal")

	lunPathPatternForVolume := helper.GetLUNPathPattern("my-Vol")
	assert.Equal(t, "/vol/*/storagePrefix_my_Vol", lunPathPatternForVolume, "Strings not equal")
}

func TestValidateLUN(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)

	isValid := helper.IsValidSnapLUNPath("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.True(t, isValid, "boolean not true")
}

func TestGetComponents_DockerContext(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)
	snapName := helper.GetSnapshotNameFromSnapLUNPath("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "mysnap", snapName, "Strings not equal")

	volName := helper.GetExternalVolumeNameFromPath("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "myLun", volName, "Strings not equal")

	bucketName := helper.GetBucketName("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "myBucket", bucketName, "Strings not equal")
}

func TestGetComponents_KubernetesContext(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	snapName := helper.GetSnapshotNameFromSnapLUNPath("/vol/myBucket/storagePrefix_myLun_snapshot_snapshot_123")
	assert.Equal(t, "snapshot_123", snapName, "Strings not equal")

	snapName2 := helper.GetSnapshotNameFromSnapLUNPath("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "mysnap", snapName2, "Strings not equal")

	volName := helper.GetExternalVolumeNameFromPath("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "myLun", volName, "Strings not equal")

	bucketName := helper.GetBucketName("/vol/myBucket/storagePrefix_myLun_snapshot_mysnap")
	assert.Equal(t, "myBucket", bucketName, "Strings not equal")
}

func TestGetComponentsNoSnapshot(t *testing.T) {
	helper := NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)
	snapName := helper.GetSnapshotNameFromSnapLUNPath("/vol/myBucket/storagePrefix_myLun")
	assert.Equal(t, "", snapName, "Strings not equal")

	volName := helper.GetExternalVolumeNameFromPath("/vol/myBucket/storagePrefix_myLun")
	assert.Equal(t, "myLun", volName, "Strings not equal")

	bucketName := helper.GetBucketName("/vol/myBucket/storagePrefix_myLun")
	assert.Equal(t, "myBucket", bucketName, "Strings not equal")

	snapName2 := helper.GetSnapshotNameFromSnapLUNPath("/vol/myBucket/storagePrefix_myLun")
	assert.Equal(t, "", snapName2, "Strings not equal")

	volName2 := helper.GetExternalVolumeNameFromPath("myBucket/storagePrefix_myLun")
	assert.NotEqual(t, "myLun", volName2, "Strings are equal")
	assert.Equal(t, "", volName2, "Strings are NOT equal")
}

func newTestOntapSanEcoDriver(
	vserverAdminHost, vserverAdminPort, vserverAggrName string, useREST bool, apiOverride api.OntapAPI,
) *SANEconomyStorageDriver {
	config := &drivers.OntapStorageDriverConfig{}
	sp := func(s string) *string { return &s }

	config.CommonStorageDriverConfig = &drivers.CommonStorageDriverConfig{}
	config.CommonStorageDriverConfig.DebugTraceFlags = make(map[string]bool)
	config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	config.CommonStorageDriverConfig.DebugTraceFlags["api"] = true

	config.ManagementLIF = vserverAdminHost + ":" + vserverAdminPort
	config.SVM = "SVM1"
	config.Aggregate = vserverAggrName
	config.Username = "ontap-san-economy-user"
	config.Password = "password1!"
	config.StorageDriverName = "ontap-san-economy"
	config.StoragePrefix = sp("test_")
	config.UseREST = useREST

	sanEcoDriver := &SANEconomyStorageDriver{}
	sanEcoDriver.Config = *config

	numRecords := api.DefaultZapiRecords
	if config.DriverContext == tridentconfig.ContextDocker {
		numRecords = api.MaxZapiRecords
	}

	var ontapAPI api.OntapAPI

	if apiOverride != nil {
		ontapAPI = apiOverride
	} else {
		if config.UseREST {
			ontapAPI, _ = api.NewRestClientFromOntapConfig(context.TODO(), config)
		} else {
			ontapAPI, _ = api.NewZAPIClientFromOntapConfig(context.TODO(), config, numRecords)
		}
	}

	sanEcoDriver.API = ontapAPI
	sanEcoDriver.telemetry = &Telemetry{
		Plugin:        sanEcoDriver.Name(),
		SVM:           "SVM1",
		StoragePrefix: *sanEcoDriver.GetConfig().StoragePrefix,
		Driver:        sanEcoDriver,
	}

	return sanEcoDriver
}

func newMockOntapSanEcoDriver(t *testing.T) (*mockapi.MockOntapAPI, *SANEconomyStorageDriver) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	driver := newTestOntapSanEcoDriver(vserverAdminHost, vserverAdminPort, vserverAggrName, false, mockAPI)
	return mockAPI, driver
}

func TestOntapSanEcoStorageDriverConfigString(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	sanEcoDrivers := []SANEconomyStorageDriver{
		*newTestOntapSanEcoDriver(vserverAdminHost, vserverAdminPort, vserverAggrName, false, mockAPI),
	}

	sensitiveIncludeList := map[string]string{
		"username":        "ontap-san-economy-user",
		"password":        "password1!",
		"client username": "client_username",
		"client password": "client_password",
	}

	externalIncludeList := map[string]string{
		"<REDACTED>":                   "<REDACTED>",
		"username":                     "Username:<REDACTED>",
		"password":                     "Password:<REDACTED>",
		"api":                          "API:<REDACTED>",
		"chap username":                "ChapUsername:<REDACTED>",
		"chap initiator secret":        "ChapInitiatorSecret:<REDACTED>",
		"chap target username":         "ChapTargetUsername:<REDACTED>",
		"chap target initiator secret": "ChapTargetInitiatorSecret:<REDACTED>",
		"client private key":           "ClientPrivateKey:<REDACTED>",
	}

	for _, sanEcoDriver := range sanEcoDrivers {
		for key, val := range externalIncludeList {
			assert.Contains(t, sanEcoDriver.String(), val,
				"ontap-san-economy driver does not contain %v", key)
			assert.Contains(t, sanEcoDriver.GoString(), val,
				"ontap-san-economy driver does not contain %v", key)
		}

		for key, val := range sensitiveIncludeList {
			assert.NotContains(t, sanEcoDriver.String(), val,
				"ontap-san-economy driver contains %v", key)
			assert.NotContains(t, sanEcoDriver.GoString(), val,
				"ontap-san-economy driver contains %v", key)
		}
	}
}

func TestOntapSanEconomyReconcileNodeAccess(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	server := api.NewFakeUnstartedVserver(ctx, vserverAdminHost, vserverAggrName)
	server.StartTLS()

	_, port, err := net.SplitHostPort(server.Listener.Addr().String())
	assert.Nil(t, err, "Unable to get Web host port %s", port)

	defer func() {
		if r := recover(); r != nil {
			server.Close()
			t.Error("Panic in fake filer", r)
		}
	}()

	cases := [][]struct {
		igroupName         string
		igroupExistingIQNs []string
		nodes              []*utils.Node
		igroupFinalIQNs    []string
	}{
		// Add a backend
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
		},
		// 2 same cluster backends/ nodes unchanged - both current
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
		// 2 different cluster backends - add node
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
					{
						Name: "node2",
						IQN:  "IQN2",
					},
				},
				igroupFinalIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
		// 2 different cluster backends - remove node
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
				nodes: []*utils.Node{
					{
						Name: "node1",
						IQN:  "IQN1",
					},
				},
				igroupFinalIQNs: []string{"IQN1"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
				nodes: []*utils.Node{
					{
						Name: "node3",
						IQN:  "IQN3",
					},
					{
						Name: "node4",
						IQN:  "IQN4",
					},
				},
				igroupFinalIQNs: []string{"IQN3", "IQN4"},
			},
		},
	}

	for _, testCase := range cases {

		api.FakeIgroups = map[string]map[string]struct{}{}

		var ontapSanDrivers []SANEconomyStorageDriver

		for _, driverInfo := range testCase {

			// simulate existing IQNs on the vserver
			igroupsIQNMap := map[string]struct{}{}
			for _, iqn := range driverInfo.igroupExistingIQNs {
				igroupsIQNMap[iqn] = struct{}{}
			}

			api.FakeIgroups[driverInfo.igroupName] = igroupsIQNMap

			sanEcoStorageDriver := newTestOntapSanEcoDriver(vserverAdminHost, port, vserverAggrName, false, nil)
			sanEcoStorageDriver.Config.IgroupName = driverInfo.igroupName
			ontapSanDrivers = append(ontapSanDrivers, *sanEcoStorageDriver)
		}

		for driverIndex, driverInfo := range testCase {
			err := ontapSanDrivers[driverIndex].ReconcileNodeAccess(ctx, driverInfo.nodes,
				uuid.New().String())
			if err != nil {
				continue
			}
		}

		for _, driverInfo := range testCase {

			assert.Equal(t, len(driverInfo.igroupFinalIQNs), len(api.FakeIgroups[driverInfo.igroupName]))

			for _, iqn := range driverInfo.igroupFinalIQNs {
				assert.Contains(t, api.FakeIgroups[driverInfo.igroupName], iqn)
			}
		}
	}
}

func TestOntapSanEconomyTerminate(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME

	server := api.NewFakeUnstartedVserver(ctx, vserverAdminHost, vserverAggrName)
	server.StartTLS()

	_, port, err := net.SplitHostPort(server.Listener.Addr().String())
	assert.Nil(t, err, "Unable to get Web host port %s", port)

	defer func() {
		if r := recover(); r != nil {
			server.Close()
			t.Error("Panic in fake filer", r)
		}
	}()

	cases := [][]struct {
		igroupName         string
		igroupExistingIQNs []string
	}{
		// 2 different cluster backends - remove backend
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{"IQN1", "IQN2"},
			},
			{
				igroupName:         "igroup2",
				igroupExistingIQNs: []string{"IQN3", "IQN4"},
			},
		},
		{
			{
				igroupName:         "igroup1",
				igroupExistingIQNs: []string{},
			},
		},
	}

	for _, testCase := range cases {

		api.FakeIgroups = map[string]map[string]struct{}{}

		var ontapSanDrivers []SANEconomyStorageDriver

		for _, driverInfo := range testCase {

			// simulate existing IQNs on the vserver
			igroupsIQNMap := map[string]struct{}{}
			for _, iqn := range driverInfo.igroupExistingIQNs {
				igroupsIQNMap[iqn] = struct{}{}
			}

			api.FakeIgroups[driverInfo.igroupName] = igroupsIQNMap

			sanEcoStorageDriver := newTestOntapSanEcoDriver(vserverAdminHost, port, vserverAggrName, false, nil)
			sanEcoStorageDriver.Config.IgroupName = driverInfo.igroupName
			sanEcoStorageDriver.telemetry = nil
			ontapSanDrivers = append(ontapSanDrivers, *sanEcoStorageDriver)
		}

		for driverIndex, driverInfo := range testCase {
			ontapSanDrivers[driverIndex].Terminate(ctx, "")
			assert.NotContains(t, api.FakeIgroups, api.FakeIgroups[driverInfo.igroupName])
		}

	}
}

func TestOntapSanEconomyTerminate_Failed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.DriverContext = tridentconfig.ContextCSI
	d.telemetry = &Telemetry{
		done: make(chan struct{}),
	}

	mockAPI.EXPECT().IgroupDestroy(ctx, gomock.Any()).Times(1).Return(fmt.Errorf("failed to destroy igroup"))

	d.Terminate(ctx, "")

	assert.False(t, d.initialized)
}

func TestEconomyGetChapInfo(t *testing.T) {
	type fields struct {
		initialized   bool
		Config        drivers.OntapStorageDriverConfig
		ips           []string
		physicalPools map[string]storage.Pool
		virtualPools  map[string]storage.Pool
	}
	type args struct {
		in0 context.Context
		in1 string
		in2 string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *utils.IscsiChapInfo
	}{
		{
			name: "driverInitialized",
			fields: fields{
				initialized: true,
				Config: drivers.OntapStorageDriverConfig{
					UseCHAP:                   true,
					ChapUsername:              "foo",
					ChapInitiatorSecret:       "bar",
					ChapTargetUsername:        "baz",
					ChapTargetInitiatorSecret: "biz",
				},
				ips:           nil,
				physicalPools: nil,
				virtualPools:  nil,
			},
			args: args{
				in0: nil,
				in1: "volume",
				in2: "node",
			},
			want: &utils.IscsiChapInfo{
				UseCHAP:              true,
				IscsiUsername:        "foo",
				IscsiInitiatorSecret: "bar",
				IscsiTargetUsername:  "baz",
				IscsiTargetSecret:    "biz",
			},
		},
		{
			name: "driverUninitialized",
			fields: fields{
				initialized: false,
				Config: drivers.OntapStorageDriverConfig{
					UseCHAP:                   true,
					ChapUsername:              "biz",
					ChapInitiatorSecret:       "baz",
					ChapTargetUsername:        "bar",
					ChapTargetInitiatorSecret: "foo",
				},
				ips:           nil,
				physicalPools: nil,
				virtualPools:  nil,
			},
			args: args{
				in0: nil,
				in1: "volume",
				in2: "node",
			},
			want: &utils.IscsiChapInfo{
				UseCHAP:              true,
				IscsiUsername:        "biz",
				IscsiInitiatorSecret: "baz",
				IscsiTargetUsername:  "bar",
				IscsiTargetSecret:    "foo",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &SANEconomyStorageDriver{
				initialized:   tt.fields.initialized,
				Config:        tt.fields.Config,
				ips:           tt.fields.ips,
				physicalPools: tt.fields.physicalPools,
				virtualPools:  tt.fields.virtualPools,
			}
			got, err := d.GetChapInfo(tt.args.in0, tt.args.in1, tt.args.in2)
			if err != nil {
				t.Errorf("GetChapInfo(%v, %v, %v)", tt.args.in0, tt.args.in1, tt.args.in2)
			}
			assert.Equalf(t, tt.want, got, "GetChapInfo(%v, %v, %v)", tt.args.in0, tt.args.in1, tt.args.in2)
		})
	}
}

func TestGetLUNPathEconomy(t *testing.T) {
	lunPathEco := GetLUNPathEconomy("ndvp_lun_pool_ciqptQLQdpAeKPep1_QRFXEBJHCT", "ciqptQLQdpAeKPep1_cFfRTWQibX")
	assert.Equal(t, "/vol/ndvp_lun_pool_ciqptQLQdpAeKPep1_QRFXEBJHCT/ciqptQLQdpAeKPep1_cFfRTWQibX", lunPathEco,
		"Incorrect lun economy path")
}

func TestGetInternalVolumeNamenoPrefix(t *testing.T) {
	helper := NewTestLUNHelper("", tridentconfig.ContextDocker)
	internalVolName := helper.GetInternalVolumeName("my-Lun")
	assert.Equal(t, "my_Lun", internalVolName, "Incorrect lun economy path")
}

func TestGetSnapshotNameFromSnapLUNPathBlank(t *testing.T) {
	helper := NewTestLUNHelper("", tridentconfig.ContextDocker)
	snapName := helper.GetSnapshotNameFromSnapLUNPath("")
	assert.Equal(t, "", snapName, "Incorrect snapshot lun name")
}

func TestGetBucketName(t *testing.T) {
	helper := NewTestLUNHelper("", tridentconfig.ContextDocker)
	snapName := helper.GetBucketName("")
	assert.Equal(t, "", snapName, "Incorrect bucket name")
}

func TestGetAPI(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	result := d.GetAPI()

	assert.True(t, reflect.DeepEqual(result, d.API), "Incorrect API returned")
}

func TestGetTelemetry(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	d.telemetry = &Telemetry{
		Plugin:        d.Name(),
		SVM:           d.GetConfig().SVM,
		StoragePrefix: *d.GetConfig().StoragePrefix,
		Driver:        d,
		done:          make(chan struct{}),
	}
	assert.True(t, reflect.DeepEqual(d.telemetry, d.GetTelemetry()), "Incorrect API returned")
}

func TestBackendNameUnset(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	d.Config.BackendName = ""
	d.ips = []string{"127.0.0.1"}
	assert.Equal(t, d.BackendName(), "ontapsaneco_127.0.0.1", "Incorrect bucket name")
}

func TestBackendNameSet(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.Config.BackendName = "ontapsaneco"
	d.ips = []string{"127.0.0.1"}

	assert.Equal(t, d.BackendName(), "ontapsaneco", "Incorrect backend name")
}

func TestDriverInitialized(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	d.initialized = true
	assert.Equal(t, d.Initialized(), d.initialized, "Incorrect initialization status")
}

func TestOntapSanEconomyTerminateCSI(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.telemetry = nil
	d.Config.DriverContext = tridentconfig.ContextCSI

	mockAPI.EXPECT().IgroupDestroy(ctx, gomock.Any()).Times(1).Return(nil)

	d.Terminate(ctx, "")
	assert.False(t, d.initialized)
}

func TestDriverValidate(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	result := d.validate(ctx)
	assert.NoError(t, result, "San Economy driver validation failed")
}

func TestDriverIgnoresDataLIF(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.Config.DataLIF = "foo"

	assert.NoError(t, d.validate(ctx), "driver validation succeeded: data LIF is ignored")
}

func TestDriverValidateInvalidPrefix(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	d.Config.StoragePrefix = ToStringPointer("B@D")

	assert.EqualError(t, d.validate(ctx), "storage prefix may only contain letters/digits/underscore/dash")
}

func TestDriverValidateInvalidPools(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[SpaceReserve] = "iaminvalid"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	result := d.validate(ctx)
	assert.EqualError(t, result, "storage pool validation failed: invalid spaceReserve iaminvalid in pool pool1")
}

func TestOntapSanEconomyVolumeCreate(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(2).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&api.Lun{Size: "1073741824"}, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeCreate_LUNExists(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}

	tests := []struct {
		message   string
		lunExists bool
	}{
		{"error fetching info", false},
		{"lun already exists", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.lunExists {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			}

			result := d.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeCreate_NoPhysicalPool(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.virtualPools = map[string]storage.Pool{}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_InvalidSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		volumeSize string
	}{
		{"invalid"},
		{"19m"},
		{"-1002947563b"},
	}
	for _, test := range tests {
		t.Run(test.volumeSize, func(t *testing.T) {
			volConfig := &storage.VolumeConfig{
				Size:       test.volumeSize,
				Encryption: "false",
				FileSystem: "xfs",
			}

			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

			result := d.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeCreate_LUNNameLimitExceeding(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
		InternalName: "thisIsATestLabelWhoseLengthShouldExceed1023Characters_AddingSomeRandomCharacters_" +
			"V88bESTQlRIWRSS40sx9ND8P9yPf0LV8jPofiqtTp2iIXgotGh83zZ1HEeFlMGxZlIcOiPdoi07cJ" +
			"bQBuHvTRNX6pHRKUXaIrjEpygM4SpaqHYdZ8O1k2meeugg7eXu4dPhqetI3Sip3W4v9QuFkh1YBaI" +
			"9sHE9w5eRxpmTv0POpCB5xAqzmN6XCkxuXKc4yfNS9PRwcTSpvkA3PcKCF3TD1TJU3NYzcChsFQgm" +
			"bAsR32cbJRdsOwx6BkHNfRCji0xSnBFUFUu1sGHfYCmzzd3OmChADIP6RwRtpnqNzvt0CU6uumBnl",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.ErrorContains(t, result, "name exceeds the limit of 254 characters")
}

func TestOntapSanEconomyVolumeCreate_InvalidSnapshotReserve(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		"tieringPolicy": "none",
		SnapshotReserve: "invalid-value",
	})
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_InvalidEncryptionValue(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "invalid-value", // invalid bool value
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_InvalidFilesystem(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "nfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_BothQosPolicies(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.SetInternalAttributes(map[string]string{
		TieringPolicy:     "",
		QosPolicy:         "fake",
		AdaptiveQosPolicy: "fake",
	})
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().TieringPolicyValue(ctx).Return("none")

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_NoAggregates(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.LimitAggregateUsage = "invalid"
	pool1 := storage.NewStoragePool(nil, "")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_NoFlexvol(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{}, fmt.Errorf("failed to list volumes"))

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_ResizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		message      string
		destroyError bool
	}{
		{"nil", false},
		{"failed to delete volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil).Times(2)
			mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
			mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("error resizing volume"))
			if test.destroyError {
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(nil)
			}

			result := d.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeCreate_LUNCreateFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		message      string
		destroyError bool
	}{
		{"nil", false},
		{"failed to delete volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(2).Return(api.Luns{}, nil)
			mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
			mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(fmt.Errorf("failed to create lun"))
			if test.destroyError {
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(nil)
			}

			result := d.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeCreate_TooManyLUNs(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(3).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(2).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(2).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(2).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(2).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(2).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(api.TooManyLunsError("too many luns"))
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&api.Lun{Size: "1073741824"}, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeCreate_GetLUNFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(2).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(nil, fmt.Errorf("failed to fetch lun"))

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeCreate_LUNSetAttributeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}

	tests := []struct {
		message   string
		errorType string
	}{
		{"nil", "None"},
		{"failed to delete volume", "Volume"},
		{"failed to delete lun", "Lun"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(2).Return(api.Luns{}, nil)
			mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
			mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
			mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
			mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&api.Lun{Size: "1073741824"}, nil)
			mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
				gomock.Any()).Times(1).Return(fmt.Errorf("failed to set attribute"))

			switch test.errorType {
			case "Lun":
				mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(fmt.Errorf(test.message))
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(nil)
			case "Volume":
				mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil)
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(fmt.Errorf(test.message))
			default:
				mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil)
				mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(nil)
			}

			result := d.Create(ctx, volConfig, pool1, volAttrs)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeCreate_Resize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}
	lun := api.Lun{
		Size: "10737418240", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(2).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil).Times(2)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&lun, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(1073741824), nil).Times(2)

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeCreate_ResizeVolumeSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}
	lun := api.Lun{
		Size: "10737418240", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&lun, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(1073741824), fmt.Errorf("failed to set size"))

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeCreate_ResizeSetSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}
	lun := api.Lun{
		Size: "10737418240", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{}, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&lun, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to set volume size"))

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeCreate_ResizeVolumeSizeFailed2(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	volAttrs := map[string]sa.Request{}
	lun := api.Lun{
		Size: "10737418240", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Times(1).Return(api.Volumes{}, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{}, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunCreate(ctx, gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&lun, nil)
	mockAPI.EXPECT().LunSetAttribute(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(),
		gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(1073741824), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(1073741824), fmt.Errorf("failed to get volume size"))

	result := d.Create(ctx, volConfig, pool1, volAttrs)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeClone(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{}, nil)
	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{api.Lun{Size: "1g", Name: "lunName", VolumeName: "volumeName"}}, nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	result := d.CreateClone(ctx, volConfig, volConfig, pool1)

	assert.Nil(t, result)
}

func TestOntapSanEconomyVolumeClone_BothQosPolicy(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[TieringPolicy] = "none"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	volConfig := &storage.VolumeConfig{
		Size:              "1g",
		Encryption:        "false",
		FileSystem:        "xfs",
		QosPolicy:         "fake",
		AdaptiveQosPolicy: "fake",
	}

	result := d.CreateClone(ctx, volConfig, volConfig, pool1)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeImport(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:       "1g",
		Encryption: "false",
		FileSystem: "xfs",
	}

	err := d.Import(ctx, volConfig, "volInternal")

	assert.EqualError(t, err, "import is not implemented")
}

func TestOntapSanEconomyVolumeRename(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	err := d.Rename(ctx, "volInternal", "newVolInternal")

	assert.EqualError(t, err, "rename is not implemented")
}

func TestOntapSanEconomyVolumeDestroy(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().LunListIgroupsMapped(ctx, gomock.Any()).Return([]string{"igroup1"}, nil)
	mockAPI.EXPECT().LunUnmap(ctx, "igroup1", gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, "volumeName").Return(flexVol, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, "volumeName", "1073741824").Return(nil).Times(2)

	result := d.Destroy(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeDestroy_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			}

			result := d.Destroy(ctx, volConfig)

			if test.expectError {
				assert.Error(t, result)
			} else {
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapSanEconomyVolumeDestroy_InvalidSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "invalid_size", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDestroy_DeleteSnapshotFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(fmt.Errorf("failed to delete lun"))

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDestroy_UnmapFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "volumeName").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, "volumeName", "1073741824").Return(nil)
	mockAPI.EXPECT().LunListIgroupsMapped(ctx, gomock.Any()).Return([]string{"igroup1"}, nil)
	mockAPI.EXPECT().LunUnmap(ctx, "igroup1", gomock.Any()).Return(fmt.Errorf("failed to unmap lun"))

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDestroy_LUNDestroyFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().LunListIgroupsMapped(ctx, gomock.Any()).Return([]string{"igroup1"}, nil)
	mockAPI.EXPECT().LunUnmap(ctx, "igroup1", gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "volumeName").Return(flexVol, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, "volumeName", "1073741824").Return(nil)
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(fmt.Errorf("failed to destroy lun"))

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDestroy_DockerContext(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}
	flexVol := &api.Volume{
		Name:    "flexvol",
		Comment: "flexvol",
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Return("node", nil)
	mockAPI.EXPECT().IscsiInterfaceGet(ctx, "SVM1").Return([]string{"node"}, nil)
	mockAPI.EXPECT().LunMapInfo(ctx, "", gomock.Any()).Return(123, nil)
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Return(nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, "volumeName").Return(flexVol, nil).Times(2)
	mockAPI.EXPECT().VolumeSetSize(ctx, "volumeName", "1073741824").Return(nil).Times(2)
	mockAPI.EXPECT().LunListIgroupsMapped(ctx, gomock.Any()).Return([]string{"igroup1"}, nil)
	mockAPI.EXPECT().LunUnmap(ctx, "igroup1", gomock.Any()).Return(nil)

	result := d.Destroy(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeDestroy_DockerContext_Failure(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Return("", fmt.Errorf("error fetching node name"))

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDestroy_DockerContext_LunMapInfoError(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.DriverContext = tridentconfig.ContextDocker
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextDocker)
	volConfig := &storage.VolumeConfig{
		InternalName: "storagePrefix_vol1",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil).AnyTimes()
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Return("node", nil)
	mockAPI.EXPECT().IscsiInterfaceGet(ctx, "SVM1").Return([]string{"node"}, nil)
	mockAPI.EXPECT().LunMapInfo(ctx, "", gomock.Any()).Return(0, fmt.Errorf("failed to fetch lun map info"))

	result := d.Destroy(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDeleteBucketIfEmpty(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeDestroy(ctx, "volumeName", true).Return(nil)

	result := d.DeleteBucketIfEmpty(ctx, "volumeName")

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeDeleteBucketIfEmpty_LUNFetchFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(nil, fmt.Errorf("error fetching lun"))

	result := d.DeleteBucketIfEmpty(ctx, "volumeName")

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDeleteBucketIfEmpty_LUNDestroyFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeDestroy(ctx, "volumeName", true).Return(fmt.Errorf("failed to destroy lun"))

	result := d.DeleteBucketIfEmpty(ctx, "volumeName")

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeDeleteBucketIfEmpty_VolumeSetSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, "volumeName").Return(nil, fmt.Errorf("failed to fetch volume info"))
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to resize the volume"))

	result := d.DeleteBucketIfEmpty(ctx, "volumeName")

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumePublish(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	publishInfo := &utils.VolumePublishInfo{
		HostName:         "bar",
		HostIQN:          []string{"host_iqn"},
		TridentUUID:      "1234",
		VolumeAccessInfo: utils.VolumeAccessInfo{PublishEnforcement: true},
		Unmanaged:        false,
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{api.Lun{Size: "1g", Name: "lunName", VolumeName: "volumeName"}}, nil)
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Times(1).Return("node1", nil)
	mockAPI.EXPECT().IscsiInterfaceGet(ctx, gomock.Any()).Return([]string{"iscsi_if"}, nil).Times(1)
	mockAPI.EXPECT().LunGetComment(ctx, "/vol/volumeName/storagePrefix_lunName")
	mockAPI.EXPECT().EnsureIgroupAdded(ctx, gomock.Any(), gomock.Any()).Times(1)
	mockAPI.EXPECT().EnsureLunMapped(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(1, nil)
	mockAPI.EXPECT().LunMapGetReportingNodes(ctx, gomock.Any(), gomock.Any()).Times(1).Return([]string{"node1"}, nil)
	mockAPI.EXPECT().GetSLMDataLifs(ctx, gomock.Any(), gomock.Any()).Times(1).Return([]string{"1.1.1.1"}, nil)

	result := d.Publish(ctx, volConfig, publishInfo)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumePublishSLMError(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	publishInfo := &utils.VolumePublishInfo{
		HostName:         "bar",
		HostIQN:          []string{"host_iqn"},
		TridentUUID:      "1234",
		VolumeAccessInfo: utils.VolumeAccessInfo{PublishEnforcement: true},
		Unmanaged:        false,
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{api.Lun{Size: "1g", Name: "lunName", VolumeName: "volumeName"}}, nil)
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Times(1).Return("node1", nil)
	mockAPI.EXPECT().IscsiInterfaceGet(ctx, gomock.Any()).Return([]string{"iscsi_if"}, nil).Times(1)
	mockAPI.EXPECT().LunGetComment(ctx, "/vol/volumeName/storagePrefix_lunName")
	mockAPI.EXPECT().EnsureIgroupAdded(ctx, gomock.Any(), gomock.Any()).Times(1)
	mockAPI.EXPECT().EnsureLunMapped(ctx, gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(1, nil)
	mockAPI.EXPECT().LunMapGetReportingNodes(ctx, gomock.Any(), gomock.Any()).Times(1).Return([]string{"node1"}, nil)
	mockAPI.EXPECT().GetSLMDataLifs(ctx, gomock.Any(), gomock.Any()).Times(1).Return([]string{}, nil)

	result := d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumePublish_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	publishInfo := &utils.VolumePublishInfo{
		HostName:         "bar",
		HostIQN:          []string{"host_iqn"},
		TridentUUID:      "1234",
		VolumeAccessInfo: utils.VolumeAccessInfo{PublishEnforcement: true},
		Unmanaged:        false,
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"volume does not exist", false},
		{"error checking for existing volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
			}

			result := d.Publish(ctx, volConfig, publishInfo)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumePublish_GetISCSITargetInfoFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}
	publishInfo := &utils.VolumePublishInfo{
		HostName:         "bar",
		HostIQN:          []string{"host_iqn"},
		TridentUUID:      "1234",
		VolumeAccessInfo: utils.VolumeAccessInfo{PublishEnforcement: true},
		Unmanaged:        false,
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{api.Lun{Size: "1g", Name: "lunName", VolumeName: "volumeName"}}, nil)
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Times(1).Return("", fmt.Errorf("failed to get name"))

	result := d.Publish(ctx, volConfig, publishInfo)

	assert.Error(t, result)
}

func TestDriverCanSnapshot(t *testing.T) {
	vserverAdminHost := ONTAPTEST_LOCALHOST
	vserverAdminPort := "0"
	vserverAggrName := ONTAPTEST_VSERVER_AGGR_NAME
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	d := newTestOntapSanEcoDriver(vserverAdminHost, vserverAdminPort, vserverAggrName, false, mockAPI)

	result := d.CanSnapshot(ctx, nil, nil)

	assert.NoError(t, result)
}

func TestOntapSanEconomyGetSnapshot(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "lunName",
		VolumeName:         "volumeName",
		Name:               "lunName",
		VolumeInternalName: "volumeName",
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1g",
			Name:       "volumeName_snapshot_lunName",
			VolumeName: "volumeName",
		},
	},
		nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&api.Lun{Size: "1073741824"}, nil)

	snap, err := d.GetSnapshot(ctx, snapConfig, nil)

	assert.NoError(t, err, "Error is not nil")
	assert.NotNil(t, snap, "snapshot is nil")
}

func TestOntapSanEconomyGetSnapshots(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "volumeName_snapshot_lunName",
			VolumeName: "volumeName",
		},
	},
		nil)
	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/volumeName/storagePrefix_LUNName_snapshot_mySnap",
			VolumeName: "volumeName",
		},
	},
		nil)

	snaps, err := d.GetSnapshots(ctx, volConfig)

	assert.NoError(t, err)
	assert.NotNil(t, snaps, "snapshots are nil")
}

func TestOntapSanEconomyGetSnapshots_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.ips = []string{"127.0.0.1"}
	volConfig := &storage.VolumeConfig{
		InternalName: "lunName",
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"volume does not exist", false},
		{"error checking for existing volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
			}

			snaps, err := d.GetSnapshots(ctx, volConfig)

			assert.Error(t, err, "Error is not nil")
			assert.Nil(t, snaps, "snapshots are nil")
		})
	}
}

func TestOntapSanEconomyCreateSnapshot(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "storagePrefix_my_Lun_my_Bucket",
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(2).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
			VolumeName: "my_Bucket",
		},
	},
		nil)
	mockAPI.EXPECT().LunGetByName(ctx,
		gomock.Any()).Times(1).Return(&api.Lun{
		Size:       "1073741824",
		Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
		VolumeName: "my_Bucket",
	},
		nil)
	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(2).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
			VolumeName: "my_Bucket",
		},
	},
		nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_snapshot_snap_1",
			VolumeName: "my_Bucket",
		},
	},
		nil)

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.NoError(t, err, "Error is not nil")
	assert.NotNil(t, snap, "snapshots are nil")
}

func TestOntapSanEconomyCreateSnapshot_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "storagePrefix_my_Lun_my_Bucket",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"volume does not exist", false},
		{"error checking for existing volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
			}

			snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

			assert.Error(t, err)
			assert.Nil(t, snap)
		})
	}
}

func TestOntapSanEconomyCreateSnapshot_LUNCreateCloneFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "vol1_snapshot_snap_1",
		VolumeInternalName: "vol1_snapshot_snap_1",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "vol/my_Bucket/storagePrefix_vol1_snapshot_snap_1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&lun, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(),
		api.QosPolicyGroup{}).Return(fmt.Errorf("failed to create lun clone"))

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.Error(t, err)
	assert.Nil(t, snap)
}

func TestOntapSanEconomyCreateSnapshot_LUNGetByNameFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "vol1",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to fetch lun"))

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.Error(t, err)
	assert.Nil(t, snap, "snapshots are nil")
}

func TestOntapSanEconomyCreateSnapshot_CreateLUNClone_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	d.Config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	policy := api.QosPolicyGroup{}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
			}

			result := d.createLUNClone(ctx, "lun_storagePrefix_vol1", "source", "snap", &d.Config, mockAPI,
				"storagePrefix", true, policy)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyCreateSnapshot_CreateLUNClone_LUNCreatedFromSnapshot(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	d.Config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	policy := api.QosPolicyGroup{}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			}

			result := d.createLUNClone(ctx, "lun_storagePrefix_vol1", "source", "snap", &d.Config, mockAPI,
				"storagePrefix", true, policy)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyCreateSnapshot_CreateLUNClone_LUNCloneCreateFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1_snapshot_snap", VolumeName: "volumeName"},
	}
	d.Config.CommonStorageDriverConfig.DebugTraceFlags["method"] = true
	policy := api.QosPolicyGroup{}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(),
		api.QosPolicyGroup{}).Return(fmt.Errorf("failed to create lun clone"))

	result := d.createLUNClone(ctx, "lun_storagePrefix_vol1_snapshot_snap", "vol1", "snap", &d.Config, mockAPI,
		"storagePrefix", true, policy)

	assert.Error(t, result)
}

func TestOntapSanEconomyCreateSnapshot_GetSnapshotsEconomyFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "vol1",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), api.QosPolicyGroup{}).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to fetch lun")).Times(2)

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.Error(t, err)
	assert.Nil(t, snap, "snapshots are nil")
}

func TestOntapSanEconomyCreateSnapshot_InvalidSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "vol1",
	}
	lun := api.Lun{
		Size: "invalid", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}
	snapLuns := []api.Lun{
		{Size: "1073741824", Name: "/vol/volumeName/storagePrefix_vol1_snapshot_mySnap", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(snapLuns, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), api.QosPolicyGroup{}).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.Error(t, err)
	assert.Nil(t, snap, "snapshots are nil")
}

func TestOntapSanEconomyCreateSnapshot_SnapshotNotFound(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.flexvolNamePrefix = "storagePrefix"
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "vol1",
	}
	lun := api.Lun{
		Size: "invalid", Name: "lun_storagePrefix_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_vol1", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().LunCloneCreate(ctx, gomock.Any(), gomock.Any(), gomock.Any(), api.QosPolicyGroup{}).Return(nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	snap, err := d.CreateSnapshot(ctx, snapConfig, nil)

	assert.Error(t, err)
	assert.Nil(t, snap, "snapshots are nil")
}

func TestOntapSanEconomyRestoreSnapshot(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "storagePrefix_my_Lun_my_Bucket",
	}

	err := d.RestoreSnapshot(ctx, snapConfig, nil)

	assert.EqualError(t, err, fmt.Sprintf("restoring snapshots is not supported by backend type %s", d.Name()))
}

func TestOntapSanEconomyVolumeDeleteSnapshot(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "storagePrefix_my_Lun_my_Bucket",
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(2).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "storagePrefix_my_Lun_my_Bucket_snapshot_snap_1",
			VolumeName: "my_Bucket",
		},
	},
		nil)
	mockAPI.EXPECT().LunDestroy(ctx, gomock.Any()).Times(1)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{api.Lun{}}, nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)

	result := d.DeleteSnapshot(ctx, snapConfig, nil)

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeDeleteSnapshot_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "snap_1",
		VolumeName:         "my_Bucket",
		Name:               "/vol/my_Bucket/storagePrefix_my_Lun",
		VolumeInternalName: "storagePrefix_my_Lun_my_Bucket",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			}

			result := d.DeleteSnapshot(ctx, snapConfig, nil)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyGet(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
			VolumeName: "my_Bucket",
		},
	},
		nil)

	result := d.Get(ctx, "my_Bucket")

	assert.NoError(t, result)
}

func TestOntapSanEconomyGet_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			}

			result := d.Get(ctx, "my_Bucket")

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyEnsureFlexvolForLUN_InvalidVolumeSize(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	d.Config.LimitVolumeSize = "invalid"

	flexVol, newly, err := d.ensureFlexvolForLUN(ctx, &api.Volume{}, uint64(1073741824), opts, d.Config, pool1,
		make(map[string]struct{}))

	assert.Equal(t, flexVol, "")
	assert.False(t, newly)
	assert.Error(t, err)
}

func TestOntapSanEconomyEnsureFlexvolForLUN_FlexvolFound(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)

	flexVol, newly, err := d.ensureFlexvolForLUN(ctx, vol, uint64(1073741824), opts, d.Config, pool1,
		make(map[string]struct{}))

	assert.NotEqual(t, flexVol, "")
	assert.False(t, newly)
	assert.NoError(t, err)
}

func TestOntapSanEconomyEnsureFlexvolForLUN_FlexvolNotFound(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name:       "",
		Aggregates: []string{"data"},
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(fmt.Errorf("failed to create volume"))

	flexVol, newly, err := d.ensureFlexvolForLUN(ctx, vol, uint64(1073741824), opts, d.Config, pool1,
		make(map[string]struct{}))

	assert.Equal(t, flexVol, "")
	assert.False(t, newly)
	assert.Error(t, err)
}

func TestOntapSanEconomyGetStorageBackendSpecs(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.ips = []string{"127.0.0.1"}
	backend := storage.StorageBackend{}

	result := d.GetStorageBackendSpecs(ctx, &backend)

	assert.Nil(t, result)
}

func TestOntapSanEconomyGetStorageBackendPhysicalPoolNames(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	poolNames := d.GetStorageBackendPhysicalPoolNames(ctx)

	assert.Equal(t, "pool1", poolNames[0], "Pool names are not equal")
}

func TestOntapSanEconomyGetInternalVolumeName(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.Config.StoragePrefix = ToStringPointer("storagePrefix_")

	internalVolName := d.GetInternalVolumeName(ctx, "my-Lun")

	assert.Equal(t, "storagePrefix_my_Lun", internalVolName, "Strings not equal")
}

func TestOntapSanEconomyGetProtocol(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	protocol := d.GetProtocol(ctx)

	assert.Equal(t, protocol, tridentconfig.Block, "Protocols not equal")
}

func TestOntapSanEconomyStoreConfig(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	persistentConfig := &storage.PersistentStorageBackendConfig{}

	d.StoreConfig(ctx, persistentConfig)

	assert.Equal(t, json.RawMessage("{}"), d.Config.CommonStorageDriverConfig.StoragePrefixRaw,
		"raw prefix mismatch")
	assert.Equal(t, d.Config, *persistentConfig.OntapConfig, "ontap config mismatch")
}

func TestGetVolumeExternal(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().LunGetByName(ctx,
		gomock.Any()).Times(1).Return(&api.Lun{
		Size:       "1073741824",
		Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
		VolumeName: "my_Bucket",
	},
		nil)

	result, resultErr := d.GetVolumeExternal(ctx, "my_Lun")

	assert.Nil(t, resultErr, "not nil")
	assert.IsType(t, &storage.VolumeExternal{}, result, "type mismatch")
	assert.Equal(t, "1", result.Config.Version)
	assert.Equal(t, "my_Lun_my_Bucket", result.Config.Name)
	assert.Equal(t, "storagePrefix_my_Lun_my_Bucket", result.Config.InternalName)
}

func TestGetVolumeExternal_MultipleAggregates(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	lun := api.Lun{
		Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap",
		VolumeName: "myLun",
	}
	vol := api.Volume{
		Aggregates: []string{"aggr1", "aggr2"},
	}

	volExt := d.getVolumeExternal(&lun, &vol)

	assert.NotNil(t, volExt)
}

func TestGetVolumeExternal_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get lun"))

	result, err := d.GetVolumeExternal(ctx, "my_Lun")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetVolumeExternal_VolumeInfoFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(nil, fmt.Errorf("failed to get volume info"))

	result, err := d.GetVolumeExternal(ctx, "my_Lun")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetVolumeExternal_LunGetByNameFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{Aggregates: []string{"data"}}, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get lun by name"))

	result, err := d.GetVolumeExternal(ctx, "my_Lun")

	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestGetVolumeExternalWrappers(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(api.Volumes{&api.Volume{}}, nil).Times(1)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(api.Luns{api.Lun{}}, nil)

	d.GetVolumeExternalWrappers(ctx, channel)

	// Read the volumes from the channel
	volumes := make([]*storage.VolumeExternal, 0)
	for wrapper := range channel {
		if wrapper.Error != nil {
			t.FailNow()
		} else {
			volumes = append(volumes, wrapper.Volume)
		}
	}

	assert.Len(t, volumes, 1, "wrong number of volumes")
}

func TestGetVolumeExternalWrappers_Failed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to list volume by prefix"))

	d.GetVolumeExternalWrappers(ctx, channel)

	// Read the volumes from the channel
	volumes := make([]*storage.VolumeExternal, 0)
	for wrapper := range channel {
		if wrapper.Error != nil {
			assert.Error(t, wrapper.Error)
		} else {
			volumes = append(volumes, wrapper.Volume)
		}
	}
}

func TestGetVolumeExternalWrappers_NoVolumes(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(nil, nil)

	d.GetVolumeExternalWrappers(ctx, channel)

	// Read the volumes from the channel
	volumes := make([]*storage.VolumeExternal, 0)
	for wrapper := range channel {
		if wrapper.Error != nil {
			t.FailNow()
		} else {
			volumes = append(volumes, wrapper.Volume)
		}
	}

	assert.Len(t, volumes, 0, "wrong number of volumes")
}

func TestGetVolumeExternalWrappers_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to fetch lun"))
	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(api.Volumes{&api.Volume{}}, nil).Times(1)

	d.GetVolumeExternalWrappers(ctx, channel)
}

func TestGetVolumeExternalWrappers_VolumeNotAttachedWithLUN(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)

	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(api.Volumes{&api.Volume{Name: "vol1"}}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{api.Lun{VolumeName: "vol2"}}, nil)

	d.GetVolumeExternalWrappers(ctx, channel)

	// Read the volumes from the channel
	volumes := make([]*storage.VolumeExternal, 0)
	for wrapper := range channel {
		if wrapper.Error != nil {
			t.FailNow()
		} else {
			volumes = append(volumes, wrapper.Volume)
		}
	}

	assert.Len(t, volumes, 0)
}

func TestGetVolumeExternalWrappers_VolumeAttachedWithLUN(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	channel := make(chan *storage.VolumeExternalWrapper, 1)
	luns := []api.Lun{
		{Size: "1g", Name: "storagePrefix_my_Lun_my_Bucket_snapshot_snap", VolumeName: "my_Bucket"},
	}

	mockAPI.EXPECT().VolumeListByPrefix(ctx, gomock.Any()).Return(api.Volumes{&api.Volume{Name: "my_Bucket"}},
		nil).Times(1)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Times(1).Return(luns, nil)

	d.GetVolumeExternalWrappers(ctx, channel)

	// Read the volumes from the channel
	volumes := make([]*storage.VolumeExternal, 0)
	for wrapper := range channel {
		if wrapper.Error != nil {
			t.FailNow()
		} else {
			volumes = append(volumes, wrapper.Volume)
		}
	}

	assert.Len(t, volumes, 0)
}

func TestGetUpdateType_OtherChanges(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)

	oldDriver := newTestOntapSanEcoDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, true, mockAPI)
	oldDriver.API = mockAPI
	prefix1 := "storagePrefix_"
	oldDriver.Config.StoragePrefix = &prefix1
	oldDriver.Config.Credentials = map[string]string{
		drivers.KeyName: "secret1",
		drivers.KeyType: string(drivers.CredentialStoreK8sSecret),
	}
	oldDriver.Config.DataLIF = "10.0.2.11"
	oldDriver.Config.Username = "oldUser"
	oldDriver.Config.Password = "oldPassword"

	newDriver := newTestOntapSanEcoDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, true, mockAPI)
	oldDriver.API = mockAPI
	prefix2 := "storagePREFIX_"

	newDriver.Config.StoragePrefix = &prefix2
	newDriver.Config.Credentials = map[string]string{
		drivers.KeyName: "secret2",
		drivers.KeyType: string(drivers.CredentialStoreK8sSecret),
	}
	newDriver.Config.DataLIF = "10.0.2.10"
	newDriver.Config.Username = "newUser"
	newDriver.Config.Password = "newPassword"

	result := newDriver.GetUpdateType(ctx, oldDriver)

	expectedBitmap := &roaring.Bitmap{}
	expectedBitmap.Add(storage.InvalidVolumeAccessInfoChange)
	expectedBitmap.Add(storage.UsernameChange)
	expectedBitmap.Add(storage.PasswordChange)
	expectedBitmap.Add(storage.PrefixChange)
	expectedBitmap.Add(storage.CredentialsChange)

	assert.Equal(t, expectedBitmap, result, "bitmap mismatch")
}

func TestGetUpdateType_Failure(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	mockAPI := mockapi.NewMockOntapAPI(mockCtrl)
	oldDriver := newTestOntapNASDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, "CSI", false)
	newDriver := newTestOntapSanEcoDriver(ONTAPTEST_LOCALHOST, "0", ONTAPTEST_VSERVER_AGGR_NAME, true, mockAPI)
	expectedBitmap := &roaring.Bitmap{}
	expectedBitmap.Add(storage.InvalidUpdate)

	result := newDriver.GetUpdateType(ctx, oldDriver)

	assert.Equal(t, expectedBitmap, result, "bitmap mismatch")
}

func TestOntapSanEconomyVolumeResize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}

	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(2).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
			VolumeName: "my_Bucket",
		},
	},
		nil)
	mockAPI.EXPECT().LunGetByName(ctx,
		gomock.Any()).Times(1).Return(&api.Lun{
		Size:       "1073741824",
		Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
		VolumeName: "my_Bucket",
	},
		nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().LunList(ctx,
		gomock.Any()).Times(1).Return(api.Luns{
		api.Lun{
			Size:       "1073741824",
			Name:       "/vol/my_Bucket/storagePrefix_my_Lun_my_Bucket",
			VolumeName: "my_Bucket",
		},
	},
		nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{Aggregates: []string{"data"}}, nil)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(2147483648), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Times(1).Return(nil)
	mockAPI.EXPECT().LunSetSize(ctx, gomock.Any(), "2147483648").Return(uint64(2147483648), nil)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Nil(t, result)
}

func TestOntapSanEconomyVolumeResize_LUNExists(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}

	tests := []struct {
		message string
		isError bool
	}{
		{"volume does not exist", false},
		{"error checking for existing volume", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.isError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(api.Luns{}, nil)
			}

			result := d.Resize(ctx, volConfig, uint64(2147483648))

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyVolumeResize_InvalidLUNSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	luns := []api.Lun{
		{
			Size:       "invalid", // invalid uint64 value
			Name:       "lun_vol1",
			VolumeName: "volumeName",
		},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_GetSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_vol1", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to fetch lun"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_VolumeInfoFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get volume info")).Times(2)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_SizeError(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}

	tests := []struct {
		message       string
		requestedSize uint64
		currentSize   string
		expectError   bool
	}{
		{"same size", 1073741824, "1073741824", false},
		{"requested size is less than current size", 1073741824, "2147483648", true},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			lun := api.Lun{
				Size: test.currentSize, Name: "lun_vol1",
				VolumeName: "volumeName",
			}
			luns := make([]api.Lun, 0)
			luns = append(luns, lun)

			mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
			mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
			mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{}, nil)

			result := d.Resize(ctx, volConfig, test.requestedSize)

			if test.expectError {
				assert.Error(t, result)
			} else {
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapSanEconomyVolumeResize_LimitVolumeSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.Config.LimitVolumeSize = "invalid-value" // invalid int value
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil).Times(2)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_LUNGetGeometryFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(2147483648),
		fmt.Errorf("failed to get lun geometry"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_LUNMaxSizeLess(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(1073741824), nil)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_VolumeSetSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(2147483648), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to set volume size"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_LunSetSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(3)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(2147483648), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunSetSize(ctx, gomock.Any(), "2147483648").Return(uint64(2147483648),
		fmt.Errorf("failed to set lun size"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.Error(t, result)
}

func TestOntapSanEconomyVolumeResize_FlexvolBiggerThanLUN(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil).Times(2)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil).Times(2)
	mockAPI.EXPECT().LunSetSize(ctx, gomock.Any(), gomock.Any()).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(2147483648), nil)

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeResize_VolumeSizeFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil).Times(2)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil).Times(2)
	mockAPI.EXPECT().LunSetSize(ctx, gomock.Any(), gomock.Any()).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSize(ctx, gomock.Any()).Return(uint64(2147483648), fmt.Errorf("failed to get volume size"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.NoError(t, result)
}

func TestOntapSanEconomyVolumeResize_VolumeSetSizeFailed2(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	volConfig := &storage.VolumeConfig{
		Size:       "1073741824",
		Encryption: "false",
		FileSystem: "xfs",
	}
	lun := api.Lun{
		Size: "1073741824", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(&api.Volume{}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Return(&api.Volume{Aggregates: []string{"aggr"}}, nil).Times(2)
	mockAPI.EXPECT().SupportsFeature(ctx, api.LunGeometrySkip).Times(1)
	mockAPI.EXPECT().LunGetGeometry(ctx, gomock.Any()).Times(1).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(nil)
	mockAPI.EXPECT().LunSetSize(ctx, gomock.Any(), gomock.Any()).Return(uint64(16106127360), nil)
	mockAPI.EXPECT().VolumeSetSize(ctx, gomock.Any(), gomock.Any()).Return(fmt.Errorf("failed to set volume size"))

	result := d.Resize(ctx, volConfig, uint64(2147483648))

	assert.NoError(t, result)
}

func TestOntapSanEconomyCreateFollowup(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.IgroupName = "igroup"
	d.ips = []string{"10.2.0.1"}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	lun := api.Lun{
		Size: "1g", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, "/vol/*/vol1").Return(luns, nil)
	mockAPI.EXPECT().EnsureLunMapped(ctx, "igroup", "/vol/volumeName/vol1", false).Return(123, nil)
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Return("iscsiNode", nil)
	mockAPI.EXPECT().LunGetByName(ctx, "/vol/volumeName/vol1").Return(&lun, nil)
	mockAPI.EXPECT().LunMapGetReportingNodes(ctx, "igroup", "/vol/volumeName/vol1").Return([]string{"iscsiNode"}, nil)
	mockAPI.EXPECT().GetSLMDataLifs(ctx, []string{"10.2.0.1"}, []string{"iscsiNode"}).Return([]string{"10.2.0.2"}, nil)

	result := d.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapSanEconomyCreateFollowup_DockerContext(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	d.Config.DriverContext = "docker"
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	result := d.CreateFollowup(ctx, volConfig)

	assert.NoError(t, result)
}

func TestOntapSanEconomyCreateFollowup_LunDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, "/vol/*/vol1").Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, "/vol/*/vol1").Return(nil, nil)
			}

			result := d.CreateFollowup(ctx, volConfig)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyCreateFollowup_LunNotMapped(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.IgroupName = "igroup"
	d.ips = []string{"10.2.0.1"}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}
	luns := []api.Lun{
		{Size: "1g", Name: "lun_vol1", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, "/vol/*/vol1").Return(luns, nil)
	mockAPI.EXPECT().EnsureLunMapped(ctx, "igroup", "/vol/volumeName/vol1", false).Return(0,
		fmt.Errorf("lun not mapped"))

	result := d.CreateFollowup(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyCreateFollowup_GetLunFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.Config.IgroupName = "igroup"
	d.ips = []string{"10.2.0.1"}
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "nfs",
		InternalName: "vol1",
	}
	lun := api.Lun{
		Size: "1g", Name: "lun_vol1",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, "/vol/*/vol1").Return(luns, nil)
	mockAPI.EXPECT().EnsureLunMapped(ctx, "igroup", "/vol/volumeName/vol1", false).Return(123, nil)
	mockAPI.EXPECT().IscsiNodeGetNameRequest(ctx).Return("iscsiNode", nil)
	mockAPI.EXPECT().LunGetByName(ctx, "/vol/volumeName/vol1").Return(&lun, fmt.Errorf("couldn't get lun"))

	result := d.CreateFollowup(ctx, volConfig)

	assert.Error(t, result)
}

func TestOntapSanEconomyInitialize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "SVM1",
	    "aggregate":         "data",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "san-eco",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}
	authResponse := api.IscsiInitiatorAuth{
		SVMName:  "SVM1",
		AuthType: "None",
	}
	d.telemetry = &Telemetry{
		Plugin: d.Name(),
		SVM:    "SVM1",
		Driver: d,
		done:   make(chan struct{}),
	}
	d.telemetry.TridentVersion = tridentconfig.OrchestratorVersion.String()
	d.telemetry.TridentBackendUUID = BackendUUID
	d.telemetry.StoragePrefix = "trident_"
	hostname, _ := os.Hostname()
	message, _ := json.Marshal(d.GetTelemetry())

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return([]string{"10.0.207.7"}, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
	mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
		map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
	)
	mockAPI.EXPECT().IgroupCreate(ctx, "trident-deadbeef-03af-4394-ace4-e177cdbcaf28", "iscsi", "linux").Return(nil)
	mockAPI.EXPECT().IscsiInitiatorGetDefaultAuth(ctx).Return(authResponse, nil)
	mockAPI.EXPECT().EmsAutosupportLog(ctx, "ontap-san-economy", "1", false, "heartbeat", hostname, string(message), 1,
		"trident", 5).AnyTimes()

	result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

	assert.NoError(t, result)
}

func TestOntapSanEconomyInitialize_InvalidConfig(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`{invalid-json}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")

	result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapSanEconomyInitialize_NoDataLIFs(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "iscsi_vs",
	    "aggregate":         "data",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "san-eco",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			if test.expectError {
				mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return(nil, nil)
			}
			result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyInitialize_NumOfLUNs(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}
	authResponse := api.IscsiInitiatorAuth{
		SVMName:  "SVM1",
		AuthType: "None",
	}
	hostname, _ := os.Hostname()

	tests := []struct {
		numOfLUNs   string
		expectError bool
	}{
		{"NaN", true},
		{"100", false},
		{"40", true},
		{"500", true},
	}
	for _, test := range tests {
		t.Run(test.numOfLUNs, func(t *testing.T) {
			commonConfigJSON := fmt.Sprintf(`
			{
			    "managementLIF":     "10.0.207.8",
			    "dataLIF":           "10.0.207.7",
			    "svm":               "iscsi_vs",
			    "aggregate":         "data",
			    "username":          "admin",
			    "password":          "password",
			    "storageDriverName": "ontap-san-economy",
			    "storagePrefix":     "san-eco",
			    "debugTraceFlags":   {"method": true, "api": true},
			    "version":1,
				"lunsPerFlexvol":    "%v"
			}`, test.numOfLUNs)

			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return([]string{"10.0.207.7"}, nil)
			mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
			mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
				map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
			)
			mockAPI.EXPECT().EmsAutosupportLog(ctx, "ontap-san-economy", "1", false, "heartbeat", hostname,
				gomock.Any(), 1,
				"trident", 5).AnyTimes()
			if !test.expectError {
				mockAPI.EXPECT().IgroupCreate(ctx, "trident-deadbeef-03af-4394-ace4-e177cdbcaf28", "iscsi",
					"linux").Return(nil)
				mockAPI.EXPECT().IscsiInitiatorGetDefaultAuth(ctx).Return(authResponse, nil)
			}

			result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

			if test.expectError {
				assert.Error(t, result)
			} else {
				assert.NoError(t, result)
			}
		})
	}
}

func TestOntapSanEconomyInitialize_OtherContext(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "iscsi_vs",
	    "aggregate":         "data",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "san-eco",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	tests := []struct {
		driverContext string
		expectError   bool
	}{
		{"docker", true},
		{"invalid", true},
	}
	for _, test := range tests {
		t.Run(test.driverContext, func(t *testing.T) {
			mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
			mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return([]string{"10.0.207.7"}, nil)
			mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
			mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
				map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
			)

			result := d.Initialize(ctx, tridentconfig.DriverContext(test.driverContext), commonConfigJSON, commonConfig,
				secrets, BackendUUID)

			assert.Error(t, result)
		})
	}
}

func TestOntapSanEconomyInitialize_NoSVMAggregates(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "iscsi_vs",
	    "aggregate":         "data",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "san-eco",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return([]string{"10.0.207.7"}, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return(nil, fmt.Errorf("error getting svm aggregate names"))

	result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapSanEconomyInitialize_IGroupCreationFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	commonConfig := &drivers.CommonStorageDriverConfig{
		Version:           1,
		StorageDriverName: "ontap-san-economy",
		BackendName:       "myOntapSanEcoBackend",
		DriverContext:     tridentconfig.ContextCSI,
		DebugTraceFlags:   debugTraceFlags,
	}
	commonConfigJSON := fmt.Sprintf(`
	{
	    "managementLIF":     "10.0.207.8",
	    "dataLIF":           "10.0.207.7",
	    "svm":               "iscsi_vs",
	    "aggregate":         "data",
	    "username":          "admin",
	    "password":          "password",
	    "storageDriverName": "ontap-san-economy",
	    "storagePrefix":     "san-eco",
	    "debugTraceFlags":   {"method": true, "api": true},
	    "version":1
	}`)
	secrets := map[string]string{
		"clientcertificate": "dummy-certificate",
	}

	mockAPI.EXPECT().SVMName().AnyTimes().Return("SVM1")
	mockAPI.EXPECT().NetInterfaceGetDataLIFs(ctx, "iscsi").Return([]string{"10.0.207.7"}, nil)
	mockAPI.EXPECT().GetSVMAggregateNames(ctx).AnyTimes().Return([]string{ONTAPTEST_VSERVER_AGGR_NAME}, nil)
	mockAPI.EXPECT().GetSVMAggregateAttributes(gomock.Any()).AnyTimes().Return(
		map[string]string{ONTAPTEST_VSERVER_AGGR_NAME: "vmdisk"}, nil,
	)
	mockAPI.EXPECT().IgroupCreate(ctx, "trident-deadbeef-03af-4394-ace4-e177cdbcaf28", "iscsi",
		"linux").Return(fmt.Errorf("igroup creation failed"))
	mockAPI.EXPECT().IgroupDestroy(ctx, "trident-deadbeef-03af-4394-ace4-e177cdbcaf28").Return(fmt.Errorf(
		"error deleting igroup"))

	result := d.Initialize(ctx, "csi", commonConfigJSON, commonConfig, secrets, BackendUUID)

	assert.Error(t, result)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(2)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(vol, nil)

	flexVol, err := d.getFlexvolForLUN(ctx, vol, uint64(1073741824), true, uint64(2147483648),
		make(map[string]struct{}))

	assert.NoError(t, err)
	assert.NotEqual(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN_InvalidLUNPath(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePref_vol1",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "myBucket/storagePref_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)

	flexVol, err := d.getFlexvolForLUN(ctx, vol, uint64(1073741824), false, uint64(2147483648),
		make(map[string]struct{}))

	assert.NoError(t, err)
	assert.Equal(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf("error fetching luns"))

	flexVol, err := d.getFlexvolForLUN(ctx, vol, uint64(1073741824), false, uint64(2147483648),
		make(map[string]struct{}))

	assert.Error(t, err)
	assert.Equal(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN_LimitFlexvolSize_Failed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(nil, fmt.Errorf("failed to get volume"))

	flexVol, err := d.getFlexvolForLUN(ctx, vol, uint64(1073741824), true, uint64(2147483648),
		make(map[string]struct{}))

	assert.NoError(t, err)
	assert.Equal(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN_LargeSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(vol, nil)

	flexVol, err := d.getFlexvolForLUN(ctx, vol, uint64(1073741824), true, uint64(1073741824),
		make(map[string]struct{}))

	assert.NoError(t, err)
	assert.Equal(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_GetFlexvolForLUN_IgnoredVols(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	d.helper = NewTestLUNHelper("storagePrefix_", tridentconfig.ContextCSI)
	d.lunsPerFlexvol = 1
	vol1 := &api.Volume{
		Name: "storagePrefix_vol1",
	}
	vol2 := &api.Volume{
		Name: "storagePrefix_vol2",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "/vol/myBucket/storagePrefix_vol1_snapshot_mySnap", VolumeName: "myLun"},
	}

	mockAPI.EXPECT().VolumeListByAttrs(ctx, gomock.Any()).Return(api.Volumes{vol1, vol2}, nil)
	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil).Times(4)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(vol1, nil)
	mockAPI.EXPECT().VolumeInfo(ctx, gomock.Any()).Times(1).Return(vol2, nil)

	flexVol, err := d.getFlexvolForLUN(ctx, vol1, uint64(1073741824), true, uint64(2147483648),
		make(map[string]struct{}))

	assert.NoError(t, err)
	assert.NotEqual(t, "", flexVol)
}

func TestOntapSanEconomyInitialize_CreateFlexvolForLUN_InvalidSnapshotReserve(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	vol := &api.Volume{
		Name: "storagePrefix_vol1",
	}
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	pool1.InternalAttributes()[SnapshotReserve] = "invalid"
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	volName, err := d.createFlexvolForLUN(ctx, vol, opts, pool1)

	assert.Error(t, err)
	assert.Equal(t, "", volName)
}

func TestOntapSanEconomyInitialize_CreateFlexvolForLUN_Failed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	vol := &api.Volume{
		Name:       "storagePrefix_vol1",
		Aggregates: []string{"data"},
	}
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(fmt.Errorf("failed to create volume"))

	volName, err := d.createFlexvolForLUN(ctx, vol, opts, pool1)

	assert.Error(t, err)
	assert.Equal(t, "", volName)
}

func TestOntapSanEconomyInitialize_CreateFlexvolForLUN_VolumeDisableSnapshotDirectoryAccessFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	vol := &api.Volume{
		Name:       "storagePrefix_vol1",
		Aggregates: []string{"data"},
	}
	opts := make(map[string]string)
	pool1 := storage.NewStoragePool(nil, "pool1")
	d.physicalPools = map[string]storage.Pool{"pool1": pool1}

	mockAPI.EXPECT().VolumeCreate(ctx, gomock.Any()).Return(nil)
	mockAPI.EXPECT().VolumeDisableSnapshotDirectoryAccess(ctx,
		gomock.Any()).Return(fmt.Errorf("failed to disable snapshot directory access"))
	mockAPI.EXPECT().VolumeDestroy(ctx, gomock.Any(), true).Return(fmt.Errorf("failed to destroy volume"))

	volName, err := d.createFlexvolForLUN(ctx, vol, opts, pool1)

	assert.Error(t, err)
	assert.Equal(t, "", volName)
}

func TestOntapSanEconomyGetSnapshotEconomy_LUNDoesNotExist(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "lunName",
		VolumeName:         "volumeName",
		Name:               "lunName",
		VolumeInternalName: "volumeName",
	}

	tests := []struct {
		message     string
		expectError bool
	}{
		{"error fetching info", true},
		{"no luns found", false},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			if test.expectError {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, fmt.Errorf(test.message))
			} else {
				mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(nil, nil)
			}

			snap, err := d.getSnapshotEconomy(ctx, snapConfig, &d.Config)

			if test.expectError {
				assert.Error(t, err)
				assert.Nil(t, snap)
			} else {
				assert.NoError(t, err)
				assert.Nil(t, snap)
			}
		})
	}
}

func TestOntapSanEconomyGetSnapshotEconomy_LunGetByNameFailed(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "lunName",
		VolumeName:         "volumeName",
		Name:               "lunName",
		VolumeInternalName: "volumeName",
	}
	luns := []api.Lun{
		{Size: "1073741824", Name: "lun_storagePrefix_volumeName_snapshot_lunName", VolumeName: "volumeName"},
	}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(nil, fmt.Errorf("failed to get lun by name"))

	snap, err := d.getSnapshotEconomy(ctx, snapConfig, &d.Config)

	assert.Error(t, err)
	assert.Nil(t, snap)
}

func TestOntapSanEconomyGetSnapshotEconomy_InvalidSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)
	snapConfig := &storage.SnapshotConfig{
		InternalName:       "lunName",
		VolumeName:         "volumeName",
		Name:               "lunName",
		VolumeInternalName: "volumeName",
	}
	lun := api.Lun{
		Size: "invalid", Name: "lun_storagePrefix_volumeName_snapshot_lunName",
		VolumeName: "volumeName",
	}
	luns := []api.Lun{lun}

	mockAPI.EXPECT().LunList(ctx, gomock.Any()).Return(luns, nil)
	mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Return(&lun, nil)

	snap, err := d.getSnapshotEconomy(ctx, snapConfig, &d.Config)

	assert.Error(t, err)
	assert.Nil(t, snap)
}

func TestOntapSanEconomyGetLUNSize(t *testing.T) {
	mockAPI, d := newMockOntapSanEcoDriver(t)

	tests := []struct {
		message string
		size    string
	}{
		{"zero size", "0"},
		{"invalid", "invalid"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			mockAPI.EXPECT().LunGetByName(ctx, gomock.Any()).Times(1).Return(&api.Lun{Size: test.size}, nil)

			size, err := d.getLUNSize(ctx, "lun", "vol1")

			assert.Equal(t, uint64(0), size)
			assert.Error(t, err)
		})
	}
}

func TestOntapSanEconomyCreatePrepare(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)
	volConfig := &storage.VolumeConfig{
		Size:         "1g",
		Encryption:   "false",
		FileSystem:   "xfs",
		InternalName: "vol1",
	}

	d.CreatePrepare(ctx, volConfig)

	assert.NotEqual(t, "", volConfig.InternalName)
}

func TestOntapSanEconomyGetCommonConfig(t *testing.T) {
	_, d := newMockOntapSanEcoDriver(t)

	config := d.GetCommonConfig(ctx)

	assert.NotNil(t, config)
}
