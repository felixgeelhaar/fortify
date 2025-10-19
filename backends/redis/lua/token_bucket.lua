--[[
Token Bucket Rate Limiting Script for Redis

This script implements atomic token bucket rate limiting with automatic refill.
It handles token refill based on elapsed time and attempts to consume tokens in a single atomic operation.

KEYS[1]: bucket key (e.g., "fortify:ratelimit:user:123")

ARGV[1]: rate (number of tokens added per interval)
ARGV[2]: burst (maximum tokens in bucket)
ARGV[3]: interval (nanoseconds between token additions)
ARGV[4]: tokens_to_take (number of tokens to consume, typically 1)
ARGV[5]: current_time_ns (current timestamp in nanoseconds)
ARGV[6]: ttl_seconds (bucket TTL in seconds for auto-cleanup)

Returns:
  1 = allowed (tokens available)
  0 = denied (insufficient tokens)

Data Structure:
  The bucket is stored as a Redis hash with two fields:
  - tokens: current number of available tokens (float)
  - last_refill: last refill timestamp in nanoseconds (integer)
--]]

local bucket_key = KEYS[1]

local rate = tonumber(ARGV[1])
local burst = tonumber(ARGV[2])
local interval_ns = tonumber(ARGV[3])
local tokens_to_take = tonumber(ARGV[4])
local current_time_ns = tonumber(ARGV[5])
local ttl_seconds = tonumber(ARGV[6])

-- Get current bucket state
local bucket = redis.call('HMGET', bucket_key, 'tokens', 'last_refill')
local current_tokens = tonumber(bucket[1])
local last_refill_ns = tonumber(bucket[2])

-- Initialize new bucket if it doesn't exist
if not current_tokens or not last_refill_ns then
    current_tokens = burst
    last_refill_ns = current_time_ns
end

-- Calculate elapsed time since last refill
local elapsed_ns = current_time_ns - last_refill_ns

-- Safety check: prevent negative elapsed time (clock skew)
if elapsed_ns < 0 then
    elapsed_ns = 0
end

-- Cap elapsed time to prevent overflow (max 1 hour)
local max_elapsed_ns = 3600 * 1000000000  -- 1 hour in nanoseconds
if elapsed_ns > max_elapsed_ns then
    elapsed_ns = max_elapsed_ns
end

-- Calculate tokens to add based on elapsed time
-- tokens_to_add = (elapsed / interval) * rate
local tokens_to_add = 0
if interval_ns > 0 and rate > 0 then
    tokens_to_add = (elapsed_ns / interval_ns) * rate
end

-- Refill tokens
current_tokens = current_tokens + tokens_to_add

-- Cap at burst limit
if current_tokens > burst then
    current_tokens = burst
end

-- Attempt to take tokens
local allowed = 0
if current_tokens >= tokens_to_take then
    current_tokens = current_tokens - tokens_to_take
    allowed = 1
end

-- Update bucket state
redis.call('HSET', bucket_key,
    'tokens', tostring(current_tokens),
    'last_refill', tostring(current_time_ns))

-- Set TTL to auto-expire unused buckets
redis.call('EXPIRE', bucket_key, ttl_seconds)

return allowed
