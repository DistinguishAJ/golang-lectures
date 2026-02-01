package main

 type Loglevel int
 const (
	LevelTrace=0 loglevel = iota
	LevelDebug=1
	LevelInfo=2
	LevelWarning=3
	LevelError=4
 )

 var  levelName = [] string{"Trace", "Debug", "Info", "Warning", "Error"} 

 func (l loglevel) string() string{
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}

	return levelNames[l]
 }

 func printloglevel(level loglevel){
	fmt.Printf("Log level: %d %s\n", level, level.String())
 }


func main(){

	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarning)
	printLogLevel(LevelError)

}