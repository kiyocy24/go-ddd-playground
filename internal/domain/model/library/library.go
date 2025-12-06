package library

import (
	"fmt"
	"slices"
	"time"
)

// BookID は書籍（タイトル単位）の識別子 (例: ISBN)
type BookID string

// BookTitle は書籍のタイトル
type BookTitle string

// Book は書籍の定義
// この構造体は「物理的な本」ではなく「作品」を表します。
type Book struct {
	id    BookID
	title BookTitle
}

func NewBook(id BookID, title BookTitle) Book {
	return Book{id: id, title: title}
}

func (b Book) ID() BookID       { return b.id }
func (b Book) Title() BookTitle { return b.title }

// BookItemID は図書館が所持する物理的な本の個体識別子
// 同じ ISBN の本が複数冊ある場合、それぞれ異なる ID を持ちます。
type BookItemID string

// BookItemState は本の現在の状態
type BookItemState int

const (
	BookItemStateInStock  BookItemState = iota + 1 // 在庫あり
	BookItemStateBorrowed                          // 貸出中
)

// BookItem は蔵書
// 図書館にある物理的な1冊の本
type BookItem struct {
	id     BookItemID
	bookID BookID
	state  BookItemState
}

func NewBookItem(id BookItemID, bookID BookID) *BookItem {
	return &BookItem{
		id:     id,
		bookID: bookID,
		state:  BookItemStateInStock, // 初期状態は在庫ありとする
	}
}

func (i *BookItem) ID() BookItemID       { return i.id }
func (i *BookItem) BookID() BookID       { return i.bookID }
func (i *BookItem) State() BookItemState { return i.state }
func (i *BookItem) ToBorrow()            { i.state = BookItemStateBorrowed }
func (i *BookItem) ToInStock()           { i.state = BookItemStateInStock }

// BorrowerID は利用者のID
type BorrowerID string

// Borrower は本を借りる利用者
type Borrower struct {
	id    BorrowerID
	items []*BookItem // 借りている本
}

func NewBorrower(id BorrowerID, items []*BookItem) *Borrower {
	return &Borrower{
		id:    id,
		items: items,
	}
}

// Stock は図書館全体の在庫管理を行う集約
type Stock struct {
	items []*BookItem
}

func NewStock(items []*BookItem) *Stock {
	return &Stock{
		items: items,
	}
}

// List は蔵書の一覧を返します
func (s *Stock) List() []*BookItem { return slices.Clone(s.items) }

// Add 蔵書を追加します
func (s *Stock) Add(item *BookItem) error {
	if _, exist := s.Find(item.id); exist {
		return fmt.Errorf("item %s already exists", item.id)
	}
	s.items = append(s.items, item)
	return nil
}

// Remove は蔵書を削除します
// もし存在しない蔵書を指定された場合何もしません
// 貸出中の蔵書を削除しようとした場合、エラーを返します
func (s *Stock) Remove(id BookItemID) error {
	item, exist := s.Find(id)
	if !exist {
		return nil
	}
	if item.state != BookItemStateInStock {
		return fmt.Errorf("item %s is already borrowed", id)
	}
	s.items = slices.DeleteFunc(s.items, func(i *BookItem) bool { return i.ID() == id })
	return nil
}

// Find は蔵書を ID で検索します
func (s *Stock) Find(id BookItemID) (*BookItem, bool) {
	for _, v := range s.items {
		if v.ID() == id {
			return v, true
		}
	}
	return nil, false
}

// LoanID は貸し出しのID
type LoanID string

// Loan は蔵書と利用者を結びつけます
type Loan struct {
	id         LoanID
	bookItemID BookItemID
	borrowerID BorrowerID
	loanedAt   time.Time
	dueDate    time.Time
	returnedAt *time.Time // 返却日はnil許容（未返却時）
}

func NewLoan(id LoanID, bookItemID BookItemID, borrowerID BorrowerID, loanedAt time.Time, dueDate time.Time) *Loan {
	return &Loan{
		id:         id,
		bookItemID: bookItemID,
		borrowerID: borrowerID,
		loanedAt:   loanedAt,
		dueDate:    dueDate,
	}
}

func (l *Loan) Return(t time.Time) { l.returnedAt = &t }
