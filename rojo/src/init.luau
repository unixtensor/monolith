--!strict

local HttpService = game:GetService("HttpService")
local Data = require(script.data)

local Monolith = {} :: Constructor
Monolith.__index = Monolith
Monolith.DEV_MODE = true

type Constructor = {
	__index: Constructor;
	DEV_MODE: boolean;
	new: (address: string) -> Monolith;
	listen: (self: Monolith) -> ()
}
type Self = {
	data: Data.Data;
	address: string;
}
export type Monolith = setmetatable<Self, typeof(Monolith)>

function Monolith.new(address)
	if not Monolith.DEV_MODE and not HttpService.HttpEnabled then
		error("Cannot proceed! Monolith REQUIRES http requests to be enabled for this game. SERVER STOPPING", 0)
	end

	local self = {}
	self.data = Data.new()
	self.address = address
	return setmetatable(self, Monolith)
end

function Monolith:listen()

end

return Monolith
