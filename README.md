# DomQuery - A jQuery like DOM search & manipulation tool

A jQuery like tool for [gopher-lua](https://github.com/yuin/gopher-lua).
It's based on [goquery](https://github.com/PuerkitoBio/goquery).

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

