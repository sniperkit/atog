package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

var text = `[Fri Apr 13 16:55:39.577818 2018] [php5:error] [pid 20066] [client 127.0.0.1:42954] PHP Fatal error:  Function name must be a string in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 149, referer: http://suaformatura//index.php?zid=X2NvZE9yZ2FuaXphY2FvPTI0
[Mon Apr 16 08:07:30.963115 2018] [php5:error] [pid 2136] [client 127.0.0.1:55984] PHP Fatal error:  Call to undefined method Zage\\App\\Log::degub() in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:07:37.418544 2018] [php5:error] [pid 3160] [client 127.0.0.1:55988] PHP Fatal error:  Call to undefined method Zage\\App\\Log::degub() in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:07:53.162884 2018] [php5:error] [pid 4808] [client 127.0.0.1:55994] PHP Fatal error:  Call to undefined method Zage\\App\\Log::degub() in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:08:07.479562 2018] [php5:error] [pid 3161] [client 127.0.0.1:55998] PHP Fatal error:  Call to undefined method Zage\\App\\Log::degub() in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:08:55.254024 2018] [php5:error] [pid 6221] [client 127.0.0.1:56002] PHP Fatal error:  Call to undefined method Zage\\App\\Log::degub() in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:10:08.498829 2018] [php5:error] [pid 2135] [client 127.0.0.1:56014] PHP Fatal error:  Maximum execution time of 30 seconds exceeded in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 08:12:20.286550 2018] [php5:error] [pid 3191] [client 127.0.0.1:56088] PHP Fatal error:  Maximum execution time of 30 seconds exceeded in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 176, referer: http://suaformatura//index.php
[Mon Apr 16 12:03:47.104939 2018] [php5:error] [pid 17752] [client 127.0.0.1:57792] PHP Fatal error:  Cannot use isset() on the result of a function call (you can use "null !== func()" instead) in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 163, referer: http://suaformatura//administracaoepsicologia201901cesmac
[Mon Apr 16 12:05:01.150296 2018] [php5:error] [pid 6221] [client 127.0.0.1:57798] PHP Fatal error:  Cannot use isset() on the result of a function call (you can use "null !== func()" instead) in /srv/www/htdocs/suaFormatura/app/view/modulos/Fmt/php/contratoFornecAlt.php on line 163, referer: http://suaformatura//administracaoepsicologia201901cesmac`

// Execute ...
func Execute() {
	tv := tview.NewTextView()
	tv.SetDynamicColors(true)
	tv.SetRegions(true)
	tv.SetBorder(true)
	tv.SetWordWrap(true)
	fmt.Fprintln(tv, processText(text))

	app := tview.NewApplication()

	flex := tview.NewFlex()
	flex.AddItem(tview.NewBox().SetBorder(true).SetTitle(" :: FILES NAMES :: "), 0, 1, false)
	flex.AddItem(tv, 0, 5, false)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}