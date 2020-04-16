package lexer

func (l *Lexer) readString() string {

	// because we do not want the double quotation here
	position := l.position + 1

	for {

		l.readChar()

		// if we encounter double quote
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}
