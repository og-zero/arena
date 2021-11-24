package tests

import (
	"github.com/og-zero/arena/pkg/quests/objectives"
	_ "github.com/og-zero/arena/pkg/server/config"
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
