package intro

import "strconv"

func GiveIntro(name string, class int) string {
	if name == "" {
		return "Name is required."
	}

	if class == 0 {
		return "Class can't be zero"
	}

	return "This is " + name + ", I study in " + strconv.Itoa(class)  + " class." 
}
