package piscine

import "github.com/01-edu/z01"

func PrintCombN(n int) {
	if n == 1 {
		for a := '0'; a <= '9'; a++ {
			z01.PrintRune(a)
			if a != '9' {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	} else if n == 2 {
		for a := '0'; a <= '8'; a++ {
			for b := a + 1; b <= '9'; b++ {
				z01.PrintRune(a)
				z01.PrintRune(b)
				if b != '9' || a != '8' {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	} else if n == 3 {
		for a := '0'; a <= '7'; a++ {
			for b := a + 1; b <= '8'; b++ {
				for c := b + 1; c <= '9'; c++ {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(c)
					if c != '9' || b != '8' || a != '7' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	} else if n == 4 {
		for a := '0'; a <= '6'; a++ {
			for b := a + 1; b <= '7'; b++ {
				for c := b + 1; c <= '8'; c++ {
					for d := c + 1; d <= '9'; d++ {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if a != '6' || b != '7' || c != '8' || d != '9' {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	} else if n == 5 {
		for a := '0'; a <= '5'; a++ {
			for b := a + 1; b <= '6'; b++ {
				for c := b + 1; c <= '7'; c++ {
					for d := c + 1; d <= '8'; d++ {
						for e := d + 1; e <= '9'; e++ {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(c)
							z01.PrintRune(d)
							z01.PrintRune(e)
							if a != '5' || b != '6' || c != '7' || d != '8' || e != '9' {
								z01.PrintRune(',')
								z01.PrintRune(' ')
							}
						}
					}
				}
			}
		}
	} else if n == 6 {
		for a := '0'; a <= '4'; a++ {
			for b := a + 1; b <= '5'; b++ {
				for c := b + 1; c <= '6'; c++ {
					for d := c + 1; d <= '7'; d++ {
						for e := d + 1; e <= '8'; e++ {
							for f := e + 1; f <= '9'; f++ {
								z01.PrintRune(a)
								z01.PrintRune(b)
								z01.PrintRune(c)
								z01.PrintRune(d)
								z01.PrintRune(e)
								z01.PrintRune(f)
								if a != '4' || b != '5' || c != '6' || d != '7' || e != '8' || f != '9' {
									z01.PrintRune(',')
									z01.PrintRune(' ')
								}
							}
						}
					}
				}
			}
		}
	} else if n == 7 {
		for a := '0'; a <= '3'; a++ {
			for b := a + 1; b <= '4'; b++ {
				for c := b + 1; c <= '5'; c++ {
					for d := c + 1; d <= '6'; d++ {
						for e := d + 1; e <= '7'; e++ {
							for f := e + 1; f <= '8'; f++ {
								for j := f + 1; j <= '9'; j++ {
									z01.PrintRune(a)
									z01.PrintRune(b)
									z01.PrintRune(c)
									z01.PrintRune(d)
									z01.PrintRune(e)
									z01.PrintRune(f)
									z01.PrintRune(j)
									if a != '3' || b != '4' || c != '5' || d != '6' || e != '7' || f != '8' || j != '9' {
										z01.PrintRune(',')
										z01.PrintRune(' ')
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 8 {
		for a := '0'; a <= '2'; a++ {
			for b := a + 1; b <= '3'; b++ {
				for c := b + 1; c <= '4'; c++ {
					for d := c + 1; d <= '5'; d++ {
						for e := d + 1; e <= '6'; e++ {
							for f := e + 1; f <= '7'; f++ {
								for j := f + 1; j <= '8'; j++ {
									for g := j + 1; g <= '9'; g++ {
										z01.PrintRune(a)
										z01.PrintRune(b)
										z01.PrintRune(c)
										z01.PrintRune(d)
										z01.PrintRune(e)
										z01.PrintRune(f)
										z01.PrintRune(j)
										z01.PrintRune(g)
										if a != '2' || b != '3' || c != '4' || d != '5' || e != '6' || f != '7' || j != '8' || g != '9' {
											z01.PrintRune(',')
											z01.PrintRune(' ')
										}
									}
								}
							}
						}
					}
				}
			}
		}
	} else if n == 9 {
		for a := '0'; a <= '1'; a++ {
			for b := a + 1; b <= '2'; b++ {
				for c := b + 1; c <= '3'; c++ {
					for d := c + 1; d <= '4'; d++ {
						for e := d + 1; e <= '5'; e++ {
							for f := e + 1; f <= '6'; f++ {
								for j := f + 1; j <= '7'; j++ {
									for g := j + 1; g <= '8'; g++ {
										for h := g + 1; h <= '9'; h++ {
											z01.PrintRune(a)
											z01.PrintRune(b)
											z01.PrintRune(c)
											z01.PrintRune(d)
											z01.PrintRune(e)
											z01.PrintRune(f)
											z01.PrintRune(j)
											z01.PrintRune(g)
											z01.PrintRune(h)
											if a != '1' || b != '2' || c != '3' || d != '4' || e != '5' || f != '6' || j != '7' || g != '8' || h != '9' {
												z01.PrintRune(',')
												z01.PrintRune(' ')
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
