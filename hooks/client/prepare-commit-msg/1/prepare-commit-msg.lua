#!/usr/bin/env lua

local arg = arg
local io = io

local commit_msg_filepath = arg[1] or ""
local handle = io.open(commit_msg_filepath, "wb")
if handle then
    handle:write("# Please include a useful commit message!")
    handle:close()
end

