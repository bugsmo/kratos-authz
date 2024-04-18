package casbin

import (
	"github.com/bugsmo/kratos-authz/engine"
	stdCasbin "github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
)

type State struct {
	model    model.Model
	policy   *Adapter
	enforcer *stdCasbin.SyncedEnforcer
	projects engine.Projects

	wildcarItem               string
	authorizedProjectsMatcher string
}


