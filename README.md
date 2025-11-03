# DomQuery - A jQuery like DOM search & manipulation tool

A **jQuery-like DOM search & manipulation tool** for [Gopher-Lua](https://github.com/yuin/gopher-lua), built on top of [goquery](https://github.com/PuerkitoBio/goquery).

DomQuery allows you to traverse and manipulate HTML documents in Lua scripts with a familiar, chainable API inspired by jQuery.

## 🚀 Features

- jQuery-like selectors (`#id`, `.class`, `tag > child`, `:nth-child()`, etc.)
- DOM traversal & manipulation (`find`, `first`, `text`, `attrOr`, etc.)
- Easy integration with [Gopher-Lua](https://github.com/yuin/gopher-lua)
- Built on the robust [goquery](https://github.com/PuerkitoBio/goquery) library
- Simple and expressive API for scraping or document parsing

## Example

```lua
local domquery = require("domquery")
local doc, err = domquery.fromString(DOMString)
local result = result()
doc:find("#screen > div:nth-child(4) > section > ol > li"):
		eachBreak(function(i, e)
			local url = (trim(e:find("header a"):
				first():attrOr("href", "")))
			local title = (trim(e:find("h2"):
				first():text()))
			local text = (trim(e:find("p"):
				first():text()))
			if empty(url, text, title) then
				return false
			end
			result:append({
				url = url;
				title = title;
				text = text;
			})
			return true
		end)
```

