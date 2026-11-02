#!/usr/bin/env bash
# Trade-A-Plane Listing API – cURL example
# Usage: PILOTERR_API_KEY=your-key ./listing.sh
set -euo pipefail

curl -sS --fail-with-body -G "https://api.piloterr.com/v2/tradeaplane/listing" \
  -H "x-api-key: ${PILOTERR_API_KEY}" \
  --data-urlencode 'query=https://www.trade-a-plane.com/search?category_level1=Piston+Engine&make=ROLLS+ROYCE&model=PACKARD+V-1650&listing_id=2419556&s-type=engine'
