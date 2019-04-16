
local path = require "path"

local info = debug.getinfo(1,"S")
--获取当前路径
local pathinfo = info.short_src
--由于获取的路径为反斜杠(\)所以用上面的函数转为正斜杠(/)
local path = string.match(path.conversion(pathinfo),"^.*/")
--添加搜索路径
package.path = string.format("%s?.lua;%s?/init.lua;%s", path, path, package.path)
--加载模块(module为lua模块文件名"module.lua")