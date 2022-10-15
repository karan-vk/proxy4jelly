package main

import "flag"

const (
	port              = "3000"
	host              = "jellyfin.org"
	officialAppURL    = "https://www.dropbox.com/s/bvmf4vu8fcbzrh4/Jelly_n.zip?dl=1"
	officialAppChksum = "eb89ce17916f874a9f5f72c53e0593ee56f8a2d34f14b882e0ad6e5f5bc52563"
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
