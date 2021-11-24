package tests

import (
	"github.com/og-zero/arena/pkg/quests/objectives"
	"github.com/og-zero/arena/pkg/zones"
)

func testImports() {
	// Test imports
	//
	objective := objectives.New()
	objective.Start()
	//
	zone := zones.New()
	zone.ID()
	//
}
