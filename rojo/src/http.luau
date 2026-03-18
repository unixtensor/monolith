--!strict

local HttpService = game:GetService("HttpService")
local Data = require(script.Parent.data)

local Http = {} :: Constructor
Http.__index = Http
Http.Enabled = HttpService.HttpEnabled
Http.ApiEndpoint = {
	CONNECT = "connect";
	UPLOAD = "upload";
	GAME_DATA = "game_data";
}

type Api = {
	["CONNECT"]: "connect";
	["UPLOAD"]: "upload";
	["GAME_DATA"]: "game_data";
}

type JsonPayload = {
	[string]: any
}

type Constructor = {
	__index: Constructor;
	Enabled: boolean;
	ApiEndpoint: Api;
	new: (address: string, compression: Enum.HttpCompression) -> Http;
	ValidAddress: (address: string) -> boolean;
	PostAsync: (self: Http, payload: JsonPayload, request_endpoint: string) -> ();
	UploadAsync: (self: Http, payload: Data.Payload) -> ();
}
type Self = {
	address: string;
	compression: Enum.HttpCompression;
}
export type Http = setmetatable<Self, typeof(Http)>

local function MakeApiRequest(address: string, request: string): string
	return `{address}/api/v1/{request}`
end

local function PostBody(self: Http, is_game_data: boolean, json_payload: string, request: string): any
	return {
		Url = MakeApiRequest(self.address, request),
		Method = "POST",
		Headers = {
			["Content-Type"] = "application/json",
		},
		Body = json_payload,
		Compress = is_game_data and self.compression or Enum.HttpCompression.None
	}
end

function Http.ValidAddress(address)
	return address:match("^https?://") ~= nil
end

function Http.new(address, compression)
	local self = {}
	self.address = address
	self.compression = compression
	return setmetatable(self, Http)
end

function Http:UploadAsync(payload)
	local json_payload = HttpService:JSONEncode(payload)
	HttpService:RequestAsync(PostBody(self, true, json_payload, Http.ApiEndpoint.UPLOAD))
end

function Http:PostAsync(payload, request_endpoint)
	local json_payload = HttpService:JSONEncode(payload)
	HttpService:RequestAsync(PostBody(self, false, json_payload, request_endpoint))
end

return Http
