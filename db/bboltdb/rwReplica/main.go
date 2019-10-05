package main

func main() {
	go fsnotify.Watch(os.Path())

	botldb.Opendatabase()
}
