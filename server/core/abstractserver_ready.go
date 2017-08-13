package core

/*****
this method is only for internal service based tasks
******/
func (as *abstractserver) onReady(ctx *serverContext) error {
	if as.securityHandlerHandle != nil {
		return as.securityHandlerHandle.(*securityHandler).onServerReady(ctx)
	}
	return nil
}
