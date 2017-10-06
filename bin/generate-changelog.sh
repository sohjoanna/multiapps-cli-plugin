#!/usr/bin/env bash

PREVIOUS_VERSION=$1
CURRENT_VERSION=$2

if [ "$#" -ne 2 ]; then
    cat <<-INFO
NAME:
  generate-changelog - Generate changelog relative to a given version
USAGE:
  generate-changelog PREVIOUS_VERSION CURRENT_VERSION
EXAMPLE:
  generate-changelog v6.2.0 v6.3.0
INFO

    exit 1
fi

git --no-pager log --grep \[.*\d*\] $PREVIOUS_VERSION..$CURRENT_VERSION --pretty='format:* %s'
