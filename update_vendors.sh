#!/usr/bin/env bash

dep ensure -update github.com/seagullbird/headr-common
dep ensure -update github.com/seagullbird/headr-repoctl
dep ensure -update github.com/seagullbird/headr-contentmgr
dep ensure -update github.com/seagullbird/headr-sitemgr
git add .
git cm -m 'updated vendors'
git push

