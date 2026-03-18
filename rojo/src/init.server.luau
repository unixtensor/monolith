local DEV_MODE: boolean = true
local HttpService = game:GetService("HttpService")

local Data = require(script:WaitForChild("data"))

if not DEV_MODE and not HttpService.HttpEnabled then
	error("Cannot proceed! Monolith REQUIRES http requests to be enabled for this game. SERVER STOPPING", 0)
end

local data_collection = Data.new()
print(data_collection:collect())
