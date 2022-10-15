package main

import "flag"

const (
	port              = "3000"
	host              = "jellyfin.org"
	officialAppURL    = "https://www.dropbox.com/s/kshh1k2tgt0yjs4/Jelly.zip?dl=1"
	officialAppChksum = "62dfce80af7a120ec63a58c9e226980ba9670288916558e48f4274ae1f2ed566"
	officialAppFile   = "Jelly"
	modifiedAppName   = "Jelly_net"
	modifiedAppFile   = "Jelly_net.zip"
)

func main() {
	disableSideloadingPtr := flag.Bool("disable-sideloading", false, "This Option will disable sideloading and prevent the use of Port 80 by the Application")
	flag.Parse()

	checkIPs()
	runProxy(*disableSideloadingPtr)
}
