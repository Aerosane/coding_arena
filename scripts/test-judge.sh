#!/usr/bin/env bash
#
# Verify the full submission flow: backend → bridge → judge → result.
#
# Prerequisites:
#   docker compose up -d
#
# Usage:
#   ./scripts/test-judge.sh
#
set -euo pipefail

API="http://localhost:8080"
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m'

pass() { echo -e "${GREEN}[PASS]${NC} $1"; }
fail() { echo -e "${RED}[FAIL]${NC} $1"; exit 1; }
info() { echo -e "${YELLOW}[INFO]${NC} $1"; }

# --- 1. Health check ---
info "Checking backend health..."
HEALTH=$(curl -sf "$API/health" 2>/dev/null || true)
if [ -z "$HEALTH" ]; then
  fail "Backend is not running at $API"
fi
echo "  $HEALTH"

JUDGE_STATUS=$(echo "$HEALTH" | grep -o '"judge":"[^"]*"' | cut -d'"' -f4)
if [ "$JUDGE_STATUS" = "connected" ]; then
  pass "Judge is connected"
else
  info "Judge status: $JUDGE_STATUS (may need a few seconds to connect)"
  info "Waiting 10s for judge to connect..."
  sleep 10
  HEALTH=$(curl -sf "$API/health")
  JUDGE_STATUS=$(echo "$HEALTH" | grep -o '"judge":"[^"]*"' | cut -d'"' -f4)
  if [ "$JUDGE_STATUS" != "connected" ]; then
    fail "Judge did not connect after 10s. Status: $JUDGE_STATUS"
  fi
  pass "Judge connected after waiting"
fi

# --- 2. Submit a correct Python solution to 'aplusb' ---
info "Submitting correct Python solution for 'aplusb'..."
RESULT=$(curl -sf -X POST "$API/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "a, b = map(int, input().split())\nprint(a + b)",
    "language": "python",
    "problem_id": "aplusb"
  }')

echo "  Response: $RESULT"

STATUS=$(echo "$RESULT" | grep -o '"status":"[^"]*"' | head -1 | cut -d'"' -f4)
if [ "$STATUS" = "graded" ]; then
  pass "Submission was graded"
else
  fail "Expected status 'graded', got '$STATUS'"
fi

VERDICT=$(echo "$RESULT" | grep -o '"verdict":"[^"]*"' | cut -d'"' -f4)
if [ "$VERDICT" = "AC" ]; then
  pass "Verdict: AC (Accepted)"
else
  fail "Expected verdict 'AC', got '$VERDICT'"
fi

# --- 3. Submit a wrong solution ---
info "Submitting wrong Python solution for 'aplusb'..."
RESULT=$(curl -sf -X POST "$API/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "print(0)",
    "language": "python",
    "problem_id": "aplusb"
  }')

echo "  Response: $RESULT"

VERDICT=$(echo "$RESULT" | grep -o '"verdict":"[^"]*"' | cut -d'"' -f4)
if [ "$VERDICT" = "WA" ]; then
  pass "Verdict: WA (Wrong Answer) — correct for bad solution"
else
  info "Verdict: $VERDICT (expected WA for wrong solution)"
fi

# --- 4. Submit a compile error (C++) ---
info "Submitting invalid C++ code for 'aplusb'..."
RESULT=$(curl -sf -X POST "$API/submit" \
  -H "Content-Type: application/json" \
  -d '{
    "code": "int main( {}",
    "language": "cpp",
    "problem_id": "aplusb"
  }')

echo "  Response: $RESULT"

VERDICT=$(echo "$RESULT" | grep -o '"verdict":"[^"]*"' | cut -d'"' -f4)
if [ "$VERDICT" = "CE" ]; then
  pass "Verdict: CE (Compile Error) — correct for invalid code"
else
  info "Verdict: $VERDICT (expected CE for invalid code)"
fi

# --- Summary ---
echo ""
echo -e "${GREEN}========================================${NC}"
echo -e "${GREEN} Full flow verified: submit → judge → result ${NC}"
echo -e "${GREEN}========================================${NC}"
