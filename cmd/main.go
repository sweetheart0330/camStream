package main

import (
	"camStream/internal/app"
	"context"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	ctx := context.Background()
	app.Run(ctx)

}
