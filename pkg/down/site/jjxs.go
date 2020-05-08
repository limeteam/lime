package site

var jyxs = SiteA{
	Name:     "精英小说",
	HomePage: "http://www.jyyxs.com/",
	Match: []string{
		`https://www\.jyyxs\.com/\d+_\d+/*`,
		`https://www\.jyyxs\.com/\d+_\d+/\d+\.html/*`,
	},
}
