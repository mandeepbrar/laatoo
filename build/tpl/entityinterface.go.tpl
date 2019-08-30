package main

import (
  "laatoo/sdk/server/components/data"
)

type I{{#type name}}{{/type}} interface {
    data.Auditable 
    {{#fieldFuncDefs fields}}{{/fieldFuncDefs}}
}

