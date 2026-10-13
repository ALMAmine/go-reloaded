package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

var marqueur = regexp.MustCompile(`\((hex|bin|up|low|cap)(?:,\s*(\d+))?\)`)

func Process(ligne string) string {
	mots := decouper(ligne)
	mots = modifierMots(mots)
	corrigerArticles(mots)
	return assembler(mots)
}

func decouper(ligne string) []string {
	ligne = marqueur.ReplaceAllStringFunc(ligne, func(m string) string {
		return " " + strings.Join(strings.Fields(m), "") + " "
	})

	var mots []string
	for _, morceau := range strings.Fields(ligne) {
		if marqueur.MatchString(morceau) {
			mots = append(mots, morceau)
			continue
		}

		mot := ""
		lettres := []rune(morceau)
		for i, lettre := range lettres {
			apostrophe := false
			if lettre == '\'' && i > 0 && i < len(lettres)-1 {
				apostrophe = estMot(string(lettres[i-1])) && estMot(string(lettres[i+1]))
			}
			if strings.ContainsRune(".,!?:;", lettre) || (lettre == '\'' && !apostrophe) {
				if mot != "" {
					mots = append(mots, mot)
					mot = ""
				}
				mots = append(mots, string(lettre))
			} else {
				mot += string(lettre)
			}
		}
		if mot != "" {
			mots = append(mots, mot)
		}
	}
	return mots
}

func modifierMots(mots []string) []string {
	var resultat []string
	for _, mot := range mots {
		commande := marqueur.FindStringSubmatch(mot)
		if commande == nil {
			resultat = append(resultat, mot)
			continue
		}

		nombre := 1
		if commande[2] != "" {
			var err error
			nombre, err = strconv.Atoi(commande[2])
			if err != nil {
				nombre = len(resultat)
			}
		}

		for i := len(resultat) - 1; i >= 0 && nombre > 0; i-- {
			if estMot(resultat[i]) {
				resultat[i] = transformer(resultat[i], commande[1])
				nombre--
			}
		}
	}
	return resultat
}

func estMot(mot string) bool {
	for _, lettre := range mot {
		if unicode.IsLetter(lettre) || unicode.IsDigit(lettre) {
			return true
		}
	}
	return false
}

func transformer(mot string, commande string) string {
	switch commande {
	case "up":
		return strings.ToUpper(mot)
	case "low":
		return strings.ToLower(mot)
	case "cap":
		lettres := []rune(strings.ToLower(mot))
		for i, lettre := range lettres {
			if unicode.IsLetter(lettre) {
				lettres[i] = unicode.ToUpper(lettre)
				break
			}
		}
		return string(lettres)
	case "hex", "bin":
		base := 16
		if commande == "bin" {
			base = 2
		}
		nombre, err := strconv.ParseInt(mot, base, 64)
		if err == nil {
			return strconv.FormatInt(nombre, 10)
		}
	}
	return mot
}

func corrigerArticles(mots []string) {
	for i, mot := range mots {
		if mot != "a" && mot != "A" {
			continue
		}
		suivant := i + 1
		if suivant < len(mots) && mots[suivant] == "'" {
			suivant++
		}
		if suivant < len(mots) {
			lettres := []rune(strings.ToLower(mots[suivant]))
			if strings.ContainsRune("aeiouh", lettres[0]) {
				mots[i] = mot + "n"
			}
		}
	}
}

func assembler(mots []string) string {
	resultat := ""
	guillemetOuvert := false
	collerSuivant := false

	for _, mot := range mots {
		if mot == "'" && guillemetOuvert {
			resultat += "'"
			guillemetOuvert = false
			collerSuivant = false
		} else if mot == "'" {
			if resultat != "" && !collerSuivant {
				resultat += " "
			}
			resultat += "'"
			guillemetOuvert = true
			collerSuivant = true
		} else if !estMot(mot) && strings.ContainsAny(mot, ".,!?:;") {
			resultat += mot
			collerSuivant = false
		} else {
			if resultat != "" && !collerSuivant {
				resultat += " "
			}
			resultat += mot
			collerSuivant = false
		}
	}
	return resultat
}
