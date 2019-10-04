#!/bin/bash

source .laatoomodulesrc

if [[ ($# < 1) ]]
  then
  echo Modules file has not been provided
  exit
  fi

modFile=$1
shift

uionly=false
verbose=false
getBuildPackages=false
nobundle=false
norust=false
overwriteJSMods=false

while [[ "$#" -gt 0 ]]; 
do case $1 in
  --uionly) uionly=true; shift;;
  --verbose) verbose=true; shift;;
  --nobundle) nobundle=true; shift;;
  --norust) norust=true; shift;;
  --getBuildPackages) getBuildPackages=true; shift;;
  --overwriteJSMods) overwriteJSMods=true; shift;;
  *) break;;
esac; 
done

compile_module() {
  docker run --rm -it -v $nodeModulesFolder:/nodemodules -v $modulesRepo:/modulesrepo -v $goModulesRepo:/laatoo/sdk/modules -v $tmpFolder:/compiletmp -v $deploy:/deploy -e name=$1 -e release=false -e norust=$norust -e uionly=$uionly -e packageFolder=$2 -e verbose=$verbose  -e nobundle=$nobundle -e overwriteJSMods=$overwriteJSMods -e getBuildPackages=$getBuildPackages laatoomodulebuilder
  echo '================================================================'
}

readarray  modules < $modFile
length=${#modules[*]}

modsToCreate=( "$@" )
#create all modules if no arguments have been provided
createAll=$(($#==0))

for ((i=0; i < $length; i++))
  do
    array=(${modules[i]//\ / })
    moduleName=${array[0]}
    moduleFolder=${array[1]}
    processModule=0
    if [[ ( $createAll == 0 )]]
    then
      for element in "${modsToCreate[@]}"; do
          if [[ $element == $moduleName ]]; then
              processModule=1
              break
          fi
      done
    else
      processModule=1
    fi
    if [[ $processModule == 1 ]]
    then
      echo Compiling module $moduleName
      compile_module  $moduleName $moduleFolder
    fi
  done
