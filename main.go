package main

import (
	"log"

	"github.com/dkiser/go-plugin-example/plugin"
)

func playWithGreeters() {
	// Grab a new manager for dealing with "greeter" type plugins
	greeters := plugin.NewManager("greeter", "greeter-*", "./plugins/built", &plugin.GreeterPlugin{})
	defer greeters.Dispose()

	// Initialize greeters manager
	err := greeters.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Launch all greeters binaries
	greeters.Launch()

	// Lets see what all the greeters say when you say "sup" to them
	for _, pluginName := range []string{"foo", "hello"} {
		// grab a plugin by its string id
		p, err := greeters.GetInterface(pluginName)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("\n\n%s: %s plugin gives me: %s\n\n", greeters.Type, pluginName, p.(plugin.Greeter).Greet())
	}

}

func playWithClubbers() {
	// Grab a new manager for dealing with "clubber" type plugins
	clubbers := plugin.NewManager("clubber", "clubber-*", "./plugins/built", &plugin.ClubberPlugin{})
	defer clubbers.Dispose()

	// Initialize clubbers manager
	err := clubbers.Init()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Launch all clubbbers binaries
	clubbers.Launch()

	// Lets see what all the clubbers do when they party hardy!
	for _, pluginName := range []string{"raver", "cowboy"} {
		// grab a plugin by its string id
		p, err := clubbers.GetInterface(pluginName)
		if err != nil {
			log.Fatal(err.Error())
		}
		log.Printf("\n%s: %s plugin gives me: %s\n", clubbers.Type, pluginName, p.(plugin.Clubber).FistPump())
	}
}

func main() {

	// excercise some greeter plugins
	playWithGreeters()

	// excercise some clubber plugins
	playWithClubbers()

}
