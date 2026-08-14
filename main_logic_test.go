package main

import (
	"strings"
	"testing"
)

func launcherByID(id string) *Launcher {
	ls := launchers()
	for i := range ls {
		if ls[i].ID == id {
			return &ls[i]
		}
	}
	return &ls[0]
}

func idxOf(cmd string) int {
	cmds := commands()
	for i := range cmds {
		if cmds[i].Command == cmd {
			return i
		}
	}
	return -1
}

func TestCombinationSteam(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, launcher: launcherByID("steam")}
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	want := "PROTON_LOG=1 mangohud %command%"
	if got != want {
		t.Fatalf("steam: got %q want %q", got, want)
	}
}

func TestCombinationNoCmd(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, launcher: launcherByID("heroic")}
	g.selected[idxOf("mangohud %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	want := "PROTON_LOG=1 mangohud"
	if got != want {
		t.Fatalf("heroic: got %q want %q", got, want)
	}
}

func TestGamescopeNoCmdDropsTrailingDashDash(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, launcher: launcherByID("heroic")}
	g.selected[idxOf("gamescope -e -f -F fsr -- %command%")] = true
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	got := g.buildCombination()
	if strings.HasSuffix(got, "--") || strings.Contains(got, " -- ") {
		t.Fatalf("heroic gamescope: linha quebrada com --: %q", got)
	}
	want := "PROTON_LOG=1 gamescope -e -f -F fsr"
	if got != want {
		t.Fatalf("heroic gamescope: got %q want %q", got, want)
	}
	gsteam := &gui{all: commands(), selected: map[int]bool{}, launcher: launcherByID("steam")}
	gsteam.selected[idxOf("gamescope -e -f -F fsr -- %command%")] = true
	if got := gsteam.buildCombination(); !strings.Contains(got, " -- ") {
		t.Fatalf("steam gamescope: esperado -- preservado, got %q", got)
	}
}

func TestDisplayCmdStrips(t *testing.T) {
	g := &gui{all: commands(), launcher: launcherByID("lutris")}
	c := commands()[idxOf("PROTON_LOG=1 %command%")]
	if got := g.displayCmd(c); got != "PROTON_LOG=1" {
		t.Fatalf("displayCmd: got %q", got)
	}
	g2 := &gui{all: commands(), launcher: launcherByID("steam")}
	if got := g2.displayCmd(c); got != "PROTON_LOG=1 %command%" {
		t.Fatalf("displayCmd steam: got %q", got)
	}
}

func TestConflictDuplicate(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, lang: "pt"}
	g.selected[idxOf("PROTON_LOG=1 %command%")] = true
	g.selected[idxOf("PROTON_LOG=warn+pipewire,warn+mmdevapi %command%")] = true
	warns := g.conflicts()
	if len(warns) == 0 {
		t.Fatal("expected duplicate-key conflict")
	}
	if !strings.Contains(warns[0], "PROTON_LOG") {
		t.Fatalf("unexpected warning: %q", warns[0])
	}
}

func TestConflictAntiLagReflex(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, lang: "pt"}
	g.selected[idxOf("LOW_LATENCY_LAYER=1 %command%")] = true
	g.selected[idxOf(`LOW_LATENCY_LAYER=1 LOW_LATENCY_LAYER_REFLEX=1 DXVK_CONFIG="dxgi.hideAmdGpu = True" %command%`)] = true
	warns := g.conflicts()
	found := false
	for _, w := range warns {
		if strings.Contains(w, "Reflex") {
			found = true
		}
	}
	if !found {
		t.Fatalf("expected anti-lag/reflex conflict, got %v", warns)
	}
}

func TestNoConflictSameValue(t *testing.T) {
	g := &gui{all: commands(), selected: map[int]bool{}, lang: "pt"}
	g.selected[idxOf("WINE_ESYNC=1 %command%")] = true
	g.selected[idxOf("WINEFSYNC=1 %command%")] = true
	if warns := g.conflicts(); len(warns) != 0 {
		t.Fatalf("expected no conflict, got %v", warns)
	}
}
