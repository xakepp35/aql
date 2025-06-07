#!/usr/bin/env bash
set -euo pipefail

############################################
# Configurable bits
############################################
DATASET="${DATASET:-1G}"          # 1G, 2G, 4G …
NAME="data_${DATASET}.ndjson"
JQ_BIN="${JQ_BIN:-jq}"            # path to jq 1.6
ZQ_BIN="${ZQ_BIN:-zq}"            # path to zq
AQ_BIN="${AQ_BIN:-aq}"            # path to your AQL CLI
############################################

bytes() { numfmt --from=iec "$1"; }

# ------------------------------------------------------------------------------
# 1. Generate dataset (if not cached)
# ------------------------------------------------------------------------------
if [ ! -f "$NAME" ]; then
  echo "Generating ${DATASET} dataset → $NAME"
  TARGET=$(bytes "$DATASET")
  rm -f "$NAME"
  i=0
  # each record is 13–17 B, so loop until file ≥ TARGET
  while [ "$(wc -c <"$NAME" 2>/dev/null || echo 0)" -lt "$TARGET" ]; do
    printf '{"value":%d}\n' "$i" >>"$NAME"
    ((i=(i+1)%1000))
  done
  du -h "$NAME"
fi

# helper: run cmd, capture time & RSS
run() {
  local label=$1; shift
  echo -e "\n▶ $label"
  /usr/bin/time -f "  ↳ %E elapsed, %M KB max RSS" "$@"
}

# ------------------------------------------------------------------------------
# 2. Benchmarks
# ------------------------------------------------------------------------------

# jq: need slurp + add
run "jq"  "$JQ_BIN"   -s 'map(.value) | add'          <"$NAME" >/dev/null

# zq: stateful aggregate
run "zq"  "$ZQ_BIN"   -j 'sum(value)'                 "$NAME"  >/dev/null

# aq: stateful aggregate (streaming)
run "aq"  "$AQ_BIN"   'sum(.value)'                   "$NAME"  >/dev/null
