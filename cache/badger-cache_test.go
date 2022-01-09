package cache

import "testing"

func TestBadgerCache_Has(t *testing.T) {
	err := testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache and it shouln't be there")
	}

	testBadgerCache.Set("foo", "bar")
	inCache, err = testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache but it should be there")
	}

	testBadgerCache.Forget("foo")
}

func TestBadgerCache_Get(t *testing.T) {
	err := testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testBadgerCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("did not get correct value from cache")
	}
}

func TestBadgerCache_Forget(t *testing.T) {
	err := testBadgerCache.Set("foo", "foo")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Forget("foo")
	if err != nil {
		t.Error(err)
	}

	inCache, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache and it shouln't be there")
	}
}

func TestBadgerCache_Empty(t *testing.T) {
	err := testBadgerCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Empty()
	if err != nil {
		t.Error(err)
	}

	hasFoo, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if hasFoo {
		t.Error("foo found in cache and it should not be there")
	}

	hasAplha, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if hasAplha {
		t.Error("alpha found in cache and it should not be there")
	}
}

func TestBadgerCache_EmptyByMatch(t *testing.T) {
	err := testBadgerCache.Set("alpha", "beta")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("alpha2", "beta2")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	err = testBadgerCache.EmptyByMatch("alph")
	if err != nil {
		t.Error(err)
	}

	hasFoo, err := testBadgerCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !hasFoo {
		t.Error("foo not found in cache but it should be there")
	}

	hasAplha, err := testBadgerCache.Has("alpha")
	if err != nil {
		t.Error(err)
	}

	if hasAplha {
		t.Error("alpha found in cache and it should not be there")
	}

	hasAplha2, err := testBadgerCache.Has("alpha2")
	if err != nil {
		t.Error(err)
	}

	if hasAplha2 {
		t.Error("alpha2 found in cache and it should not be there")
	}
}
