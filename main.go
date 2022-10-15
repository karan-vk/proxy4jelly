package main

import "flag"

const (
	port              = "3000"
	host              = "jellyfin.org"
	officialAppURL    = "https://www.dropbox.com/s/zn66d1f0mj2ok8h/Jelly_n.zip?dl=1"
	officialAppChksum = "d24a668ca5639dac6bbe75bb85eec743292b8096da67b2df5ebd8304bd3f2b4e"
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
