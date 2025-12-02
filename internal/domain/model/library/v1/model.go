package v1

// BookID は Book の識別子
type BookID string

// BookTitle は Book のタイトル
type BookTitle string

// Book は本
type Book struct {
	id    BookID
	title BookTitle
}

// Stock は図書館の在庫
type Stock struct {
	books []*StockedBook
}

// Take は在庫から本を取り出します
func (s *Stock) Take(id BookID) (*Book, bool) { return nil, false }

// Add は本を追加します
func (s *Stock) Add(book *Book) {}

// StockedBookID は Stock に保存されている Book を識別するための ID
type StockedBookID string

// StockedBook は Stock に保存されている Book
type StockedBook struct {
	id   StockedBookID
	book *Book
}

// BorrowedBookID は 借りている本を識別するための ID
type BorrowedBookID string

// BorrowedBook は書庫から一時的に別の場所で管理されている本
type BorrowedBook struct {
	id   BorrowedBookID
	book *Book
}

// BorrowerID は本を借りる人の識別するための ID
type BorrowerID string

// Borrower は本を借りる人
type Borrower struct {
	id    BorrowerID
	books []*BorrowedBook
}

// Borrow は本を借ります
func (b *Borrower) Borrow(book *BorrowedBook) {}

// Return は借りている本を返します
func (b *Borrower) Return(book *BorrowedBook) {}
