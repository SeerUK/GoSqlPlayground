package database

import (
	"fmt"
	"sort"
)

var versionKeys []int
var versions map[int]Version = make(map[int]Version)

func RegisterVersion(version Version) {
	versionNo := version.Number()

	versionKeys = append(versionKeys, versionNo)
	versions[versionNo] = version
}

// @todo: Needs sql.DB
func Migrate() {
	sort.Ints(versionKeys)

	fmt.Println(versionKeys)
}
