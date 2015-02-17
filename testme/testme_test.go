package testme

import (
    "testing"
    "strconv"
)

// This is an example test suite for the testme package.
// These tests run because the word "Test" is part of their function names.

// this test should pass
func TestMeTest1(t *testing.T) {

    result := TestMe(2, 3)
    expected_result := 6

    if result != expected_result {
        t.Error("Expected result: " + strconv.FormatInt(int64(expected_result), 10) + "  Actual result: " + strconv.FormatInt(int64(result), 10) + "\n")
    }
}

// this test should fail
func TestMeTest2(t *testing.T) {

    result := TestMe(2, 3)
    expected_result := 7

    if result != expected_result {
        t.Error("Expected result: " + strconv.FormatInt(int64(expected_result), 10) + "  Actual result: " + strconv.FormatInt(int64(result), 10) + "\n")
    }
}

