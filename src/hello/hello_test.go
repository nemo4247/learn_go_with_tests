package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// 告诉测试套件，这个方法是辅助函数（helper）
		// 如此，当测试失败时所报告的行号将在函数调用中，而不是在辅助函数内部
		t.Helper()
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello world when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Foo", "French")
		want := "Bonjour, Foo"
		assertCorrectMessage(t, got, want)
	})
}
