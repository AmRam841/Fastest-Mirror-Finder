package main

import (
	file_reader "Fastest_Mirror_Finder/internal"
)

// new extions better comments guide :
//** this is important : im going to use the mfor things to remmember
// ! this is for alerts !
// TODO
// ?

//import "Fastest_Mirror_Finder/cli"
func main(){
// TODO CLI Goes here ! 
//cli.Arg_reader()
// cli.Tables()
// cli.Lists()


// TODO make the logic first : 
// TODO logic :  
// TODO first read from a file . we will do this for arch ! how can we read from a file ? and not only a simple file , how can we read the arch mirror files ? 
// this is the url for arch https://archlinux.org/mirrorlist/all/
// * OPTION : ask the user if he/she wants to cache the list
path := "/home/archedlinux/mirrorlists/arch.txt"
file_reader.ReadLinesFromFile(path)










}