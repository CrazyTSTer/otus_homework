package hw04_lru_cache //nolint:golint,stylecheck

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("Push front", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushBack(20)
		l.PushFront(30)
		first := l.Front()
		require.Equal(t, 30, first.Value)
		require.Equal(t, first.Next.Value, 10)
	})

	t.Run("Push back", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushBack(20)
		l.PushBack(30)
		first := l.Back()
		require.Equal(t, 30, first.Value)
		require.Equal(t, first.Prev.Value, 20)
	})

	t.Run("Remove", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		itemToRemove := l.PushBack(20)
		l.PushBack(30)
		require.Equal(t, 3, l.Len())
		l.Remove(itemToRemove)
		require.Equal(t, 2, l.Len())
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 30, l.Back().Value)
	})

	t.Run("Move to front", func(t *testing.T) {
		l := NewList()
		l.PushFront(10)
		l.PushBack(20)
		itemToMove := l.PushBack(30)
		l.MoveToFront(itemToMove)
		require.Equal(t, 3, l.Len())
		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 10, l.Front().Next.Value)
		require.Equal(t, 20, l.Back().Value)
	})

	t.Run("Check for Nil in Next/Prev", func(t *testing.T) {
		l := NewList()
		firstItem := l.PushFront(10)
		middleItem := l.PushBack(20)
		lastItem := l.PushBack(30)
		require.Nil(t, firstItem.Prev)
		require.NotNil(t, firstItem.Next)

		require.NotNil(t, middleItem.Prev)
		require.NotNil(t, middleItem.Next)

		require.NotNil(t, lastItem.Prev)
		require.Nil(t, lastItem.Next)
	})

}
