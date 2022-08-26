package main

import (
	"fmt"
	"image/color"
	"io/ioutil"
	"net/http"
	"task/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TITLE")
	img := canvas.NewImageFromFile("foto6.jpg")
	tab1 := container.NewTabItem("Hello", merhaba())
	tab3 := container.NewTabItemWithIcon("Weather", theme.VisibilityIcon(), weather(myWindow))
	tabs := container.NewAppTabs(
		tab1,
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), img),
		container.NewTabItemWithIcon("Open Image", theme.FileImageIcon(), openImage(myWindow)),
		tab3,
	)

	tabs.SetTabLocation(container.TabLocationLeading)
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}

func merhaba() fyne.CanvasObject {
	hello := widget.NewLabel("Hello Abdullah!")
	return container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Welcome Abdullahcığım")
		}),
	)
}
func openImage(w fyne.Window) *widget.Button {
	btn := widget.NewButton("Open foto", func() {
		fileDialog := dialog.NewFileOpen(
			func(uc fyne.URIReadCloser, _ error) {
				data, _ := ioutil.ReadAll(uc)
				res := fyne.NewStaticResource(uc.URI().Name(), data)
				img := canvas.NewImageFromResource(res)
				w := fyne.CurrentApp().NewWindow(uc.URI().Name())
				w.SetContent(img)
				w.Resize(fyne.NewSize(400, 400))
				w.Show()
			}, w)
		fileDialog.SetFilter(storage.NewExtensionFileFilter([]string{".png", ".jpg", ".jpeg"}))
		fileDialog.Show()
	})
	return btn

}

func weather(w fyne.Window) *fyne.Container {
	res, err := http.Get("https://api.openweathermap.org/data/2.5/weather?q=denizli&APPID=5bca4a5be15f6ce7f9c0793fea113e41")
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()
	byte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)

	}
	weather, err := models.UnmarshalWelcome(byte)
	if err != nil {
		fmt.Println(err)

	}
	label := canvas.NewText("WEATHER API", color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff})
	label.TextStyle = fyne.TextStyle{Bold: true}
	label.TextSize = 50
	country := canvas.NewText(fmt.Sprintf("COUNTRY: %s", weather.Sys.Country), color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff})
	province := canvas.NewText(fmt.Sprintf("PROVINCE: %s", weather.Name), color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff})
	sky := canvas.NewText(fmt.Sprintf("SKY: %s", weather.Weather[0].Main), color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff})
	temp := canvas.NewText(fmt.Sprintf("TEMP %.2f", weather.Main.Temp), color.NRGBA{R: 0xf4, G: 0x43, B: 0x36, A: 0xff})

	box := container.NewVBox()
	box.Add(label)
	box.Add(country)
	box.Add(province)
	box.Add(sky)
	box.Add(temp)
	w.SetContent(box)
	return box

}

