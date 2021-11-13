package terminal

import "os"

// from https://github.com/watson/ci-info/blob/master/vendors.json
var ciEnvs = []string{
	"APPVEYOR",
	"SYSTEM_TEAMFOUNDATIONCOLLECTIONURI",
	"bamboo_planKey",
	"BITBUCKET_COMMIT",
	"BITRISE_IO",
	"BUDDY_WORKSPACE_ID",
	"BUILDKITE",
	"CIRCLECI",
	"CIRRUS_CI",
	"CODEBUILD_BUILD_ARN",
	"CI_NAME",
	"DRONE",
	"DSARI",
	"GITLAB_CI",
	"GO_PIPELINE_LABEL",
	"HUDSON_URL",
	"JENKINS_URL",
	"BUILD_ID",
	"MAGNUM",
	"NETLIFY_BUILD_BASE",
	"NEVERCODE",
	"SAILCI",
	"SEMAPHORE",
	"SHIPPABLE",
	"TDDIUM",
	"STRIDER",
	"TASK_ID",
	"RUN_ID",
	"TEAMCITY_VERSION",
	"TRAVIS",
	"GITHUB_ACTIONS",
	"NOW_BUILDER",
	"APPCENTER_BUILD_ID",
}

func IsCI() bool {
	if os.Getenv("DEBIAN_FRONTEND") == "noninteractive" {
		return true
	}

	for _, env := range ciEnvs {
		if _, hasEnv := os.LookupEnv(env); hasEnv {
			return true
		}
	}

	return false
}
