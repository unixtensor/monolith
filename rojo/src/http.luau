--!strict

local HttpService = game:GetService("HttpService")

local Http = {} :: Constructor
Http.__index = Http
Http.Enabled = HttpService.HttpEnabled
Http.ApiEndpoint = {
	CONNECT = "connect";
	UPLOAD = "upload";
}

type Api = {
	["CONNECT"]: "connect";
	["UPLOAD"]: "upload";
}

type Constructor = {
	__index: Constructor;
	Enabled: boolean;
	ApiEndpoint: Api;
	new: (address: string, compression: Enum.HttpCompression) -> Http;
	ValidAddress: (address: string) -> boolean;
	PostAsync: (self: Http, json_payload: string, request: string) -> ();
}
type Self = {
	address: string;
	compression: Enum.HttpCompression;
}
export type Http = setmetatable<Self, typeof(Http)>

local function MakeApiRequest(address: string, request: string): string
	return `{address}/api/v1/{request}`
end

local function PostBody(self: Http, game_data: boolean, json_payload: string, request: string): any
	return {
		Url = MakeApiRequest(self.address, request),
		Method = "POST",
		Headers = {
			["Content-Type"] = "application/json",
		},
		Body = json_payload,
		Compress = game_data and self.compression or Enum.HttpCompression.None
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

function Http:PostAsync(json_payload, request)
	local CreatePostBody = PostBody(self, false, json_payload, request)
	HttpService:RequestAsync(CreatePostBody)
end

return Http
