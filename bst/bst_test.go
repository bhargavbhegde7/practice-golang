package main

import (
  "testing"
  "strconv"
)

func TestBst(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	}

  t.Run("In order traversal recursive", func(t *testing.T) {
    got := getInOrderString(getExampleBST())
  	want := "1 3 4 6 7 8 10 13 14"

  	assertCorrectMessage(t, got, want)
	})

  t.Run("Inserting in somewhere in the middle", func(t *testing.T) {
    root := getExampleBST()
    insert(9, root)
    got := getInOrderString(root)
  	want := "1 3 4 6 7 8 9 10 13 14"

  	assertCorrectMessage(t, got, want)
	})

  t.Run("Pre order traversal recursive", func(t *testing.T) {
    got := getPreOrderString(getExampleBST())
  	want := "831647101413"

  	assertCorrectMessage(t, got, want)
	})

  t.Run("Post order traversal recursive", func(t *testing.T) {
    got := getPostOrderString(getExampleBST())
  	want := "147631314108"

  	assertCorrectMessage(t, got, want)
	})

  // t.Run("Level order traversal iterative", func(t *testing.T) {
  //   got := getLevelOrderString(getExampleBST())
  // 	want := "831016144713"
  //
  // 	assertCorrectMessage(t, got, want)
	// })

  t.Run("Tree Height check", func(t *testing.T) {
    root := getExampleBST()
    insert(12, root)
    got := strconv.Itoa(getHeight(root))
  	want := "5"

  	assertCorrectMessage(t, got, want)
	})
}
