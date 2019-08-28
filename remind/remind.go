package remind

import "github.com/gen2brain/beeep"

func Notify(title, message string) {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)

	if err != nil {
		panic(err)
	}

	err = beeep.Notify(title, message, "assets/information.png")

	if err != nil {
		panic(err)
	}
}

func Alert() {
	err := beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)

	if err != nil {
		panic(err)
	}

	err = beeep.Alert("Title", "Warning message", "assets/warning.png")

	if err != nil {
		panic(err)
	}
}
