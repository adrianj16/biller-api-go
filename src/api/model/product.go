package model

type Product struct {
	ID           int              `db:"ID" json:"id"`
	Img          *[]Image         `json:"img,omitempty"`
	Title        string           `db:"Title" json:"title"`
	Price        *float64         `json:"price,omitempty"`
	PriceOld     *float64         `json:"price_old,omitempty"`
	Discount     *ProductDiscount `json:"discount,omitempty"`
	New          bool             `db:"New" json:"new"`
	InStock      *bool            `json:"in_stock,omitempty"`
	Multiproduct *bool            `db:"MultiProduct" json:"multiproduct"`
	Description  *string          `db:"Description" json:"description"`
	HasReview    *HasReview       `json:"has_review,omitempty"`
	Tags         *[]Tag           `json:"tags,omitempty"`
	Categories   *[]Category      `json:"categories,omitempty"`
	Specs        *Specs           `json:"specs,omitempty"`
	Details      *Details         `json:"details,omitempty"`
	Reviews      *Reviews         `json:"reviews,omitempty"`
	Related      *[]Related       `json:"related,omitempty"`
}

type Image struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
}

type ProductDiscount struct {
	Visible bool         `json:"visible"`
	Title   string       `json:"title"`
	Type    DiscountType `json:"type"`
	Value   float64      `json:"value"`
}

type DiscountType struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type HasReview struct {
	Visible bool   `json:"visible"`
	Count   int    `json:"count"`
	Link    string `json:"link"`
}

type Tag struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type Specs struct {
	Visible bool   `json:"visible"`
	Title   string `json:"title"`
	Data    []Spec `json:"data"`
}

type Spec struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int    `json:"value"`
}

type Details struct {
	Visible bool   `json:"visible"`
	Title   string `json:"title"`
	Data    string `json:"data"`
}

type Reviews struct {
	Visible        bool            `json:"visible"`
	Title          string          `json:"title"`
	AverageMark    int             `json:"average_mark"`
	Qualifications []Qualification `json:"qualifications"`
	Data           []Review        `json:"data"`
}

type Qualification struct {
	Stars int `json:"stars"`
	Count int `json:"count"`
}

type Review struct {
	UserID  int    `json:"user_id"`
	Alias   string `json:"alias"`
	Date    string `json:"date"`
	Rate    int    `json:"rate"`
	Comment string `json:"comment"`
}

type Related struct {
	ID       string          `json:"id"`
	Img      string          `json:"img"`
	Discount RelatedDiscount `json:"discount"`
	New      bool            `json:"new"`
	Category Category        `json:"category"`
	Title    string          `json:"title"`
	Price    float64         `json:"price"`
	OldPrice float64         `json:"old_price"`
}

type RelatedDiscount struct {
	Visible bool `json:"visible"`
	Percent int  `json:"percent"`
}
