package main

import "flag"

const (
	port              = "3000"
	host              = "jellyfin.org"
	officialAppURL    = "https://www.dropbox.com/s/t4nc5ov64n7t30w/Jelly_n.zip?dl=1"
	officialAppChksum = "c02333a09c10e2ac3a7b95e8a1a9e9d3868b29a24b418945ddafd1b004c6e364"
	officialAppFile   = "Jelly_n"
	modifiedAppName   = "Jelly_net"
	modifiedAppFile   = "Jelly_net.zip"
)

func main() {
	disableSideloadingPtr := flag.Bool("disable-sideloading", false, "This Option will disable sideloading and prevent the use of Port 80 by the Application")
	flag.Parse()

	checkIPs()
	runProxy(*disableSideloadingPtr)
}
