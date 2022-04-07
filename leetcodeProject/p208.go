package leetcodeProject

type Trie struct {
	Next    []*Trie //这里需要存放指针数组而不是对象数组
	Endflag bool    //如果是对象（[]Trie），则Insert中，改变EndFlag不会起作用（改变Next仍会起作用,因为数组存的是指针，Endflag存的是对象）
}

/*
    输入ab的情况
root.Next    a b c ... z		root 指向 a
	         |
a.Next  a b c ... z 			root 指向 b,b.EndFlag = true
		  |
b.Next a b c ... z

	没有指向下一层的，均为*Trie = nil,即空指针

*/
func Constructor2() Trie {
	return Trie{Next: make([]*Trie, 26), Endflag: false}
}

func (trie *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ { //遍历字符串
		if trie.Next[word[i]-'a'] == nil { //每一层会生成至多26*26个*Trie节点
			var temp Trie = Constructor2()
			trie.Next[word[i]-'a'] = &temp
		}
		trie = trie.Next[word[i]-'a']
		//Trie跟随字符串进行树的生成  trie（创建trie.Next[0]，trie指向trie.Next[0]，即a）->a->p->p（EndFlag=true）
	}
	trie.Endflag = true
}

func (trie *Trie) Search(word string) bool {
	for i := 0; i < len(word); i++ {
		if trie.Next[word[i]-'a'] == nil { //遍历不完字符串，说明没有该字符串
			return false
		}
		trie = trie.Next[word[i]-'a']
	}
	return trie.Endflag //遍历完成字符串，判断是否为终点，如果是存放的是apple，则app虽然可以遍历完但依旧返回false
}

func (trie *Trie) StartsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if trie.Next[prefix[i]-'a'] == nil { //前缀字符串，只要可以遍历完，都返回true
			return false
		}
		trie = trie.Next[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
