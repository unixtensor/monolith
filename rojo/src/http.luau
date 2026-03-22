--!strict

local HttpService = game:GetService("HttpService")
local Data = require(script.Parent.data)

local Http = {} :: Constructor
Http.__index = Http
Http.__bypassInsecure = false
Http.Insecure = Http.__bypassInsecure
Http.Enabled = HttpService.HttpEnabled

type Json = {
	[string]: any
}
type RestDetails = {
	method: "POST" | "GET";
	json: string;
	request: string;
}

type Constructor = {
	__index: Constructor;
	__bypassInsecure: boolean;
	Enabled: boolean;
	Insecure: boolean;
	new: (address: string, token: string, enable_compression: boolean?) -> Http;
	ValidAddress: (address: string) -> string?;
	PostAsync: (self: Http, payload: Json, request_endpoint: string) -> ();
	UploadAsync: (self: Http, payload: Data.Payload) -> ();
}
type Self = {
	address: string;
	token: string;
	compression: Enum.HttpCompression;
}
export type Http = setmetatable<Self, typeof(Http)>

local api_root: string = "/api/v1"

local function restful_body(self: Http, is_game_data: boolean, rest: RestDetails): Json
	local payload = {
		Url = self.address .. api_root .. rest.request;
		Method = rest.method;
		Headers = {
			["Content-Type"] = "application/json";
			["x-api-key"] = self.token;
		},
		Body = rest.json;
	}
	if is_game_data then
		print(self.compression)
		payload.Compress = self.compression
	end
	return payload
end

local function local_address(host: string): string?
	return host:match("^localhost")
		or host:match("^0%.0%.0%.0$")
		or host:match("^192%.168%.")
		or host:match("^10%.")
		or host:match("^127%.")
		or host:match("^172%.%d+%.")
end

function Http.ValidAddress(address)
	local http_compliant = address:match("^https?://")
	if http_compliant then
		local domain = address:match("https?://([^/]+)")
		if not domain then
			return nil
		end
		if not local_address(domain) and http_compliant == "http://" then
			Http.Insecure = true
		end
		--Trim trailing slash
		return (address:gsub("/+$", ""))
	end
	return nil
end

function Http.new(address, token, enable_compression)
	if Http.Insecure then
		if not Http.__bypassInsecure then
			error("An insecure address is being used. STOPPING", 0)
		end
		warn("Monolith has been forcefully connected using a public insecure address. You're open to man-in-the-middle attacks!")
	end

	local self = {}
	self.address = address
	self.token = token
	self.compression = enable_compression and Enum.HttpCompression.Gzip or Enum.HttpCompression.None
	return setmetatable(self, Http)
end



return Http
