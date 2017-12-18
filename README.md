# bravebin

bravebin is like ansible, but with Go.  bravebin generates and compiles Go code into binaries to distribute and execute on hosts from yaml recipes.


## Project Goals

Open Source, have not chosen a license to release code under yet but will choose a common one as project develops.

Compatible with existing ansible recipes by using same yaml naming for bravebin modules

Compatible with Linux, OSX, and Windows


## Tasks

+ Go code generator built with Go
+ add debug module to print out a line
+ YAML parsing
+ SSH setup
+ Compiling cross-platform binaries based on the target os, need to add a gather facts here?
+ Start building out more useful modules
