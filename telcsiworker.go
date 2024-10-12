package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gocolly/colly/v2"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

type Phone struct {
	title string
	price int
	city  string
	link  string
	ram   string
}

type PhoneBrand struct {
	name        string
	phones      []Phone
	excel_color string
}

type PhoneRegExp struct {
	name     string
	p_regexp *regexp.Regexp
}

type PhoneRegExpCatalog struct {
	regexps []PhoneRegExp
}

// type PhoneCatalog struct {
// 	flip    []Phone
// 	oneplus []Phone
// 	sony    []Phone
// 	nothing []Phone
// 	xiaomi  []Phone
// 	huawei  []Phone
// 	pixel   []Phone
// 	honor   []Phone
// 	other   []Phone
// 	samsung []Phone
// 	apple   []Phone
// }

type PhoneCatalog struct {
	phonebrands []PhoneBrand
}

func telcsiworker(telcsiFarmolasDef *TelcsiFarmolas) (excelName string) {

	c_jofog := colly.NewCollector()

	c_hardapro := colly.NewCollector()

	// const MAX_PRICE = 100000
	// const MIN_PRICE = 20000

	const DEPTH = 50
	currentDepth := 0

	const FOLDER = "scrapings/"

	var foundPhones PhoneCatalog
	var p_regexps PhoneRegExpCatalog

	// var phoneBrands []string

	// INIT PHONEBRANDS
	foundPhones = decidePhoneCatalog(telcsiFarmolasDef)

	// INIT PHONEBRANDS
	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "flip",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "GOOD",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "fold",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "GOOD",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "samsung",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "GOOD",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "sony",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "GOOD",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "oneplus",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "pixel",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "honor",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "nothing",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "huawei",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	// foundPhones.phonebrands = append(foundPhones.phonebrands,
	// 	PhoneBrand{
	// 		name:        "xiaomi",
	// 		phones:      make([]Phone, 0),
	// 		excel_color: "NEUTRAL",
	// 	})

	foundPhones.phonebrands = append(foundPhones.phonebrands,
		PhoneBrand{
			name:        "other",
			phones:      make([]Phone, 0),
			excel_color: "NEUTRAL",
		})

	foundPhones.phonebrands = append(foundPhones.phonebrands,
		PhoneBrand{
			name:        "apple",
			phones:      make([]Phone, 0),
			excel_color: "BAD",
		})

	// INIT REGEXPS
	p_regexps = decideRegexp(foundPhones)

	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "flip",
	// 	p_regexp: regexp.MustCompile("FLIP"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "fold",
	// 	p_regexp: regexp.MustCompile("FOLD"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "samsung",
	// 	p_regexp: regexp.MustCompile("SAMSUNG|GALAXY"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "apple",
	// 	p_regexp: regexp.MustCompile("APPLE|IPHONE"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "huawei",
	// 	p_regexp: regexp.MustCompile("HUAWEI|HAUWEI"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "xiaomi",
	// 	p_regexp: regexp.MustCompile("XIAOMI|XAOMI"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "sony",
	// 	p_regexp: regexp.MustCompile("SONY"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "nothing",
	// 	p_regexp: regexp.MustCompile("NOTHING"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "oneplus",
	// 	p_regexp: regexp.MustCompile("ONEPLUS|ONE PLUS"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "pixel",
	// 	p_regexp: regexp.MustCompile("PIXEL"),
	// })
	// p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
	// 	name:     "honor",
	// 	p_regexp: regexp.MustCompile("HONOR"),
	// })

	/*
		----------------------------------------------------------
		----------------------------------------------------------
		||                       JÓFOGÁS                        ||
		----------------------------------------------------------
		----------------------------------------------------------
	*/
	c_jofog.OnHTML(".contentArea", func(e *colly.HTMLElement) {
		currentTitle := e.ChildText("section.subjectWrapper h3.item-title a.subject")
		currentLink := e.ChildAttr("section.subjectWrapper h3.item-title a.subject", "href")
		currentPrice := strings.Replace(e.ChildText("section.price div.priceBox h3.item-price span.price-value"), " ", "", -1)
		currentCity := strings.Replace(e.ChildText("section.cityname"), "  ,", ",", 1)
		currentRam := getRamClassification(currentTitle)

		//currentRam = strconv.FormatBool(re512GB.MatchString(currentTitle)) + strconv.FormatBool(re256GB.MatchString(currentTitle)) + strconv.FormatBool(re128GB.MatchString(currentTitle)) + strconv.FormatBool(re512GBprob.MatchString(currentTitle)) + strconv.FormatBool(re256GBprob.MatchString(currentTitle)) + strconv.FormatBool(re128GBprob.MatchString(currentTitle))
		//"a[href].subject"

		// fmt.Printf("%v : %s \n", reSamsung.MatchString(currentTitle), currentTitle)
		currentPriceInt, err := strconv.Atoi(currentPrice)
		if err != nil {
			currentPriceInt = -1
		}

		var foundbrand = false

		for _, v := range p_regexps.regexps {
			if v.p_regexp.MatchString(strings.ToUpper(currentTitle)) {
				phone := slices.IndexFunc(foundPhones.phonebrands, func(p PhoneBrand) bool { return strings.ToUpper(p.name) == strings.ToUpper(v.name) })
				foundPhones.phonebrands[phone].phones = append(foundPhones.phonebrands[phone].phones,
					Phone{
						title: currentTitle,
						price: currentPriceInt,
						city:  currentCity,
						link:  currentLink,
						ram:   currentRam,
					})
				foundbrand = true
				break
			}
		}
		if foundbrand == false {
			phone := slices.IndexFunc(foundPhones.phonebrands, func(p PhoneBrand) bool { return strings.ToUpper(p.name) == strings.ToUpper("other") })
			foundPhones.phonebrands[phone].phones = append(foundPhones.phonebrands[phone].phones,
				Phone{
					title: currentTitle,
					price: currentPriceInt,
					city:  currentCity,
					link:  currentLink,
					ram:   currentRam,
				})
		}

	})

	c_jofog.OnHTML(".ad-list-pager-item-next", func(n *colly.HTMLElement) {
		nextLink := n.Attr("href")
		phonesCount := 0
		for _, v := range foundPhones.phonebrands {
			phonesCount += len(v.phones)
		}
		//fmt.Printf("Jófogás Visiting:  %s\n Length Phones: %d\n", nextLink, phonesCount)
		currentDepth++
		if currentDepth >= DEPTH || nextLink == "" {
			//letsExcelize(foundPhones)
			currentDepth = 0
			return
		}
		c_jofog.Visit(nextLink)
	})

	/*
		----------------------------------------------------------
		----------------------------------------------------------
		||                     HARDVERAPRÓ                      ||
		----------------------------------------------------------
		----------------------------------------------------------
	*/
	c_hardapro.OnHTML(".media-body", func(e *colly.HTMLElement) {
		currentTitle := e.ChildText("div.uad-title h1 a")
		currentLink := e.ChildAttr("div.uad-title h1 a", "href")
		currentPrice := strings.Replace(strings.Replace(e.ChildText("div.uad-info div.uad-price"), " ", "", -1), "Ft", "", -1)
		currentCity := e.ChildText("div.uad-info div.uad-light")
		currentRam := getRamClassification(currentTitle)

		//currentRam = strconv.FormatBool(re512GB.MatchString(currentTitle)) + strconv.FormatBool(re256GB.MatchString(currentTitle)) + strconv.FormatBool(re128GB.MatchString(currentTitle)) + strconv.FormatBool(re512GBprob.MatchString(currentTitle)) + strconv.FormatBool(re256GBprob.MatchString(currentTitle)) + strconv.FormatBool(re128GBprob.MatchString(currentTitle))
		//"a[href].subject"

		//fmt.Println(currentTitle)
		// fmt.Printf("%v : %s \n", reSamsung.MatchString(currentTitle), currentTitle)
		currentPriceInt, err := strconv.Atoi(currentPrice)
		if err != nil {
			currentPriceInt = -1
		}

		var foundbrand = false

		for _, v := range p_regexps.regexps {
			if v.p_regexp.MatchString(strings.ToUpper(currentTitle)) {
				phone := slices.IndexFunc(foundPhones.phonebrands, func(p PhoneBrand) bool { return strings.ToUpper(p.name) == strings.ToUpper(v.name) })
				foundPhones.phonebrands[phone].phones = append(foundPhones.phonebrands[phone].phones,
					Phone{
						title: currentTitle,
						price: currentPriceInt,
						city:  currentCity,
						link:  currentLink,
						ram:   currentRam,
					})
				foundbrand = true
				break
			}
		}
		if foundbrand == false {
			phone := slices.IndexFunc(foundPhones.phonebrands, func(p PhoneBrand) bool { return strings.ToUpper(p.name) == strings.ToUpper("other") })
			foundPhones.phonebrands[phone].phones = append(foundPhones.phonebrands[phone].phones,
				Phone{
					title: currentTitle,
					price: currentPriceInt,
					city:  currentCity,
					link:  currentLink,
					ram:   currentRam,
				})
		}

	})

	c_hardapro.OnHTML("#forum-nav-top ~ ul.mr-md-auto > li.nav-arrow > a", func(n *colly.HTMLElement) {
		if n.Attr("rel") == "next" {
			nextLink := "https://hardverapro.hu" + n.Attr("href")
			phonesCount := 0
			for _, v := range foundPhones.phonebrands {
				phonesCount += len(v.phones)
			}
			//fmt.Printf("Hardverapró Visiting:  %s\n Length Phones: %d\n", nextLink, phonesCount)
			currentDepth++
			if currentDepth >= DEPTH || nextLink == "" {
				//letsExcelize(foundPhones)
				return
			}
			c_hardapro.Visit(nextLink)
		}
	})

	jofog_link := ""
	hardapro_link := ""
	if telcsiFarmolasDef.IsPhone == "on" {
		jofog_link = fmt.Sprintf("https://www.jofogas.hu/magyarorszag/mobiltelefon?max_price=%d&min_price=%d&mobile_memory=3,4,5,6,7,8&mobile_os=1&sp=2", telcsiFarmolasDef.MaxPrice, telcsiFarmolasDef.MinPrice)
		hardapro_link = fmt.Sprintf("https://hardverapro.hu/aprok/mobil/mobil/android/keres.php?stext=&stcid_text=&stcid=&stmid_text=&stmid=&minprice=%d&maxprice=%d&cmpid_text=&cmpid=&usrid_text=&usrid=&__buying=0&__buying=1&stext_none=", telcsiFarmolasDef.MinPrice, telcsiFarmolasDef.MaxPrice)
	} else {
		jofog_link = fmt.Sprintf("https://www.jofogas.hu/magyarorszag/tablet-e-book-olvaso-es-kiegeszitok?max_price=%d&min_price=%d&tablet_os=1&tablet_size=5,6,7&tablet_type=1", telcsiFarmolasDef.MaxPrice, telcsiFarmolasDef.MinPrice)
		hardapro_link = fmt.Sprintf("https://hardverapro.hu/aprok/mobil/tablet/android_tablet/keres.php?stext=&stcid_text=&stcid=&stmid_text=&stmid=&minprice=%d&maxprice=%d&cmpid_text=&cmpid=&usrid_text=&usrid=&__buying=0&__buying=1&stext_none=", telcsiFarmolasDef.MinPrice, telcsiFarmolasDef.MaxPrice)

	}

	c_jofog.Visit(jofog_link)
	c_jofog.Wait()
	fmt.Println("Done with Jófogás")
	c_hardapro.Visit(hardapro_link)
	c_hardapro.Wait()
	fmt.Println("Done with Hardverapró")
	//	fmt.Println("Done scraping\nStarting Sorting")

	sortPhones(foundPhones)
	//	fmt.Println("Done sorting\nStarting Excel")
	timeZoneString := time.Now().Format("20060102150405")
	excelName = FOLDER + "phones_" + timeZoneString + ".xlsx"
	excelNameReturn := "phones_" + timeZoneString + ".xlsx"
	letsExcelize(foundPhones, excelName)
	//exportToCsv(foundPhones, FOLDER+"phones_raw_"+timeZoneString+".csv")

	fmt.Println("Done with Excel")
	return excelNameReturn
}

func letsExcelize(phones PhoneCatalog, filename string) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, err := f.NewSheet("Phones")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	/*
		// LOOPING \\
	*/

	for i, p := range phones.phonebrands {
		//fmt.Println("Starting: " + p.name + "\n")
		title_cell1, err := excelize.CoordinatesToCellName(1+(i*5), 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		title_cell4, err := excelize.CoordinatesToCellName(4+(i*5), 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		subtitle_cell1, err := excelize.CoordinatesToCellName(1+(i*5), 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		subtitle_cell2, err := excelize.CoordinatesToCellName(2+(i*5), 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		subtitle_cell3, err := excelize.CoordinatesToCellName(3+(i*5), 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		subtitle_cell4, err := excelize.CoordinatesToCellName(4+(i*5), 2)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetCellValue("Phones", title_cell1, strings.ToUpper(p.name))

		f.MergeCell("Phones", title_cell1, title_cell4)
		styleColorofCol(f, "Phones", title_cell1[:len(title_cell1)-1]+":"+title_cell4[:len(title_cell4)-1], p.excel_color)
		styleTitle(f, "Phones", 1+(i*5), 1, 4+(i*5), 1, p.excel_color)
		f.SetCellValue("Phones", subtitle_cell1, "Cím")
		styleTitle(f, "Phones", 1+(i*5), 2, 1+(i*5), 2, p.excel_color)
		f.SetCellValue("Phones", subtitle_cell2, "Ár")
		styleTitle(f, "Phones", 2+(i*5), 2, 2+(i*5), 2, p.excel_color)
		f.SetCellValue("Phones", subtitle_cell3, "Város")
		styleTitle(f, "Phones", 3+(i*5), 2, 3+(i*5), 2, p.excel_color)
		f.SetCellValue("Phones", subtitle_cell4, "Link")
		styleTitle(f, "Phones", 4+(i*5), 2, 4+(i*5), 2, p.excel_color)
		for j, v := range p.phones {
			//fmt.Println("Writing: " + v.title + "\n")
			cell, err := excelize.CoordinatesToCellName(1+(i*5), j+3)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetCellValue("Phones", cell, v.title)

			cell, err = excelize.CoordinatesToCellName(2+(i*5), j+3)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetCellValue("Phones", cell, v.price)

			cell, err = excelize.CoordinatesToCellName(3+(i*5), j+3)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetCellValue("Phones", cell, v.city)

			cell, err = excelize.CoordinatesToCellName(4+(i*5), j+3)
			if err != nil {
				fmt.Println(err)
				return
			}
			f.SetCellValue("Phones", cell, v.link)
			f.SetCellHyperLink("Phones", cell, v.link, "External")
		}
	}

	autoFitCols(f, "Phones")

	// Save spreadsheet by the given path.
	if err := f.SaveAs(filename); err != nil {
		fmt.Println(err)
	}
}

func autoFitCols(f *excelize.File, sheetName string) {
	// Autofit all columns according to their text content
	cols, err := f.GetCols(sheetName)
	if err != nil {
		fmt.Println(err)
		return
	}
	for idx, col := range cols {
		largestWidth := 0
		for _, rowCell := range col {
			cellWidth := utf8.RuneCountInString(rowCell) + 2 // + 2 for margin
			if cellWidth > largestWidth {
				largestWidth = cellWidth
			}
		}
		name, err := excelize.ColumnNumberToName(idx + 1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetColWidth(sheetName, name, name, float64(largestWidth))
	}
}

func styleTitle(f *excelize.File, sheetName string, cellPositionx int, cellPositiony int, cellPositionx2 int, cellPositiony2 int, colorgoodorbad string) {
	fontsize := 11.0
	if cellPositiony+cellPositiony2 == 2 {
		fontsize = 14.0
	}

	cell, err := excelize.CoordinatesToCellName(cellPositionx, cellPositiony)
	if err != nil {
		fmt.Println(err)
		return
	}
	cell2, err := excelize.CoordinatesToCellName(cellPositionx2, cellPositiony2)
	if err != nil {
		fmt.Println(err)
		return
	}

	var style int
	var styleerr error
	if colorgoodorbad == "GOOD" { // GOOD
		style, styleerr = f.NewStyle(&excelize.Style{
			Fill:      excelize.Fill{Type: "pattern", Color: []string{"C6EFCE"}, Pattern: 1},
			Font:      &excelize.Font{Bold: true, Size: fontsize},
			Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
			Border: []excelize.Border{
				{Type: "top", Style: 2, Color: "000000"},
				{Type: "bottom", Style: 2, Color: "000000"},
				{Type: "left", Style: 2, Color: "000000"},
				{Type: "right", Style: 2, Color: "000000"},
			},
		})
	} else if colorgoodorbad == "BAD" { // BAD
		style, styleerr = f.NewStyle(&excelize.Style{
			Fill:      excelize.Fill{Type: "pattern", Color: []string{"FFC7CE"}, Pattern: 1},
			Font:      &excelize.Font{Bold: true, Size: fontsize},
			Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
			Border: []excelize.Border{
				{Type: "top", Style: 2, Color: "000000"},
				{Type: "bottom", Style: 2, Color: "000000"},
				{Type: "left", Style: 2, Color: "000000"},
				{Type: "right", Style: 2, Color: "000000"},
			},
		})
	} else { // NEUTRAL
		style, styleerr = f.NewStyle(&excelize.Style{
			Font:      &excelize.Font{Bold: true, Size: fontsize},
			Alignment: &excelize.Alignment{Vertical: "center", Horizontal: "center"},
			Border: []excelize.Border{
				{Type: "top", Style: 2, Color: "000000"},
				{Type: "bottom", Style: 2, Color: "000000"},
				{Type: "left", Style: 2, Color: "000000"},
				{Type: "right", Style: 2, Color: "000000"},
			},
		})
	}
	if styleerr != nil {
		fmt.Println(styleerr)
		return
	}
	f.SetCellStyle(sheetName, cell, cell2, style)
}

func styleColorofCol(f *excelize.File, sheetName string, cols string, color string) {

	if color == "GOOD" {
		style, err := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{Type: "pattern", Color: []string{"C6EFCE"}, Pattern: 1},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetColStyle(sheetName, cols, style)

	} else if color == "BAD" {
		style, err := f.NewStyle(&excelize.Style{
			Fill: excelize.Fill{Type: "pattern", Color: []string{"FFC7CE"}, Pattern: 1},
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetColStyle(sheetName, cols, style)
	}

}

func sortPhones(phones PhoneCatalog) {

	for _, v := range phones.phonebrands {
		sort.Slice(v.phones, func(i, j int) bool {
			return v.phones[i].price > v.phones[j].price //&& v.phones[i].ram > v.phones[j].ram
		})
	}

}

func getRamClassification(title string) (ram string) {
	ram = ""

	re512GB := regexp.MustCompile("512GB|512 GB|512")
	re512GBprob := regexp.MustCompile("512")
	re256GB := regexp.MustCompile("256GB|256 GB|256")
	re256GBprob := regexp.MustCompile("256")
	re128GB := regexp.MustCompile("128GB|128 GB|128")
	re128GBprob := regexp.MustCompile("128")

	if re512GB.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	if re256GB.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	if re128GB.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	if re512GBprob.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	if re256GBprob.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	if re128GBprob.MatchString(title) {
		ram += "1"
	} else {
		ram += "0"
	}
	return
}

func decidePhoneCatalog(telcsi *TelcsiFarmolas) PhoneCatalog {
	var foundPhones PhoneCatalog

	//var phoneBrands []string

	// Important
	for _, p := range telcsi.ImportantPhones {
		foundPhones.phonebrands = append(foundPhones.phonebrands,
			PhoneBrand{
				name:        p,
				phones:      make([]Phone, 0),
				excel_color: "GOOD",
			})
	}

	// Neutral
	for _, p := range telcsi.NeutralPhones {
		foundPhones.phonebrands = append(foundPhones.phonebrands,
			PhoneBrand{
				name:        p,
				phones:      make([]Phone, 0),
				excel_color: "NEUTRAL",
			})
	}

	return foundPhones
}

func decideRegexp(fp PhoneCatalog) PhoneRegExpCatalog {
	var p_regexps PhoneRegExpCatalog

	p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
		name:     "apple",
		p_regexp: regexp.MustCompile("APPLE|IPHONE"),
	})

	if checkPhoneCatalogForBrand(fp, "flip") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "flip",
			p_regexp: regexp.MustCompile("FLIP"),
		})
	}

	if checkPhoneCatalogForBrand(fp, "fold") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "fold",
			p_regexp: regexp.MustCompile("FOLD"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "samsung") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "samsung",
			p_regexp: regexp.MustCompile("SAMSUNG|GALAXY|SAMSUBG"),
		})
	}

	if checkPhoneCatalogForBrand(fp, "huawei") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "huawei",
			p_regexp: regexp.MustCompile("HUAWEI|HAUWEI|HUANWEI"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "xiaomi") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "xiaomi",
			p_regexp: regexp.MustCompile("XIAOMI|XAOMI"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "sony") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "sony",
			p_regexp: regexp.MustCompile("SONY"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "nothing") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "nothing",
			p_regexp: regexp.MustCompile("NOTHING"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "oneplus") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "oneplus",
			p_regexp: regexp.MustCompile("ONEPLUS|ONE PLUS"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "pixel") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "pixel",
			p_regexp: regexp.MustCompile("PIXEL"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "honor") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "honor",
			p_regexp: regexp.MustCompile("HONOR"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "oppo") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "oppo",
			p_regexp: regexp.MustCompile("OPPO"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "redmi") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "redmi",
			p_regexp: regexp.MustCompile("REDMI"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "realme") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "realme",
			p_regexp: regexp.MustCompile("REALME"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "poco") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "poco",
			p_regexp: regexp.MustCompile("POCO"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "nokia") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "nokia",
			p_regexp: regexp.MustCompile("NOKIA"),
		})
	}
	if checkPhoneCatalogForBrand(fp, "motorola") {
		p_regexps.regexps = append(p_regexps.regexps, PhoneRegExp{
			name:     "motorola",
			p_regexp: regexp.MustCompile("MOTOROLA"),
		})
	}

	return p_regexps
}

func checkPhoneCatalogForBrand(fp PhoneCatalog, brand string) bool {
	found := false
	for _, pb := range fp.phonebrands {
		if pb.name == brand {
			found = true
			return found
		}
	}
	return found
}

/*

#
#	STOLENED !!!!!!
#

*/

func exportToCsv(phones PhoneCatalog, filename string) {
	records := [][]string{
		{"Marka", "Cim", "Ar", "Varos", "Link"},
		//		{"Rob", "Pike", "rob"},
	}

	for _, pb := range phones.phonebrands {

		for _, p := range pb.phones {
			records = append(records, []string{
				strings.ToUpper(pb.name),
				p.title,
				strconv.Itoa(p.price),
				p.city,
				p.link,
			},
			)
		}
	}

	csvFile, err := os.Create(filename)
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	t := transform.NewWriter(csvFile, unicode.UTF16(unicode.LittleEndian, unicode.UseBOM).NewEncoder())

	w := csv.NewWriter(t)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	t.Close()
	// Output:
	// first_name,last_name,username
	// Rob,Pike,rob
	// Ken,Thompson,ken
	// Robert,Griesemer,gri
}
