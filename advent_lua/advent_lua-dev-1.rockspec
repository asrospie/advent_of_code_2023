package = "advent_lua"
version = "dev-1"
source = {
   url = "git+https://github.com/asrospie/advent_of_code_2023.git"
}
description = {
   homepage = "https://github.com/asrospie/advent_of_code_2023",
   license = "*** please specify a license ***"
}
dependencies = {
   "lua ~> 5.4"
}
build = {
   type = "builtin",
   modules = {
       main = "src/main.lua"
   }
}
