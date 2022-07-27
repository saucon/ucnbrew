package boilerplate

func createFiles(pkgname string, appname string) (err error) {
	if err = createMain(pkgname, appname); err != nil {
		return err
	}
	if err = createPkgEnv(); err != nil {
		return err
	}
	if err = createPkgEnvConfig(); err != nil {
		return err
	}
	if err = createPkgLogConfig(pkgname); err != nil {
		return err
	}
	if err = createPkgLogDbConfig(); err != nil {
		return err
	}
	if err = createPkgDbConfig(pkgname); err != nil {
		return err
	}
	if err = createRouter(); err != nil {
		return err
	}
	if err = createEnv(); err != nil {
		return err
	}
	return nil
}
