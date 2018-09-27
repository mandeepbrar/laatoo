package core

/*
//service method for doing various tasks
func NewService(ctx ServerContext, alias string, servFunc ServiceFunc, svcInfo ServiceInfo) Service {
	if servFunc != nil {
		return &serviceImpl{name: alias, servFunc: servFunc, info: svcInfo}
	}
	return nil
}

type serviceImpl struct {
	Service
	name     string
	servFunc ServiceFunc
	info     ServiceInfo
}

func (svc *serviceImpl) GetName() string {
	return svc.name
}
func (svc *serviceImpl) Info() ServiceInfo {
	return svc.info
}

func (svc *serviceImpl) Invoke(ctx RequestContext, request Request) (*Response, error) {
	return svc.servFunc(ctx, request)
}

/*
func (gc *serviceImpl) GetResponseHandler() core.ServiceResponseHandler {
	return nil
}
func createMW(conf config.Config, parentMW *Middleware) *Middleware {
	var retVal Middleware
	retVal = []string{}
	if parentMW != nil {
		retVal = append(retVal, *parentMW...)
	}
	mw, ok := conf.GetStringArray(CONF_MIDDLEWARE)
	if ok {
		retVal = append(retVal, mw...)
	}
	return &retVal
}



func (env *Environment) processMiddleware(ctx *serverContext, servicesStore map[string]core.Service, serviceMiddleware map[string]*Middleware) error {
	//process middleware
	for svcAlias, mw := range serviceMiddleware {
		listmw := *mw
		lenmw := len(listmw)
		if lenmw > 0 { // only if there is a middleware configured
			svc := servicesStore[svcAlias]
			targetSvc, ok := svc.(core.GenericService)
			if !ok {
				return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Middleware supported only for Generic Service", svcAlias)
			}
			gen_mw_svcs := make([]core.GenericService, lenmw)
			for i := 0; i < lenmw; i++ {
				mwname := listmw[i]
				mwsvc, ok := servicesStore[mwname]
				if !ok {
					return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Could not find middleware service for alias", svcAlias, "MW", mwname)
				}
				gen_mw_svc, ok := mwsvc.(core.GenericService)
				if !ok {
					return errors.ThrowError(ctx, CORE_ERROR_SERVICE_CREATION, "Middleware supported only for Generic Service", svcAlias, "MW", mwname)
				}
				gen_mw_svcs[i] = gen_mw_svc
			}

			//add the service to the environment
			app.ServicesStore[svcAlias] = func(mwsvcs []core.GenericService, gensvc core.GenericService) core.Service {
				log.Debug(ctx, "Registered service", "service name", svcAlias)
				return func(svcctx core.RequestContext) error {
					for _, mwsvc := range mwsvcs {
						err := mwsvc(svcctx)
						if err != nil {
							return errors.WrapError(ctx, err)
						}
					}
					return gensvc(svcctx)
				}
			}(gen_mw_svcs, targetSvc)
		}
	}
	return nil
}


*/