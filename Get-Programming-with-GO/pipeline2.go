package main

func sourceGopher(downstream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goobye all"}
	{
		downstream <- v
	}
	close(downstream)
}
func filterGopher(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	for item := range upstream {
		if !strings.Contains(item, "bad") {
			downstream <- item
		}
	}
	close(downstream)
}
func printGopher(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}