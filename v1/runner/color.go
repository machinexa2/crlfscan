package runner

type ColoredObject struct {
	Information	string
        Good		string
        Bad		string
	Vuln		string
}

var Color ColoredObject = ColoredObject{}

func init(){
	color := &Color
	color.Information = "[\x1b[33mINF\x1b[0m]"
        color.Good = "[\x1b[32mYES\x1b[0m]"
        color.Vuln = "[\x1b[32mVLN\x1b[0m]"
        color.Bad = "[\x1b[31mERR\x1b[0m]"
}

func (color ColoredObject) Green(text string) string{
	return "\x1b[32m" + text + "\x1b[0m"
}

func (color ColoredObject) Blue(text string) string{
	return "\x1b[34m" + text + "\x1b[0m"
}

func (color ColoredObject) Red(text string) string{
	return "\x1b[31m" + text + "\x1b[0m"
}

func (color ColoredObject) Yellow(text string) string{
	return "\x1b[33m" + text + "\x1b[0m"
}
