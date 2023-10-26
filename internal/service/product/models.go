package product

var allProduct = []Product{
	{Title: "Курьер Яндекс.ЕДА"},
	{Title: "Курьер Яндекс.Лавка"},
	{Title: "Курьер Яндекс.Доставка"},
}

type Product struct {
	Title string
}
