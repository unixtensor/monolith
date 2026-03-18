--!strict

local HttpService = game:GetService("HttpService")

local Http = {} :: Constructor
Http.__index = Http

type Constructor = {
	__index: Constructor;
	new: (address: string) -> Http;
}

export type Http = setmetatable<Constructor, typeof(Http)>

function Http.new()

end

return Http
