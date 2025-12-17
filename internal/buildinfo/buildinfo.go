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
	Version   = value(VersionLd, os.Getenv("BACK_VERSION"), "unknown")
	BuildTime = value(BuildTimeLd, os.Getenv("BACK_BUILD_TIME"), "unknown")
	CommitSHA = value(CommitSHALd, os.Getenv("BACK_COMMIT_SHA"), "unknown")
)

func value(values ...string) string {
	for _, v := range values {
		if v != "" && v != "unknown" {
			return v
		}
	}
	return "unknown"
}