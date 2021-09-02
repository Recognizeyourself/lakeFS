package settings

import (
	"context"
	"io/ioutil"
	"sync"
	"testing"

	"github.com/go-test/deep"
	"github.com/golang/mock/gomock"
	"github.com/treeverse/lakefs/pkg/block"
	"github.com/treeverse/lakefs/pkg/block/mem"
	"github.com/treeverse/lakefs/pkg/graveler"
	"github.com/treeverse/lakefs/pkg/graveler/mock"
	"github.com/treeverse/lakefs/pkg/testutil"
	"google.golang.org/protobuf/proto"
)

func TestSaveAndGet(t *testing.T) {
	ctx := context.Background()
	m := prepareTest(t, ctx)
	expectedSettings := &ExampleSettings{ExampleInt: 5, ExampleStr: "hello", ExampleMap: map[string]int32{"boo": 6}}
	err := m.Save(ctx, "example-repo", "settingKey", expectedSettings)
	testutil.Must(t, err)
	gotSettings := &ExampleSettings{}
	err = m.Get(ctx, "example-repo", "settingKey", gotSettings)
	testutil.Must(t, err)
	if diff := deep.Equal(expectedSettings, gotSettings); diff != nil {
		t.Fatal("got unexpected settings:", diff)
	}
}

func TestUpdateWithLock(t *testing.T) {
	ctx := context.Background()
	m := prepareTest(t, ctx)
	var lockStartWaitGroup sync.WaitGroup
	var lock sync.Mutex
	const IncrementCount = 20
	lockStartWaitGroup.Add(IncrementCount)
	mockBranchLocker := m.branchLock.(*mock.MockBranchLocker)
	mockBranchLocker.EXPECT().MetadataUpdater(ctx, gomock.Eq(graveler.RepositoryID("example-repo")), graveler.BranchID("main"), gomock.Any()).DoAndReturn(func(_ context.Context, _ graveler.RepositoryID, _ graveler.BranchID, f func() (interface{}, error)) (interface{}, error) {
		lockStartWaitGroup.Done()
		lockStartWaitGroup.Wait() // wait until all goroutines ask for the lock
		lock.Lock()
		retVal, err := f()
		lock.Unlock()
		return retVal, err
	}).Times(IncrementCount)
	expectedSettings := &ExampleSettings{ExampleInt: 5, ExampleStr: "hello", ExampleMap: map[string]int32{"boo": 6}}
	err := m.Save(ctx, "example-repo", "settingKey", expectedSettings)
	testutil.Must(t, err)
	settingsToEdit := &ExampleSettings{}
	update := func() {
		settingsToEdit.ExampleInt++
		settingsToEdit.ExampleMap["boo"]++
	}
	var wg sync.WaitGroup
	wg.Add(IncrementCount)
	for i := 0; i < IncrementCount; i++ {
		go func() {
			testutil.Must(t, m.UpdateWithLock(ctx, "example-repo", "settingKey", settingsToEdit, update))
			wg.Done()
		}()
	}
	wg.Wait()
	testutil.Must(t, err)
	gotSettings := &ExampleSettings{}
	err = m.Get(ctx, "example-repo", "settingKey", gotSettings)
	testutil.Must(t, err)
	expectedSettings.ExampleInt += IncrementCount
	expectedSettings.ExampleMap["boo"] += IncrementCount
	if diff := deep.Equal(expectedSettings, gotSettings); diff != nil {
		t.Fatal("got unexpected settings:", diff)
	}
}

func TestStoredSettings(t *testing.T) {
	ctx := context.Background()
	m := prepareTest(t, ctx)
	expectedSettings := &ExampleSettings{ExampleInt: 5, ExampleStr: "hello", ExampleMap: map[string]int32{"boo": 6}}
	err := m.Save(ctx, "example-repo", "settingKey", expectedSettings)
	testutil.Must(t, err)
	reader, err := m.blockAdapter.Get(ctx, block.ObjectPointer{
		StorageNamespace: "mem://my-storage",
		Identifier:       "/_lakefs/settings/settingKey.json",
		IdentifierType:   block.IdentifierTypeRelative,
	}, -1)
	testutil.Must(t, err)
	bytes, err := ioutil.ReadAll(reader)
	testutil.Must(t, err)
	gotSettings := &ExampleSettings{}
	testutil.Must(t, proto.Unmarshal(bytes, gotSettings))
	if diff := deep.Equal(expectedSettings, gotSettings); diff != nil {
		t.Fatal("got unexpected settings:", diff)
	}
}

func prepareTest(t *testing.T, ctx context.Context) *Manager {
	ctrl := gomock.NewController(t)
	refManager := mock.NewMockRefManager(ctrl)
	repo := &graveler.Repository{
		StorageNamespace: "mem://my-storage",
		DefaultBranchID:  "main",
	}
	blockAdapter := mem.New()
	branchLock := mock.NewMockBranchLocker(ctrl)
	m := &Manager{
		refManager:                  refManager,
		branchLock:                  branchLock,
		blockAdapter:                blockAdapter,
		committedBlockStoragePrefix: "_lakefs",
	}
	refManager.EXPECT().GetRepository(ctx, gomock.Eq(graveler.RepositoryID("example-repo"))).AnyTimes().Return(repo, nil)
	return m
}
