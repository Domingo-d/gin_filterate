/**
 * @Author: DaiGuanYu
 * @Desc:
 * @Date: 2024/9/30 11:57
 */

package model

import (
	"bufio"
	"os"
	"strings"
)

type (
	TrieNode struct {
		children map[rune]*TrieNode
		fail     *TrieNode
		isEnd    bool
		pattern  string
	}

	AhoCorasick struct {
		root *TrieNode
	}
)

// NewAhoCorasick 初始化
func NewAhoCorasick() *AhoCorasick {
	return &AhoCorasick{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

func (ac *AhoCorasick) NewAhoCorasick() *AhoCorasick {
	return NewAhoCorasick()
}

// AddPattern 添加敏感字
func (ac *AhoCorasick) AddPattern(pattern string) {
	node := ac.root
	for _, ch := range pattern {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}

		node = node.children[ch]
	}

	node.isEnd = true
	node.pattern = pattern
}

// Build 构建失败指针
func (ac *AhoCorasick) Build() {
	var queue []*TrieNode
	ac.root.fail = nil

	for _, node := range ac.root.children {
		node.fail = ac.root
		queue = append(queue, node)
	}

	for 0 < len(queue) {
		currentNode := queue[0]
		queue = queue[1:]

		for ch, childNode := range currentNode.children {
			failNode := currentNode.fail
			for failNode != nil {
				if nextNode, exists := failNode.children[ch]; exists {
					childNode.fail = nextNode
					break
				}

				failNode = failNode.fail
			}

			if childNode.fail == nil {
				childNode.fail = ac.root
			}

			queue = append(queue, childNode)
		}
	}
}

// ReadPattern 读取文件
func (ac *AhoCorasick) ReadPattern(fileName string) error {
	file, err := os.Open(fileName)
	if nil != err {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if "" != line {
			ac.AddPattern(line)
		}
	}

	return scanner.Err()
}

func (ac *AhoCorasick) SearchAndReplace(text string) string {
	node := ac.root
	runes := []rune(text)
	replacements := make([]bool, len(runes))

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		for node != nil && node != ac.root && node.children[ch] == nil {
			node = node.fail
		}

		if nextNode, exists := node.children[ch]; exists {
			node = nextNode
		} else {
			node = ac.root
		}

		tempNode := node
		for tempNode != ac.root {
			if tempNode.isEnd {
				start := i - len(tempNode.pattern) + 1
				for j := start; j <= i; j++ {
					replacements[j] = true
				}
			}

			tempNode = tempNode.fail
		}
	}

	for i := 0; i < len(runes); i++ {
		if replacements[i] {
			runes[i] = '*'
		}
	}

	return string(runes)
}
