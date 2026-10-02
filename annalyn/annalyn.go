package annalyn

var knightIsAwake = false
var archerIsAwake = true
var prisonerIsAwake = false
var petDogIsPresent = false

func CanFastAttack(knightIsAwake bool) bool {
	return !knightIsAwake 
}


func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
	return knightIsAwake || archerIsAwake || prisonerIsAwake
}



func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
	return !archerIsAwake && prisonerIsAwake
}


func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
	return (petDogIsPresent && !archerIsAwake) ||  (!knightIsAwake && !archerIsAwake && prisonerIsAwake)
}
