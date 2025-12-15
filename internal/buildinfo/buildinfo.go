package buildinfo

import "os"

var (
	AppEnvLd    = ""
	VersionLd   = ""
	BuildTimeLd = ""
	CommitSHALd = ""
)

var (
	AppEnv    = value(AppEnvLd, os.Getenv("APP_ENV"), "unknown")
	Version   = value(VersionLd, os.Getenv("VERSION"), "unknown")
	BuildTime = value(BuildTimeLd, os.Getenv("BUILD_TIME"), "unknown")
	CommitSHA = value(CommitSHALd, os.Getenv("GIT_COMMIT"), "unknown")
)

func value(values ...string) string {
	for _, v := range values {
		if v != "" && v != "unknown" {
			return v
		}
	}
	return "unknown"
}