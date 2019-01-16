package wifi

import (
	"sort"
	"strings"
)

// GetAPList gets a list of access points from the output string
func GetAPList(output string) ([]string, error) {
	retval := []string{}

	//	First, split on newline:
	outputlines := strings.Split(output, "\n")

	//	Then, for each string, look for 'SSID' and add it
	for _, value := range outputlines {
		//	Look for SSID
		location := strings.Index(value, "SSID:")

		if location > -1 {
			//	See if it's an SSID
			ap := value[location+len("SSID:"):]

			//	Trim the string:
			ap = strings.TrimSpace(ap)

			//	If we have something ...
			if ap != "" {

				//	... and we haven't already added the item...
				if !apListContains(retval, ap) {
					//	Add the AP to the list
					retval = append(retval, ap)
				}
			}
		}
	}

	//	Sort the list:
	sort.Strings(retval)

	//	Return the list:
	return retval, nil
}

func apListContains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
