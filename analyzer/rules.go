package analyzer

var suspiciousKeywords = []string{
	"passwd", "root", "injection", "exploit",
	"admin", "sudo", "unauthorized", "forbidden",
}

const (
	burstThreshold  = 10 // errors in this many seconds
	burstWindow     = 30
	repeatThreshold = 5 // same message this many times
)
