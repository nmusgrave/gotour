package main

import (
  "golang.org/x/tour/wc"
  "strings"
)

/*
 * Counts occurences of each word in string s
 */
func WordCount(s string) map[string]int {
  m := make(map[string]int)
  fields := strings.Fields(s)
  for _, word := range fields {
    m[word] += 1
  }
  return m
}

func runWordCount() {
  wc.Test(WordCount)
}