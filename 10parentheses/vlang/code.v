struct Stack {
mut:
	data []byte
}

fn (s mut Stack) push(b byte) {
	s.data << b
}

fn (s mut Stack) pop() {
	mut data := []byte
	for i := 0 ; i < s.data.len - 1; i++ {
		data << s.data[i]
	}
	s.data = data
}

fn (s Stack) top() byte {
	if s.data.len == 0 {
		return byte(0)
	}
	return s.data[s.data.len-1]
}

fn (s Stack) is_empty() bool {
	return s.data.len == 0
}

fn valid_char(c byte) bool {
	return match c {
		`{`, `}`, `[`, `]`, `(`, `)` {
			true
		}
		else {
			false
		}
	}
}

fn invert_char(c byte) byte {
	return match c {
		`{` {
			`}`
		}
		`(` {
			`)`
		}
		`[` {
			`]`
		}
		else {
			0
		}
	}
}

fn is_valid(s string) bool {
	mut stack := Stack{}
	for c in s {
		if valid_char(c) {
			if invert_char(stack.top()) == c {
				stack.pop()
			} else {
				stack.push(c)
			}
		} else {
			return false
		}
	}
	return stack.is_empty()
}


fn main() {
	// mut s := &Stack{
	// 	data: [`(`, `)`]
	// }
	// s.push(`[`)
	// s.push(`]`)
	// s.pop()
	// s.pop()
	// println(s.top())
	// println(s.data)
	// println("is empty ${s.is_empty()}")
	assert is_valid('()[]{}') == true
	assert is_valid('{}{{{{((()))}}}}') == true
	assert is_valid('[{{}}((({[]})))]') == true
	assert is_valid('[{{}}((({[]})))]') == true
	assert is_valid('[[[[[[]][[][][][][[]]') == false
	assert is_valid('([)]') == false
	assert is_valid('[)]') == false
	println('done')
}
