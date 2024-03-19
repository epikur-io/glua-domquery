package domquery

import (
	"github.com/PuerkitoBio/goquery"
	lua "github.com/yuin/gopher-lua"
)

type Lselection struct {
	Sel *goquery.Selection
}

func Lselection_New(L *lua.LState, gs *goquery.Selection) *lua.LUserData {
	ud := L.NewUserData()
	ud.Value = gs
	L.SetMetatable(ud, L.GetTypeMetatable("selection"))
	return ud
}

func check_Lselection(L *lua.LState, index int) *goquery.Selection {

	ud := L.CheckUserData(index)
	if v, ok := ud.Value.(*goquery.Selection); ok {
		return v
	}
	L.ArgError(1, "domquery.selection object expected")
	return nil
}

func Lselection_Html(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html, err := sel.Html()
	if err != nil {
		L.Push(lua.LNil)
		L.Push(lua.LString(err.Error()))
		return 2
	}

	L.Push(lua.LString(html))
	return 1
}

func Lselection_Text(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	text := sel.Text()
	L.Push(lua.LString(text))
	return 1
}

func Lselection_First(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Last()))
	return 1
}

func Lselection_Last(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.First()))
	return 1
}

func Lselection_Index(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(lua.LNumber(sel.Index()))
	return 1
}

func Lselection_IndexSelector(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(lua.LNumber(sel.IndexSelector(qry)))
	return 1
}

func Lselection_Slice(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	start := L.CheckInt(2)
	end := L.CheckInt(3)
	L.Push(Lselection_New(L, sel.Slice(start, end)))
	return 1
}

/*func Lselection_Get(L *lua.LState) int {
	// !TODO
	sel := check_Lselection(L, 1)
	i := L.CheckNumber(2)
	L.Push(Lselection_New(L, sel.Get(i)))
	return 1
}*/

func Lselection_Children(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Children()))
	return 1
}

func Lselection_Parent(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Parent()))
	return 1
}

func Lselection_Parents(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Parents()))
	return 1
}

func Lselection_Each(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	f := L.CheckFunction(2)
	L.Push(Lselection_New(L, sel.Each(func(i int, s *goquery.Selection) {
		if err := L.CallByParam(lua.P{
			Fn:      f,
			NRet:    0,
			Protect: true,
		}, lua.LNumber(i), Lselection_New(L, s)); err != nil {
			panic(err)
		}
	})))
	return 1
}

func Lselection_EachBreak(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	f := L.CheckFunction(2)
	L.Push(Lselection_New(
		L,
		sel.EachWithBreak(func(i int, s *goquery.Selection) bool {
			if err := L.CallByParam(lua.P{
				Fn:      f,
				NRet:    1,
				Protect: true,
			}, lua.LNumber(i), Lselection_New(L, s)); err != nil {
				panic(err)
			}
			ret := L.Get(-1) // returned value
			L.Pop(1)         // remove received value
			if ret.Type() == lua.LTBool && ret.String() == "false" {
				return false
			}
			return true
		}),
	))
	return 1
}

func Lselection_Filter(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Filter(qry)))
	return 1
}

func Lselection_Find(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Find(qry)))
	return 1
}

func Lselection_PrependSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 1)

	L.Push(Lselection_New(L, sel.PrependSelection(sel2)))
	return 1
}

func Lselection_Eq(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckInt(2)
	L.Push(Lselection_New(L, sel.Eq(qry)))
	return 1
}

func Lselection_Not(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Not(qry)))
	return 1
}

func Lselection_Has(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Has(qry)))
	return 1
}

func Lselection_HasClass(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(lua.LBool(sel.HasClass(qry)))
	return 1
}

func Lselection_Is(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(lua.LBool(sel.Is(qry)))
	return 1
}

func Lselection_HasSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.HasSelection(sel2)))
	return 1
}

func Lselection_Closest(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Closest(qry)))
	return 1
}

func Lselection_Attr(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	name := L.CheckString(2)
	val, exists := sel.Attr(name)
	L.Push(lua.LString(val))
	L.Push(lua.LBool(exists))
	return 2
}

func Lselection_AttrOr(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	name := L.CheckString(2)
	def := L.CheckString(3)
	val := sel.AttrOr(name, def)
	L.Push(lua.LString(val))
	return 1
}

func Lselection_Append(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Append(qry)))
	return 1
}

func Lselection_AppendSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 2)
	L.Push(Lselection_New(L, sel.AppendSelection(sel2)))
	return 1
}

func Lselection_AppendHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.AppendHtml(html)))
	return 1
}

func Lselection_After(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.After(qry)))
	return 1
}

func Lselection_AfterHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.AfterHtml(html)))
	return 1
}

func Lselection_AfterSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 2)
	L.Push(Lselection_New(L, sel.AfterSelection(sel2)))
	return 1
}

func Lselection_Before(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Before(qry)))
	return 1
}

func Lselection_BeforeSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 2)
	L.Push(Lselection_New(L, sel.BeforeSelection(sel2)))
	return 1
}

func Lselection_BeforeHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.BeforeHtml(html)))
	return 1
}

func Lselection_Clone(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Clone()))
	return 1
}

func Lselection_Empty(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Empty()))
	return 1
}

func Lselection_Prepend(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.Prepend(qry)))
	return 1
}

func Lselection_PrependHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.PrependHtml(html)))
	return 1
}

/*func Lselection_PrependSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 2)
	L.Push(Lselection_New(L, sel.PrependSelection(sel2)))
	return 1
}*/

func Lselection_Remove(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	L.Push(Lselection_New(L, sel.Remove()))
	return 1
}

func Lselection_RemoveFiltered(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.RemoveFiltered(qry)))
	return 1
}

func Lselection_SetAttr(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	name := L.CheckString(2)
	val := L.CheckString(3)
	L.Push(Lselection_New(L, sel.SetAttr(name, val)))
	return 1
}

func Lselection_SetHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.SetHtml(html)))
	return 1
}

func Lselection_SetText(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	txt := L.CheckString(2)
	L.Push(Lselection_New(L, sel.SetText(txt)))
	return 1
}

func Lselection_AddClass(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	classes := []string{}
	argCnt := L.GetTop()
	if argCnt > 1 {
		for i := 2; i < argCnt; i++ {
			classes = append(classes, L.CheckString(i))
		}
	}
	L.Push(Lselection_New(L, sel.AddClass(classes...)))
	return 1
}

func Lselection_ToggleClass(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	classes := []string{}
	argCnt := L.GetTop()
	if argCnt > 1 {
		for i := 2; i < argCnt; i++ {
			classes = append(classes, L.CheckString(i))
		}
	}
	L.Push(Lselection_New(L, sel.ToggleClass(classes...)))
	return 1
}

func Lselection_ReplaceWith(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	qry := L.CheckString(2)
	L.Push(Lselection_New(L, sel.ReplaceWith(qry)))
	return 1
}

func Lselection_ReplaceWithHtml(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	html := L.CheckString(2)
	L.Push(Lselection_New(L, sel.ReplaceWithHtml(html)))
	return 1
}

func Lselection_ReplaceWithSelection(L *lua.LState) int {
	sel := check_Lselection(L, 1)
	sel2 := check_Lselection(L, 2)
	L.Push(Lselection_New(L, sel.ReplaceWithSelection(sel2)))
	return 1
}
