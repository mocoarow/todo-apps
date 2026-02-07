#!/bin/bash
set -euo pipefail

USER_ID="${RANDOM}"
LOGIN_ID="user${USER_ID}" PASSWORD="password${USER_ID}" runn run runn.yml --debug --scopes run:exec
