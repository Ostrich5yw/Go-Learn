package leetcodeProject

type WordDictionary struct {
	Next    []*WordDictionary
	Endflag bool
}

func Constructor3() WordDictionary {
	return WordDictionary{Next: make([]*WordDictionary, 26), Endflag: false}
}

func (wd *WordDictionary) AddWord(word string) {
	for i := 0; i < len(word); i++ {
		if wd.Next[word[i]-'a'] == nil {
			temp := Constructor3()
			wd.Next[word[i]-'a'] = &temp
		}
		wd = wd.Next[word[i]-'a']
	}
	wd.Endflag = true
}

func (wd *WordDictionary) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		// fmt.Println(word[i])
		if word[i] == 46 { //这里比207要多一步，用来判断作为通配符的'.'的特殊情况
			for j := 0; j < 26; j++ { //遍历26个子节点，只要有一个可以完成search，则返回true
				if wd.Next[j] != nil && wd.Next[j].Search(word[i+1:]) { //用子节点和剩余字符串完成余下的search
					return true
				}
			}
			return false
		}
		if wd.Next[word[i]-'a'] == nil {
			return false
		}
		wd = wd.Next[word[i]-'a']
	}
	return wd.Endflag
}
