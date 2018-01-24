@echo off

set nodeModulesFolder=f:/goprogs/src/laatoo/ui/nodemodules
set pluginsRoot=f:/goprogs/
set deploy=f:/goprogs/src/laatoo/modules

setlocal EnableDelayedExpansion 
set argC=0
set modCount=0
for %%x in (%*) do ( 
	if !argC! GTR 0 (
		set modsToProcess[!argC!]=%%x
		set /A modCount+=1
	)
	set /A argC+=1
)

IF %argC% LSS 1 (
	echo Modules file has not been provided.
	exit /B 1
)
	

set modFile=%1
shift

set /A i=0

for /F "usebackq delims=" %%a in ("%modFile%") do (
	set modEntry=%%a
	for /f "tokens=1" %%G IN ("!modEntry!") DO set moduleName=%%G 
	for /f "tokens=2" %%G IN ("!modEntry!") DO set modulePath=%%G 
	set moduleName=!moduleName: =!
	if %argC%==1 (
		call :compilemodule !moduleName! !modulePath!
	) else (
		for /l %%n in (1,1,%modCount%) do ( 
			set mod=!modsToProcess[%%n]!
			if !moduleName!==!mod! (
				call :compilemodule !moduleName! !modulePath!
			)
		)
	)
)


exit /B 0

:compilemodule
echo Compiling module %1
docker run --rm -it -v %nodeModulesFolder%:/nodemodules -v %pluginsRoot%:/plugins -v %deploy%:/deploy -e name=%1 -e packageFolder=%2 -e verbose=true  laatoomodulebuilder
echo ================================================================
EXIT /B 0