package passports

import (
	"fmt"
	"regexp"
)

type Passport struct {
	BirthYear      int
	IssueYear      int
	ExpirationYear int
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
	CountryID      string
}

func (p Passport) IsValid() (bool, error) {
	// validating BirthYear field
	if p.BirthYear < 1920 || 2002 < p.BirthYear {
		return false, fmt.Errorf("invalid birth year found: %d", p.BirthYear)
	}

	// validating IssueYear field
	if p.IssueYear < 2010 || 2020 < p.IssueYear {
		return false, fmt.Errorf("invalid issue year found: %d", p.IssueYear)
	}

	// validating ExpirationYear field
	if p.ExpirationYear < 2020 || 2030 < p.ExpirationYear {
		return false, fmt.Errorf("invalid expiration year found: %d", p.ExpirationYear)
	}

	// validating Height field
	var height int
	var heightUnit string
	n, err := fmt.Sscanf(p.Height, "%d%s", &height, &heightUnit)
	if n != 2 || err != nil {
		return false, fmt.Errorf("failed to parse height: n=%d; err=%v", n, err)
	}

	switch heightUnit {
	case "cm":
		if height < 150 || 193 < height {
			return false, fmt.Errorf("invalid height found: %s", p.Height)
		}
	case "in":
		if height < 59 || 76 < height {
			return false, fmt.Errorf("invalid height found: %s", p.Height)
		}
	default:
		return false, fmt.Errorf("unknown height unit found: %s", p.Height)
	}

	// validating HairColor
	re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
	s := re.FindAllString(p.HairColor, -1)
	if len(s) != 1 {
		return false, fmt.Errorf("invalid hair color found: %s", p.HairColor)
	}

	// validating EyeColor
	switch p.EyeColor {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
	default:
		return false, fmt.Errorf("invalid eye color found: %s", p.EyeColor)
	}

	// validating PassportID
	re = regexp.MustCompile(`^[0-9]{9}$`)
	s = re.FindAllString(p.PassportID, -1)
	if len(s) != 1 {
		return false, fmt.Errorf("invalid passport ID found: %s", p.PassportID)
	}

	return true, nil
}
