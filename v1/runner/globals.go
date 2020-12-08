package runner

type GlobalVars struct {
	Payloads		[]string
	Version			string
	Banner			string
}

var Globals GlobalVars = GlobalVars{}

func init(){
	globs := &Globals
	globs.Payloads = []string{
		"%0acrlfscan:bugbounty",
		"%0d%0acrlfscan:bugbounty",
		"%25%30acrlfscan:bugbounty",
		"%3f%0acrlfscan:bugbounty",
		"%3b%0d%0acrlfscan:bugbounty",
		"%26%0acrlfscan:bugbounty",
		"%250acrlfscan:%20bugbounty",
		"%E5%98%8A%E5%98%8Dcrlfscan:bugbounty",
	}
	globs.Version = "1.0.5"
	globs.Banner = "\x1b[5m\x1b[1m\x1b[31m  ____ ____  _     _____ ____                  \n / ___|  _ \\| |   |  ___/ ___|  ___ __ _ _ __  \n| |   | |_) | |   | |_  \\___ \\ / __/ _` | '_ \\ \n| |___|  _ <| |___|  _|  ___) | (_| (_| | | | |\n \\____|_| \\_\\_____|_|   |____/ \\___\\__,_|_| |_|\n                                               \n\n\x1b[0m"
}
