package models

type FaindaIndexes struct {
	PrimaryKey           string
	Uid                  string
	Filterableattributes *[]string
	Searchableattributes *[]string
	Sortableattributes   *[]string
	Rankingrules         *[]string
}
