package library_test

import (
	"math/rand"
	"testing"

	"github.com/kiyocy24/go-ddd-playground/internal/domain/model/library"
	"github.com/kiyocy24/go-ddd-playground/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBook(t *testing.T) {
	bookID := newBookID()
	bookTitle := newBookTitle()
	book := library.NewBook(bookID, bookTitle)
	assert.Equal(t, bookID, book.ID())
	assert.Equal(t, bookTitle, book.Title())
}

func TestBookItem(t *testing.T) {
	itemID := newBookItemID()
	bookID := newBookID()
	item := library.NewBookItem(itemID, bookID)
	assert.Equal(t, itemID, item.ID())
	assert.Equal(t, bookID, item.BookID())
	assert.Equal(t, library.BookItemStateInStock, item.State())
	item.ToBorrow()
	assert.Equal(t, library.BookItemStateBorrowed, item.State())
	item.ToInStock()
	assert.Equal(t, library.BookItemStateInStock, item.State())
}

func TestStock(t *testing.T) {
	stocked := []*library.BookItem{
		newBookItem(),
		newBookItem(),
		newBookItem(),
	}
	stock := library.NewStock(stocked)

	t.Run("List", func(t *testing.T) {
		assert.Equal(t, stocked, stock.List(), "コンストラクトで渡した内容がリストで取れる")
	})
	t.Run("Add", func(t *testing.T) {
		newItem := newBookItem()
		err := stock.Add(newItem)
		require.NoError(t, err, "新しい蔵書")

		list := stock.List()
		assert.Contains(t, list, newItem, "新しい蔵書がリストに追加されている")

		existItem := list[rand.Intn(len(list))]
		err = stock.Add(existItem)
		require.Error(t, err, "すでに登録済みの蔵書")
	})
	t.Run("Remove", func(t *testing.T) {
		target := stock.List()[rand.Intn(len(stock.List()))]
		err := stock.Remove(target.ID())
		require.NoError(t, err)
		assert.NotContains(t, stock.List(), target, "リストから蔵書が削除されている")
	})
	t.Run("Find", func(t *testing.T) {
		expected := stock.List()[rand.Intn(len(stock.List()))]
		actual, ok := stock.Find(expected.ID())
		require.True(t, ok)
		assert.Equal(t, expected, actual, "蔵書が取得できる")

		_, ok = stock.Find(newBookItemID())
		assert.False(t, ok, "存在しない ID")
	})
}

func newBookID() library.BookID         { return library.BookID(testutil.RandString(10)) }
func newBookTitle() library.BookTitle   { return library.BookTitle(testutil.RandString(10)) }
func newBookItemID() library.BookItemID { return library.BookItemID(testutil.RandString(10)) }
func newBookItem() *library.BookItem    { return library.NewBookItem(newBookItemID(), newBookID()) }
