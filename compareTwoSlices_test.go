package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCompareTwoSlices_데이터_추가_신규(t *testing.T) {
	//given
	oldSlice := make([]int64, 0)
	newSlice := []int64{1, 2}
	//when
	result := CompareTwoSlices(oldSlice, newSlice)
	// then
	assert.Equal(t, 0, len(result[SAME]))
	assert.Equal(t, 2, len(result[ADD]))
	assert.Equal(t, 0, len(result[DELETE]))
}
func TestCompareTwoSlices_데이터_추가(t *testing.T) {
	//given
	oldSlice := []int64{1, 2}
	newSlice := []int64{1, 2, 3}
	//when
	result := CompareTwoSlices(oldSlice, newSlice)
	// then
	assert.Equal(t, 2, len(result[SAME]))
	assert.Equal(t, 1, len(result[ADD]))
	assert.Equal(t, 0, len(result[DELETE]))
}
func TestCompareTwoSlices_데이터_일치(t *testing.T) {
	//given
	oldSlice := []int64{1, 2, 3}
	newSlice := []int64{1, 2, 3}
	//when
	result := CompareTwoSlices(oldSlice, newSlice)
	// then
	assert.Equal(t, 3, len(result[SAME]))
	assert.Equal(t, 0, len(result[ADD]))
	assert.Equal(t, 0, len(result[DELETE]))
}
func TestCompareTwoSlices_데이터_삭제(t *testing.T) {
	//given
	oldSlice := []int64{1, 2, 3}
	newSlice := []int64{1, 2}
	//when
	result := CompareTwoSlices(oldSlice, newSlice)
	// then
	assert.Equal(t, 2, len(result[SAME]))
	assert.Equal(t, 0, len(result[ADD]))
	assert.Equal(t, 1, len(result[DELETE]))
}
