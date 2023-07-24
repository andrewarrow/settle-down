package app

import (
	"github.com/andrewarrow/feedback/router"
	"github.com/andrewarrow/feedback/util"
)

func HandleInvoices(c *router.Context, second, third string) {
	if NotLoggedIn(c) {
		return
	}
	if second == "" && third == "" && c.Method == "GET" {
		handleInvoiceIndex(c)
		return
	}
	if second != "" && third == "" && c.Method == "GET" {
		//handleInvoiceShow(c, second)
		return
	}
	if second != "" && third == "" && c.Method == "POST" {
		//handleInvoiceShowPost(c, second)
		return
	}
	c.NotFound = true
}

func handleInvoiceIndex(c *router.Context) {
	list := c.All("invoice", "order by created_at desc", "")

	colAttributes := map[int]string{}
	//colAttributes[0] = "w-1/2"

	m := map[string]any{}
	headers := []string{"name", "street", "city/state/zip", "created"}

	params := map[string]any{}
	m["headers"] = headers
	m["cells"] = c.MakeCells(util.ToAnyArray(list), headers, params, "_invoice")
	m["col_attributes"] = colAttributes

	topVars := map[string]any{}
	send := map[string]any{}
	send["bottom"] = c.Template("table_show.html", m)
	send["top"] = c.Template("invoices_list_top.html", topVars)
	c.SendContentInLayout("generic_top_bottom.html", send, 200)
}
