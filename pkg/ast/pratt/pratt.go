package pratt

// type Pratt struct {
// 	lbp [maxTok]int
// 	nud [maxTok]Nud
// 	led [maxTok]Led
// }

// func (p *Pratt) registerTok(id int, nud Nud, led Led, bindingPower int) {
// 	if nud != nil {
// 		p.nud[id] = nud
// 	}
// 	if led != nil {
// 		p.led[id] = led
// 		p.lbp[id] = bindingPower
// 	}
// 	if bindingPower > 0 && led == nil {
// 		p.lbp[id] = bindingPower
// 	}
// }

// func buildPratt(y yamlRoot, b Builder, lx *lexer) (*Pratt, error) {
// 	p := &Pratt{}
// 	id := lx.id

// 	// 1) приоритеты
// 	for pr, ent := range y.Precedence {
// 		val := pr * 10
// 		for _, name := range append([]string{ent.Op}, ent.Set...) {
// 			if name != "" {
// 				p.lbp[id(name)] = val
// 			}
// 		}
// 	}

// 	// 2) NU D
// 	p.registerTok(id("NUMBER"), NumNud{b}, nil, 0)
// 	p.registerTok(id("STRING"), StringNud{b}, nil, 0)
// 	p.registerTok(id("."), DupNud{b}, nil, 0)
// 	p.registerTok(id("-"), PrefixNud{b: b, op: "Not"}, nil, 0)

// 	// 3) LE D (просто перебираем precedence set)
// 	binaries := map[string]string{
// 		"+": "Add", "-": "Sub",
// 		"*": "Mul", "/": "Div", "%": "Mod",
// 		"==": "Eq", "!=": "Neq",
// 		"<": "Lt", "<=": "Le", ">": "Gt", ">=": "Ge",
// 		"&&": "And", "||": "Or", "|": "Pipe",
// 	}
// 	for sym, op := range binaries {
// 		idTok := id(sym)
// 		lbp := p.lbp[idTok]
// 		if sym == "|" { // особый case → Pipe()
// 			// p.led[idTok] = PipeLed{b, lbp}
// 		} else {
// 			p.led[idTok] = BinaryLed{b: b, op: op, lbp: lbp}
// 		}
// 	}

// 	// 4) постфиксы
// 	p.registerTok(id("."), nil, FieldLed{b}, 200)
// 	// p.registerTok(id("["), nil, IndexLed{b}, 200)
// 	// p.registerTok(id("("), nil, CallLed{b}, 200)

// 	return p, nil
// }

// func (pr *Parser) expr(rbp int) (asi.AST, error) {
// 	t := pr.cur.kind
// 	nud := pr.pratt.nud[t]
// 	if nud == nil {
// 		return nil, pr.err("nud missing")
// 	}
// 	left, err := nud.Parse(pr)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for rbp < pr.pratt.lbp[pr.peek().kind] {
// 		led := pr.pratt.led[pr.peek().kind]
// 		if led == nil {
// 			break
// 		}
// 		pr.next()
// 		left, err = led.Parse(pr, left)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}
// 	return left, nil
// }
